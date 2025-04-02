# # Specify the version of Go to use
#
# # Copy all the files from the host into the container
# WORKDIR /src
# COPY . .
#
# # Enable Go modules
# ENV GO111MODULE=on
#
# # Compile the action
# RUN go build -o /bin/action
#
# # Specify the container's entrypoint as the action
# ENTRYPOINT ["/bin/action"]
# your-repo/Dockerfile
FROM --platform=linux/amd64 golang:1.24

WORKDIR /src

COPY . .

RUN go build -o /bin/action .

ENTRYPOINT ["/bin/action"]
