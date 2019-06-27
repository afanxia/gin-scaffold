.PHONY: start build

PROJECT_NAME = github.com/afanxia/gohspt

all: start

build:
	@go build

start: 
	@go build
	./gin-scaffold -i=.go -i=.html -i=Makefile -i=.sql -i=.toml -i=.conf -i=.md -i=.gitignore -i=go.mod -t=portal generate -d="portal_db" -u="root" -p="asdfasdf" $(PROJECT_NAME)

clean:
	rm ./gin-scaffold
	rm -rf ~/go/src/$(PROJECT_NAME)
