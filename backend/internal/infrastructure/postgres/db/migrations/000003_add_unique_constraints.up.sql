ALTER TABLE ingredient_categories
ADD CONSTRAINT unique_ingredient_category_name UNIQUE (name);

ALTER TABLE ingredients
ADD CONSTRAINT unique_ingredient_name UNIQUE (name);