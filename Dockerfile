FROM golang:1.23 AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/app ./cmd/api

FROM gcr.io/distroless/base-debian12
ENV GIN_MODE=release
WORKDIR /app
COPY --from=build /bin/app /app/app
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app/app"]