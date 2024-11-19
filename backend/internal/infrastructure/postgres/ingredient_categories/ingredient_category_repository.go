package ingredient_categories

import (
	"github.com/RyotaAbe1014/Pastapal/internal/domain/ingredient_categories"
	"github.com/RyotaAbe1014/Pastapal/internal/infrastructure/postgres/db/dbgen"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/net/context"
)

// 抽象クラス
type IngredientCategoryRepository struct {
	ctx   context.Context
	query *dbgen.Queries
	conn  *pgx.Conn
}

func NewIngredientCategoryRepository(ctx context.Context, query *dbgen.Queries, conn *pgx.Conn) ingredient_categories.IIngredientCategoryRepository {
	return &IngredientCategoryRepository{
		ctx:   ctx,
		query: query,
		conn:  conn,
	}
}

func (r *IngredientCategoryRepository) CreateIngredientCategory(ingredientCategory ingredient_categories.IngredientCategory) (ingredient_categories.IngredientCategory, error) {
	defer r.conn.Close(r.ctx)
	uuidBytes, err := uuid.Parse(ingredientCategory.ID())
	if err != nil {
		return ingredient_categories.IngredientCategory{}, err
	}
	// postgresのuuid型に変換
	pgUUID := pgtype.UUID{
		Bytes: uuidBytes,
		Valid: true,
	}
	result, err := r.query.CreateIngredientCategory(r.ctx, dbgen.CreateIngredientCategoryParams{
		ID:   pgUUID,
		Name: ingredientCategory.Name(),
	})

	if err != nil {
		return ingredient_categories.IngredientCategory{}, err
	}

	return ingredient_categories.NewIngredientCategory(result.ID.String, result.Name), nil
}

func (r *IngredientCategoryRepository) GetIngredientCategoryByID(id int) (ingredient_categories.IngredientCategory, error) {
	return ingredient_categories.IngredientCategory{}, nil
}

func (r *IngredientCategoryRepository) GetAllIngredientCategories() ([]ingredient_categories.IngredientCategory, error) {
	return []ingredient_categories.IngredientCategory{}, nil
}

func (r *IngredientCategoryRepository) UpdateIngredientCategory(ingredientCategory ingredient_categories.IngredientCategory) (ingredient_categories.IngredientCategory, error) {
	return ingredientCategory, nil
}

func (r *IngredientCategoryRepository) DeleteIngredientCategory(id int) error {
	return nil
}
