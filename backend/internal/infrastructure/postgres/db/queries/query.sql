-- name: ListIngredientsCategory :many
SELECT * FROM ingredient_categories
ORDER BY name;

-- name: GetIngredientCategory :one
SELECT * FROM ingredient_categories
WHERE id = $1 LIMIT 1;

-- name: FindIngredientCategoryByName :one
SELECT * FROM ingredient_categories
WHERE name = $1 LIMIT 1;

-- name: CreateIngredientCategory :one
INSERT INTO ingredient_categories (
  id, name
) VALUES (
  $1,
  $2
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

-- name: GetIngredient :one
SELECT * FROM ingredients
WHERE id = $1 LIMIT 1;

-- name: FindIngredientByName :one
SELECT * FROM ingredients
WHERE name = $1 LIMIT 1;

-- name: CreateIngredient :one
INSERT INTO ingredients (
  id, name, ingredient_category_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteIngredient :exec
DELETE FROM ingredients
WHERE id = $1;

-- name: UpdateIngredient :one
UPDATE ingredients
  set name = $2,
  ingredient_category_id = $3
WHERE
  id = $1
RETURNING *;