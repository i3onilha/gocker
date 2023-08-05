package entity

import "time"

type LabelDTO struct {
	ID         int32
	Customer   string
	Family     string
	Model      string
	PartNumber string
	Station    string
	Label      string
	Author     string
	CreatedAt  time.Time
}

type LabelUpdateDTO struct {
	ID         int32
	Customer   string
	Family     string
	Model      string
	PartNumber string
	Station    string
	Label      string
	Author     string
}
