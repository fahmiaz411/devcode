# start with base image
# FROM mysql:latest
# import data into container
# All scripts in docker-entrypoint-initdb.d/ are automatically executed during container startup
# COPY ./sql/*.sql /docker-entrypoint-initdb.d/

# Docker image use golang latest version @1.20.
FROM golang:1.20-alpine
# Specify that we now to execute any commands in this working directory.
WORKDIR /app
# prevent the re-installation of vendors at every change in the source code
COPY ./go.mod go.sum ./
RUN go mod download && go mod verify
# Copy everything from this project into the system directory of container.
COPY . .
# Compile and build binary file for our server.
RUN go build -o server ./app
# Start the server
CMD ["./server"]