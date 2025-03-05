package main

import (
	"data-integration-service/internal/database"
	pb "data-integration-service/internal/proto"
	"data-integration-service/internal/repository"
	"data-integration-service/internal/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	db, err := database.ConnectClickHouse()
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	repo := repository.NewDataIntegrationRepository(db)
	service := service.NewDataProcessingService(repo)

	grpcServer := grpc.NewServer()

	pb.RegisterDataProcessingServiceServer(grpcServer, service)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}

	log.Println("gRPC сервер запущен на порту 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Ошибка запуска gRPC сервера: %v", err)
	}
}
