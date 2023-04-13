package service

import (
	"demo04/internal/repository"
	"demo04/pkg/api"
	"demo04/pkg/util/jwt"
	"errors"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) CheckUserExist(id string) bool {
	user := &repository.User{
		Model: repository.Model{
			Id: id,
		},
	}
	return user.CheckUserIdExist()
}

func (*UserService) UserRegister(req *api.UserAuthRequest) (resp *api.UserAuthResponse, err error) {
	user := &repository.User{
		Username: req.Username,
		Password: req.Password,
	}
	if user.CheckUsernameExist() {
		return nil, errors.New("用户名已存在")
	}
	err = user.SetPassword()
	user, err = user.UserRegister()
	token, err := jwt.GenerateToken(user.Model.Id)
	resp = api.BuildUserAuthResponse(user, token)
	return resp, err
}

// UserLogin	用户登录
func (*UserService) UserLogin(req *api.UserAuthRequest) (resp *api.UserAuthResponse, err error) {
	user := &repository.User{
		Username: req.Username,
		Password: req.Password,
	}
	user, err = user.UserLogin()
	if err != nil {
		return nil, err
	}
	token, err := jwt.GenerateToken(user.Model.Id)
	resp = api.BuildUserAuthResponse(user, token)
	return resp, err
}
