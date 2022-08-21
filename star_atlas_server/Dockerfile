FROM golang:1.18 as builder

# 启用go module
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go mod tidy && go build . \
    && mkdir publish && cp star_atlas_server publish && cp config.yaml publish \
    && cp utest publish && cp gen_udp.sh publish && cp trans publish && cp -R conf publish

FROM centos

WORKDIR /app

COPY --from=builder /app/publish .
ENV GIN_MODE=release 

EXPOSE 9999/tcp 9191/udp 9192/tcp

ENTRYPOINT ["./star_atlas_server", "-log_dir=./", "-alsologtostderr"]
