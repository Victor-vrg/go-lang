package queue

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/Victor-vrg/go-lang/monitoring"
)

type Task struct {
	ID      int
	StartAt time.Time `json:"start_at"` // Data de início da tarefa
	EndAt   time.Time `json:"end_at"`   // Data de fim da tarefa
	Retries int
}

var (
	taskQueue  = make(chan Task, 100) // Tamanho da fila
	maxRetries = 3                    // Máximo de tentativas
)

// Inicia os workers que processam as tarefas na fila
func StartWorkerPool(numWorkers int) {
	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
}

// Adicionar uma nova tarefa à fila
func AddTask(task Task) {
	taskQueue <- task
}

// Worker que processa as tarefas
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range taskQueue {
		processTask(task, id)
	}
}

// Lógica de processamento com retries
func processTask(task Task, workerID int) {
	fmt.Printf("Worker %d processando tarefa ID %d\n", workerID, task.ID)

	// Calcular SLA
	duration := task.EndAt.Sub(task.StartAt)
	fmt.Printf("SLA da tarefa %d: %v\n", task.ID, duration)

	// Métrica para Prometheus
	monitoring.ObserveTaskSLA(duration.Seconds())

	// Simular uma falha aleatória
	if rand.Float32() < 0.3 {
		fmt.Printf("Worker %d falhou na tarefa ID %d\n", workerID, task.ID)
		monitoring.IncrementFailedTasks() // Incrementa falhas
		if task.Retries < maxRetries {
			task.Retries++
			fmt.Printf("Retrying task ID %d (Tentativa %d)\n", task.ID, task.Retries)
			time.Sleep(time.Second * time.Duration(task.Retries)) // Backoff exponencial
			taskQueue <- task                                     // Reenfileirar a tarefa
		} else {
			fmt.Printf("Máximo de tentativas atingido para a tarefa ID %d\n", task.ID)
		}
	} else {
		fmt.Printf("Worker %d completou a tarefa ID %d com sucesso\n", workerID, task.ID)
		monitoring.IncrementProcessedTasks() // Incrementa tarefas processadas
	}
}
