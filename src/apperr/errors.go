package apperr

import "fmt"

type Error struct {
	Code    int    `json:"code"`
	Name    string `json:"name"`
	Message string `json:"msg,omitempty"`
}

const (
	AppError           = "AppError"
	BadRequest         = "BadRequest"
	NoDataFound        = "NoDataFound"
	DuplicateRecord    = "DuplicateRecord"
	InvalidPermissions = "InvalidPermissions"
)

var ErrMap = map[string]*Error{
	AppError:           {Code: 101, Name: AppError},
	BadRequest:         {Code: 102, Name: BadRequest},
	NoDataFound:        {Code: 103, Name: NoDataFound},
	DuplicateRecord:    {Code: 104, Name: DuplicateRecord},
	InvalidPermissions: {Code: 105, Name: InvalidPermissions},
}

func (err Error) Error() string {
	return fmt.Sprintf("%d: %s - %s", err.Code, err.Name, err.Message)
}

func (err Error) AddMsg(msg string) Error {
	return Error{Code: err.Code, Name: err.Name, Message: msg}
}
