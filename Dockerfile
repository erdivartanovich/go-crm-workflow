# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.8

# Copy the local package files to the container's workspace.
ADD . $GOPATH/src/github.com/kwri/go-workflow

#  Build the go-workflow command inside the container.
WORKDIR $GOPATH/src/github.com/kwri/go-workflow

RUN go get -v . && go build


CMD ["go-workflow", "api"]

RUN echo 'Exposing client port for Workflow 8001'
EXPOSE 8001
