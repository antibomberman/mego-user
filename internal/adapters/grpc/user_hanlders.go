package grpc

import (
	"context"
	"fmt"
	pb "github.com/antibomberman/mego-protos/gen/go/user"
	"github.com/antibomberman/mego-user/internal/dto"
	"github.com/antibomberman/mego-user/internal/models"
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
func (s serverAPI) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserDetails, error) {
	user := models.CreateUserRequest{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
	}
	userDetails, err := s.service.Create(&user)
	if err != nil {
		return nil, err
	}
	return dto.ToPbUserDetail(userDetails), nil
}
func (s serverAPI) Update(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserDetails, error) {
	user := models.UpdateUserRequest{
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
	if req.Avatar != nil {
		user.Avatar = &models.NewAvatar{
			FileName: req.Avatar.FileName,
			Data:     req.Avatar.Data,
		}
	}
	userDetails, err := s.service.Update(req.Id, &user)
	if err != nil {
		return nil, err
	}
	return dto.ToPbUserDetail(userDetails), nil
}
func (s serverAPI) Delete(ctx context.Context, req *pb.Id) (*pb.UserDetails, error) {
	err := s.service.Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
