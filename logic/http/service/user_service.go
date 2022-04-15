/**
** @创建时间 : 2022/1/6 10:48
** @作者 : fzy
 */
package service

import (
	"github.com/nelsonkti/echo-framework/logic/http/model"
	"github.com/nelsonkti/echo-framework/logic/http/repository"
	"github.com/nelsonkti/echo-framework/logic/http/validator"
	"github.com/nelsonkti/echo-framework/util/xrsp"
)

type UserService struct {
}

func (us *UserService) Create(requestData validator.UserRequest) xrsp.Response {
	if requestData.Username == "sb" {
		return xrsp.ErrorText("操作失败，命名的名称不规范")
	}

	var userRepository repository.UserRepository

	var userModel model.UserModel
	userModel.Username = requestData.Username
	err := userRepository.Create(userModel)
	if err != nil {
		return xrsp.Error(err)
	}

	return xrsp.Nil()
}
