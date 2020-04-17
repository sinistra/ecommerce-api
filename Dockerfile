######## Start a new stage from scratch #######
FROM alpine
# FROM scratch

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

# copy api into build
WORKDIR /
COPY ./api .
COPY ./.env .
COPY ./templates ./templates
COPY ./certs ./certs
# COPY /etc/ssl /etc/ssl

EXPOSE 8000
ENV GIN_MODE=release

CMD ["/api"]
