package service

import (
	"errors"

	"github.com/RyotaAbe1014/Pastapal/internal/domain/ingredients"
)

func IngredientDuplicateCheck(ingredientRepository ingredients.IIngredientRepository, name string) error {
	result, err := ingredientRepository.FindByName(name)
	if err != nil {
		return err
	}

	if result == nil {
		return nil
	}

	return errors.New("ingredient already exists")
}
