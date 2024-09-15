package controllers

import (
	"time"

	"github.com/Victor-vrg/go-lang/queue"
	"github.com/gofiber/fiber/v2"
)

// Estrutura para o payload da requisição
type TaskPayload struct {
	ID      int    `json:"id"`
	StartAt string `json:"start_at"` // A data de início será recebida como string (ISO 8601)
	EndAt   string `json:"end_at"`   // A data de fim também será recebida como string
}

// Adicionar uma nova tarefa à fila
func AddTask(c *fiber.Ctx) error {
	// Parse do body JSON
	payload := new(TaskPayload)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Erro ao parsear o payload: " + err.Error(),
		})
	}

	// Conversão das strings de data para time.Time
	startAt, err := time.Parse(time.RFC3339, payload.StartAt)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data de início inválida",
		})
	}

	endAt, err := time.Parse(time.RFC3339, payload.EndAt)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data de fim inválida",
		})
	}

	// Criar a nova tarefa
	task := queue.Task{
		ID:      payload.ID,
		StartAt: startAt,
		EndAt:   endAt,
		Retries: 0,
	}

	// Adicionar a tarefa à fila
	queue.AddTask(task)

	return c.JSON(fiber.Map{
		"message": "Tarefa adicionada à fila",
		"id":      payload.ID,
	})
}
