package task

import (
	"context"
	"time"

	"github.com/budhalantara/filebag/pkg"
)

type TaskRepo struct{}

var taskRepo = TaskRepo{}

func (TaskRepo) Create(ctx context.Context, task Repo_CreateParams) *pkg.AppError {
	if task.CreatedAt == 0 {
		task.CreatedAt = time.Now().Unix()
	}

	if task.Status == "" {
		task.Status = TaskStatusPending
	}

	_, err := pkg.DB.NamedExecContext(ctx, `
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
		pkg.Log.Trace(err)
		return pkg.NewAppError()
	}

	return nil
}

func (TaskRepo) FindAll(ctx context.Context) ([]Repo_Task, *pkg.AppError) {
	res := []Repo_Task{}
	err := pkg.DB.SelectContext(ctx, &res, `
		SELECT
			id,
			url,
			raw_url,
			file_name,
			file_size,
			connection_count,
			status,
			created_at
		FROM tasks
		ORDER BY id DESC
	`)
	if err != nil {
		pkg.Log.Trace(err)
		return res, pkg.NewAppError()
	}

	return res, nil
}
