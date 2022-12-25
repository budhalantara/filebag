package main

const (
	TaskStatusPending  = "pending"
	TaskStatusStarted  = "started"
	TaskStatusError    = "error"
	TaskStatusFinished = "finished"
)

type (
	TaskRequest struct {
		Url             string `json:"url"`
		RawUrl          string `json:"raw_url"`
		FileName        string `json:"file_name"`
		ConnectionCount int    `json:"connection_count"`
	}

	TaskRepo_CreateParams struct {
		Url             string `db:"url"`
		RawUrl          string `db:"raw_url"`
		FileName        string `db:"file_name"`
		FileSize        uint64 `db:"file_size"`
		ConnectionCount int    `db:"connection_count"`
		Status          string `db:"status"`
		CreatedAt       int64  `db:"created_at"`
	}
)

func (req TaskRequest) Validate() *AppError {
	errors := []string{}

	if req.Url == "" {
		errors = append(errors, "url is a required field")
	}

	if req.RawUrl == "" {
		errors = append(errors, "raw_url is a required field")
	}

	if len(errors) > 0 {
		return NewAppError_BadRequest("ValidationError", errors)
	}

	return nil
}
