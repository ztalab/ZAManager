FROM golang:1.17.8-alpine AS builder

WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 go build -o ZAManager .


FROM alpine AS final
WORKDIR /app
COPY --from=builder /build/ZAManager /app/
COPY --from=builder /build/config.yaml /app/
COPY --from=builder /build/db /app/db

ENTRYPOINT ["/app/ZAManager"]
