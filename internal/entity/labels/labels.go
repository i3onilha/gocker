package labels

import "time"

type CreateDTO struct {
	ID         int32  `validate:"gte=0"`
	Customer   string `validate:"required,max=32"`
	Model      string `validate:"required,max=16"`
	PartNumber string `validate:"required,max=16"`
	Station    string `validate:"required,max=32"`
	Dpi        int32  `validate:"required,gte=0,lte=600"`
	Label      string `validate:"required,min=6,max=65535"`
	Setup      string `validate:"required,min=6,max=65535"`
	SqlQueries string `validate:"required,min=6,max=65535"`
	Author     string `validate:"required"`
	CreatedAt  time.Time
}

type UpdateDTO struct {
	ID         int32  `validate:"gte=0"`
	Customer   string `validate:"required,max=32"`
	Model      string `validate:"required,max=16"`
	PartNumber string `validate:"required,max=16"`
	Station    string `validate:"required,max=32"`
	Dpi        int32  `validate:"required,gte=0,lte=600"`
	Label      string `validate:"required,min=6,max=65535"`
	Setup      string `validate:"required,min=6,max=65535"`
	SqlQueries string `validate:"required,min=6,max=65535"`
	Author     string `validate:"required"`
}
