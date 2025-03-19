package dto

import "time"

type DataDTO struct {
	SourceID     string `json:"source_id"`
	Payload      string `json:"payload"`
	Status       string `json:"status"`
	Format       string `json:"format"`
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Size         int64  `json:"size"`
	Category     string `json:"category"`
}

type MetadataDTO struct {
	DataID    string    `json:"data_id"`
	Source    string    `json:"source"`
	Format    string    `json:"format"`
	Size      int64     `json:"size"`
	Timestamp time.Time `json:"timestamp"`
	Category  string    `json:"category"`
}

type DataResponseDTO struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ErrorReportDTO struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	DataID       string `json:"data_id"`
}
