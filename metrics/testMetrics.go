package metrics

type TestMetrics struct {
}

func NewTestMetrics() Metrics {
	metrics := &TestMetrics{}
	return metrics
}

func (p *TestMetrics) RequestReceivedDetail(path string, method string, code int, elapsedDuration float64) {
}
