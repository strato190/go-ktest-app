FROM golang:1.14 AS buildenv

ARG VERSION
ARG GIT_COMMIT

ENV VERSION=${VERSION}
ENV GIT_COMMIT=${GIT_COMMIT}
ENV GO111MODULE=on
ENV CGO_ENABLED=0

RUN mkdir -p /app

COPY . /app/

WORKDIR /app
RUN go mod download
RUN go mod verify

RUN go build \
    -mod=readonly \
    -ldflags "-X main.Version=$VERSION -X main.GitCommit=$GIT_COMMIT -X 'main.BuildTime=$(date -u '+%Y-%m-%d %H:%M:%S')'" \
    -a -o /go/bin/app .

FROM scratch
ENV ENV PRODUCTION
COPY --from=buildenv /go/bin/app /go/bin/app
EXPOSE 8080 9781
ENTRYPOINT ["/go/bin/app"]
