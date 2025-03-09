package main

import (
	pb "cat_alog/src/api/grpc"
	"cat_alog/src/infrastructure/cassandra"
	"cat_alog/src/internal/handler"
	"cat_alog/src/internal/service"
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
	cassandraCheck, err := cassandra.CheckCassandraConnection()
	if !cassandraCheck || err != nil {
		log.Fatalf("Cassandra connection failed: %v", err)
	}
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
