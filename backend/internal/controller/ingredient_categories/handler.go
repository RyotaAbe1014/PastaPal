package ingredient_categories

import "github.com/labstack/echo/v4"

type IngredientCategoryHandler interface {
	CreateIngredientCategory(c echo.Context) error
	GetIngredientCategories(c echo.Context) error
	UpdateIngredientCategory(c echo.Context) error
	DeleteIngredientCategory(c echo.Context) error
}

type ingredientCategoryHandler struct {
	// ingredientCategoryUseCase IngredientCategoryUseCase
}

// constructorを使用して、controllerの構造体を生成
func New() IngredientCategoryHandler {
	return &ingredientCategoryHandler{}
}

func (h *ingredientCategoryHandler) CreateIngredientCategory(c echo.Context) error {
	return nil
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
