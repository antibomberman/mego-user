package main

import (
	"context"
	"fmt"
	adapter "github.com/antibomberman/mego-user/internal/adapters/grpc"
	"github.com/antibomberman/mego-user/internal/clients"
	"github.com/antibomberman/mego-user/internal/config"
	"github.com/antibomberman/mego-user/internal/database"
	"github.com/antibomberman/mego-user/internal/repositories"
	"github.com/antibomberman/mego-user/internal/services"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cfg := config.Load()
	db, err := database.ConnectToDB(cfg)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: "",
		DB:       0,
	})

	err = rdb.Ping(context.Background()).Err()
	if err != nil {
		log.Fatal("redis connection error")
	}
	log.Println("Connected to Redis")
	authClient, err := clients.NewPostClient(cfg.AuthServerAddress)

	userRepository := repositories.NewUserRepository(db, rdb)
	userService := services.NewUserService(userRepository, authClient)

	l, err := net.Listen("tcp", ":"+cfg.UserServiceServerPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gRPC := grpc.NewServer()
	adapter.Register(gRPC, cfg, userService)
	log.Println("Server started on port", cfg.UserServiceServerPort)
	if err := gRPC.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("Server stopped")
}
