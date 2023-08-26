################################################################
# Builds a Docker image for the Picasso server.
################################################################
FROM golang:latest as build

# Set working directory for all commands
WORKDIR /picasso
# Initialize the Golang project
RUN go mod init main
# Copy source code
COPY main.go .
COPY lib ./lib
# Update all project's required packages
RUN go mod tidy

# Compile the source code
RUN CGO_ENABLED=0 GOOS=linux go build -o picasso

#--------------------------------------------------------------
FROM alpine:latest as service

# Set working directory for all commands
WORKDIR /picasso
# Copy artifacts from build
COPY --from=build /picasso/picasso .
# Copy all static resources
COPY images/ /picasso/images/
COPY lib/fonts/ /picasso/fonts/

# Entry point for container
CMD ["/picasso/picasso"]
