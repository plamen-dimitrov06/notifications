# :latest
FROM golang

# creates the dir since it doesn't exist
WORKDIR /usr/src/notifications

COPY . .

RUN go run cmd/notifications/*.go

EXPOSE 8080