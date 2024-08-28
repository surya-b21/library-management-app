FROM golang:1.22-alpine as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /book-service

FROM gcr.io/distroless/base-debian12

COPY --from=build /book-service /book-service

EXPOSE 8080

CMD ["/book-service"]