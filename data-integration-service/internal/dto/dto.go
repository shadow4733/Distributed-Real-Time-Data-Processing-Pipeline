package dto

type DataDTO struct {
	SourceID     string `json:"source_id"`
	Payload      string `json:"payload"`
	Status       string `json:"status"`
	Format       string `json:"format"`
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Size         int64  `json:"size"`
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
