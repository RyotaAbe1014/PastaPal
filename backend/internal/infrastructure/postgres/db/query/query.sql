-- name: ListIngredientsCategory :many
SELECT * FROM ingredient_categories
ORDER BY name;

-- name: GetIngredientCategory :one
SELECT * FROM ingredient_categories
WHERE id = $1 LIMIT 1;

-- name: CreateIngredientCategory :one
INSERT INTO ingredient_categories (
  name
) VALUES (
  $1
)
RETURNING *;

-- name: UpdateIngredientCategory :exec
UPDATE ingredient_categories
  set name = $2
WHERE id = $1;

-- name: DeleteIngredientCategory :exec
DELETE FROM ingredient_categories
WHERE id = $1;

-- name: ListIngredients :many
SELECT * FROM ingredients
ORDER BY name;

-- name: CreateIngredient :one
INSERT INTO ingredients (
  name, category_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: deleteIngredient :exec
DELETE FROM ingredients
WHERE id = $1;