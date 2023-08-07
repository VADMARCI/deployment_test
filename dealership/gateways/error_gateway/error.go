package error_gateway

import (
	"github.com/pepusz/go_redirect/models"
)

type Errors struct {
	Errors []models.Error `json:"errors"`
}

func (e *Errors) AddError(error models.Error) {
	e.Errors = append(e.Errors, error)
}

func (e *Errors) GetErrors() []models.Error {
	return e.Errors
}

func (e *Errors) HasError() bool {
	return len(e.Errors) > 0
}

func (e *Errors) Clear() {
	e.Errors = nil
}
