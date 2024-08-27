FROM golang:1.22-alpine as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /auth-service

FROM gcr.io/distroless/base-debian12

COPY --from=build /auth-service /auth-service

EXPOSE 8080

CMD ["/auth-service"]