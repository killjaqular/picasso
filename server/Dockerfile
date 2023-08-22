################################################################
# Builds a Docker image for the Picasso server.
################################################################
FROM golang:latest as build

WORKDIR /picasso
# Copy source code
COPY main.go .
COPY go.mod .
COPY lib ./lib
# Install go required packages
RUN go get github.com/lib/pq
RUN go get gopkg.in/yaml.v2

# Compile the source code
RUN CGO_ENABLED=0 GOOS=linux go build -o picasso

#--------------------------------------------------------------
FROM alpine:latest as service

WORKDIR /picasso
# Copy artifacts from build
COPY --from=build /picasso/picasso .
COPY images/ /picasso/images/
COPY lib/fonts/ /picasso/fonts/

# Entry point for container
CMD ["/picasso/picasso"]
