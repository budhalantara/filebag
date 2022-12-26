package task

import (
	"context"

	"github.com/budhalantara/filebag/pkg"
)

type TaskService struct{}

var taskService = TaskService{}

func (TaskService) Create(ctx context.Context, req Request) *pkg.AppError {
	fm, err := pkg.GetFileMetadata(req.Url)
	if err != nil {
		pkg.Log.Trace(err)
		return pkg.NewAppError()
	}

	if req.FileName == "" {
		req.FileName = fm.FileName
	}

	if req.ConnectionCount == 0 {
		req.ConnectionCount = 1
	}

	task := Repo_CreateParams{
		Url:             req.Url,
		RawUrl:          req.RawUrl,
		FileName:        req.FileName,
		FileSize:        fm.ContentLength,
		ConnectionCount: req.ConnectionCount,
	}

	ae := taskRepo.Create(ctx, task)
	if ae != nil {
		return ae
	}

	return nil
}

func (TaskService) GetAll(ctx context.Context) ([]Service_Task, *pkg.AppError) {
	res := []Service_Task{}
	tasks, ae := taskRepo.FindAll(ctx)
	if ae != nil {
		return res, ae
	}

	for _, task := range tasks {
		res = append(res, Service_Task(task))
	}

	return res, nil
}
