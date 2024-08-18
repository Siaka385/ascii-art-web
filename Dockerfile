# Start from the official Go image
FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the entire current directory into the container
COPY . .

# Expose port 8080
EXPOSE 8080

# Run the Go application
CMD ["go", "run", "main.go"]