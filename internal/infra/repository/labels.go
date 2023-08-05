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

func NewLabels(queries *mysql.MySQL) *Labels {
	return &Labels{
		queries:        queries.Labels,
		dataSourceName: queries.GetDataSourceName(),
	}
}

func (l *Labels) CreateLabel(dto *entity.LabelDTO) (*entity.LabelDTO, error) {
	ctx := context.Background()
	data := labels.CreateLabelParams{}
	id, err := l.queries.CreateAndUpdateLabel(ctx, l.dataSourceName, data)
	if err != nil {
		return &entity.LabelDTO{}, err
	}
	label, err := l.queries.GetLabel(ctx, int32(id))
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

func (l *Labels) DeleteLabel(id int) error {
	return l.queries.DeleteLabel(context.Background(), int32(id))
}

func (l *Labels) GetLabel(id int) (*entity.LabelDTO, error) {
	label, err := l.queries.GetLabel(context.Background(), int32(id))
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

func (l *Labels) GetLabelList() ([]*entity.LabelDTO, error) {
	arg := labels.GetLabelListParams{
		Limit:  10,
		Offset: 0,
	}
	list, err := l.queries.GetLabelList(context.Background(), arg)
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

func (l *Labels) UpdateLabel(dto *entity.LabelUpdateDTO) (*entity.LabelDTO, error) {
	ctx := context.Background()
	arg := labels.UpdateLabelParams{
		ID:         dto.ID,
		Customer:   dto.Customer,
		Family:     dto.Family,
		Model:      dto.Model,
		PartNumber: dto.PartNumber,
		Station:    dto.Station,
		Label:      dto.Label,
		Author:     dto.Author,
	}
	result, err := l.queries.UpdateLabel(ctx, arg)
	if err != nil {
		return &entity.LabelDTO{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return &entity.LabelDTO{}, err
	}
	label, err := l.queries.GetLabel(ctx, int32(id))
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
