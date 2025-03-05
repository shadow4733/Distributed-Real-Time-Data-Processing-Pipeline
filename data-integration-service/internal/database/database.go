package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/ClickHouse/clickhouse-go/v2"
)

func ConnectClickHouse() (*sql.DB, error) {
	host := os.Getenv("CLICKHOUSE_HOST")
	port := os.Getenv("CLICKHOUSE_PORT")
	user := os.Getenv("CLICKHOUSE_USER")
	password := os.Getenv("CLICKHOUSE_PASSWORD")
	database := os.Getenv("CLICKHOUSE_DATABASE")

	dsn := fmt.Sprintf("clickhouse://%s:%s@%s:%s/%s", user, password, host, port, database)
	db, err := sql.Open("clickhouse", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Не удалось подключиться к ClickHouse: %v", err)
		return nil, err
	}

	log.Println("Успешное подключение к ClickHouse")
	return db, nil
}
