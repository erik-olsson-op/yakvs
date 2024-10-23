dockertag = erikolssonop/yakvs

proto-compile:
	 protoc --proto_path=proto --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. proto/model/*.proto

# Build the docker image
docker-build:
	docker build -t $(dockertag) .

# Set exposed port 8080 default and pass the .env file
docker-run:
	docker run --env-file .env -it --rm -p 8080:8080 -p 8081:8081AQW2 C $(dockertag)