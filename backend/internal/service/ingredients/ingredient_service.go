package ingredients

import (
	"context"
	"errors"

	"github.com/RyotaAbe1014/Pastapal/internal/domain/ingredients"
	"github.com/google/uuid"
)

type CreateIngredientRequestDTO struct {
	ID                   string
	Name                 string
	IngredientCategoryID string
}

type IngredientService interface {
	CreateIngredient(ctx context.Context, requestDTO CreateIngredientRequestDTO) (ingredients.Ingredient, error)
	GetIngredientByID(ctx context.Context, id string) (ingredients.Ingredient, error)
	GetIngredients(ctx context.Context) ([]ingredients.Ingredient, error)
	UpdateIngredient(ctx context.Context, requestDTO CreateIngredientRequestDTO) (ingredients.Ingredient, error)
	DeleteIngredient(ctx context.Context, id string) error
}

type ingredientService struct {
	ir ingredients.IIngredientRepository // 抽象リポジトリ
}

func NewIngredientService(ir ingredients.IIngredientRepository) IngredientService {
	return &ingredientService{ir: ir}
}

func (is *ingredientService) CreateIngredient(ctx context.Context, requestDTO CreateIngredientRequestDTO) (ingredients.Ingredient, error) {
	id := uuid.New().String()
	ingredient, err := ingredients.NewIngredient(id, requestDTO.Name, requestDTO.IngredientCategoryID)

	if err != nil {
		return ingredients.Ingredient{},
			errors.New("failed to create ingredient category")
	}

	createdIngredient, createErr := is.ir.CreateIngredient(ctx, ingredient)

	if createErr != nil {
		return ingredients.Ingredient{}, createErr
	}

	return createdIngredient, nil
}

func (is *ingredientService) GetIngredientByID(ctx context.Context, id string) (ingredients.Ingredient, error) {
	return is.ir.GetIngredientByID(ctx, id)
}

func (is *ingredientService) GetIngredients(ctx context.Context) ([]ingredients.Ingredient, error) {
	return is.ir.GetIngredients(ctx)
}

func (is *ingredientService) UpdateIngredient(ctx context.Context, requestDTO CreateIngredientRequestDTO) (ingredients.Ingredient, error) {
	targetIngredient, err := is.GetIngredientByID(ctx, requestDTO.ID)

	if err != nil {
		return ingredients.Ingredient{}, err
	}

	err = targetIngredient.UpdateName(requestDTO.Name)
	err = targetIngredient.UpdateIngredientCategoryID(requestDTO.IngredientCategoryID)

	if err != nil {
		return ingredients.Ingredient{}, err
	}

	updatedIngredient, updateErr := is.ir.UpdateIngredient(ctx, targetIngredient)

	if updateErr != nil {
		return ingredients.Ingredient{}, updateErr
	}

	return updatedIngredient, nil
}

func (is *ingredientService) DeleteIngredient(ctx context.Context, id string) error {
	return is.ir.DeleteIngredient(ctx, id)
}
