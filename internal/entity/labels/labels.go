package labels

import "time"

type CreateLabelDTO struct {
	Customer   string  `json:"customer"`
	Model      string  `json:"model"`
	PartNumber string  `json:"part_number"`
	Station    string  `json:"station"`
	Dpi        int32   `json:"dpi"`
	Label      string  `json:"label"`
	Setup      []Setup `json:"setup"`
	Author     string  `json:"author"`
}

type UpdateLabelDTO struct {
	ID         string  `json:"id"`
	Customer   string  `json:"customer"`
	Model      string  `json:"model"`
	PartNumber string  `json:"part_number"`
	Station    string  `json:"station"`
	Dpi        int32   `json:"dpi"`
	Label      string  `json:"label"`
	Setup      []Setup `json:"setup"`
	Author     string  `json:"author"`
}

type Setup struct {
	Variable   string `json:"variable"`
	ReportID   string `json:"report_id"`
	ReportName string `json:"report_name"`
	Start      string `json:"start"`
	X          string `json:"x"`
	LoopVar    bool   `json:"loop_var"`
}
type ZplDTO struct {
	Label      string  `validate:"required,min=6,max=65535" json:"label"`
	SqlQueries string  `validate:"required,min=6,max=65535" json:"sql_queries"`
	Setup      []Setup `json:"setup"`
}

type CreateDTO struct {
	ID         int32     `validate:"gte=0" json:"id"`
	Customer   string    `validate:"required,max=32" json:"customer"`
	Model      string    `validate:"required,max=16" json:"model"`
	PartNumber string    `validate:"max=16" json:"part_number"`
	Station    string    `validate:"required,max=32" json:"station"`
	Dpi        int32     `validate:"required,gte=0,lte=600" json:"dpi"`
	Label      string    `validate:"required,min=6,max=65535" json:"label"`
	Setup      []Setup   `json:"setup"`
	Author     string    `validate:"required" json:"author"`
	SqlQueries string    `validate:"required,min=6,max=65535" json:"-"`
	CreatedAt  time.Time `json:"created_at"`
}

type UpdateDTO struct {
	ID         int32  `validate:"gte=0"`
	Customer   string `validate:"required,max=32"`
	Model      string `validate:"required,max=16"`
	PartNumber string `validate:"max=16"`
	Station    string `validate:"required,max=32"`
	Dpi        int32  `validate:"required,gte=0,lte=600"`
	Label      string `validate:"required,min=6,max=65535"`
	Setup      []Setup
	Author     string `validate:"required"`
	SqlQueries string `validate:"required,min=6,max=65535"`
}
