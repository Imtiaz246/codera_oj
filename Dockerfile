FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o main -a

// TODO: needs to change
# second stage #
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/config/app.yaml .
EXPOSE 3000
CMD ["./main", "-p=3000"]