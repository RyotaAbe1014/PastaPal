package ingredient_categories

import (
	"errors"

	"github.com/RyotaAbe1014/Pastapal/internal/domain/ingredient_categories"
	"github.com/google/uuid"
)

type CreateIngredientCategoryRequestDTO struct {
	Name string
}

type IngredientCategoryService interface {
	CreateIngredientCategory(requestDTO CreateIngredientCategoryRequestDTO) (ingredient_categories.IngredientCategory, error)
	GetIngredientCategoryByID(id int) (ingredient_categories.IngredientCategory, error)
	GetAllIngredientCategories() ([]ingredient_categories.IngredientCategory, error)
	UpdateIngredientCategory(id int, requestDTO CreateIngredientCategoryRequestDTO) (ingredient_categories.IngredientCategory, error)
	DeleteIngredientCategory(id int) error
}

type ingredientCategoryService struct {
	ir ingredient_categories.IIngredientCategoryRepository // 抽象リポジトリ
}

func NewIngredientCategoryService(ir ingredient_categories.IIngredientCategoryRepository) IngredientCategoryService {
	return &ingredientCategoryService{ir: ir}
}

func (is *ingredientCategoryService) CreateIngredientCategory(requestDTO CreateIngredientCategoryRequestDTO) (ingredient_categories.IngredientCategory, error) {
	id := uuid.New().String()
	ingredientCategory, err := ingredient_categories.NewIngredientCategory(id, requestDTO.Name)

	if err != nil {
		return ingredient_categories.IngredientCategory{},
			errors.New("failed to create ingredient category")
	}

	is.ir.CreateIngredientCategory(ingredientCategory)

	return ingredientCategory, nil

}

func (is *ingredientCategoryService) GetIngredientCategoryByID(id int) (ingredient_categories.IngredientCategory, error) {
	panic("implement me")
}

func (is *ingredientCategoryService) GetAllIngredientCategories() ([]ingredient_categories.IngredientCategory, error) {
	panic("implement me")
}

func (is *ingredientCategoryService) UpdateIngredientCategory(id int, requestDTO CreateIngredientCategoryRequestDTO) (ingredient_categories.IngredientCategory, error) {
	panic("implement me")
}

func (is *ingredientCategoryService) DeleteIngredientCategory(id int) error {
	panic("implement me")
}
