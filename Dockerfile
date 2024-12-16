FROM golang:1.22-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /routines main.go

# Final stage: Minimal image
FROM scratch

WORKDIR /app

COPY --from=build-stage /routines /app/routines
COPY --from=build-stage /app/.env /app
COPY --from=build-stage /app/db/migration /app/db/migration

EXPOSE 8000

ENTRYPOINT ["./routines"]