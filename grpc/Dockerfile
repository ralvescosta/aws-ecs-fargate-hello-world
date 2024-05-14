FROM golang:1.21.6 as build

WORKDIR /build
COPY . .
RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -o=myapp -ldflags "-s -w"

FROM scratch

WORKDIR /app
COPY .env* ./
COPY --from=build /build/myapp ./
EXPOSE 5000

CMD ["./myapp", "grpc"]