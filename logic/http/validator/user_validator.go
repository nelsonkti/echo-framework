// Package validator
// @Author fuzengyao
// @Date 2022-11-09 11:15:24
package validator

type UserRequest struct {
	Username string `form:"user_name" json:"user_name" query:"user_name" comment:"用户名" validate:"required"`
}
