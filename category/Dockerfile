FROM golang:1.22-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /category-service

FROM gcr.io/distroless/base-debian12

COPY --from=build /category-service /category-service

EXPOSE 8080

CMD ["/category-service"]