package labels

import "time"

type CreateDTO struct {
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

type UpdateDTO struct {
	ID         int32
	Customer   string
	Family     string
	Model      string
	PartNumber string
	Station    string
	Label      string
	Author     string
}
