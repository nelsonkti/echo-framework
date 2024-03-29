package routes

import (
	"context"
	"github.com/nelsonkti/echo-framework/lib/logger"
	"github.com/nelsonkti/echo-framework/logic/http/controller"
	"github.com/labstack/echo/v4"
	"time"
)

func Register(api *echo.Echo) {

	g := api.Group("/api")
	g.GET("/hello", controller.GetHello)
	g.GET("/hello2", controller.GetHello2)
	g.POST("/user", controller.CreateUser)

	//userGroup := g.Group("/user", middleware.Auth)
	//userGroup.GET("/list2", controllers.GetHello2)

}

// 结束router
func CancelRoute(e *echo.Echo) {
	if e == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		logger.Sugar.Fatal(err)
	}
	logger.Sugar.Info("stop router")
}
