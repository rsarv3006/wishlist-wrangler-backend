package coresharedexceptions

import (
	"fmt"
)

type EntityNotFoundException struct {
	message string
}

func NewEntityNotFoundException() *EntityNotFoundException {
	return &EntityNotFoundException{
		message: "No Entry Found",
	}
}

func (e *EntityNotFoundException) Error() string {
	return fmt.Sprintf("EntityNotFoundException: %s", e.message)
}
