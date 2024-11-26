package ingredients

import (
	"context"

	"github.com/RyotaAbe1014/Pastapal/internal/domain/ingredients"
)

type CreateIngredientRequestDTO struct {
	Name                 string
	IngredientCategoryID string
}

type IngredientService interface {
	CreateIngredient(ctx context.Context, requestDTO CreateIngredientRequestDTO) (ingredients.Ingredient, error)
	GetIngredientByID(ctx context.Context, id int) (ingredients.Ingredient, error)
	GetIngredients(ctx context.Context) ([]ingredients.Ingredient, error)
	UpdateIngredient(ctx context.Context, id int, requestDTO CreateIngredientRequestDTO) (ingredients.Ingredient, error)
	DeleteIngredient(ctx context.Context, id int) error
}

type ingredientService struct {
	ir ingredients.IIngredientRepository // 抽象リポジトリ
}

func NewIngredientService(ir ingredients.IIngredientRepository) IngredientService {
	return &ingredientService{ir: ir}
}

func (is *ingredientService) CreateIngredient(ctx context.Context, requestDTO CreateIngredientRequestDTO) (ingredients.Ingredient, error) {
	return ingredients.Ingredient{}, nil
}

func (is *ingredientService) GetIngredientByID(ctx context.Context, id int) (ingredients.Ingredient, error) {
	return ingredients.Ingredient{}, nil
}

func (is *ingredientService) GetIngredients(ctx context.Context) ([]ingredients.Ingredient, error) {
	return []ingredients.Ingredient{}, nil
}

func (is *ingredientService) UpdateIngredient(ctx context.Context, id int, requestDTO CreateIngredientRequestDTO) (ingredients.Ingredient, error) {
	return ingredients.Ingredient{}, nil
}

func (is *ingredientService) DeleteIngredient(ctx context.Context, id int) error {
	return nil
}
