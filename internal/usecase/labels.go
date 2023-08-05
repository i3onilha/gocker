package usecase

import "github.com/i3onilha/MESEnterpriseSmart/internal/entity"

type Repository interface {
	CreateLabel(dto *entity.LabelDTO) (*entity.LabelDTO, error)
	DeleteLabel(id int) error
	GetLabel(id int) (*entity.LabelDTO, error)
	GetLabelList() ([]*entity.LabelDTO, error)
	UpdateLabel(dto *entity.LabelUpdateDTO) (*entity.LabelDTO, error)
}

type Validator interface {
	ValidateDTO(dto *entity.LabelDTO) error
	ValidateUpdateDTO(dto *entity.LabelUpdateDTO) error
	ValidateID(id int) error
}

type labels struct {
	repository Repository
	validator  Validator
}

func NewLabels(r Repository, v Validator) *labels {
	return &labels{
		repository: r,
		validator:  v,
	}
}

func (l *labels) CreateLabel(dto *entity.LabelDTO) (*entity.LabelDTO, error) {
	err := l.validator.ValidateDTO(dto)
	if err != nil {
		return &entity.LabelDTO{}, err
	}
	return l.repository.CreateLabel(dto)
}

func (l *labels) DeleteLabel(id int) error {
	err := l.validator.ValidateID(id)
	if err != nil {
		return err
	}
	return l.repository.DeleteLabel(id)
}

func (l *labels) GetLabel(id int) (*entity.LabelDTO, error) {
	err := l.validator.ValidateID(id)
	if err != nil {
		return nil, err
	}
	return l.repository.GetLabel(id)
}

func (l *labels) GetLabelList() ([]*entity.LabelDTO, error) {
	return l.repository.GetLabelList()
}

func (l *labels) UpdateLabel(dto *entity.LabelUpdateDTO) (*entity.LabelDTO, error) {
	err := l.validator.ValidateUpdateDTO(dto)
	if err != nil {
		return &entity.LabelDTO{}, err
	}
	return l.repository.UpdateLabel(dto)
}
