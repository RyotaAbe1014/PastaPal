package v1

import (
	"net/http"

	"github.com/RyotaAbe1014/Pastapal/internal/infrastructure/oauth/github"
	"github.com/labstack/echo/v4"
)

type GenerateAuthURLResponse struct {
	AuthURL string `json:"url"`
}

type GenerateTokenRequest struct {
	Code string `json:"code"`
}

type GenerateTokenResponse struct {
	UserID    string `json:"userId"`
	AvatarUrl string `json:"avatarUrl"`
}

type AuthStatusResponse struct {
	IsAuthenticated bool `json:"isAuthenticated"`
}

func AuthRouter(g *echo.Group) {
	g.GET("/github/url", func(c echo.Context) error {
		// 認証URLを生成し、GitHubにリダイレクト
		authURL, err := github.GenerateAuthURL("state") // TODO: stateはCSRF対策のための値、実際にはランダムな値を設定するのが良い
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, GenerateAuthURLResponse{AuthURL: authURL})
	})

	g.POST("/github/token", func(c echo.Context) error {
		// GitHubからリダイレクトされてきたときに呼び出される
		var req GenerateTokenRequest
		if err := c.Bind(&req); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		code := req.Code
		if code == "" {
			return c.String(http.StatusBadRequest, "code is required")
		}

		// 認証コードをトークンに交換
		token, err := github.ExchangeCodeForToken(c.Request().Context(), code)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		user, err := github.GetUser(c.Request().Context(), token.AccessToken)

		if err != nil || user == nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// ここでRyotaAbe1014でないと認証エラーを返すようにする
		if user.Login != "RyotaAbe1014" {
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}
		cookie := &http.Cookie{
			Name:     "accessToken",
			Value:    token.AccessToken,
			Expires:  token.Expiry,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
			Path:     "/",
		}
		c.SetCookie(cookie)

		return c.JSON(http.StatusOK, GenerateTokenResponse{UserID: user.Login, AvatarUrl: user.AvatarUrl})
	})

	g.GET("/github/status", func(c echo.Context) error {
		// cookieに保存されているtokenを取得できれば認証済みとする
		cookie, err := c.Cookie("accessToken")
		if err != nil || cookie.Value == "" {
			return c.JSON(http.StatusOK, AuthStatusResponse{IsAuthenticated: false})
		}

		return c.JSON(http.StatusOK, AuthStatusResponse{IsAuthenticated: true})
	})

	g.POST("/github/logout", func(c echo.Context) error {
		cookie := &http.Cookie{
			Name:     "accessToken",
			Value:    "",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
			Path:     "/",
		}
		c.SetCookie(cookie)

		return c.JSON(http.StatusOK, AuthStatusResponse{IsAuthenticated: false})
	})
}
