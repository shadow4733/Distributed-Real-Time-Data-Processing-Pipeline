package service

import (
	"context"
	"data-integration-service/internal/dto"
	"data-integration-service/internal/model/enum"
	pb "data-integration-service/internal/proto"
	"data-integration-service/internal/repository"
	"encoding/json"
	"encoding/xml"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func detectFormat(payload []byte) string {
	trimmedPayload := strings.TrimSpace(string(payload))

	if strings.HasPrefix(trimmedPayload, "{") && strings.HasSuffix(trimmedPayload, "}") ||
		strings.HasPrefix(trimmedPayload, "[") && strings.HasSuffix(trimmedPayload, "]") {
		var js interface{}
		if json.Unmarshal(payload, &js) == nil {
			return "json"
		}
	}

	if strings.HasPrefix(trimmedPayload, "<") && strings.HasSuffix(trimmedPayload, ">") {
		var x interface{}
		if xml.Unmarshal(payload, &x) == nil {
			return "xml"
		}
	}

	return "unknown"
}

func (s *DataProcessingService) ProcessData(ctx context.Context, req *pb.DataRecord) (*pb.ProcessResponse, error) {
	dataID := uuid.New().String()

	timestamp := time.Now()
	processedAt := time.Now()

	log.Printf("Received DataRecord: DataID=%s, SourceID=%s, Status=%s, ErrorCode=%d, ErrorMessage=%s",
		req.DataID, req.SourceId, req.Status, req.ErrorCode, req.ErrorMessage)

	if len(req.DataPayload) == 0 {
		log.Printf("Ошибка: dataPayload пуст")
		return &pb.ProcessResponse{Success: false, Message: "dataPayload пуст"}, nil
	}

	format := detectFormat(req.DataPayload)
	log.Printf("Определен формат данных: %s", format)

	categoryMap := map[string]enum.Category{
		"weather":    enum.Weather,
		"finance":    enum.Finance,
		"health":     enum.Health,
		"technology": enum.Technology,
	}

	category, exists := categoryMap[req.Category]
	if !exists {
		log.Printf("Неизвестная категория: %s", req.Category)
		category = enum.Unknown
	}

	data := dto.DataDTO{
		SourceID:     req.SourceId,
		Payload:      string(req.DataPayload),
		Status:       req.Status,
		Format:       format,
		ErrorCode:    string(int32(req.ErrorCode)),
		ErrorMessage: req.ErrorMessage,
		Size:         int64(len(req.DataPayload)),
		Category:     string(category),
	}

	err := s.repo.InsertData(dataID, timestamp, processedAt, data)
	if err != nil {
		log.Printf("Ошибка при обработке данных: %v", err)
		return &pb.ProcessResponse{Success: false, Message: "Ошибка при обработке данных"}, err
	}

	log.Printf("Данные успешно обработаны: DataID=%s", dataID)
	return &pb.ProcessResponse{Success: true, Message: "Данные успешно обработаны"}, nil
}

func (s *DataProcessingService) GetMetadata(ctx context.Context, req *pb.MetadataRequest) (*pb.MetadataResponse, error) {
	log.Printf("Запрос метаданных для DataID: %s", req.DataID)

	metadata, err := s.repo.GetMetadata(req.DataID)
	if err != nil {
		log.Printf("Ошибка при получении метаданных: %v", err)
		return nil, err
	}

	response := &pb.MetadataResponse{
		DataID:    metadata.DataID,
		Source:    metadata.Source,
		Format:    metadata.Format,
		Size:      metadata.Size,
		Timestamp: timestamppb.New(metadata.Timestamp),
		Category:  metadata.Category,
	}

	log.Printf("Метаданные успешно получены для DataID: %s", req.DataID)
	return response, nil
}
