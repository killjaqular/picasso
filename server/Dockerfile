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
# Copy artifacts from build
COPY --from=build /picasso/picasso .
# Entry point for container
CMD ["/picasso/picasso"]
