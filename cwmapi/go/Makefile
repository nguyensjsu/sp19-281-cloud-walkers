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
       	 	-e "MONGO1=10.0.3.78" \
        	-e "MONGO2=10.0.3.218" \
        	-e "MONGO3=10.0.3.224" \
		-e "MONGO_ADMIN=admin" \
		-e "MONGO_PASSWORD=query" \
		-e "USER_ACTIVITY_PROXY=http://10.0.2.43:3000" \
		cwmapi-app

test-ping:
	curl localhost:3000/ping

test-get-topics:
	curl -X GET \
	http://localhost:3000/msgstore/v1/topics \
	-H 'Postman-Token: 6a1657aa-9305-422c-816a-817f449b4cf4' \
	-H 'cache-control: no-cache'


