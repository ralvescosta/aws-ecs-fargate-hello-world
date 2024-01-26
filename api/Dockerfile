FROM golang as build

WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o=/build/myapp -ldflags "-s -w"

FROM scratch

WORKDIR /app
COPY .env* ./
COPY --from=build /build/myapp ./
EXPOSE 3333

CMD ["./myapp", "api"]