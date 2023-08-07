package usecase

import "github.com/i3onilha/MESEnterpriseSmart/internal/entity"

type Repository interface {
	Create(dto *entity.LabelDTO) (*entity.LabelDTO, error)
	DeleteByID(id int) error
	GetByID(id int) (*entity.LabelDTO, error)
	List() ([]*entity.LabelDTO, error)
	Update(dto *entity.LabelUpdateDTO) (*entity.LabelDTO, error)
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

func New(r Repository, v Validator) *labels {
	return &labels{
		repository: r,
		validator:  v,
	}
}

func (l *labels) Create(dto *entity.LabelDTO) (*entity.LabelDTO, error) {
	err := l.validator.ValidateDTO(dto)
	if err != nil {
		return &entity.LabelDTO{}, err
	}
	return l.repository.Create(dto)
}

func (l *labels) DeleteByID(id int) error {
	err := l.validator.ValidateID(id)
	if err != nil {
		return err
	}
	return l.repository.DeleteByID(id)
}

func (l *labels) GetByID(id int) (*entity.LabelDTO, error) {
	err := l.validator.ValidateID(id)
	if err != nil {
		return nil, err
	}
	return l.repository.GetByID(id)
}

func (l *labels) List() ([]*entity.LabelDTO, error) {
	return l.repository.List()
}

func (l *labels) Update(dto *entity.LabelUpdateDTO) (*entity.LabelDTO, error) {
	err := l.validator.ValidateUpdateDTO(dto)
	if err != nil {
		return &entity.LabelDTO{}, err
	}
	return l.repository.Update(dto)
}
