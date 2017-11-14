# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.8

# Copy the local package files to the container's workspace.
COPY . $GOPATH/src/github.com/kwri/go-workflow

#  Build the go-workflow command inside the container.
WORKDIR $GOPATH/src/github.com/kwri/go-workflow

RUN go get .
RUN go build

# Document that the service listens on port 8080
EXPOSE 8001

# Run the go-workflow command by default when the container starts.
ENTRYPOINT ["go-workflow", "api"]
