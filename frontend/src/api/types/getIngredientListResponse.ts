export type GetIngredientListResponse = {
	data: Ingredient[];
};

export type Ingredient = {
	id: string;
	name: string;
	ingredientCategoryId: string;
};
