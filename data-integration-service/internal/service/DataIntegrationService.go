package service

import (
	"context"
	"data-integration-service/internal/dto"
	pb "data-integration-service/internal/proto"
	"data-integration-service/internal/repository"
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

// detectFormat определяет формат данных (json, xml и т.д.)
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
	// Генерация dataID
	dataID := uuid.New().String()

	// Текущее время
	timestamp := time.Now()
	processedAt := time.Now()

	// Логирование входящих данных
	log.Printf("Received DataRecord: DataID=%s, SourceID=%s, Status=%s, ErrorCode=%d, ErrorMessage=%s",
		req.DataID, req.SourceId, req.Status, req.ErrorCode, req.ErrorMessage)

	// Проверка данных перед обработкой
	if len(req.DataPayload) == 0 {
		log.Printf("Ошибка: dataPayload пуст")
		return &pb.ProcessResponse{Success: false, Message: "dataPayload пуст"}, nil
	}

	// Определение формата данных
	format := detectFormat(req.DataPayload)
	log.Printf("Определен формат данных: %s", format)

	// Создание DTO
	data := dto.DataDTO{
		SourceID:     req.SourceId,
		Payload:      string(req.DataPayload), // Передаем сырые данные как строку
		Status:       req.Status,
		Format:       format,
		ErrorCode:    string(int32(req.ErrorCode)), // Преобразуем int32 в string (если нужно)
		ErrorMessage: req.ErrorMessage,
		Size:         int64(len(req.DataPayload)), // Размер данных
	}

	// Вставка данных в репозиторий
	err := s.repo.InsertData(dataID, timestamp, processedAt, data)
	if err != nil {
		log.Printf("Ошибка при обработке данных: %v", err)
		return &pb.ProcessResponse{Success: false, Message: "Ошибка при обработке данных"}, err
	}

	log.Printf("Данные успешно обработаны: DataID=%s", dataID)
	return &pb.ProcessResponse{Success: true, Message: "Данные успешно обработаны"}, nil
}
