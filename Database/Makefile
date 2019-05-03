clean:
	sudo docker stop userprofile
	sudo docker rm userprofile
	sudo docker rmi userprofile

gopath:
	GOPATH=$(PWD)
	echo $GOPATH
start: 
	sudo systemctl start docker

network:
	sudo docker network create --driver bridge cloudwalker

build:
	sudo docker build -t userprofile .

run:
	sudo docker run --name userprofile --network cloudwalker -p 3001:3001 -td userprofile

mongo:
	sudo docker run --name mongodb --network cloudwalker -p 27017:27017 -d mongo:3.7

test-ping:
	curl localhost:3001/ping














