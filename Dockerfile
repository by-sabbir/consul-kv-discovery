FROM golang:latest AS builder

RUN mkdir /app

ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

FROM alpine:latest AS prod

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/app .
CMD [ "./app" ]
