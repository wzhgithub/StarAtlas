FROM scratch

WORKDIR /app

COPY start_atlas_server /app/
COPY config.yaml /app/
ENV GIN_MODE=release 

EXPOSE 9999

ENTRYPOINT ["./start_atlas_server"]
