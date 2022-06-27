FROM centos

WORKDIR /app

COPY star_atlas_server /app/
COPY config.yaml /app/
ENV GIN_MODE=release 

#EXPOSE 9999

ENTRYPOINT ["./star_atlas_server"]
