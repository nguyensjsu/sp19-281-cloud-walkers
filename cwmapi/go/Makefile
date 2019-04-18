clean:
	docker stop cwmapi
	docker rm cwmapi

gopath:
	GOPATH=$(PWD)
	echo $GOPATH
start: 
	sudo systemctl start docker

build:
	docker build -t cwmapi-app .

start-gumballs:
	docker run -d  --name cwmapi -p 3000:3000 cwmapi-app

test-ping:
	curl localhost:3000/ping


