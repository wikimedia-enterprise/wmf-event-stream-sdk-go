package eventstream

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	TotalErrorMetric    string = "eventstream_total_errors"
	TotalEventsMetric   string = "eventstream_total_events"
	AverageEventsMetric string = "eventstream_average_events_per_second"

	MetricLabelStream string = "stream"

	SeverityLabelValueHigh   string = "high"
	SeverityLabelValueMedium string = "medium"
	SeverityLabelValueLow    string = "low"
)

type Metrics struct {
	enabled bool
	Opts    map[string]any
}

func (m *Metrics) SetConstantLabelValues(stream string) {

}

func (m *Metrics) Enable() {
	m.enabled = true

	m.Opts = make(map[string]any)

	m.Opts[TotalErrorMetric] = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: TotalErrorMetric,
		Help: "Total number of errors",
	}, []string{MetricLabelStream, "severity", "action_required", "data_loss"})

	m.Opts[TotalEventsMetric] = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: TotalEventsMetric,
		Help: "Total number of events",
	}, []string{MetricLabelStream})

	m.Opts[AverageEventsMetric] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: AverageEventsMetric,
		Help: "Average number of events per second.",
	}, []string{MetricLabelStream})
}

func (m *Metrics) Disable() {
	m.enabled = false

	for k := range m.Opts {
		delete(m.Opts, k)
	}
}

func (m *Metrics) IsEnabled() bool {
	return m.enabled
}

func (m *Metrics) IncTotalErrors(stream string, severity string, actionRequired string, dataLost string) {
	if m.enabled {
		v, _ := m.Opts[TotalErrorMetric].(prometheus.CounterVec)
		v.WithLabelValues(stream, severity, actionRequired).Inc()
	}
}

func (m *Metrics) IncTotalEvents(stream string) {
	if m.enabled {
		v, _ := m.Opts[TotalEventsMetric].(prometheus.CounterVec)
		v.WithLabelValues(stream).Inc()
	}
}

func (m *Metrics) IncAverageEvents(stream string) {
	if m.enabled {
		v, _ := m.Opts[AverageEventsMetric].(prometheus.GaugeVec)
		v.WithLabelValues(stream).Inc()
	}
}
