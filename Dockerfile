# Use an official Golang runtime as a parent image
FROM golang:latest
# Set ENV variables based on the ARGs
ENV SERVER_PORT=$SERVER_PORT
ENV SERVER_ADDRESS=$SERVER_ADDRESS
ENV DB_NAME=$DB_NAME
ENV DB_USER=$DB_USER
ENV DB_PASS=$DB_PASS
ENV DB_HOST=$DB_HOST
ENV DB_PORT=$DB_PORT
ENV HMAC_SAMPLE_SECRET=$HMAC_SAMPLE_SECRET
ENV DATABASE_URL=$DATABASE_URL
# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Download and install any required dependencies
RUN go mod download

# Build the Go app
RUN go build -o main .

# Expose port 8080 for incoming traffic
EXPOSE 8000

# Define the command to run the app when the container starts
CMD ["/app/main"]