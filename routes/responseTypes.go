package routes

import (

)

type CommitValidationErrorResponse struct {
	Status string `json:"status"`
	Errors []string `json:"errors"`
}