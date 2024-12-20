package ingredient_categories

import (
	"context"

	"github.com/RyotaAbe1014/Pastapal/internal/service/ingredient_categories"
)

type IngredientCategoryController interface {
	CreateIngredientCategory(ctx context.Context, request CreateIngredientCategoryRequest) (CreateIngredientCategoryResponse, error)
	GetIngredientCategories(ctx context.Context) (GetIngredientCategoriesResponse, error)
	UpdateIngredientCategory(ctx context.Context) error
	DeleteIngredientCategory(ctx context.Context) error
}

type ingredientCategoryController struct {
	ingredientCategoryService ingredient_categories.IngredientCategoryService
}

// constructorを使用して、controllerの構造体を生成
func NewIngredientCategoryController(is ingredient_categories.IngredientCategoryService) IngredientCategoryController {
	return &ingredientCategoryController{ingredientCategoryService: is}
}

func (h *ingredientCategoryController) CreateIngredientCategory(ctx context.Context, request CreateIngredientCategoryRequest) (CreateIngredientCategoryResponse, error) {
	ingredientCategory, err := h.ingredientCategoryService.CreateIngredientCategory(ctx, ingredient_categories.CreateIngredientCategoryRequestDTO{
		Name: request.Name,
	})

	if err != nil {
		return CreateIngredientCategoryResponse{}, err
	}

	return CreateIngredientCategoryResponse{
		ID:   ingredientCategory.ID(),
		Name: ingredientCategory.Name(),
	}, nil
}

func (h *ingredientCategoryController) GetIngredientCategories(ctx context.Context) (GetIngredientCategoriesResponse, error) {
	ingredientCategories, err := h.ingredientCategoryService.GetIngredientCategories(ctx)

	if err != nil {
		return GetIngredientCategoriesResponse{}, err
	}

	var ingredientCategoryResponses GetIngredientCategoriesResponse
	// 0件の場合は空のスライスが返却されるようにする
	ingredientCategoryResponses.Data = make([]CreateIngredientCategoryResponse, 0)

	for _, ingredientCategory := range ingredientCategories {
		ingredientCategoryResponses.Data = append(ingredientCategoryResponses.Data, CreateIngredientCategoryResponse{
			ID:   ingredientCategory.ID(),
			Name: ingredientCategory.Name(),
		})
	}

	return ingredientCategoryResponses, nil
}

func (h *ingredientCategoryController) UpdateIngredientCategory(ctx context.Context) error {
	return nil
}

func (h *ingredientCategoryController) DeleteIngredientCategory(ctx context.Context) error {
	return nil
}
