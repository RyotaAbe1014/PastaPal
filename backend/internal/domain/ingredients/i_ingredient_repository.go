package ingredients

import "context"

// 抽象クラス
type IIngredientRepository interface {
	CreateIngredient(ctx context.Context, ingredient Ingredient) (Ingredient, error)
	GetIngredientByID(ctx context.Context, id int) (Ingredient, error)
	GetIngredients(ctx context.Context) ([]Ingredient, error)
	UpdateIngredient(ctx context.Context, ingredient Ingredient) (Ingredient, error)
	DeleteIngredient(ctx context.Context, id int) error
}
