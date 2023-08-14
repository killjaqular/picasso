################################################################
# Builds a Docker image for the Picasso server.
################################################################
FROM golang:latest as build

WORKDIR /picasso
# Copy source code
COPY *.go .
COPY go.mod .
# Compile the source code
RUN CGO_ENABLED=0 GOOS=linux go build -o picasso

#--------------------------------------------------------------
FROM alpine:latest as service

WORKDIR /picasso
# Create the directory to store images on local file disk instead of the database
RUN ["mkdir", "/picasso/images"]
# Copy artifacts from build
COPY --from=build /picasso/picasso .
# Entry point for container
CMD ["/picasso/picasso"]
