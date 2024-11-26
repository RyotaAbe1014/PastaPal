package repository

import (
	"github.com/RyotaAbe1014/Pastapal/internal/domain/ingredients"
	db "github.com/RyotaAbe1014/Pastapal/internal/infrastructure/postgres"
	"github.com/RyotaAbe1014/Pastapal/internal/infrastructure/postgres/db/dbgen"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/net/context"
)

// 抽象クラス
type IngredientRepository struct {
}

func NewIngredientRepository() ingredients.IIngredientRepository {
	return &IngredientRepository{}
}

func (r *IngredientRepository) CreateIngredient(ctx context.Context, ingredient ingredients.Ingredient) (ingredients.Ingredient, error) {
	query := db.GetQuery(ctx)
	ingredientUuidBytes, err := uuid.Parse(ingredient.ID())
	if err != nil {
		return ingredients.Ingredient{}, err
	}

	ingredientCategoryUuidBytes, err := uuid.Parse(ingredient.IngredientCategoryID())

	if err != nil {
		return ingredients.Ingredient{}, err
	}

	// postgresのuuid型に変換
	ingredientPgUUID := pgtype.UUID{
		Bytes: ingredientUuidBytes,
		Valid: true,
	}

	ingredientCategoryPgUUID := pgtype.UUID{
		Bytes: ingredientCategoryUuidBytes,
		Valid: true,
	}

	result, err := query.CreateIngredient(ctx, dbgen.CreateIngredientParams{
		ID:                   ingredientPgUUID,
		Name:                 ingredient.Name(),
		IngredientCategoryID: ingredientCategoryPgUUID,
	})

	if err != nil {
		return ingredients.Ingredient{}, err
	}

	newIngredient, err := ingredients.NewIngredient(uuid.UUID(result.ID.Bytes).String(), result.Name, uuid.UUID(result.IngredientCategoryID.Bytes).String())
	if err != nil {
		return ingredients.Ingredient{}, err
	}

	return newIngredient, nil
}

func (r *IngredientRepository) GetIngredientByID(ctx context.Context, id int) (ingredients.Ingredient, error) {
	return ingredients.Ingredient{}, nil
}

func (r *IngredientRepository) GetIngredients(ctx context.Context) ([]ingredients.Ingredient, error) {
	return []ingredients.Ingredient{}, nil
}

func (r *IngredientRepository) UpdateIngredient(ctx context.Context, ingredientCategory ingredients.Ingredient) (ingredients.Ingredient, error) {
	return ingredientCategory, nil
}

func (r *IngredientRepository) DeleteIngredient(ctx context.Context, id int) error {
	return nil
}
