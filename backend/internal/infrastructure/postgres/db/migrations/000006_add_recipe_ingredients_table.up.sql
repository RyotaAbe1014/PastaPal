CREATE TABLE IF NOT EXISTS recipe_ingredients (
  recipe_id uuid NOT NULL,
  ingredient_id uuid NOT NULL,
  quantity NUMERIC(10, 2) NOT NULL, -- 数量を数値で保存
  unit_id INT NOT NULL,            -- 単位を参照
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (recipe_id, ingredient_id),
  FOREIGN KEY (recipe_id) REFERENCES recipes(id),
  FOREIGN KEY (ingredient_id) REFERENCES ingredients(id) ON DELETE CASCADE,
  FOREIGN KEY (unit_id) REFERENCES units(id)
);
