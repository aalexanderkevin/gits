dep:
	@echo ">> Downloading Dependencies"
	@go mod download

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./transport/grpc/gitspb/gits.proto

run-grpc:
	env $$(cat .env | xargs) go run test3/cmd server
	
migrate:
	eval $$(egrep -v '^#' .env | xargs -0) go run test3/cmd migrate

test-all: test-unit test-integration-with-infra

test-unit: dep
	@echo ">> Running Unit Test"
	@env $$(cat .env.testing | xargs) go test -tags=unit -failfast -cover -covermode=atomic ./...

test-integration: dep
	@echo ">> Running Integration Test"
	@env $$(cat .env.testing | xargs) env POSTGRES_MIGRATION_PATH=$$(pwd)/database/migrations go test -tags=integration -failfast -cover -covermode=atomic ./...

test-integration-with-infra: test-infra-up test-integration test-infra-down

test-infra-up:
	$(MAKE) test-infra-down
	@echo ">> Starting Test DB"
	docker run -d --rm --name test-postgres -p 5431:5432 --env-file .env.testing postgres:12

test-infra-down:
	@echo ">> Shutting Down Test DB"
	@-docker kill test-postgres