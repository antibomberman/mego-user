package dto

import (
	pb "github.com/antibomberman/mego-protos/gen/go/user"
	"github.com/antibomberman/mego-user/internal/models"
)

func UserDetail(details *models.UserDetails) *pb.UserDetails {
	pbUserDetails := &pb.UserDetails{
		Id:    details.Id,
		Email: details.Email,
	}

	return pbUserDetails
}
