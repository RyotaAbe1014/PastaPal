package recipes

import "context"

type IRecipeRepository interface {
	CreateRecipe(ctx context.Context, recipe Recipe) (Recipe, error)
	GetRecipeByID(ctx context.Context, id string) (Recipe, error)
	GetRecipes(ctx context.Context) ([]Recipe, error)
	UpdateRecipe(ctx context.Context, recipe Recipe) (Recipe, error)
	DeleteRecipe(ctx context.Context, id string) error
	FindByName(name string) (*Recipe, error)
}
