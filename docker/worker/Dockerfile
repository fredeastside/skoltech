FROM golang:latest as worker
WORKDIR /app
COPY ./go.mod go.sum ./
RUN go mod download
COPY ./ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/worker/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=worker /app/main .
CMD ["./main"]