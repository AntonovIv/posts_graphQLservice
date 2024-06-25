FROM golang:1.21-alpine AS builder

WORKDIR /user/local/src
RUN apk --no-cache add bash make # gcc gettext musl-dev

# dependencies
COPY go.mod go.sum ./
RUN go mod download

# build
COPY . .
RUN go build -o ./bin/app cmd/main.go

# run
FROM alpine AS runner

COPY --from=builder /app/bin/app /
COPY docker.env .env
CMD ["/app"]