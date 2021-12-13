#Docker multi-stage builds

# ------------------------------------------------------------------------------
# Development image
# ------------------------------------------------------------------------------

#Builder stage
FROM golang:1.14.1-alpine3.11 as builder

# Force the go compiler to use modules
ENV GO111MODULE=on

# Indicates module paths that are not publicly available
ENV GOPRIVATE=*github.com/upload-file

# Update OS package and install Git
RUN apk update && apk add git openssh && apk add build-base

# Set working directory
WORKDIR /go/src/upload-file

# Setup github credential
ADD ./resources/docker/keys/id_rsa /root/.ssh/id_rsa
ADD ./resources/docker/keys/id_rsa.pub /root/.ssh/id_rsa.pub
RUN chmod 600 /root/.ssh/id_rsa

# make sure your domain is accepted
RUN touch /root/.ssh/known_hosts
RUN ssh-keyscan github.com >> /root/.ssh/known_hosts
RUN git config --global url."git@github.com:".insteadOf "https://github.com"

# Install wait-for
RUN wget https://raw.githubusercontent.com/eficode/wait-for/master/wait-for -O /usr/local/bin/wait-for &&\
    chmod +x /usr/local/bin/wait-for

# Copy Go dependency file
ADD go.mod go.mod
ADD go.sum go.sum
ADD app app
ADD Makefile Makefile

RUN go mod download

# Install Fresh for local development
RUN go get github.com/pilu/fresh

# Install go tool for convert go test output to junit xml
RUN go get -u github.com/jstemmer/go-junit-report
RUN go get github.com/axw/gocov/gocov
RUN go get github.com/AlekSi/gocov-xml

# Set Docker's entry point commands
RUN cd app/ && go build -o /go/bin/app.bin

# ------------------------------------------------------------------------------
# Deployment image
# ------------------------------------------------------------------------------

#App stage
FROM golang:1.14.1-alpine3.11

RUN apk add --no-cache tini tzdata
RUN addgroup -g 211000 -S appgroup && adduser -u 211000 -S appuser -G appgroup

# Set working directory
WORKDIR /app

#Get artifact from buider stage
RUN mkdir -p migrations

#Get artifact from buider stage
COPY --from=builder /go/bin/app.bin /app/app.bin
COPY --from=builder /go/src/upload-file/app/migrations/ migrations/
COPY --from=builder /go/pkg/mod/github.com/upload-file/ /go/pkg/mod/github.com/upload-file/

# Set Docker's entry point commands
RUN chown -R appuser:appgroup /go/pkg/mod/github.com/upload-file/ /app
USER appuser
ENTRYPOINT ["/sbin/tini","-sg","--","/app/app.bin"]