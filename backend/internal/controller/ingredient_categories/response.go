package ingredient_categories

type CreateIngredientCategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetIngredientCategoriesResponse struct {
	Data []CreateIngredientCategoryResponse `json:"data"`
}

type IngredientCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateIngredientCategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
