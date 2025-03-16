-- init_clickhouse.sql

CREATE DATABASE IF NOT EXISTS data_pipeline;

CREATE TABLE IF NOT EXISTS data_pipeline.DataIngestionModel (
    DataID String,
    SourceID String,
    DataPayload String,
    Timestamp DateTime,
    Status String,
    ProcessedAt DateTime,
    ErrorCode Nullable(String) DEFAULT NULL,
    ErrorMessage Nullable(String) DEFAULT NULL
) ENGINE = MergeTree()
ORDER BY (DataID);

CREATE TABLE IF NOT EXISTS data_pipeline.DataMetadataModel (
    DataID String,
    Source String,
    Format String,
    Size Int64,
    Timestamp DateTime
) ENGINE = MergeTree()
ORDER BY (DataID);

CREATE TABLE IF NOT EXISTS data_pipeline.ErrorReportModel (
    ErrorCode String,
    ErrorMessage String,
    DataID String,
    Timestamp DateTime
) ENGINE = MergeTree()
ORDER BY (DataID);
