package ingredient_categories

import "time"

type IngredientCategory struct {
	id         string
	name       string
	created_at time.Time
	updated_at time.Time
}

func NewIngredientCategory(id string, name string) (IngredientCategory, error) {
	// TODO: add validation
	return IngredientCategory{id: id, name: name}, nil
}

func (i IngredientCategory) ID() string {
	return i.id
}

func (i IngredientCategory) Name() string {
	return i.name
}

// TODO: add private validate method
