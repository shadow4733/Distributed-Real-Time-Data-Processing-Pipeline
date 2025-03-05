package models

import "time"

type DataIngestionModel struct {
	DataID       string    `json:"DataID" ch:"DataID"`
	SourceID     string    `json:"SourceID" ch:"SourceID"`
	DataPayload  string    `json:"DataPayload" ch:"DataPayload"`
	Timestamp    time.Time `json:"Timestamp" ch:"Timestamp"`
	Status       string    `json:"Status" ch:"Status"`
	ProcessedAt  time.Time `json:"ProcessedAt" ch:"ProcessedAt"`
	ErrorCode    string    `json:"ErrorCode,omitempty" ch:"ErrorCode"`
	ErrorMessage string    `json:"ErrorMessage,omitempty" ch:"ErrorMessage"`
}

type DataMetadataModel struct {
	DataID    string    `json:"DataID" ch:"DataID"`
	Source    string    `json:"Source" ch:"Source"`
	Format    string    `json:"Format" ch:"Format"`
	Size      int64     `json:"Size" ch:"Size"`
	Timestamp time.Time `json:"Timestamp" ch:"Timestamp"`
}

type ErrorReportModel struct {
	ErrorCode    string    `json:"ErrorCode" ch:"ErrorCode"`
	ErrorMessage string    `json:"ErrorMessage" ch:"ErrorMessage"`
	DataID       string    `json:"DataID" ch:"DataID"`
	Timestamp    time.Time `json:"Timestamp" ch:"Timestamp"`
}
