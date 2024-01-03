# start golang service
FROM golang

WORKDIR /app

# Copy every files into docker working dire
COPY . /app

# Download dependencies
RUN go mod download    

COPY src/ src/

WORKDIR /app/src/cmd

# Expose port in main.go
EXPOSE 9000


CMD ["go", "run", "main.go"]
# Run executable
