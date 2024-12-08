package recipes

import (
	"errors"
	"time"
)

type Recipe struct {
	id           string
	name         string
	instructions string
	units        []Unit
	createdAt    time.Time
	updatedAt    time.Time
}

// constructor
func NewRecipe(id string, name string, instructions string, units []Unit) (Recipe, error) {
	if err := validateName(name); err != nil {
		return Recipe{}, err
	}
	return Recipe{id: id, name: name, instructions: instructions, units: units}, nil
}

func (r Recipe) ID() string {
	return r.id
}

func (r Recipe) Name() string {
	return r.name
}

func (r Recipe) Instructions() string {
	return r.instructions
}

func (r Recipe) Units() []Unit {
	return r.units
}

func validateName(name string) error {
	if name == "" {
		return errors.New("レシピ名は必須です")
	}

	if len(name) > 90 {
		return errors.New("レシピ名は90文字以内で入力してください")
	}
	return nil
}
