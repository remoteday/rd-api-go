build:
	go build ./cmd/api

dev-dependencies:
	./scripts/dev-dependencies.sh

list-packages:
	go list ./...

test:
	GO_ENV=test ./scripts/test.sh

lint:
	./scripts/lint.sh

coverage:
	./scripts/coverage.sh

coverage.html:
	./scripts/coverage.sh --html

coverage.coveralls:
	./scripts/coverage.sh --coveralls

doc-gen:
	swag init -o ./src/docs -g ./cmd/api/routes/route.go

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go get -u -t -d -v ./...

deps-cleancache:
	go clean -modcache

server:
	go run ./cmd/api

dev:
	reflex -s -r '\.go$'' go run ./cmd/api

server-with-migrate:
	go run ./cmd/api with-migrate

migrate:
	go run ./cmd/api migrate

seed:
	go run ./cmd/api seed

wire:
	wire gen ./cmd/api