package controller

import (
	"github.com/labstack/echo/v4"
	"echo-framework/logic/http/model"
	"echo-framework/logic/http/response"
	"net/http"
)

// Handler
func GetHello(context echo.Context) error {

	type getUserRequest struct {
		PageRequest
	}

	var data getUserRequest
	_ = context.Bind(&data)

	//pageRequest := DefaultPageRequest(data.PageRequest)
	pageResponse := NewPageResponse(data.PageRequest)

	var res []model.EmployeesBase
	//
	EmployeesBase := model.EmployeesBase{}
	query := EmployeesBase.Model()

	query = query.
		Limit(10).
		Order("id asc").
		Where("id = ?", 817).
		Preload("EmployeesDetail").
		Find(&res)


	pageResponse.Data = res

	return context.JSON(http.StatusOK, response.Success(pageResponse))
}

func GetHello2(ctx echo.Context) error {
	type signinStaffRequest struct {
		Username uint64 `form:"user_name" json:"user_name" query:"user_name" comment:"用户名" validate:"required"`
		Nickname uint64 `form:"nick_name" json:"nick_name" query:"nick_name" comment:"昵称" validate:"required"`
	}
	var data signinStaffRequest
	err := ctx.Bind(&data)

	if err != nil {
		return ctx.JSON(http.StatusLocked, response.Fail(err.Error()))
	}

	return ctx.String(http.StatusOK, "Hello, World!23334XDS")
}
