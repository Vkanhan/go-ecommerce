# syntax=docker/dockerfile:1

FROM golang:1.23 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download 

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /blog-agg .

FROM alpine:latest
COPY --from=build-stage /app/blog-agg .

EXPOSE 8080

CMD [ "./blog-agg" ]