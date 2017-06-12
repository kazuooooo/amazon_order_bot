# Use the official go docker image built on debian.
FROM golang:1.8.3

# Grab the source code and add it to the workspace.
ADD . /go/src/amazon_order_bot/revel

# Install revel and the revel CLI.
RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel
RUN go get github.com/sclevine/agouti
RUN go get -u github.com/nlopes/slack

# Use the revel CLI to start up our application.
ENTRYPOINT revel run amazon_order_bot/revel
