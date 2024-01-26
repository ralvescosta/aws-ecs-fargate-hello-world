FROM golang as build

WORKDIR /build
COPY . .
RUN go build -o=/build/myapp -ldflags "-s -w"

FROM alpine

RUN adduser -D appuser
WORKDIR /
COPY .env* /
COPY --from=build /build/myapp /myapp
RUN chmod +x ./myapp
USER appuser
EXPOSE 5000

CMD ["./myapp", "grpc"]