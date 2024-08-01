package main

import (
	adapter "github.com/antibomberman/mego-user/internal/adapters/grpc"
	"github.com/antibomberman/mego-user/internal/config"
	"github.com/antibomberman/mego-user/internal/repositories"
	"github.com/antibomberman/mego-user/internal/services"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cfg := config.Load()
	log.Printf("Config: %+v", cfg)
	db, err := sqlx.Open("postgres", cfg.DatabaseURL)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)

	l, err := net.Listen("tcp", ":"+cfg.ServerPort)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gRPC := grpc.NewServer()
	adapter.Register(gRPC, cfg, userService)
	if err := gRPC.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
