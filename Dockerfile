FROM golang:1.25.0-alpine3.22
  
# Working directory inside container
WORKDIR /app

# Install air for live reload
RUN go install github.com/air-verse/air@latest

# Copy dependencies description and install them
COPY go.mod ./
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Expose service port
EXPOSE 8192