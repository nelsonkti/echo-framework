package controller

import (
	"echo-framework/logic/http/model"
	"echo-framework/logic/http/service"
	"echo-framework/logic/http/validator"
	"echo-framework/util/xrsp"
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

	var res []model.EmployeesBase
	EmployeesBase := model.EmployeesBase{}
	query := EmployeesBase.Model()

	query = query.
		Where("id = ?", 1).
		Limit(10).
		Order("id asc").
		Find(&res)

	pageResponse.Data = res

	return context.JSON(http.StatusOK, xrsp.Data(pageResponse))
}

func GetHello2(ctx echo.Context) error {

	type signinStaffRequest struct {
		Username uint64 `form:"user_name" json:"user_name" query:"user_name" comment:"用户名" validate:"required"`
		Nickname uint64 `form:"nick_name" json:"nick_name" query:"nick_name" comment:"昵称" validate:"required"`
	}
	var data signinStaffRequest
	err := ctx.Bind(&data)

	if err != nil {
		return ctx.JSON(http.StatusLocked, xrsp.Error(err))
	}

	return ctx.String(http.StatusOK, "Hello, World")
}

func CreateUser(ctx echo.Context) error {
	var requestData validator.UserRequest
	err := ctx.Bind(&requestData)

	if err != nil {
		return ctx.JSON(http.StatusLocked, xrsp.Error(err))
	}

	var userService service.UserService
	res := userService.Create(requestData)

	return ctx.JSON(res.Status, res)
}