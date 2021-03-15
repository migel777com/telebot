FROM golang

ENV GO111MODULE=on

RUN mkdir /monjjubot

COPY . /monjjubot

ENV CGO_ENABLED=0
WORKDIR /monjjubot/auth/auth_server
RUN go mod vendor
RUN go build -o main .


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR ./
COPY --from=0 /monjjubot /monjjubot

CMD ["/monjjubot/auth/auth_server/main"]