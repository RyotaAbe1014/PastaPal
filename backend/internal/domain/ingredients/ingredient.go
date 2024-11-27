package ingredients

import (
	"errors"
	"time"
)

type Ingredient struct {
	id                     string
	ingredient_category_id string
	name                   string
	created_at             time.Time
	updated_at             time.Time
}

// constructor
func NewIngredient(id string, name string, ingredient_category_id string) (Ingredient, error) {
	if err := validateName(name); err != nil {
		return Ingredient{}, err
	}
	return Ingredient{id: id, name: name, ingredient_category_id: ingredient_category_id}, nil
}

// 永続化層から取得したデータをドメイン層に変換する
func NewIngredientFromRepository(id string, name string, ingredient_category_id string, created_at time.Time, updated_at time.Time) (Ingredient, error) {
	if err := validateName(name); err != nil {
		return Ingredient{}, err
	}
	return Ingredient{id: id, name: name, created_at: created_at, updated_at: updated_at}, nil
}

// getter
func (i Ingredient) ID() string {
	return i.id
}

func (i Ingredient) Name() string {
	return i.name
}

func (i Ingredient) IngredientCategoryID() string {
	return i.ingredient_category_id
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
