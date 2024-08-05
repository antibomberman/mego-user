package dto

import (
	pb "github.com/antibomberman/mego-protos/gen/go/user"
	"github.com/antibomberman/mego-user/internal/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToPbUserDetail(details *models.UserDetails) *pb.UserDetails {
	pbUserDetails := &pb.UserDetails{
		Id:         details.Id,
		FirstName:  details.FirstName,
		MiddleName: details.MiddleName,
		LastName:   details.LastName,
		Email:      details.Email,
		Phone:      details.Phone,
		Avatar:     details.Avatar,
		CreatedAt:  timestamppb.New(details.CreatedAt),
		UpdatedAt:  timestamppb.New(details.UpdatedAt),
		DeletedAt:  timestamppb.New(details.DeletedAt),
	}
	return pbUserDetails
}
func ToUserDetail(user *models.User) *models.UserDetails {
	return &models.UserDetails{
		Id:         user.Id,
		FirstName:  user.FirstName.String,
		MiddleName: user.MiddleName.String,
		LastName:   user.LastName.String,
		Email:      user.Email.String,
		Phone:      user.Phone.String,
		Avatar:     user.Avatar.String,
		DeletedAt:  user.DeletedAt.Time,
		CreatedAt:  user.CreatedAt.Time,
		UpdatedAt:  user.UpdatedAt.Time,
	}
}
