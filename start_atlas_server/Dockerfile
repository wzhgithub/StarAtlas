FROM scratch

WORKDIR /app

COPY start_atlas_server /app/
ENV GIN_MODE=release 

ENTRYPOINT ["./start_atlas_server"]