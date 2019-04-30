clean:
	docker stop loginsystem
	docker rm loginsystem

gopath:
	GOPATH=$(PWD)
	echo $GOPATH
start: 
	sudo systemctl start docker

build:
	docker build -t loginsystem-app .

run:
	docker run -d  --name loginsystem -p 3001:3001 \
       	 	-e "MONGO1=127.0.0.1" \
        	# -e "MONGO2=54.148.30.107" \
        	# -e "MONGO3=54.203.107.100" \
		-e "MONGO_ADMIN=admin" \
		-e "MONGO_PASSWORD=query" \
		loginsystem-app

test-ping:
	curl localhost:3001/ping




