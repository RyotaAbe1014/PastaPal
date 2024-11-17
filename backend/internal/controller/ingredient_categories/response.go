package ingredient_categories

type CreateIngredientCategoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GetIngredientCategoriesResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UpdateIngredientCategoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
