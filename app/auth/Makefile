.PHONY: build run all
all: build run
build:
	docker build -t noah-huppert/apollo-auth:dev-latest .
run:
	docker run --rm --name apollo-auth -p 3001:80 noah-huppert/apollo-auth:dev-latest
