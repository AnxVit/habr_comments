FROM golang:alpine

RUN apk add --no-cache gcc musl-dev

WORKDIR /cmd

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o ./bin/api ./cmd/web \
    && go build -o ./bin/migrate ./cmd/migrate

CMD ["/cmd/bin/api"]