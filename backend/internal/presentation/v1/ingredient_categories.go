package v1

import (
	"net/http"

	ingredient_categories_controller "github.com/RyotaAbe1014/Pastapal/internal/controller/ingredient_categories"
	"github.com/RyotaAbe1014/Pastapal/internal/infrastructure/postgres/repository"

	ingredient_categories_service "github.com/RyotaAbe1014/Pastapal/internal/service/ingredient_categories"
	"github.com/labstack/echo/v4"
)

func IngredientCategoriesRouter(g *echo.Group) {
	// 食材種別管理
	g.POST("", func(c echo.Context) error {
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

	g.GET((""), func(c echo.Context) error {
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

	g.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "ingredient-categories")
	})

	g.PUT("/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "ingredient-categories")
	})

	g.DELETE("/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "ingredient-categories")
	})

}
