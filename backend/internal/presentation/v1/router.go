package v1

import (
	"net/http"

	ingredient_categories_controller "github.com/RyotaAbe1014/Pastapal/internal/controller/ingredient_categories"
	ingredients_controller "github.com/RyotaAbe1014/Pastapal/internal/controller/ingredients"
	"github.com/RyotaAbe1014/Pastapal/internal/infrastructure/oauth/github"
	"github.com/RyotaAbe1014/Pastapal/internal/infrastructure/postgres/repository"

	ingredient_categories_service "github.com/RyotaAbe1014/Pastapal/internal/service/ingredient_categories"
	ingredients_service "github.com/RyotaAbe1014/Pastapal/internal/service/ingredients"

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
		ctx := c.Request().Context()
		// TODO: ここのDIはもっとスマートに書けるはず、他のingredientCategory関連の処理と共通化できるはず
		ingredientCategoryRepository := repository.NewIngredientCategoryRepository()
		ingredientCategoryService := ingredient_categories_service.NewIngredientCategoryService(ingredientCategoryRepository)
		ingredientCategoryController := ingredient_categories_controller.NewIngredientCategoryController(ingredientCategoryService)

		req := new(ingredient_categories_controller.CreateIngredientCategoryRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		result, err := ingredientCategoryController.CreateIngredientCategory(ctx, ingredient_categories_controller.CreateIngredientCategoryRequest{
			Name: req.Name,
		})

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, result)
	})

	g.GET(("/ingredient-categories"), func(c echo.Context) error {
		ctx := c.Request().Context()
		ingredientCategoryRepository := repository.NewIngredientCategoryRepository()
		ingredientCategoryService := ingredient_categories_service.NewIngredientCategoryService(ingredientCategoryRepository)
		ingredientCategoryController := ingredient_categories_controller.NewIngredientCategoryController(ingredientCategoryService)

		result, err := ingredientCategoryController.GetIngredientCategories(ctx)

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, result)
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
		ctx := c.Request().Context()
		ingredientRepository := repository.NewIngredientRepository()
		ingredientService := ingredients_service.NewIngredientService(ingredientRepository)
		ingredientController := ingredients_controller.NewIngredientController(ingredientService)

		result, err := ingredientController.GetIngredients(ctx)

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, result)
	})

	g.GET(("/ingredients/:id"), func(c echo.Context) error {
		return c.String(http.StatusOK, "ingredients")
	})

	g.POST("/ingredients", func(c echo.Context) error {
		ctx := c.Request().Context()
		ingredientRepository := repository.NewIngredientRepository()
		ingredientService := ingredients_service.NewIngredientService(ingredientRepository)
		ingredientController := ingredients_controller.NewIngredientController(ingredientService)

		req := new(ingredients_controller.CreateIngredientRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		result, err := ingredientController.CreateIngredient(ctx, ingredients_controller.CreateIngredientRequest{
			Name:                 req.Name,
			IngredientCategoryID: req.IngredientCategoryID,
		})

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, result)
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
