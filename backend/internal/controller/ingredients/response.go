package ingredients

type CreateIngredientResponse struct {
	ID                   string `json:"id"`
	IngredientCategoryID string `json:"ingredientCategoryId"`
	Name                 string `json:"name"`
}

type GetIngredientsResponse struct {
	Data []CreateIngredientResponse `json:"data"`
}

type Ingredient struct {
	ID                   string `json:"id"`
	IngredientCategoryID string `json:"ingredientCategoryId"`
	Name                 string `json:"name"`
}

type UpdateIngredientResponse struct {
	ID                   string `json:"id"`
	IngredientCategoryID string `json:"ingredientCategoryId"`
	Name                 string `json:"name"`
}
