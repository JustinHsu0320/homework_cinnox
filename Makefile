
network:
	docker network create cinnox

mongodb:
	docker run -d --name mongodb -p 27017:27017 --network cinnox -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=secret mongo:4.4-rc


createcollection:
	docker exec -it mongodb mongo -u root -p secret --authenticationDatabase admin cinnox --eval "db.createCollection('users')"

server:
	go run main.go

.PHONY: network mongodb createcollection server