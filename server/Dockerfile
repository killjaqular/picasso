################################################################
# Builds a Docker image for the Picasso server.
################################################################
# stage1, compile
FROM golang:latest as stage1

WORKDIR /picasso
# Copy source code
COPY main.go .
# Compile the source code
RUN CGO_ENABLED=0 GOOS=linux go build -o picasso

#--------------------------------------------------------------
# stage2, run the service
FROM alpine:latest as stage2

WORKDIR /picasso
# Copy binary
COPY --from=stage1 /picasso/picasso .
# Copy configuration files
COPY ./configs/server.yml /picasso/config.yml
# Entry point for container
CMD ["/picasso/picasso"]
