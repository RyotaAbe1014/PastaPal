# PastaPal Backend

## 環境構築
build/ディレクトリに移動してください。
```bash
docker-compose up -d
```

## コマンド
Makefileに移行予定ですが、現状は以下のコマンドを使用してください。
### appコンテナ内に入る
```bash
docker-compose exec api sh
```

### マイグレーションコンテナに入る
```bash
docker-compose exec migrate sh
```


### マイグレーション
#### マイグレーションファイルを作成する
```bash
migrate create -ext sql -dir /migrations -seq xxx
```

#### マイグレーションを実行する
```bash
migrate -path /migrations -database postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable up
```

#### マイグレーションを戻す
```bash
migrate -path /migrations -database postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable down
```
