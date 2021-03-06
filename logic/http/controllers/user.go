package controllers

import (
	"echo-framework/logic/http/models"
	"echo-framework/logic/http/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Handler
func GetHello(context echo.Context) error {

	type getUserRequest struct {
		PageRequest
	}

	var data getUserRequest
	_ = context.Bind(&data)

	pageResponse := NewPageResponse(data.PageRequest)

	var res []models.EmployeesBase
	EmployeesBase := models.EmployeesBase{}
	query := EmployeesBase.Model()

	query = query.
		Where("id = ?", 1).
		Limit(10).
		Order("id asc").
		Find(&res)

	pageResponse.Data = res

	return context.JSON(http.StatusOK, responses.Success(pageResponse))
}

func GetHello2(ctx echo.Context) error {

	type signinStaffRequest struct {
		Username uint64 `form:"user_name" json:"user_name" query:"user_name" comment:"用户名" validate:"required"`
		Nickname uint64 `form:"nick_name" json:"nick_name" query:"nick_name" comment:"昵称" validate:"required"`
	}
	var data signinStaffRequest
	err := ctx.Bind(&data)

	if err != nil {
		return ctx.JSON(http.StatusLocked, responses.Fail(err.Error()))
	}

	return ctx.String(http.StatusOK, "Hello, World")
}
