package repository

import (
	recipesDomain "github.com/RyotaAbe1014/Pastapal/internal/domain/recipes"
	"golang.org/x/net/context"
)

type RecipeRepository struct {
}

func NewRecipeRepository() recipesDomain.IRecipeRepository {
	return &RecipeRepository{}
}

func (r *RecipeRepository) CreateRecipe(ctx context.Context, recipe recipesDomain.Recipe) (recipesDomain.Recipe, error) {
	return recipesDomain.Recipe{}, nil
}

func (r *RecipeRepository) GetRecipeByID(ctx context.Context, id string) (recipesDomain.Recipe, error) {
	return recipesDomain.Recipe{}, nil
}

func (r *RecipeRepository) GetRecipes(ctx context.Context) ([]recipesDomain.Recipe, error) {
	return nil, nil
}

func (r *RecipeRepository) UpdateRecipe(ctx context.Context, recipe recipesDomain.Recipe) (recipesDomain.Recipe, error) {
	return recipesDomain.Recipe{}, nil
}

func (r *RecipeRepository) DeleteRecipe(ctx context.Context, id string) error {
	return nil
}

func (r *RecipeRepository) FindByName(ctx context.Context, name string) (*recipesDomain.Recipe, error) {
	return nil, nil
}
