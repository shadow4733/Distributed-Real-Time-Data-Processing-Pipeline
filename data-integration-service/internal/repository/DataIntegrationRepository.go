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
    INSERT INTO data_pipeline."DataIngestionModel" (DataID, SourceID, DataPayload, Status, Timestamp, ProcessedAt, ErrorCode, ErrorMessage) 
    VALUES (?, ?, ?, ?, ?, ?, ?, ?)
    `
	_, err := r.db.Exec(query, dataID, data.SourceID, data.Payload, data.Status, timestamp, processedAt, data.ErrorCode, data.ErrorMessage)
	if err != nil {
		log.Printf("Ошибка вставки данных в DataIngestionModel в ClickHouse: %v", err)
		return err
	}

	metadataQuery := `
    INSERT INTO data_pipeline."DataMetadataModel" (DataID, Source, Format, Size, Timestamp) 
    VALUES (?, ?, ?, ?, ?)
    `
	_, err = r.db.Exec(metadataQuery, dataID, data.SourceID, data.Format, data.Size, timestamp)
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
