go mod tidy
go build
docker rm -f debug-mongodb-1
docker compose -f debug/docker-compose-local-mongo.yaml up -d
./star_atlas_server -alsologtostderr -path ./debug/config.yaml