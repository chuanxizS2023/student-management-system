# start golang service
FROM golang

WORKDIR /app

# Copy every files into docker working dire
COPY . /app

# Download dependencies
RUN go mod download

COPY src/ src/

WORKDIR /app/src/cmd

# Build project
RUN go build -o main .

# Expose port in main.go
EXPOSE 9000

# Run executable
CMD ["./main"]