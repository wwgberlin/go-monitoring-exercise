FROM golang:latest
WORKDIR /go/src/github.com/wwgberlin/go-monitoring-exercise
RUN go get github.com/prometheus/client_golang/prometheus
RUN go get github.com/prometheus/client_golang/prometheus/promhttp
CMD go install github.com/wwgberlin/go-monitoring-exercise && /go/bin/go-monitoring-exercise
EXPOSE 8080