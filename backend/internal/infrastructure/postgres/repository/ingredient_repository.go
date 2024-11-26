package repository

import (
	"github.com/RyotaAbe1014/Pastapal/internal/domain/ingredients"
	"golang.org/x/net/context"
)

// 抽象クラス
type IngredientRepository struct {
}

func NewIngredientRepository() ingredients.IIngredientRepository {
	return &IngredientRepository{}
}

func (r *IngredientRepository) CreateIngredient(ctx context.Context, ingredientCategory ingredients.Ingredient) (ingredients.Ingredient, error) {
	return ingredientCategory, nil
}

func (r *IngredientRepository) GetIngredientByID(ctx context.Context, id int) (ingredients.Ingredient, error) {
	return ingredients.Ingredient{}, nil
}

func (r *IngredientRepository) GetIngredients(ctx context.Context) ([]ingredients.Ingredient, error) {
	return []ingredients.Ingredient{}, nil
}

func (r *IngredientRepository) UpdateIngredient(ctx context.Context, ingredientCategory ingredients.Ingredient) (ingredients.Ingredient, error) {
	return ingredientCategory, nil
}

func (r *IngredientRepository) DeleteIngredient(ctx context.Context, id int) error {
	return nil
}
