package monitoring

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	taskProcessed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "tasks_processed_total",
		Help: "Número total de tarefas processadas",
	})
	taskFailed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "tasks_failed_total",
		Help: "Número total de falhas nas tarefas",
	})
	taskSLA = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "task_sla_seconds",
		Help:    "SLA das tarefas em segundos",
		Buckets: prometheus.LinearBuckets(0, 10, 10), // Faixas de 10s até 100s
	})
)

func init() {
	prometheus.MustRegister(taskProcessed)
	prometheus.MustRegister(taskFailed)
	prometheus.MustRegister(taskSLA)
}

// SetupPrometheus inicia o servidor HTTP para expor as métricas
func SetupPrometheus() {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Fatal(http.ListenAndServe(":2112", nil)) // Porta para o Prometheus
	}()
}

// Funções para incrementar as métricas
func IncrementProcessedTasks() {
	taskProcessed.Inc()
}

func IncrementFailedTasks() {
	taskFailed.Inc()
}

// Função para observar o tempo de SLA
func ObserveTaskSLA(slaSeconds float64) {
	taskSLA.Observe(slaSeconds)
}
