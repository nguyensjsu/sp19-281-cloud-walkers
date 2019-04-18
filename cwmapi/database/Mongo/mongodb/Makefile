clean:
	docker stop mongodb
	docker rm mongodb

network-create:
	docker network create --driver bridge cmpe281-network

network-inspect:
	docker network inspect cmpe281-network

run:
	docker run -d --name mongodb -p 27017:27017 \
	--network cmpe281-network \
	-v $(PWD)/data:/data/db mongo:latest

bash:
	docker exec -it mongodb bash 