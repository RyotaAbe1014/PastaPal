package ingredient_categories

// 抽象クラス
type IIngredientCategoryRepository interface {
	CreateIngredientCategory(ingredientCategory IngredientCategory) (IngredientCategory, error)
	GetIngredientCategoryByID(id int) (IngredientCategory, error)
	GetAllIngredientCategories() ([]IngredientCategory, error)
	UpdateIngredientCategory(ingredientCategory IngredientCategory) (IngredientCategory, error)
	DeleteIngredientCategory(id int) error
}
