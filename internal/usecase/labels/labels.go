package usecase

import (
	entity "github.com/i3onilha/MESEnterpriseSmart/internal/entity/labels"
)

type Repository interface {
	Create(dto *entity.CreateDTO) (*entity.CreateDTO, error)
	DeleteByID(id int) error
	GetOracleDataSource(customer string) (string, error)
	GetByID(id int) (*entity.CreateDTO, error)
	ListPaginate(limit, offset int) ([]*entity.CreateDTO, error)
	ListByModelAndStationAndDpi(customer, model, station string, dpi int) ([]*entity.CreateDTO, error)
	ListByPartsAndStationAndDpi(customer, partNumber, station string, dpi int) ([]*entity.CreateDTO, error)
	ListZPLByModelAndStationAndDpi(customer, partNumber, station string, dpi int) ([]*entity.ZplDTO, error)
	ListZPLByPartsAndStationAndDpi(customer, partNumber, station string, dpi int) ([]*entity.ZplDTO, error)
	ListByModel(customer, partNumber string) ([]*entity.CreateDTO, error)
	ListByParts(customer, partNumber string) ([]*entity.CreateDTO, error)
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

func (l *labels) GetOracleDataSource(customer string) (string, error) {
	return l.repository.GetOracleDataSource(customer)
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

func (l *labels) ListByModel(customer, model string) ([]*entity.CreateDTO, error) {
	return l.repository.ListByModel(customer, model)
}

func (l *labels) ListByParts(customer, partNumber string) ([]*entity.CreateDTO, error) {
	return l.repository.ListByParts(customer, partNumber)
}

func (l *labels) ListByModelAndStationAndDpi(customer, model, station string, dpi int) ([]*entity.CreateDTO, error) {
	return l.repository.ListByModelAndStationAndDpi(customer, model, station, dpi)
}

func (l *labels) ListByPartsAndStationAndDpi(customer, partNumber, station string, dpi int) ([]*entity.CreateDTO, error) {
	return l.repository.ListByPartsAndStationAndDpi(customer, partNumber, station, dpi)
}

func (l *labels) ListZPLByModelAndStationAndDpi(customer, model, station string, dpi int) ([]*entity.ZplDTO, error) {
	return l.repository.ListZPLByModelAndStationAndDpi(customer, model, station, dpi)
}

func (l *labels) ListZPLByPartsAndStationAndDpi(customer, partNumber, station string, dpi int) ([]*entity.ZplDTO, error) {
	return l.repository.ListZPLByPartsAndStationAndDpi(customer, partNumber, station, dpi)
}

func (l *labels) Update(dto *entity.UpdateDTO) (*entity.CreateDTO, error) {
	err := l.validator.ValidateDataUpdate(dto)
	if err != nil {
		return &entity.CreateDTO{}, err
	}
	return l.repository.Update(dto)
}
