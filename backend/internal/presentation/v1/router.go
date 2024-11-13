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

func Router(g *echo.Group) {
	g.GET("/auth/github/url", func(c echo.Context) error {
		// 認証URLを生成し、GitHubにリダイレクト
		authURL, err := github.GenerateAuthURL("state") // TODO: stateはCSRF対策のための値、実際にはランダムな値を設定するのが良い
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, GenerateAuthURLResponse{AuthURL: authURL})
	})

	g.POST("/auth/github/token", func(c echo.Context) error {
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

	g.GET("/auth/github/status", func(c echo.Context) error {
		// cookieに保存されているtokenを取得できれば認証済みとする
		cookie, err := c.Cookie("accessToken")
		if err != nil || cookie.Value == "" {
			return c.JSON(http.StatusOK, AuthStatusResponse{IsAuthenticated: false})
		}

		return c.JSON(http.StatusOK, AuthStatusResponse{IsAuthenticated: true})
	})

	g.POST("/auth/github/logout", func(c echo.Context) error {
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

	// TODO: 必要なapiをとりあえず定義しているので、実装時に修正する
	// echo.Groupはネストすることができるので、この関数内でさらにGroupを作成してルーティングを定義することができる
	// ミドルウェアを設定する場合など、複雑になりそうであればモジュールを分割する

	// 食材種別管理
	g.POST("/ingredient-categories", func(c echo.Context) error {
		return c.String(http.StatusOK, "ingredient-categories")
	})
	g.GET(("/ingredient-categories"), func(c echo.Context) error {
		return c.String(http.StatusOK, "ingredient-categories")
	})
	g.GET("/ingredient-categories/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "ingredient-categories")
	})

	g.PUT("/ingredient-categories/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "ingredient-categories")
	})

	g.DELETE("/ingredient-categories/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "ingredient-categories")
	})

	// 食材管理
	g.GET(("/ingredients"), func(c echo.Context) error {
		return c.String(http.StatusOK, "ingredients")
	})
	g.GET(("/ingredients/:id"), func(c echo.Context) error {
		return c.String(http.StatusOK, "ingredients")
	})

	g.POST("/ingredients", func(c echo.Context) error {
		return c.String(http.StatusOK, "ingredients")
	})

	g.PUT("/ingredients/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "ingredients")
	})

	g.DELETE("/ingredients/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "ingredients")
	})

	// レシピ管理
	g.POST("/recipes", func(c echo.Context) error {
		return c.String(http.StatusOK, "resipes")
	})
	g.GET("/recipes", func(c echo.Context) error {
		return c.String(http.StatusOK, "resipes")
	})
	g.GET("/recipes/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "resipes")
	})
	g.PUT("/recipes/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "resipes")
	})
	g.DELETE("/recipes/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "resipes")
	})
}
