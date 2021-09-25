FROM golang:1.17-alpine

LABEL org.opencontainers.image.source="https://github.com/ServerBoiOrg/ServerBoi-Status-Monitor-Container"

WORKDIR /status

COPY status-monitor/go.mod ./
COPY status-monitor/go.sum ./

RUN git clone https://github.com/ServerBoiOrg/ServerBoi-Lambdas-Go ./ServerBoi-Lambdas-Go

RUN go mod download

COPY status-monitor/*.go ./

RUN export CGO_ENABLED=0
RUN go build -o /serverboi-status-monitor

EXPOSE 7032/tcp

CMD [ "/serverboi-status-monitor" ]