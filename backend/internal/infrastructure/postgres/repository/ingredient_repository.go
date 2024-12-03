package repository

import (
	ingredientsDomain "github.com/RyotaAbe1014/Pastapal/internal/domain/ingredients"
	db "github.com/RyotaAbe1014/Pastapal/internal/infrastructure/postgres"
	"github.com/RyotaAbe1014/Pastapal/internal/infrastructure/postgres/db/dbgen"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/net/context"
)

// 抽象クラス
type IngredientRepository struct {
}

func NewIngredientRepository() ingredientsDomain.IIngredientRepository {
	return &IngredientRepository{}
}

func (r *IngredientRepository) CreateIngredient(ctx context.Context, ingredient ingredientsDomain.Ingredient) (ingredientsDomain.Ingredient, error) {
	query := db.GetQuery(ctx)
	ingredientUuidBytes, err := uuid.Parse(ingredient.ID())
	if err != nil {
		return ingredientsDomain.Ingredient{}, err
	}

	ingredientCategoryUuidBytes, err := uuid.Parse(ingredient.IngredientCategoryID())

	if err != nil {
		return ingredientsDomain.Ingredient{}, err
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
		return ingredientsDomain.Ingredient{}, err
	}

	newIngredient, err := ingredientsDomain.NewIngredient(uuid.UUID(result.ID.Bytes).String(), result.Name, uuid.UUID(result.IngredientCategoryID.Bytes).String())
	if err != nil {
		return ingredientsDomain.Ingredient{}, err
	}

	return newIngredient, nil
}

func (r *IngredientRepository) GetIngredientByID(ctx context.Context, id string) (ingredientsDomain.Ingredient, error) {
	query := db.GetQuery(ctx)

	ingredientUuidBytes, err := uuid.Parse(id)

	if err != nil {
		return ingredientsDomain.Ingredient{}, err
	}

	ingredientPgUUID := pgtype.UUID{
		Bytes: ingredientUuidBytes,
		Valid: true,
	}

	result, err := query.GetIngredient(ctx, ingredientPgUUID)
	if err != nil {
		return ingredientsDomain.Ingredient{}, err
	}

	ingredient, err := ingredientsDomain.NewIngredientFromRepository(uuid.UUID(result.ID.Bytes).String(), result.Name, uuid.UUID(result.IngredientCategoryID.Bytes).String(), result.CreatedAt.Time, result.UpdatedAt.Time)

	if err != nil {
		return ingredientsDomain.Ingredient{}, err
	}

	return ingredient, nil
}

func (r *IngredientRepository) GetIngredients(ctx context.Context) ([]ingredientsDomain.Ingredient, error) {
	query := db.GetQuery(ctx)
	result, err := query.ListIngredients(ctx)
	if err != nil {
		return nil, err
	}

	ingredients := make([]ingredientsDomain.Ingredient, 0)
	for _, ingredient := range result {
		ingredientModel, err := ingredientsDomain.NewIngredientFromRepository(uuid.UUID(ingredient.ID.Bytes).String(), ingredient.Name, uuid.UUID(ingredient.IngredientCategoryID.Bytes).String(), ingredient.CreatedAt.Time, ingredient.UpdatedAt.Time)
		if err != nil {
			return nil, err
		}
		ingredients = append(ingredients, ingredientModel)
	}

	return ingredients, nil
}

func (r *IngredientRepository) UpdateIngredient(ctx context.Context, ingredientCategory ingredientsDomain.Ingredient) (ingredientsDomain.Ingredient, error) {
	query := db.GetQuery(ctx)
	ingredientUuidBytes, err := uuid.Parse(ingredientCategory.ID())
	if err != nil {
		return ingredientsDomain.Ingredient{}, err
	}

	ingredientCategoryUuidBytes, err := uuid.Parse(ingredientCategory.IngredientCategoryID())

	if err != nil {
		return ingredientsDomain.Ingredient{}, err
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

	result, err := query.UpdateIngredient(ctx, dbgen.UpdateIngredientParams{
		ID:                   ingredientPgUUID,
		Name:                 ingredientCategory.Name(),
		IngredientCategoryID: ingredientCategoryPgUUID,
	})
	if err != nil {
		return ingredientsDomain.Ingredient{}, err
	}

	updatedIngredient, err := ingredientsDomain.NewIngredient(uuid.UUID(result.ID.Bytes).String(), result.Name, uuid.UUID(result.IngredientCategoryID.Bytes).String())
	if err != nil {
		return ingredientsDomain.Ingredient{}, err
	}

	return updatedIngredient, nil
}

func (r *IngredientRepository) DeleteIngredient(ctx context.Context, id string) error {
	query := db.GetQuery(ctx)
	ingredientUuidBytes, err := uuid.Parse(id)

	if err != nil {
		return err
	}

	ingredientPgUUID := pgtype.UUID{
		Bytes: ingredientUuidBytes,
		Valid: true,
	}

	err = query.DeleteIngredient(ctx, ingredientPgUUID)
	if err != nil {
		return err
	}

	return nil
}

func (r *IngredientRepository) FindByName(name string) (*ingredientsDomain.Ingredient, error) {
	query := db.GetQuery(context.Background())
	result, err := query.FindIngredientByName(context.Background(), name)
	if err != nil {
		return nil, err
	}

	if uuid.UUID(result.ID.Bytes).String() == "" {
		return nil, nil
	}

	ingredient, err := ingredientsDomain.NewIngredientFromRepository(uuid.UUID(result.ID.Bytes).String(), result.Name, uuid.UUID(result.IngredientCategoryID.Bytes).String(), result.CreatedAt.Time, result.UpdatedAt.Time)
	if err != nil {
		return nil, err
	}

	return &ingredient, nil
}
