package grpc

import (
	pb "github.com/antibomberman/mego-protos/gen/go/user"
	"github.com/antibomberman/mego-user/internal/config"
	"github.com/antibomberman/mego-user/internal/services"
	"google.golang.org/grpc"
)

type serverAPI struct {
	pb.UnimplementedUserServiceServer

	service services.UserService
	cfg     *config.Config
}

func Register(gRPC *grpc.Server, cfg *config.Config, service services.UserService) {
	pb.RegisterUserServiceServer(gRPC, &serverAPI{
		service: service,
		cfg:     cfg,
	})

}
