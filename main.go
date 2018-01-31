package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// Task:

// In order to let prometheus start consuming metrics from our app
// we need to provide a /metrics/ endpoint.
// The package github.com/prometheus/client_golang/prometheus/promhttp
// already provides us with an http.Handler that we can use to do this.
// Add the package to the imports and add call to
// http.Handle("/metrics/", promhttp.Handler())
// just before the call to ListenAndServe.
// Checkout 127.0.0.1:8080/metrics.
func main() {
	reg := prometheus.DefaultRegisterer

	go background(reg)

	newDemoAPI(reg).register(http.DefaultServeMux)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type demoAPI struct {
	requestDurations *prometheus.HistogramVec
}

// Task:

// this function receives a prometheus Registerer.

// Add a prometheus Counter (not vector) to count the total runs
// using prometheus.NewCounter
// Name: background_total_runs
// register it to the registerer
// look at the grafana dashboard

// create a prometheus Counter vector using prometheus.NewCounterVec:
// with the name "background_task_results"
// with the labels ["results"].
// Where we encounter timeout, call vec.WithLabelValues with "timeout" and
// Where we are successful, call WithLabelValues with "success".
// Do not forget to register your vector to the prometheus registerer to
// see the results in the grafana dashboard

// Task (advanced):
// This function also has a go routine leak, fix the leak
// to see that the number of goroutines remains ~constant in the
// grafana dashboard

func background(registerer prometheus.Registerer) {
	for {
		ch := make(chan string)
		go func() {
			ch <- foo()
		}()
		select {
		case str := <-ch:
			log.Println(str)
		case <-time.After(time.Millisecond * 50):
			log.Println("gave up!")
		}
		time.Sleep(1 * time.Second)
	}

}

// Task:

// newDemoAPI is a constructor function that returns a pointer to a demoApi
// The demoApi requires a histogram vector.
// This is the histogram that will record our endpoints speed per endpoint.
// Instantiate a new histogram vector (similarly to how we did for counters
// with the name endpoint_duration and the labels ["endpoint"].
// check out prometheus documentation to decide to configure your buckets
// register the histogram to the registerer and return a demoApi object
// with the histogram.
func newDemoAPI(reg prometheus.Registerer) *demoAPI {
	return &demoAPI{}
}

// Task:

// The following function is missing some code.
// inside this function we declared another anonymous function called instr
// Wrap the call to fn(w,r) with timestamps to obtain the time it took to
// call fn (time.Now() and time.Since() are your friends).
// Then call a.requestDurations.WithLabelValues(endpoint).Observe(duration) to
// send the time it took to execute fn() to prometheus.
func (a demoAPI) register(mux *http.ServeMux) {
	instr := func(endpoint string, fn http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			fn(w, r)
		}
	}

	mux.HandleFunc("/foo/", instr("foo", serveFunc(foo)))
	mux.HandleFunc("/bar/", instr("bar", serveFunc(bar)))
}

func serveFunc(fn func() string) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(fn()))
	}
}

func foo() string {
	time.Sleep(time.Millisecond * time.Duration(rand.Int63n(100)))
	return "foo done"
}

func bar() string {
	time.Sleep(time.Millisecond * time.Duration(rand.Int63n(100)))
	return "bar done"
}
