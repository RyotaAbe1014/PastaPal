package ingredient_categories

import (
	"github.com/RyotaAbe1014/Pastapal/internal/service/ingredient_categories"
	"github.com/labstack/echo/v4"
)

type IngredientCategoryHandler interface {
	CreateIngredientCategory(c echo.Context) (CreateIngredientCategoryResponse, error)
	GetIngredientCategories(c echo.Context) error
	UpdateIngredientCategory(c echo.Context) error
	DeleteIngredientCategory(c echo.Context) error
}

type ingredientCategoryHandler struct {
	ingredientCategoryService ingredient_categories.IngredientCategoryService
}

// constructorを使用して、controllerの構造体を生成
func New(is ingredient_categories.IngredientCategoryService) IngredientCategoryHandler {
	return &ingredientCategoryHandler{ingredientCategoryService: is}
}

func (h *ingredientCategoryHandler) CreateIngredientCategory(c echo.Context) (CreateIngredientCategoryResponse, error) {
	ingredientCategory, err := h.ingredientCategoryService.CreateIngredientCategory(ingredient_categories.CreateIngredientCategoryRequestDTO{
		Name: c.FormValue("name"),
	})

	if err != nil {
		return CreateIngredientCategoryResponse{}, err
	}

	return CreateIngredientCategoryResponse{
		ID:   ingredientCategory.ID(),
		Name: ingredientCategory.Name(),
	}, nil
}

func (h *ingredientCategoryHandler) GetIngredientCategories(c echo.Context) error {
	return nil
}

func (h *ingredientCategoryHandler) UpdateIngredientCategory(c echo.Context) error {
	return nil
}

func (h *ingredientCategoryHandler) DeleteIngredientCategory(c echo.Context) error {
	return nil
}
