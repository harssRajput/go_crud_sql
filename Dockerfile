FROM golang:1.21.10-alpine

# Install CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon@latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Set the PORT environment variable
ENV PORT 8080

# Expose the port specified by the PORT environment variable
EXPOSE ${PORT}

# Command to run CompileDaemon
CMD ["CompileDaemon", "--build=go build -o main", "--command=./main"]
