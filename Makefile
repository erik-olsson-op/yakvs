dockertag = erikolssonop/yakvs

# Compile the proto file and out go file and go_grpc file to /model
proto-compile:
	 protoc --proto_path=proto --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. proto/model/*.proto

# Build the docker image
docker-build:
	docker build -t $(dockertag) .

# Push the image to public repository
docker-push:
	docker push $(dockertag)

# Set exposed port 8080 and 8081 default and pass the .env file
docker-run:
	docker run --env-file .env -it --rm -p 8443:8443 -p 7433:7443 $(dockertag)

benchmark-http:
	cd test/http && ./set.sh
	cd test/http && ./get.sh
	cd test/http && ./delete.sh

benchmark-grpc:
	cd test/grpc && ./set.sh
	cd test/grpc && ./get.sh
	cd test/grpc && ./delete.sh