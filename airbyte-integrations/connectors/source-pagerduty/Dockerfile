FROM golang:1.17-alpine3.14 AS builder

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

FROM alpine:3.14

LABEL io.airbyte.name=airbyte/source-pagerduty
LABEL io.airbyte.version=0.1.0

COPY --from=builder /go/bin/source-pagerduty /bin/source-pagerduty

ENV AIRBYTE_ENTRYPOINT="source-pagerduty"

ENTRYPOINT ["source-pagerduty"]
