package v1

import (
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

	rg := g.Group("/recipes")
	rg.Use(settings.GitHubAuthMiddleware)
	RecipesRouter(rg)
}
