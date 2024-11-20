package ingredient_categories

import "context"

// 抽象クラス
type IIngredientCategoryRepository interface {
	CreateIngredientCategory(ctx context.Context, ingredientCategory IngredientCategory) (IngredientCategory, error)
	GetIngredientCategoryByID(ctx context.Context, id int) (IngredientCategory, error)
	GetIngredientCategories(ctx context.Context) ([]IngredientCategory, error)
	UpdateIngredientCategory(ctx context.Context, ingredientCategory IngredientCategory) (IngredientCategory, error)
	DeleteIngredientCategory(ctx context.Context, id int) error
}
