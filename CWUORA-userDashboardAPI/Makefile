clean: 
	docker stop cloudwalker
	docker rm cloudwalker
	docker rmi cloudwalker

build:
	docker build -t cloudwalker .

run:
	docker run --name cloudwalker --network cloudwalker -p 3000:3000 -td cloudwalker

mongo:
	docker run --name mongodb --network cloudwalker -p 27017:27017 -d mongo:3.7

