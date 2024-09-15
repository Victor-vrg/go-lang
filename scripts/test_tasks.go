package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Victor-vrg/go-lang/monitoring"
	"github.com/Victor-vrg/go-lang/queue"
)

func main() {
	// Inicializar Prometheus
	monitoring.SetupPrometheus()

	// Inicializar pool de workers
	go queue.StartWorkerPool(5)

	// Gerar tarefas de teste
	for i := 1; i <= 10; i++ {
		start := time.Now()
		// Randomizar o fim da tarefa para simular diferentes SLAs
		end := start.Add(time.Duration(rand.Intn(100)) * time.Second)

		task := queue.Task{
			ID:      i,
			StartAt: start,
			EndAt:   end,
			Retries: 0,
		}

		fmt.Printf("Adicionando tarefa ID %d com início em %s e fim em %s\n", task.ID, start, end)
		queue.AddTask(task)

		// Espera para criar uma nova tarefa
		time.Sleep(2 * time.Second)
	}

	// Manter o programa rodando para permitir que Prometheus colete métricas
	select {}
}
