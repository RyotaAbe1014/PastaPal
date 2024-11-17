package ingredient_categories

type CreateIngredientCategoryRequest struct {
	Name string `json:"name"`
}
type UpdateIngredientCategoryRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type DeleteIngredientCategoryRequest struct {
	ID int `json:"id"`
}
