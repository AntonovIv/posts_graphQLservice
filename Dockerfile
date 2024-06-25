FROM golang:1.21-alpine AS builder

WORKDIR /usr/local/src
RUN apk --no-cache add bash # gcc gettext musl-dev

# dependencies
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

# build
COPY . ./
RUN go build -o ./bin/postapp cmd/main.go

# run
FROM alpine AS runner

COPY --from=builder usr/local/src/bin/postapp /
COPY .env /.env
CMD ["/postapp"]