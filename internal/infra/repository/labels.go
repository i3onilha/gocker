package labels

import (
	"context"

	"github.com/i3onilha/MESEnterpriseSmart/internal/entity"
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

func (l *Labels) Create(dto *entity.LabelDTO) (*entity.LabelDTO, error) {
	ctx := context.Background()
	data := labels.CreateParams{}
	id, err := l.queries.CreateAndUpdate(ctx, l.dataSourceName, data)
	if err != nil {
		return &entity.LabelDTO{}, err
	}
	label, err := l.queries.GetByID(ctx, int32(id))
	if err != nil {
		return &entity.LabelDTO{}, err
	}
	result := entity.LabelDTO{
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

func (l *Labels) GetByID(id int) (*entity.LabelDTO, error) {
	label, err := l.queries.GetByID(context.Background(), int32(id))
	if err != nil {
		return nil, err
	}
	result := &entity.LabelDTO{
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

func (l *Labels) List() ([]*entity.LabelDTO, error) {
	arg := labels.ListParams{
		Limit:  10,
		Offset: 0,
	}
	list, err := l.queries.List(context.Background(), arg)
	if err != nil {
		return nil, err
	}
	result := make([]*entity.LabelDTO, len(list))
	for i, label := range list {
		result[i] = &entity.LabelDTO{
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

func (l *Labels) Update(dto *entity.LabelUpdateDTO) (*entity.LabelDTO, error) {
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
	result, err := l.queries.Update(ctx, arg)
	if err != nil {
		return &entity.LabelDTO{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return &entity.LabelDTO{}, err
	}
	label, err := l.queries.GetByID(ctx, int32(id))
	if err != nil {
		return &entity.LabelDTO{}, err
	}
	resultDTO := entity.LabelDTO{
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
