.PHONY: all clean build run stop
all: clean build run
clean:
	docker stop apollo-frontend || :
	docker rm apollo-frontend || :
build: 
	docker build -t noah-huppert/apollo-frontend:dev-latest .
run:
	docker run --name apollo-frontend -p 3000:80 -d noah-huppert/apollo-frontend:dev-latest || docker start apollo-frontend
stop:
	docker stop apollo-frontend || :
