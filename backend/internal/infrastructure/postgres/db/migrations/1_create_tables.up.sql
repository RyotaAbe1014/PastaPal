-- 食材種別
CREATE TABLE IF NOT EXISTS ingredient_categories (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

-- 食材
CREATE TABLE IF NOT EXISTS ingredients (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  name VARCHAR(255) NOT NULL,
  category_id uuid NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY (category_id) REFERENCES ingredient_categories(id)
);