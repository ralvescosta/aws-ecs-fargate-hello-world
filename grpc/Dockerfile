FROM golang as build

ENV APP_USER=appuser 

WORKDIR /build
COPY . .
RUN go build -o=exec -ldflags "-s -w"

FROM alpine

WORKDIR /app
COPY .env* ./
COPY --from=build /build/exec ./exec
USER $APP_USER
EXPOSE 5000

CMD ["sh", "-c", "./exec"]