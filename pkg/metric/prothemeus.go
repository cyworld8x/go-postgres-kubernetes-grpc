package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

const (
	// PROMETHEUS_PUSHGATEWAY pushgateway url
	PROMETHEUS_PUSHGATEWAY = "http://localhost:9091"
)

type Option func(*metricService)

// metricService implements Service interface
type metricService struct {
	pHistogram           *prometheus.HistogramVec
	httpRequestHistogram *prometheus.HistogramVec
	pushGatewayURL       string
}

// NewPrometheusService create a new prometheus service
func NewPrometheusService() (*metricService, error) {
	cli := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "pushgateway",
		Name:      "cmd_duration_seconds",
		Help:      "CLI application execution in seconds",
		Buckets:   prometheus.DefBuckets,
	}, []string{"name"})
	http := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "http",
		Name:      "request_duration_seconds",
		Help:      "The latency of the HTTP requests.",
		Buckets:   prometheus.DefBuckets,
	}, []string{"handler", "method", "code"})

	s := &metricService{
		pHistogram:           cli,
		httpRequestHistogram: http,
	}
	err := prometheus.Register(s.pHistogram)
	if err != nil && err.Error() != "duplicate metrics collector registration attempted" {
		return nil, err
	}
	err = prometheus.Register(s.httpRequestHistogram)
	if err != nil && err.Error() != "duplicate metrics collector registration attempted" {
		return nil, err
	}
	return s, nil
}

func WithPushGatewayURL(url string) Option {
	return func(s *metricService) {
		s.pushGatewayURL = url
	}
}

func (s *metricService) Configure(opts ...Option) metricService {
	for _, opt := range opts {
		opt(s)
	}
	return *s
}

func (s *metricService) Close() {
	prometheus.Unregister(s.pHistogram)
	prometheus.Unregister(s.httpRequestHistogram)
}

// SaveCLI send metrics to server
func (s *metricService) SaveCLI(c *CLI) error {
	gatewayURL := s.pushGatewayURL
	if gatewayURL == "" {
		gatewayURL = PROMETHEUS_PUSHGATEWAY
	}
	s.pHistogram.WithLabelValues(c.Name).Observe(c.Duration)
	return push.New(gatewayURL, "cmd_job").Collector(s.pHistogram).Push()
}

// SaveHTTP send metrics to server
func (s *metricService) SaveHTTP(h *HTTP) {
	s.httpRequestHistogram.WithLabelValues(h.Handler, h.Method, h.StatusCode).Observe(h.Duration)
}
