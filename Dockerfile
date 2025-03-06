FROM golang:alpine AS builder

WORKDIR /src

ADD go.mod .

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build  -ldflags "-s -w" -a -installsuffix cgo -o=micro ./main.go

FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates

WORKDIR /src

COPY --from=builder /src/micro ./micro
COPY --from=builder /src/config ./config
COPY --from=builder /src/VERSION ./VERSION

ENV GATE_IN_TITLE Yo-ho-ho

CMD ["./micro"]