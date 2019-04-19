clean: 
	docker stop cloudwalker
	docker rm cloudwalker
	docker rmi cloudwalker

build:
	docker build -t cloudwalker .

run:
	docker run --name cloudwalker --network cloudwalker -p 3000:3000 -td cloudwalker

