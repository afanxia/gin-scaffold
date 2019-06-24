.PHONY: start build

NOW = $(shell date -u '+%Y%m%d%I%M%S')

SERVER_BIN = "./cmd/[[.projectName]]/[[.projectName]]"
RELEASE_ROOT = "release"
RELEASE_SERVER = "release/[[.projectName]]"

all: start

build:
	@go build -ldflags "-w -s" -o $(SERVER_BIN) ./cmd/[[.projectName]]

start: 
	@go build -o $(SERVER_BIN) ./cmd/[[.projectName]]
	$(SERVER_BIN) -c ./configs/[[.projectName]]/config.toml -m ./configs/[[.projectName]]/model.conf -swagger ./internal/app/swagger

swagger:
	swaggo -s ./internal/app/swagger.go -p . -o ./internal/app/swagger

test:
	@go test -cover -race ./...

clean:
	rm -rf data release $(SERVER_BIN) ./internal/app/admin/test/data ./internal/app/test/data

pack: build
	rm -rf $(RELEASE_ROOT)
	mkdir -p $(RELEASE_SERVER)
	cp -r $(SERVER_BIN) configs $(RELEASE_SERVER)
	cd $(RELEASE_ROOT) && zip -r [[.projectName]].$(NOW).zip "[[.projectName]]"
