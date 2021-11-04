MIGRATIONS_DIR := "db/migrations"

.PHONY: migrate-generate

migration-create:
	$(GOPATH)/bin/goose -dir $(MIGRATIONS_DIR) create $(name) sql