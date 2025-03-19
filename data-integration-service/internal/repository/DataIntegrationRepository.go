package repository

import (
	"data-integration-service/internal/dto"
	"database/sql"
	"log"
	"time"
)

type DataIntegrationRepository struct {
	db *sql.DB
}

func NewDataIntegrationRepository(db *sql.DB) *DataIntegrationRepository {
	return &DataIntegrationRepository{db: db}
}

func (r *DataIntegrationRepository) InsertData(dataID string, timestamp time.Time, processedAt time.Time, data dto.DataDTO) error {
	query := `
    INSERT INTO data_pipeline."DataIngestionModel" (DataID, SourceID, DataPayload, Status, Timestamp, ProcessedAt, ErrorCode, ErrorMessage, Category) 
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
    `
	_, err := r.db.Exec(query, dataID, data.SourceID, data.Payload, data.Status, timestamp, processedAt, data.ErrorCode, data.ErrorMessage, data.Category)
	if err != nil {
		log.Printf("Ошибка вставки данных в DataIngestionModel в ClickHouse: %v", err)
		return err
	}

	metadataQuery := `
    INSERT INTO data_pipeline."DataMetadataModel" (DataID, Source, Format, Size, Timestamp, Category) 
    VALUES (?, ?, ?, ?, ?, ?)
    `
	_, err = r.db.Exec(metadataQuery, dataID, data.SourceID, data.Format, data.Size, timestamp, data.Category)
	if err != nil {
		log.Printf("Ошибка вставки метаданных в DataMetadataModel в ClickHouse: %v", err)
		return err
	}

	if data.ErrorCode != "" && data.ErrorMessage != "" {
		errorReportQuery := `
        INSERT INTO data_pipeline."ErrorReportModel" (DataID, ErrorCode, ErrorMessage, Timestamp) 
        VALUES (?, ?, ?, ?)
        `
		_, err = r.db.Exec(errorReportQuery, dataID, data.ErrorCode, data.ErrorMessage, timestamp)
		if err != nil {
			log.Printf("Ошибка вставки отчета об ошибке в ErrorReportModel в ClickHouse: %v", err)
			return err
		}
	}

	return nil
}

func (r *DataIntegrationRepository) GetMetadata(dataId string) (dto.MetadataDTO, error) {
	var metadata dto.MetadataDTO
	query := `
    SELECT DataID, Source, Format, Size, Timestamp, Category 
    FROM data_pipeline."DataMetadataModel" 
    WHERE DataID = ?
    `
	err := r.db.QueryRow(query, dataId).Scan(&metadata.DataID, &metadata.Source, &metadata.Format, &metadata.Size, &metadata.Timestamp, &metadata.Category)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Метаданные для DataID=%s не найдены", dataId)
			return dto.MetadataDTO{}, nil
		}
		log.Printf("Ошибка при получении метаданных: %v", err)
		return dto.MetadataDTO{}, err
	}

	return metadata, nil
}
