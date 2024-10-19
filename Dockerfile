FROM golang:1.23.0 AS build-stage 

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /blog-agg .

FROM alpine:latest  
WORKDIR /root/
COPY --from=build-stage /blog-agg .
CMD ["./blog-agg"]
