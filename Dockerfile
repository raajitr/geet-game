# build stage
FROM golang:1.21.0-alpine3.17 AS build
WORKDIR /app
COPY . .

RUN go build -o main main.go

# Run stage
FROM alpine:3.18.3
WORKDIR /app
COPY --from=build /app/main .

EXPOSE 8080
CMD [ "/app/main" ]