package ingredient_categories

import (
	"context"
	"errors"

	"github.com/RyotaAbe1014/Pastapal/internal/domain/ingredient_categories"
	"github.com/google/uuid"
)

type CreateIngredientCategoryRequestDTO struct {
	Name string
}

type IngredientCategoryService interface {
	CreateIngredientCategory(ctx context.Context, requestDTO CreateIngredientCategoryRequestDTO) (ingredient_categories.IngredientCategory, error)
	GetIngredientCategoryByID(ctx context.Context, id int) (ingredient_categories.IngredientCategory, error)
	GetIngredientCategories(ctx context.Context) ([]ingredient_categories.IngredientCategory, error)
	UpdateIngredientCategory(ctx context.Context, id int, requestDTO CreateIngredientCategoryRequestDTO) (ingredient_categories.IngredientCategory, error)
	DeleteIngredientCategory(ctx context.Context, id int) error
}

type ingredientCategoryService struct {
	ir ingredient_categories.IIngredientCategoryRepository // 抽象リポジトリ
}

func NewIngredientCategoryService(ir ingredient_categories.IIngredientCategoryRepository) IngredientCategoryService {
	return &ingredientCategoryService{ir: ir}
}

func (is *ingredientCategoryService) CreateIngredientCategory(ctx context.Context, requestDTO CreateIngredientCategoryRequestDTO) (ingredient_categories.IngredientCategory, error) {
	id := uuid.New().String()
	ingredientCategory, err := ingredient_categories.NewIngredientCategory(id, requestDTO.Name)

	if err != nil {
		return ingredient_categories.IngredientCategory{},
			errors.New("failed to create ingredient category")
	}

	is.ir.CreateIngredientCategory(ctx, ingredientCategory)

	return ingredientCategory, nil
}

func (is *ingredientCategoryService) GetIngredientCategoryByID(ctx context.Context, id int) (ingredient_categories.IngredientCategory, error) {
	panic("implement me")
}

func (is *ingredientCategoryService) GetIngredientCategories(ctx context.Context) ([]ingredient_categories.IngredientCategory, error) {
	return is.ir.GetIngredientCategories(ctx)
}

func (is *ingredientCategoryService) UpdateIngredientCategory(ctx context.Context, id int, requestDTO CreateIngredientCategoryRequestDTO) (ingredient_categories.IngredientCategory, error) {
	panic("implement me")
}

func (is *ingredientCategoryService) DeleteIngredientCategory(ctx context.Context, id int) error {
	panic("implement me")
}
