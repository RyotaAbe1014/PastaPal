package ingredients

type CreateIngredientRequest struct {
	Name                 string `json:"name"`
	IngredientCategoryID string `json:"ingredient_category_id"`
}
type UpdateIngredientCategoryRequest struct {
	ID                   int    `json:"id"`
	IngredientCategoryID string `json:"ingredient_category_id"`
	Name                 string `json:"name"`
}
type DeleteIngredientCategoryRequest struct {
	ID int `json:"id"`
}
