# Specify the version of Go to use
FROM golang:1.24

# Copy all the files from the host into the container
WORKDIR /src
COPY . .

# Enable Go modules
ENV GO111MODULE=on

# Compile the action
RUN go build -o /bin/action

# Install Node.js (required to run the JavaScript script)
RUN apt-get update && apt-get install -y nodejs npm

# Specify the container's entrypoint as the action
ENTRYPOINT ["/bin/action"]

# Set the entrypoint to run the Go action first, then the JavaScript script
CMD ["/bin/action"] && node /src/scripts/post-pr-comment.js
