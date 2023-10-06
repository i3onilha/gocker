package usecase

import entity "github.com/i3onilha/MESEnterpriseSmart/internal/entity/labels"

type Repository interface {
	Create(dto *entity.CreateDTO) (*entity.CreateDTO, error)
	DeleteByID(id int) error
	GetByID(id int) (*entity.CreateDTO, error)
	ListPaginate(limit, offset int) ([]*entity.CreateDTO, error)
	ListByModelAndStationAndDpi(model, station string, dpi int) ([]*entity.CreateDTO, error)
	ListByPartsAndStationAndDpi(partNumber, station string, dpi int) ([]*entity.CreateDTO, error)
	ListZPLByPartsAndStationAndDpi(partNumber, station string, dpi int) ([]*entity.ZplDTO, error)
	ListByModel(partNumber string) ([]*entity.CreateDTO, error)
	ListByParts(partNumber string) ([]*entity.CreateDTO, error)
	Update(dto *entity.UpdateDTO) (*entity.CreateDTO, error)
}

type Validator interface {
	ValidateDataCreate(dto *entity.CreateDTO) error
	ValidateDataUpdate(dto *entity.UpdateDTO) error
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

func (l *labels) Create(dto *entity.CreateDTO) (*entity.CreateDTO, error) {
	err := l.validator.ValidateDataCreate(dto)
	if err != nil {
		return &entity.CreateDTO{}, err
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

func (l *labels) GetByID(id int) (*entity.CreateDTO, error) {
	err := l.validator.ValidateID(id)
	if err != nil {
		return nil, err
	}
	return l.repository.GetByID(id)
}

func (l *labels) List(limit int, offset int) ([]*entity.CreateDTO, error) {
	return l.repository.ListPaginate(limit, offset)
}

func (l *labels) ListByModel(model string) ([]*entity.CreateDTO, error) {
	return l.repository.ListByModel(model)
}

func (l *labels) ListByParts(partNumber string) ([]*entity.CreateDTO, error) {
	return l.repository.ListByParts(partNumber)
}

func (l *labels) ListByModelAndStationAndDpi(model, station string, dpi int) ([]*entity.CreateDTO, error) {
	return l.repository.ListByModelAndStationAndDpi(model, station, dpi)
}

func (l *labels) ListByPartsAndStationAndDpi(partNumber, station string, dpi int) ([]*entity.CreateDTO, error) {
	return l.repository.ListByPartsAndStationAndDpi(partNumber, station, dpi)
}

func (l *labels) ListZPLByPartsAndStationAndDpi(partNumber, station string, dpi int) ([]*entity.ZplDTO, error) {
	return l.repository.ListZPLByPartsAndStationAndDpi(partNumber, station, dpi)
}

func (l *labels) Update(dto *entity.UpdateDTO) (*entity.CreateDTO, error) {
	err := l.validator.ValidateDataUpdate(dto)
	if err != nil {
		return &entity.CreateDTO{}, err
	}
	return l.repository.Update(dto)
}
