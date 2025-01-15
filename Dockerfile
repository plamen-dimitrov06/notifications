# :latest
FROM golang

# creates the dir since it doesn't exist
WORKDIR /usr/src/notifications

COPY . .

RUN go mod download