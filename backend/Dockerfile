# Stage 1: Build the application
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install Git so Go can download modules directly if GOPROXY=direct is used
RUN apk add --no-cache git
# ---------------------------

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# # Tidy the modules to ensure all imports are resolved and stored correctly
RUN go mod tidy
# # ---------------------------

# Copy the rest of the source code
COPY . .

# Build the executable
# -o main tells Go to name the output file 'main'
# ./cmd/api is the path to your main package (we'll create this structure later)
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go


# Stage 2: Create the final, minimal image
FROM alpine:latest

# Set the working directory to the path where we placed the executable
WORKDIR /root/

# Copy the built executable from the builder stage
COPY --from=builder /app/main .

# Expose the port (8080) that the Go application will listen on
EXPOSE 8080

# Run the executable when the container starts
CMD ["./main"]
