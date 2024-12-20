package ingredients

type CreateIngredientRequest struct {
	Name                 string `json:"name"`
	IngredientCategoryID string `json:"ingredientCategoryId"`
}
type UpdateIngredientRequest struct {
	IngredientCategoryID string `json:"ingredientCategoryId"`
	Name                 string `json:"name"`
}
type DeleteIngredientRequest struct {
	ID string `json:"id"`
}
