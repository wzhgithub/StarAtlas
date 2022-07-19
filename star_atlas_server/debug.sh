go mod tidy
go build
docker compose rm -f
docker compose -f debug/docker-compose-local-mongo.yaml up -d
./star_atlas_server -alsologtostderr -path ./debug/config.yaml
