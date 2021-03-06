package middleware

import (
	"echo-framework/lib/jwt"
	"echo-framework/logic/http/responses"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("token")
		if token == "" {
			token = c.QueryParam("token")
		}

		tokenData, err := jwt.ParseToken(token)
		if err != nil || tokenData.UserId == 0 || tokenData.ExpireAt < time.Now().UnixNano() {
			return c.JSON(http.StatusUnauthorized, responses.Fail("invalid token"))
		}

		c.Set("user_id", tokenData.UserId)

		return next(c)
	}
}
