FROM golang:1.21-alpine

WORKDIR /app

COPY . ./

RUN go mod download && go mod verify


CMD ["go", "run", "."]
EXPOSE 8080