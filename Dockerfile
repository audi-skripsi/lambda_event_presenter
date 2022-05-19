FROM golang:1.17 as build

WORKDIR /app

COPY go.mod /app/
COPY go.sum /app/

RUN go mod download

COPY . /app/

RUN go build -a -o /app/main

# --------

FROM debian:stable-slim

WORKDIR /app

ENV APP_ADDRESS=:8080
ENV APP_NAME=lambda_event_presenter

ENV KAFKA_ADDRESS=172.17.0.1:29092,172.17.0.1:29093,172.17.0.1:29094
ENV KAFKA_CONSUMER_GROUP=audi.skripsi.lambda_speed_event_presenter_group_tes
ENV KAFKA_IN_TOPIC=audi.skripsi.lambda_speed_event_level_standardizer

ENV MONGODB_DB_NAME=audi_skripsi_lambda
ENV MONGODB_ADDRESS=172.17.0.1:27017

ENV REDIS_ADDRESS=172.17.0.1:6379
ENV REDIS_PASSWORD=

COPY --from=build /app/main /app/main

CMD ["./main"]