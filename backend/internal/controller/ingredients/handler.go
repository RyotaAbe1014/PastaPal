package ingredients

import (
	"context"

	"github.com/RyotaAbe1014/Pastapal/internal/service/ingredients"
)

type IngredientController interface {
	CreateIngredient(ctx context.Context, request CreateIngredientRequest) (CreateIngredientResponse, error)
	GetIngredients(ctx context.Context) (GetIngredientsResponse, error)
	UpdateIngredient(ctx context.Context, req UpdateIngredientRequest) (UpdateIngredientResponse, error)
	DeleteIngredient(ctx context.Context, id string) error
}

type ingredientCategoryController struct {
	ingredientCategoryService ingredients.IngredientService
}

// constructorを使用して、controllerの構造体を生成
func NewIngredientController(is ingredients.IngredientService) IngredientController {
	return &ingredientCategoryController{ingredientCategoryService: is}
}

func (h *ingredientCategoryController) CreateIngredient(ctx context.Context, request CreateIngredientRequest) (CreateIngredientResponse, error) {
	ingredientCategory, err := h.ingredientCategoryService.CreateIngredient(ctx, ingredients.CreateIngredientRequestDTO{
		Name:                 request.Name,
		IngredientCategoryID: request.IngredientCategoryID,
	})

	if err != nil {
		return CreateIngredientResponse{}, err
	}

	return CreateIngredientResponse{
		ID:                   ingredientCategory.ID(),
		Name:                 ingredientCategory.Name(),
		IngredientCategoryID: ingredientCategory.IngredientCategoryID(),
	}, nil
}

func (h *ingredientCategoryController) GetIngredients(ctx context.Context) (GetIngredientsResponse, error) {
	ingredients, err := h.ingredientCategoryService.GetIngredients(ctx)
	if err != nil {
		return GetIngredientsResponse{}, err
	}

	var response GetIngredientsResponse
	// 0件の場合は空のスライスが返却されるようにする
	response.Data = make([]CreateIngredientResponse, 0)
	for _, ingredient := range ingredients {
		response.Data = append(response.Data, CreateIngredientResponse{
			ID:                   ingredient.ID(),
			Name:                 ingredient.Name(),
			IngredientCategoryID: ingredient.IngredientCategoryID(),
		})
	}

	return response, nil
}

func (h *ingredientCategoryController) UpdateIngredient(ctx context.Context, req UpdateIngredientRequest) (UpdateIngredientResponse, error) {
	ingredient, err := h.ingredientCategoryService.UpdateIngredient(ctx, ingredients.CreateIngredientRequestDTO{
		ID:                   req.ID,
		Name:                 req.Name,
		IngredientCategoryID: req.IngredientCategoryID,
	})

	if err != nil {
		return UpdateIngredientResponse{}, err
	}

	return UpdateIngredientResponse{
		ID:                   ingredient.ID(),
		Name:                 ingredient.Name(),
		IngredientCategoryID: ingredient.IngredientCategoryID(),
	}, nil
}

func (h *ingredientCategoryController) DeleteIngredient(ctx context.Context, id string) error {
	return h.ingredientCategoryService.DeleteIngredient(ctx, id)
}
