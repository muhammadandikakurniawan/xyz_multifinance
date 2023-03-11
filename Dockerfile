FROM golang:1.18-alpine AS BUILDER
RUN apk add --no-cache gcc g++ git openssh-client

WORKDIR /app

COPY . .
COPY cmd/app/docs ./docs

RUN go mod download
RUN go mod tidy
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-w -s" -o binary cmd/app/main.go

ENTRYPOINT ["/app/binary"]