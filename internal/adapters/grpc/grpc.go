package grpc

import (
	"context"
	userGrpc "github.com/antibomberman/mego-protos/gen/go/user"
	"github.com/antibomberman/mego-user/internal/config"
	"github.com/antibomberman/mego-user/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type serverAPI struct {
	userGrpc.UnimplementedUserServiceServer
	service services.UserService
	cfg     *config.Config
}

func Register(gRPC *grpc.Server, cfg *config.Config, service services.UserService) {
	userGrpc.RegisterUserServiceServer(gRPC, &serverAPI{
		service: service,
		cfg:     cfg,
	})
}
func (s serverAPI) Find(ctx context.Context, req *userGrpc.FindUserRequest) (*userGrpc.FindUserResponse, error) {
	users, nextPageToken, err := s.service.Find(int(req.PageSize), req.PageToken, req.Search)
	if err != nil {
		log.Printf("Error getting posts: %v", err)
		return nil, status.Error(codes.Internal, "Failed to retrieve posts")
	}

	userResponses := make([]*userGrpc.UserDetails, len(users))
	for i, user := range users {
		userResponses[i] = &userGrpc.UserDetails{
			FirstName: user.FirstName,
			//CreatedAt: post.CreatedAt.Unix(),
		}
	}

	return &userGrpc.FindUserResponse{
		Users:        userResponses,
		NexPageToken: nextPageToken,
	}, nil
}

func (s serverAPI) GetById(ctx context.Context, req *userGrpc.Id) (*userGrpc.UserDetails, error) {
	userDetails, err := s.service.GetById(req.Id)
	log.Printf("Error getting user: %v", err)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	return &userGrpc.UserDetails{
		FirstName: userDetails.FirstName,
		//CreatedAt: post.CreatedAt.Unix(),
	}, nil

}
func (s serverAPI) GetByEmail(context.Context, *userGrpc.Email) (*userGrpc.UserDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByEmail not implemented")
}
func (s serverAPI) GetByToken(context.Context, *userGrpc.Token) (*userGrpc.UserDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByToken not implemented")
}
func (s serverAPI) Create(context.Context, *userGrpc.CreateUserRequest) (*userGrpc.UserDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (s serverAPI) Update(context.Context, *userGrpc.UpdateUserRequest) (*userGrpc.UserDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (s serverAPI) Delete(context.Context, *userGrpc.Id) (*userGrpc.UserDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
