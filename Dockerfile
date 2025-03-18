# Specify the version of Go to use
# FROM golang:1.24
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
#

FROM golang:1.24 as builder

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

RUN apt-get -qq update && \
  apt-get -qq -y install upx

WORKDIR /src
COPY . .

RUN go build \
  -ldflags "-s -w -extldflags '-static'" \
  -o /bin/action \
  . \
  && strip /bin/action \
  && upx -q -9 /bin/action

RUN echo "nobody:x:65534:65534:Nobody:/:" > /etc_passwd

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc_passwd /etc/passwd
COPY --from=builder --chown=65534:0 /bin/action /action

USER nobody
ENTRYPOINT ["/action"]
