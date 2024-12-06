-- 食材の単位を管理するテーブル
CREATE TABLE IF NOT EXISTS units (
  id SERIAL PRIMARY KEY,         -- ユニークな単位ID
  name VARCHAR(50) NOT NULL,     -- 表示名 (例: "g", "ml", "個", "本")
  description VARCHAR(255)       -- 説明 (例: "グラム単位", "ミリリットル単位")
);

-- 初期データの投入
INSERT INTO units (name, description)
VALUES
  ('g', 'グラム'),
  ('ml', 'ミリリットル'),
  ('個', '個数'),
  ('本', '本数');