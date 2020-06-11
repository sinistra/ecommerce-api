FROM golang:1.14 AS builder

RUN mkdir /app
#add source code to /app folder
ADD . /app
WORKDIR /app
# Call go mod command to pull in any dependencies
RUN go mod download
# Project will now successfully build with the necessary libraries included.
# RUN go build -o api .
RUN CGO_ENABLED=0 GOOS=linux go build -o api .

#lightweight container to start with
FROM alpine:latest AS production
#set pwd to /
WORKDIR /
# Copy the compiled app from the builder to production
COPY --from=builder /app/api .
# copy the production env file from dev
COPY ./prod.env .env
# create public folder and copy html files needed
RUN mkdir /public
COPY ./public/index.html /public/index.html

EXPOSE 8000
# set GIN-GONIC to release mode
ENV GIN_MODE=release

# Start command which kicks off binary executable
CMD ["/api"]
