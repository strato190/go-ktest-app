FROM golang:1.14 AS buildenv

ARG VERSION
ARG GIT_COMMIT

ENV VERSION=${VERSION}
ENV GIT_COMMIT=${GIT_COMMIT}

ENV GO111MODULE=on
ENV CGO_ENABLED=0

# Create a location in the container for the source code.
RUN mkdir -p /app

# Copy the module files first and then download the dependencies. If this
# doesn't change, we won't need to do this again in future builds.
COPY . /app/

WORKDIR /app
RUN go mod download
RUN go mod verify

RUN go build \
    -mod=readonly \
    -ldflags "-X main.Version=$VERSION -X main.GitCommit=$GIT_COMMIT -X 'main.BuildTime=$(date -u '+%Y-%m-%d %H:%M:%S')'" \
    -a -o /go/bin/app ./cmd/ktest

FROM scratch
COPY --from=buildenv /go/bin/app /go/bin/app
ENTRYPOINT ["/go/bin/app"]
