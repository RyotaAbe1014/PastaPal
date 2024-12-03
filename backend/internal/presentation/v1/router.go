package v1

import (
	"net/http"

	"github.com/RyotaAbe1014/Pastapal/internal/presentation/settings"
	"github.com/labstack/echo/v4"
)

func Router(g *echo.Group) {
	ag := g.Group("/auth")
	AuthRouter(ag)

	icg := g.Group("/ingredient-categories")
	icg.Use(settings.GitHubAuthMiddleware)
	IngredientCategoriesRouter(icg)

	ig := g.Group("/ingredients")
	ig.Use(settings.GitHubAuthMiddleware)
	IngredientsRouter(ig)

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
