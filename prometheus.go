package inverter_exporter

import (
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"github.com/prologic/bitcask"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type PrometheusConfig struct {
	Prometheus []*PrometheusMetrics `yaml:"prometheus"`
}

type PrometheusMetrics struct {
	Type   string                `yaml:"type"`
	Key    string                `yaml:"key"`
	Config PrometheusMetricsOpts `yaml:"config"`
}

type PrometheusMetricsOpts struct {
	Namespace   string              `yaml:"namespace,omitempty"`
	Subsystem   string              `yaml:"subsystem,omitempty"`
	Name        string              `yaml:"name"`
	Help        string              `yaml:"help"`
	ConstLabels map[string]string   `yaml:"const_labels,omitempty"`
	Objectives  map[float64]float64 `yaml:"objectives,omitempty"`
	MaxAge      time.Duration       `yaml:"max_age,omitempty"`
	AgeBuckets  uint32              `yaml:"age_buckets,omitempty"`
	BufCap      uint32              `yaml:"buf_cap,omitempty"`
	Buckets     []float64           `yaml:"buckets,omitempty"`
}

type PrometheusMetricsType string

const (
	PROMETHEUS_METRICS_COUNTER   PrometheusMetricsType = "Counter"
	PROMETHEUS_METRICS_GAUGE     PrometheusMetricsType = "Gauge"
	PROMETHEUS_METRICS_SUMMARY   PrometheusMetricsType = "Summary"
	PROMETHEUS_METRICS_HISTOGRAM PrometheusMetricsType = "Histogram"
)

type PrometheusExporter interface{}

func NewPrometheusExporter(config *Config, exporterConfig *ExporterConfigPrometheusExporter, schema *PrometheusConfig, sensors *bitcask.Bitcask) (PrometheusExporter, error) {
	promMetrics := map[string]interface{}{}
	promMetricsType := map[string]PrometheusMetricsType{}

	pe := &prometheusexporter{
		config:          config,
		exporterConfig:  exporterConfig,
		schema:          schema,
		sensors:         *sensors,
		promMetrics:     promMetrics,
		promMetricsType: promMetricsType,
	}

	log.Infof(" - Host: %s", exporterConfig.Host)
	log.Infof(" - Port: %d", exporterConfig.Port)
	log.Infof(" - Path: %s", exporterConfig.Path)

	if err := pe.Run(); err != nil {
		return nil, err
	}

	return pe, nil
}

type prometheusexporter struct {
	config          *Config
	exporterConfig  *ExporterConfigPrometheusExporter
	schema          *PrometheusConfig
	sensors         bitcask.Bitcask
	promMetrics     map[string]interface{}
	promMetricsType map[string]PrometheusMetricsType
}

func (p *prometheusexporter) metricsHandler() {
	go func() {
		for range time.NewTicker(time.Second * 10).C {
			for key, metrics := range p.promMetrics {
				value, err := p.sensors.Get([]byte(key))
				if err != nil {
					log.Warnf("unable to get sensor value: %v", err)
				}
				switch p.promMetricsType[key] {
				case PROMETHEUS_METRICS_GAUGE:
					metrics.(prometheus.Gauge).Set(gconv.Float64(string(value)))
				case PROMETHEUS_METRICS_COUNTER:
					metrics.(prometheus.Counter).Add(gconv.Float64(string(value)))
				default:
					log.Fatalf("unsupported metrics type: %s", p.promMetricsType[key])
				}
			}
		}
	}()
}

func (p *prometheusexporter) Run() error {
	for _, metrics := range p.schema.Prometheus {
		switch PrometheusMetricsType(metrics.Type) {
		case PROMETHEUS_METRICS_GAUGE:
			p.promMetrics[metrics.Key] = newGauge(metrics.Config)
		case PROMETHEUS_METRICS_COUNTER:
			p.promMetrics[metrics.Key] = newCounter(metrics.Config)
		case PROMETHEUS_METRICS_HISTOGRAM:
			p.promMetrics[metrics.Key] = newHistogram(metrics.Config)
		case PROMETHEUS_METRICS_SUMMARY:
			p.promMetrics[metrics.Key] = newSummary(metrics.Config)
		default:
			return fmt.Errorf("unsupported metrics type: %s", metrics.Type)
		}
		p.promMetricsType[metrics.Key] = PrometheusMetricsType(metrics.Type)
	}

	p.metricsHandler()

	hostPort := fmt.Sprintf("%s:%d", p.exporterConfig.Host, p.exporterConfig.Port)
	pathPattern := p.exporterConfig.Path

	http.Handle(pathPattern, promhttp.Handler())
	return http.ListenAndServe(hostPort, nil)
}

func newGauge(opts PrometheusMetricsOpts) prometheus.Gauge {
	return promauto.NewGauge(prometheus.GaugeOpts{
		Namespace:   opts.Namespace,
		Subsystem:   opts.Subsystem,
		Name:        opts.Name,
		Help:        opts.Help,
		ConstLabels: opts.ConstLabels,
	})
}

func newCounter(opts PrometheusMetricsOpts) prometheus.Counter {
	return promauto.NewCounter(prometheus.CounterOpts{
		Namespace:   opts.Namespace,
		Subsystem:   opts.Subsystem,
		Name:        opts.Name,
		Help:        opts.Help,
		ConstLabels: opts.ConstLabels,
	})
}

func newSummary(opts PrometheusMetricsOpts) prometheus.Summary {
	return promauto.NewSummary(prometheus.SummaryOpts{
		Namespace:   opts.Namespace,
		Subsystem:   opts.Subsystem,
		Name:        opts.Name,
		Help:        opts.Help,
		ConstLabels: opts.ConstLabels,
		Objectives:  opts.Objectives,
		MaxAge:      opts.MaxAge,
		AgeBuckets:  opts.AgeBuckets,
		BufCap:      opts.BufCap,
	})
}

func newHistogram(opts PrometheusMetricsOpts) prometheus.Histogram {
	return promauto.NewHistogram(prometheus.HistogramOpts{
		Namespace:   opts.Namespace,
		Subsystem:   opts.Subsystem,
		Name:        opts.Name,
		Help:        opts.Help,
		ConstLabels: opts.ConstLabels,
		Buckets:     opts.Buckets,
	})
}
