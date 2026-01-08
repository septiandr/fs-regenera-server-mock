# Stage 1: build Go binary
FROM golang:1.25-alpine AS build

WORKDIR /app

# Copy go.mod & go.sum dulu
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh kode
COPY . .

# Build aplikasi
RUN go build -o fs-regenera .

# Stage 2: runtime minimal
FROM alpine:latest
WORKDIR /app

# Copy binary dari stage build
COPY --from=build /app/fs-regenera .

# Copy seluruh folder src/data ke container
COPY --from=build /app/src/data ./src/data

# Expose port backend
EXPOSE 9070

# Jalankan binary
CMD ["./fs-regenera"]
