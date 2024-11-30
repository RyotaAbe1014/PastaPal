package v1

import (
	"net/http"

	ingredients_controller "github.com/RyotaAbe1014/Pastapal/internal/controller/ingredients"
	"github.com/RyotaAbe1014/Pastapal/internal/infrastructure/postgres/repository"
	ingredients_service "github.com/RyotaAbe1014/Pastapal/internal/service/ingredients"
	"github.com/labstack/echo/v4"
)

func IngredientsRouter(g *echo.Group) {
	ingredientRepository := repository.NewIngredientRepository()
	ingredientService := ingredients_service.NewIngredientService(ingredientRepository)
	ingredientController := ingredients_controller.NewIngredientController(ingredientService)

	// 食材管理
	g.GET((""), func(c echo.Context) error {
		ctx := c.Request().Context()
		result, err := ingredientController.GetIngredients(ctx)

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, result)
	})

	g.GET((":id"), func(c echo.Context) error {
		return c.String(http.StatusOK, "ingredients")
	})

	g.POST("", func(c echo.Context) error {
		ctx := c.Request().Context()
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

	g.PUT("/:id", func(c echo.Context) error {
		ctx := c.Request().Context()
		req := new(ingredients_controller.UpdateIngredientRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		result, err := ingredientController.UpdateIngredient(ctx, ingredients_controller.UpdateIngredientRequest{
			ID:                   req.ID,
			Name:                 req.Name,
			IngredientCategoryID: req.IngredientCategoryID,
		})

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, result)
	})

	g.DELETE("/:id", func(c echo.Context) error {
		ctx := c.Request().Context()
		id := c.Param("id")
		err := ingredientController.DeleteIngredient(ctx, id)

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	})
}
