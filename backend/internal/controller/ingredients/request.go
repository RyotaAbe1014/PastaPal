package ingredients

type CreateIngredientRequest struct {
	Name                 string `json:"name"`
	IngredientCategoryID string `json:"ingredientCategoryId"`
}
type UpdateIngredientCategoryRequest struct {
	ID                   int    `json:"id"`
	IngredientCategoryID string `json:"ingredientCategoryId"`
	Name                 string `json:"name"`
}
type DeleteIngredientCategoryRequest struct {
	ID int `json:"id"`
}
