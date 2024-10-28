# Official Golanf Image
FROM golang:1.23-alpine AS builder

# WorkDir
WORKDIR /app

# Copy mod files
COPY go.mod go.sum ./
RUN go mod download

# Now copy the whole code
COPY . .

# Copy the .env file
COPY .env ./

# Build
RUN go build -o rest-auth main.go

# For the final stage using alpine
FROM alpine:latest

# Working Directory
WORKDIR /app

# Copy Binary file
COPY --from=builder /app/rest-auth .

# Port 8080
EXPOSE 8080

# Run the app
CMD [ "./rest-auth" ]
