package repository

import (
	"github.com/RyotaAbe1014/Pastapal/internal/domain/ingredient_categories"
	db "github.com/RyotaAbe1014/Pastapal/internal/infrastructure/postgres"
	"github.com/RyotaAbe1014/Pastapal/internal/infrastructure/postgres/db/dbgen"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/net/context"
)

// 抽象クラス
type IngredientCategoryRepository struct {
}

func NewIngredientCategoryRepository() ingredient_categories.IIngredientCategoryRepository {
	return &IngredientCategoryRepository{}
}

func (r *IngredientCategoryRepository) CreateIngredientCategory(ctx context.Context, ingredientCategory ingredient_categories.IngredientCategory) (ingredient_categories.IngredientCategory, error) {
	query := db.GetQuery(ctx)
	uuidBytes, err := uuid.Parse(ingredientCategory.ID())
	if err != nil {
		return ingredient_categories.IngredientCategory{}, err
	}
	// postgresのuuid型に変換
	pgUUID := pgtype.UUID{
		Bytes: uuidBytes,
		Valid: true,
	}
	result, err := query.CreateIngredientCategory(ctx, dbgen.CreateIngredientCategoryParams{
		ID:   pgUUID,
		Name: ingredientCategory.Name(),
	})

	if err != nil {
		return ingredient_categories.IngredientCategory{}, err
	}

	newIngredientCategory, err := ingredient_categories.NewIngredientCategory(uuid.UUID(result.ID.Bytes).String(), result.Name)
	if err != nil {
		return ingredient_categories.IngredientCategory{}, err
	}

	return newIngredientCategory, nil
}

func (r *IngredientCategoryRepository) GetIngredientCategoryByID(ctx context.Context, id int) (ingredient_categories.IngredientCategory, error) {
	return ingredient_categories.IngredientCategory{}, nil
}

func (r *IngredientCategoryRepository) GetIngredientCategories(ctx context.Context) ([]ingredient_categories.IngredientCategory, error) {
	query := db.GetQuery(ctx)
	result, err := query.ListIngredientsCategory(ctx)
	if err != nil {
		return nil, err
	}

	var ingredientCategories []ingredient_categories.IngredientCategory
	for _, ingredientCategory := range result {
		newIngredientCategory, err := ingredient_categories.NewIngredientCategory(uuid.UUID(ingredientCategory.ID.Bytes).String(), ingredientCategory.Name)
		if err != nil {
			return nil, err
		}
		ingredientCategories = append(ingredientCategories, newIngredientCategory)
	}

	return ingredientCategories, nil
}

func (r *IngredientCategoryRepository) UpdateIngredientCategory(ctx context.Context, ingredientCategory ingredient_categories.IngredientCategory) (ingredient_categories.IngredientCategory, error) {
	return ingredientCategory, nil
}

func (r *IngredientCategoryRepository) DeleteIngredientCategory(ctx context.Context, id int) error {
	return nil
}
