package ingredients

import (
	"errors"
	"time"
)

type Ingredient struct {
	id                   string
	ingredientCategoryID string
	name                 string
	createdAt            time.Time
	updatedAt            time.Time
}

// constructor
func NewIngredient(id string, name string, ingredientCategoryID string) (Ingredient, error) {
	if err := validateName(name); err != nil {
		return Ingredient{}, err
	}
	return Ingredient{id: id, name: name, ingredientCategoryID: ingredientCategoryID}, nil
}

// 永続化層から取得したデータをドメイン層に変換する
func NewIngredientFromRepository(id string, name string, ingredientCategoryID string, createdAt time.Time, updatedAt time.Time) (Ingredient, error) {
	if err := validateName(name); err != nil {
		return Ingredient{}, err
	}
	return Ingredient{id: id, name: name, ingredientCategoryID: ingredientCategoryID, createdAt: createdAt, updatedAt: updatedAt}, nil
}

// getter
func (i Ingredient) ID() string {
	return i.id
}

func (i Ingredient) Name() string {
	return i.name
}

func (i Ingredient) IngredientCategoryID() string {
	return i.ingredientCategoryID
}

// validation
func validateName(name string) error {
	if name == "" {
		return errors.New("食材名は必須です")
	}

	if len(name) > 90 {
		return errors.New("食材名は90文字以内で入力してください")
	}
	return nil
}

// setter
func (i *Ingredient) UpdateName(name string) error {
	if err := validateName(name); err != nil {
		return err
	}
	i.name = name
	return nil
}

func (i *Ingredient) UpdateIngredientCategoryID(ingredientCategoryID string) error {
	i.ingredientCategoryID = ingredientCategoryID
	return nil
}
