# 環境変数
DOCKER_COMPOSE = docker-compose
API_CONTAINER = api
MIGRATE_CONTAINER = migrate
DB_URL = postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable

help: # コマンド確認
	@echo "\033[32mAvailable targets:\033[0m"
	@grep "^[a-zA-Z\-]*:" Makefile | grep -v "grep" | sed -e 's/^/make /' | sed -e 's/://'

# コマンドの例
# make api-shell
app-shell: # apiコンテナに入る
	$(DOCKER_COMPOSE) exec $(API_CONTAINER) sh

# コマンドの例
# make migrate-shell
migrate-shell: # migrationコンテナに入る
	$(DOCKER_COMPOSE) exec $(MIGRATE_CONTAINER) sh

# コマンドの例
# make sqlc-generate
sqlc-generate: # sqlcを実行する
	$(DOCKER_COMPOSE) exec $(API_CONTAINER) sqlc generate

# コマンドの例
# make migrate-create name=xxx_create_table
migrate-create: # マイグレーションファイルを作成
	$(DOCKER_COMPOSE) exec $(MIGRATE_CONTAINER) migrate create -ext sql -dir /migrations -seq $(name)

# コマンドの例
# make migrate-up
migrate-up: # マイグレーションを実行する
	$(DOCKER_COMPOSE) exec $(MIGRATE_CONTAINER) migrate -path /migrations -database $(DB_URL) up

# コマンドの例
# make migrate-down
migrate-down: # マイグレーションをロールバック
	$(DOCKER_COMPOSE) exec $(MIGRATE_CONTAINER) migrate -path /migrations -database $(DB_URL) down 1