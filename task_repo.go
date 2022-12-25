package main

import (
	"context"
	"time"
)

type TaskRepo struct{}

var taskRepo = TaskRepo{}

func (TaskRepo) Create(ctx context.Context, task TaskRepo_CreateParams) *AppError {
	if task.CreatedAt == 0 {
		task.CreatedAt = time.Now().Unix()
	}

	if task.Status == "" {
		task.Status = TaskStatusPending
	}

	_, err := db.NamedExecContext(ctx, `
		INSERT INTO tasks (
			url,
			raw_url,
			file_name,
			file_size,
			connection_count,
			status,
			created_at
		) VALUES (
			:url,
			:raw_url,
			:file_name,
			:file_size,
			:connection_count,
			:status,
			:created_at
		)
	`, task)
	if err != nil {
		logger.Trace(err)
		return NewAppError()
	}

	return nil
}
