# Dev stage
FROM golang:1.18 AS dev

ENV APP_HOME /go/src/github.com/michelsazevedo/kuala/
WORKDIR $APP_HOME

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

# Builder stage
FROM dev AS builder

ENV APP_HOME /go/src/github.com/michelsazevedo/kuala/
WORKDIR $APP_HOME

RUN CGO_ENABLED=0 GOOS=linux go build -o kuala .

# Production stage
FROM alpine:latest AS production

ENV APP_HOME /go/src/github.com/michelsazevedo/kuala/

COPY --from=builder $APP_HOME .

EXPOSE 8080

CMD ["./kuala"]
