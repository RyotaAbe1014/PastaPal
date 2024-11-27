package ingredient_categories

import (
	"errors"
	"time"
)

type IngredientCategory struct {
	id        string
	name      string
	createdAt time.Time
	updatedAt time.Time
}

// constructor
func NewIngredientCategory(id string, name string) (IngredientCategory, error) {
	if err := validateName(name); err != nil {
		return IngredientCategory{}, err
	}
	return IngredientCategory{id: id, name: name}, nil
}

// 永続化層から取得したデータをドメイン層に変換する
func NewIngredientCategoryFromRepository(id string, name string, createdAt time.Time, updatedAt time.Time) (IngredientCategory, error) {
	if err := validateName(name); err != nil {
		return IngredientCategory{}, err
	}
	return IngredientCategory{id: id, name: name, createdAt: createdAt, updatedAt: updatedAt}, nil
}

// getter
func (i IngredientCategory) ID() string {
	return i.id
}

func (i IngredientCategory) Name() string {
	return i.name
}

// validation
func validateName(name string) error {
	if name == "" {
		return errors.New("食材種別名は必須です")
	}

	if len(name) > 30 {
		return errors.New("食材種別は30文字以内で入力してください")
	}
	return nil
}
