// Package service
// @Author fuzengyao
// @Date 2022-11-09 11:15:15
package service

import (
	"github.com/nelsonkti/echo-framework/logic/http/repository"
	"github.com/nelsonkti/echo-framework/logic/http/validator"
	"github.com/nelsonkti/echo-framework/util/xrsp"
)

type UserService interface {
	Create(validator.UserRequest) xrsp.Response
}

type userService struct {
	repo *repository.UserRepository
}

func NewUserService() UserService {
	return &userService{
		repo: repository.NewUserRepository(),
	}
}

func (s *userService) Create(requestData validator.UserRequest) xrsp.Response {
	if requestData.Username == "sb" {
		return xrsp.ErrorText("操作失败，命名的名称不规范")
	}

	userModel := *s.repo.Model
	userModel.Username = requestData.Username
	err := s.repo.Create(userModel)
	if err != nil {
		return xrsp.Error(err)
	}

	return xrsp.Nil()
}
