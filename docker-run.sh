docker build -t parking-service .
docker run --rm -p 8080:8080 parking-service
docker image rm parking-service
