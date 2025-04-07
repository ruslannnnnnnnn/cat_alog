package main

import (
	"cat_alog/internal/api/grpc"
	"cat_alog/internal/domain/service"
	"cat_alog/internal/infrastructure/cassandra"
	"cat_alog/internal/interfaces/grpc/handler"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

const envPath = "/app/.env"

func main() {
	if err := godotenv.Load(envPath); err != nil {
		log.Fatal("Error loading .env file")
	}
	connection, err := cassandra.GetCassandraSession()
	if err != nil {
		log.Fatalf("Cassandra connection failed: %v", err)
	}
	connection.Close()
	fmt.Println("Cassandra connection successful")

	repo := cassandra.NewCatRepository()
	catService := service.NewCatService(repo)
	grpcCatHandler := handler.NewGrpcCatHandler(catService)

	grpcServer := grpc.NewServer()
	pb.RegisterCatServiceServer(grpcServer, grpcCatHandler)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("gRPC server starting on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
