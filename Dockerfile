# Use Go 1.23 bookworm as base imag
FROM golang:1.23-bookworm AS base

# Builder stage
# ======================================================
# Create a builder stage based on base Image
FROM base AS build-stage

WORKDIR /build

# Copy the go.mod and go.sum files to the /build directory
COPY . .

# Install dependencies
RUN go mod download

# Copy the entire source code into the container


# Build the application
# Turn off CGO to ensure static binaries
RUN CGO_ENABLED=0 go build -o app .

# Test stage
#=======================================================
# Run the tests in the container
FROM build-stage AS run-test-stage

RUN go test -v ./...

# Production stage
# ======================================================
# Create a production stage to run the application binary
FROM scratch AS build-release-stage

# Move to working directory /prod
WORKDIR /

# Copy binary from builder stage
COPY --from=build-stage /build/app /app

# Document the port that may need to be published
EXPOSE 8000

# Start the application
CMD ["/app"]