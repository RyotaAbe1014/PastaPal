package ingredients

import (
	"context"

	"github.com/RyotaAbe1014/Pastapal/internal/service/ingredients"
)

type IngredientController interface {
	CreateIngredient(ctx context.Context, request CreateIngredientRequest) (CreateIngredientResponse, error)
	GetIngredients(ctx context.Context) (GetIngredientsResponse, error)
	UpdateIngredient(ctx context.Context, id string, req UpdateIngredientRequest) (UpdateIngredientResponse, error)
	DeleteIngredient(ctx context.Context, id string) error
}

type ingredientController struct {
	ingredientService ingredients.IngredientService
}

// constructorを使用して、controllerの構造体を生成
func NewIngredientController(is ingredients.IngredientService) IngredientController {
	return &ingredientController{ingredientService: is}
}

func (h *ingredientController) CreateIngredient(ctx context.Context, request CreateIngredientRequest) (CreateIngredientResponse, error) {
	ingredientCategory, err := h.ingredientService.CreateIngredient(ctx, ingredients.CreateIngredientRequestDTO{
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

func (h *ingredientController) GetIngredients(ctx context.Context) (GetIngredientsResponse, error) {
	ingredients, err := h.ingredientService.GetIngredients(ctx)
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

func (h *ingredientController) UpdateIngredient(ctx context.Context, id string, req UpdateIngredientRequest) (UpdateIngredientResponse, error) {
	ingredient, err := h.ingredientService.UpdateIngredient(ctx, ingredients.CreateIngredientRequestDTO{
		ID:                   id,
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

func (h *ingredientController) DeleteIngredient(ctx context.Context, id string) error {
	return h.ingredientService.DeleteIngredient(ctx, id)
}
