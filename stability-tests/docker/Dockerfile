ARG GORD_IMAGE
ARG GORDMINER_IMAGE

FROM ${GORD_IMAGE} as gord
FROM ${GORDMINER_IMAGE} as gorminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/gordanet/gord

WORKDIR /go/src/github.com/gordanet/gord

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=gord /app/ /app/
COPY --from=gorminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/gordanet/gord/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]
