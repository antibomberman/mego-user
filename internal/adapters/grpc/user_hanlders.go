package grpc

import (
	"context"
	"fmt"
	pb "github.com/antibomberman/mego-protos/gen/go/user"
	"github.com/antibomberman/mego-user/internal/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s serverAPI) Find(ctx context.Context, req *pb.FindUserRequest) (*pb.FindUserResponse, error) {
	users, nextPageToken, err := s.service.Find(int(req.PageSize), req.PageToken, req.Sort, req.Search)
	if err != nil {
		log.Printf("Error getting posts: %v", err)
		return nil, status.Error(codes.Internal, "Failed to retrieve posts")
	}

	userResponses := make([]*pb.UserDetails, len(users))
	for i, user := range users {
		userResponses[i] = dto.ToPbUserDetail(user)
	}

	return &pb.FindUserResponse{
		Success:      true,
		Message:      "",
		Users:        userResponses,
		NexPageToken: nextPageToken,
	}, nil
}
func (s serverAPI) GetById(ctx context.Context, req *pb.Id) (*pb.UserDetails, error) {
	log.Println("GetById", req.Id)
	if req.Id == "" {
		return nil, fmt.Errorf("invalid id")
	}

	userDetails, err := s.service.GetById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	return dto.ToPbUserDetail(userDetails), nil
}
func (s serverAPI) GetByEmail(ctx context.Context, req *pb.Email) (*pb.UserDetails, error) {
	if req.Email == "" {
		return nil, fmt.Errorf("invalid email")
	}
	userDetails, err := s.service.GetByEmail(req.Email)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	return dto.ToPbUserDetail(userDetails), nil
}
func (s serverAPI) GetByPhone(ctx context.Context, req *pb.Phone) (*pb.UserDetails, error) {
	if req.Phone == "" {
		return nil, fmt.Errorf("invalid phone number")
	}
	userDetails, err := s.service.GetByPhone(req.Phone)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	return dto.ToPbUserDetail(userDetails), nil
}
func (s serverAPI) GetByToken(ctx context.Context, req *pb.Token) (*pb.UserDetails, error) {
	token := req.Token
	userDetails, err := s.service.GetByToken(token)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return dto.ToPbUserDetail(userDetails), nil
}

func (s serverAPI) Create(context.Context, *pb.CreateUserRequest) (*pb.UserDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (s serverAPI) Update(context.Context, *pb.UpdateUserRequest) (*pb.UserDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (s serverAPI) Delete(context.Context, *pb.Id) (*pb.UserDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
