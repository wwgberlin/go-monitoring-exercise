# go-monitoring-exercise

To begin this exerice you will need:

* Prometheus:
    * If you have docker and docker-compose you can use the `docker` branch of this repository to complete the challenge.
    * On mac (with homebrew): `brew install prometheus`
    * Linux: `apt-get install prometheus`
    * Download for [linux and windows](https://prometheus.io/download/)
* Prometheus go client - run `go get github.com/prometheus/client_golang/prometheus`
* Prometheus http - run `go get github.com/prometheus/client_golang/prometheus/promhttp`

If you completed the installation successfuly you should be able to start prometheus with the yml file in the repository
`/path/to/prometheus --config.file="prometheus.yml"`



