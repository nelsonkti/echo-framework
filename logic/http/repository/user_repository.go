// Package repository
// @Author fuzengyao
// @Date 2022-11-09 11:15:09
package repository

import "github.com/nelsonkti/echo-framework/logic/http/model"

type UserRepository struct {
	Model *model.UserModel
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		Model: &model.UserModel{},
	}
}

func (ur *UserRepository) Create(data model.UserModel) error {
	var userModel model.UserModel
	err := userModel.Model().Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}
