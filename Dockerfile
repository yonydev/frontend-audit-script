# Specify the version of Go to use
FROM --platform=linux/amd64 golang:1.24

# Copy all the files from the host into the container
WORKDIR /src
COPY . .
RUN mkdir -p /scripts
COPY scripts/postPRComment.js /home/runner/work/clipmx-frontend-libs/clipmx-frontend-libs/scripts/postPRComment.js

# Enable Go modules
ENV GO111MODULE=on

# Compile the action
RUN go build -o /bin/action

# Specify the container's entrypoint as the action
ENTRYPOINT ["/bin/action"]
