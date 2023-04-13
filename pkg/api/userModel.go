package api

import (
	"demo04/internal/repository"
	"demo04/pkg/errno"
)

type UserAuthRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type UserAuthResponse struct {
	User  *repository.User
	Token string
	errno.ErrNo
}

func BuildUserAuthResponse(user *repository.User, token string) *UserAuthResponse {
	return &UserAuthResponse{
		User:  user,
		Token: token,
	}
}
