package middleware

import (
    "github.com/labstack/echo/v4"
    "echo-framework/lib/jwt"
    "echo-framework/logic/http/response"
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
        if err != nil || tokenData.StaffId == 0 || tokenData.ExpireAt < time.Now().UnixNano() {
            return c.JSON(http.StatusUnauthorized, response.Fail("invalid token"))
        }

        c.Set("staff_id", tokenData.StaffId)

        return next(c)
    }
}
