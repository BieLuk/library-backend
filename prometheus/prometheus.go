package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

// we create a new custom metric of type counter
var BookStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_get_book_status_count", // metric name
		Help: "Count of status returned by book.",
	},
	[]string{"book", "status"}, // labels
)

func init() {
	// we need to register the counter so prometheus can collect this metric
	prometheus.MustRegister(BookStatus)
}
