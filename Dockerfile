FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o main -a

// TODO: needs to change
# second stage #
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /internal/app/main .
COPY --from=builder /internal/codera_server/config/app.yaml .
EXPOSE 3000
CMD ["./main", "-p=3000"]