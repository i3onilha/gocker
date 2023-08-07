package labels

import (
	"context"

	entity "github.com/i3onilha/MESEnterpriseSmart/internal/entity/labels"
	"github.com/i3onilha/MESEnterpriseSmart/internal/infra/mysql"
	"github.com/i3onilha/MESEnterpriseSmart/internal/infra/mysql/labels"
)

type Labels struct {
	queries        *labels.Queries
	dataSourceName string
}

func New(queries *mysql.MySQL) *Labels {
	return &Labels{
		queries:        queries.Labels,
		dataSourceName: queries.GetDataSourceName(),
	}
}

func (l *Labels) Create(dto *entity.CreateDTO) (*entity.CreateDTO, error) {
	ctx := context.Background()
	data := labels.CreateParams{
		Customer:   dto.Customer,
		Family:     dto.Family,
		Model:      dto.Model,
		PartNumber: dto.PartNumber,
		Station:    dto.Station,
		Label:      dto.Label,
		Author:     dto.Author,
	}
	id, err := l.queries.CreateAndUpdate(ctx, l.dataSourceName, data)
	if err != nil {
		return &entity.CreateDTO{}, err
	}
	label, err := l.queries.GetByID(ctx, int32(id))
	if err != nil {
		return &entity.CreateDTO{}, err
	}
	result := entity.CreateDTO{
		ID:         label.ID,
		Customer:   label.Customer,
		Family:     label.Family,
		Model:      label.Model,
		PartNumber: label.PartNumber,
		Station:    label.Station,
		Label:      label.Label,
		Author:     label.Author,
		CreatedAt:  label.CreatedAt,
	}
	return &result, nil
}

func (l *Labels) DeleteByID(id int) error {
	return l.queries.DeleteByID(context.Background(), int32(id))
}

func (l *Labels) GetByID(id int) (*entity.CreateDTO, error) {
	label, err := l.queries.GetByID(context.Background(), int32(id))
	if err != nil {
		return nil, err
	}
	result := &entity.CreateDTO{
		ID:         label.ID,
		Customer:   label.Customer,
		Family:     label.Family,
		Model:      label.Model,
		PartNumber: label.PartNumber,
		Station:    label.Station,
		Label:      label.Label,
		Author:     label.Author,
		CreatedAt:  label.CreatedAt,
	}
	return result, nil
}

func (l *Labels) List(limit, offset int) ([]*entity.CreateDTO, error) {
	arg := labels.ListParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	list, err := l.queries.List(context.Background(), arg)
	if err != nil {
		return nil, err
	}
	result := make([]*entity.CreateDTO, len(list))
	for i, label := range list {
		result[i] = &entity.CreateDTO{
			ID:         label.ID,
			Customer:   label.Customer,
			Family:     label.Family,
			Model:      label.Model,
			PartNumber: label.PartNumber,
			Station:    label.Station,
			Label:      label.Label,
			Author:     label.Author,
			CreatedAt:  label.CreatedAt,
		}
	}
	return result, nil
}

func (l *Labels) Update(dto *entity.UpdateDTO) (*entity.CreateDTO, error) {
	ctx := context.Background()
	arg := labels.UpdateParams{
		ID:         dto.ID,
		Customer:   dto.Customer,
		Family:     dto.Family,
		Model:      dto.Model,
		PartNumber: dto.PartNumber,
		Station:    dto.Station,
		Label:      dto.Label,
		Author:     dto.Author,
	}
	_, err := l.queries.Update(ctx, arg)
	if err != nil {
		return &entity.CreateDTO{}, err
	}
	label, err := l.queries.GetByID(ctx, dto.ID)
	if err != nil {
		return &entity.CreateDTO{}, err
	}
	resultDTO := entity.CreateDTO{
		ID:         label.ID,
		Customer:   label.Customer,
		Family:     label.Family,
		Model:      label.Model,
		PartNumber: label.PartNumber,
		Station:    label.Station,
		Label:      label.Label,
		Author:     label.Author,
		CreatedAt:  label.CreatedAt,
	}
	return &resultDTO, nil
}
