include .env

.PHONY: proto

init:
	@make up
	@make db-recreate

rebuild-up:
	@make down
	@make build
	@make up

reup:
	@make down
	@make up

restart:
	@make stop
	@make up

build:
	docker compose build
	@echo "\ncontainer build done🎉\n"

up:
	docker compose up -d
	@echo "\ncontainer up done🎉\n"

stop:
	docker compose stop
	@echo "\ncontainer stop done🎉\n"

down:
	docker compose down
	@echo "\ncontainer down done🎉\n"

ps:
	docker ps
	@echo "\ncontainer ps done🎉\n"

exec:
	docker compose exec $(APP_NAME) /bin/bash
	@echo "\napp container exec done🎉\n"

log:
	docker compose logs -f $(APP_NAME)
	@echo "\napp container log done🎉\n"

db-recreate:
	@make db-drop
	@make db-migrate-diff
	@make db-migrate-apply

db-migrate-diff:
	docker pull arigaio/atlas:latest
	docker run --rm --net=host -v $(CURDIR):/app arigaio/atlas \
		migrate diff \
			--dir "file://app/migrations" \
			--to "file://app/atlas.hcl" \
			--dev-url "$(DB_DRIVER)://$(DB_USERNAME):$(DB_PASSWORD)@127.0.0.1:$(DB_PORT)/$(DB_NAME)" \
			--format '{{ sql . "  " }}'
	@echo "\ndb migrate diff done🎉\n"

db-migrate-apply:
	docker pull arigaio/atlas:latest
	docker run --rm --net=host -v $(CURDIR):/app arigaio/atlas \
		migrate apply --dir "file://app/migrations" --url "$(DB_DRIVER)://$(DB_USERNAME):$(DB_PASSWORD)@127.0.0.1:$(DB_PORT)/$(DB_NAME)"
	@echo "\ndb migrate apply done🎉\n"

db-drop:
	docker run --rm --net=host arigaio/atlas schema clean --auto-approve -u "$(DB_DRIVER)://$(DB_USERNAME):$(DB_PASSWORD)@127.0.0.1:$(DB_PORT)/$(DB_NAME)"
	@echo "\ndb drop done🎉\n"

go-check:
	@make go-fmt
	@make go-lint

go-fmt:
	docker-compose exec $(APP_NAME) go fmt ./...
	@echo "\ngo fmt done🎉\n"

go-lint:
	docker-compose exec $(APP_NAME) staticcheck -f stylish ./... 
	@echo "\ngo lint done🎉\n"

go-mod-tidy:
	docker-compose exec $(APP_NAME) go mod tidy
	@echo "\ngo mod tidy done🎉\n"

go-test:
	docker-compose exec $(APP_NAME) go test ./... -cover -coverprofile=../cover.out
	@echo "\ngo test done🎉\n"

go-update-schema:
	oapi-codegen -package "openapi" ../doc/openapi.yaml > internal/gen/openapi/doc.go
