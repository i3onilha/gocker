package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	entity "github.com/i3onilha/MESEnterpriseSmart/internal/entity/labels"
)

type Labels struct{}

func (l *Labels) ValidateDataCreate(dto *entity.CreateDTO) error {
	return validator.New().Struct(dto)
}

func (l *Labels) ValidateDataUpdate(dto *entity.UpdateDTO) error {
	return validator.New().Struct(dto)
}

func (l *Labels) ValidateID(id int) error {
	if id <= 0 {
		return fmt.Errorf("id must be greater than 0")
	}
	return nil
}
