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

run:
	docker run -d  --name cwmapi -p 3000:3000 \
       	 	-e "MONGO1=52.10.186.169" \
        	-e "MONGO2=54.148.30.107" \
        	-e "MONGO3=54.203.107.100" \
		-e "MONGO_ADMIN=admin" \
		-e "MONGO_PASSWORD=query" \
		cwmapi-app

test-ping:
	curl localhost:3000/ping

test-get-topics:
	curl -X GET \
	http://localhost:3000/msgstore/v1/topics \
	-H 'Postman-Token: 6a1657aa-9305-422c-816a-817f449b4cf4' \
	-H 'cache-control: no-cache'


