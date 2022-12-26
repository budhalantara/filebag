package task

import "github.com/budhalantara/filebag/pkg"

const (
	TaskStatusPending  = "pending"
	TaskStatusStarted  = "started"
	TaskStatusError    = "error"
	TaskStatusFinished = "finished"
)

type (
	Request struct {
		Url             string `json:"url"`
		RawUrl          string `json:"raw_url"`
		FileName        string `json:"file_name"`
		ConnectionCount int    `json:"connection_count"`
	}

	Repo_CreateParams struct {
		Url             string `db:"url"`
		RawUrl          string `db:"raw_url"`
		FileName        string `db:"file_name"`
		FileSize        uint64 `db:"file_size"`
		ConnectionCount int    `db:"connection_count"`
		Status          string `db:"status"`
		CreatedAt       int64  `db:"created_at"`
	}

	Repo_Task struct {
		ID              int    `db:"id"`
		Url             string `db:"url"`
		RawUrl          string `db:"raw_url"`
		FileName        string `db:"file_name"`
		FileSize        uint64 `db:"file_size"`
		ConnectionCount int    `db:"connection_count"`
		Status          string `db:"status"`
		CreatedAt       int64  `db:"created_at"`
	}

	Service_Task struct {
		ID              int    `json:"id"`
		Url             string `json:"url"`
		RawUrl          string `json:"raw_url"`
		FileName        string `json:"file_name"`
		FileSize        uint64 `json:"file_size"`
		ConnectionCount int    `json:"connection_count"`
		Status          string `json:"status"`
		CreatedAt       int64  `json:"created_at"`
	}
)

func (req Request) Validate() *pkg.AppError {
	errors := []string{}

	if req.Url == "" {
		errors = append(errors, "url is a required field")
	}

	if req.RawUrl == "" {
		errors = append(errors, "raw_url is a required field")
	}

	if len(errors) > 0 {
		return pkg.NewAppError_BadRequest("ValidationError", errors)
	}

	return nil
}
