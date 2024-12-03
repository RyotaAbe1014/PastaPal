package settings

import (
	"net/http"

	"github.com/RyotaAbe1014/Pastapal/internal/infrastructure/oauth/github"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CorsMiddleware(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))
}

func GitHubAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// GitHub認証の処理
		cookie, err := c.Cookie("accessToken")
		if err != nil {
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}

		user, err := github.GetUser(c.Request().Context(), cookie.Value)

		if err != nil || user == nil {
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}

		// ここでRyotaAbe1014でないと認証エラーを返すようにする
		if user.Login != "RyotaAbe1014" {
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}

		return next(c)
	}
}
