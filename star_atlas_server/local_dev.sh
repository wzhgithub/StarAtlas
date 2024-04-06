rm -fr ./log/*
cp go.mod.bk go.mod
go mod tidy
go build
# docker compose rm -f
# docker compose -f debug/docker-compose-local-mongo.yaml up -d
./star_atlas_server -log_dir=./log -alsologtostderr -path ./debug/config.yaml
