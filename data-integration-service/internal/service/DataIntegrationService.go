package service

import (
	"context"
	"data-integration-service/internal/dto"
	pb "data-integration-service/internal/proto"
	"data-integration-service/internal/repository"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"github.com/google/uuid"
	"log"
	"strings"
	"time"
)

type DataProcessingService struct {
	repo *repository.DataIntegrationRepository
	pb.UnimplementedDataProcessingServiceServer
}

func NewDataProcessingService(repo *repository.DataIntegrationRepository) *DataProcessingService {
	return &DataProcessingService{repo: repo}
}

func detectFormat(payload string) string {
	trimmedPayload := strings.TrimSpace(payload)

	if strings.HasPrefix(trimmedPayload, "{") && strings.HasSuffix(trimmedPayload, "}") ||
		strings.HasPrefix(trimmedPayload, "[") && strings.HasSuffix(trimmedPayload, "]") {
		var js interface{}
		if json.Unmarshal([]byte(trimmedPayload), &js) == nil {
			return "json"
		}
	}

	if strings.HasPrefix(trimmedPayload, "<") && strings.HasSuffix(trimmedPayload, ">") {
		var x interface{}
		if xml.Unmarshal([]byte(trimmedPayload), &x) == nil {
			return "xml"
		}
	}

	return "unknown"
}

func (s *DataProcessingService) ProcessData(ctx context.Context, req *pb.DataRecord) (*pb.ProcessResponse, error) {
	dataID := uuid.New().String()

	// Текущее время
	timestamp := time.Now()
	processedAt := time.Now()

	decodedPayload, err := base64.StdEncoding.DecodeString(string(req.DataPayload))
	if err != nil {
		log.Printf("Ошибка декодирования base64: %v", err)
		return &pb.ProcessResponse{Success: false, Message: "Ошибка декодирования данных"}, err
	}

	format := detectFormat(string(decodedPayload))

	size := int64(len(decodedPayload))

	data := dto.DataDTO{
		SourceID:     req.SourceId,
		Payload:      string(decodedPayload),
		Status:       req.Status,
		Format:       format,
		ErrorCode:    string(req.ErrorCode),
		ErrorMessage: req.ErrorMessage,
		Size:         size,
	}

	err = s.repo.InsertData(dataID, timestamp, processedAt, data)
	if err != nil {
		log.Printf("Ошибка при обработке данных: %v", err)
		return &pb.ProcessResponse{Success: false, Message: "Ошибка при обработке данных"}, err
	}

	return &pb.ProcessResponse{Success: true, Message: "Данные успешно обработаны"}, nil
}

func (s *DataProcessingService) GetMetadata(ctx context.Context, req *pb.MetadataRequest) (*pb.MetadataResponse, error) {

}
