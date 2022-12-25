package main

import "context"

type TaskService struct{}

var taskService = TaskService{}

func (TaskService) Create(ctx context.Context, req TaskRequest) *AppError {
	fm, err := GetFileMetadata(req.Url)
	if err != nil {
		logger.Trace(err)
		return NewAppError()
	}

	if req.FileName == "" {
		req.FileName = fm.FileName
	}

	if req.ConnectionCount == 0 {
		req.ConnectionCount = 1
	}

	task := TaskRepo_CreateParams{
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
