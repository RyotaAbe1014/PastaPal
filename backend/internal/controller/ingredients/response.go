package ingredients

type CreateIngredientResponse struct {
	ID                   string `json:"id"`
	IngredientCategoryID string `json:"ingredient_category_id"`
	Name                 string `json:"name"`
}

type GetIngredientsResponse struct {
	Data []CreateIngredientResponse `json:"data"`
}

type Ingredient struct {
	ID                   string `json:"id"`
	IngredientCategoryID string `json:"ingredient_category_id"`
	Name                 string `json:"name"`
}

type UpdateIngredientCategoryResponse struct {
	ID                   string `json:"id"`
	IngredientCategoryID string `json:"ingredient_category_id"`
	Name                 string `json:"name"`
}
