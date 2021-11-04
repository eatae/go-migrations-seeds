
include .env
MIGRATIONS_DIR := "db/migrations"
DB_CONNECTION := "user=${DOCKER_DB_USER} password=${DOCKER_DB_PASSWORD} port=${DOCKER_DB_PORT} dbname=${DOCKER_DB_NAME} sslmode=disable"


.PHONY: migration-connect
migration-connect:
	${GOPATH}/bin/goose -dir ${MIGRATIONS_DIR} postgres ${DB_CONNECTION} status

.PHONY: migration-create
migration-create:
	${GOPATH}/bin/goose -dir ${MIGRATIONS_DIR} create ${name} sql

.PHONY: migration-up
migration-up:
	${GOPATH}/bin/goose -dir ${MIGRATIONS_DIR} postgres ${DB_CONNECTION} up

.PHONY: migration-down
migration-down:
	${GOPATH}/bin/goose -dir ${MIGRATIONS_DIR} postgres ${DB_CONNECTION} down


.PHONY: debug
debug:
	echo ${DB_CONNECTION}
	echo ${GOPATH}