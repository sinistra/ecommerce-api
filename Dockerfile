FROM golang:1.14-alpine3.12

RUN mkdir /app
ADD . /app
WORKDIR /app
# Add this go mod download command to pull in any dependencies
RUN go mod download
# Our project will now successfully build with the necessary go libraries included.
RUN go build -o api .

COPY prod.env .env
EXPOSE 8000
ENV GIN_MODE=release

# Our start command which kicks off
# our newly created binary executable
CMD ["/app/api"]
