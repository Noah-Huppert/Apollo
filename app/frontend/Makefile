.PHONY: all build run
all: build run
build: 
	docker build -t noah-huppert/apollo-frontend:dev-latest .
run:
	docker run --rm --name apollo-frontend -p 3000:80 noah-huppert/apollo-frontend:dev-latest
