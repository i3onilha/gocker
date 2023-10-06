package labels

import (
	"context"
	"encoding/json"

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
	setupInput, err := json.Marshal(dto.Setup)
	if err != nil {
		return &entity.CreateDTO{}, err
	}
	data := labels.CreateParams{
		Customer:   dto.Customer,
		Model:      dto.Model,
		PartNumber: dto.PartNumber,
		Station:    dto.Station,
		Dpi:        dto.Dpi,
		Label:      dto.Label,
		Setup:      string(setupInput),
		SqlQueries: dto.SqlQueries,
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
	var setupOutput []entity.Setup
	err = json.Unmarshal([]byte(label.Setup), &setupOutput)
	if err != nil {
		return &entity.CreateDTO{}, err
	}
	result := entity.CreateDTO{
		ID:         label.ID,
		Customer:   label.Customer,
		Model:      label.Model,
		PartNumber: label.PartNumber,
		Station:    label.Station,
		Dpi:        label.Dpi,
		Label:      label.Label,
		Setup:      setupOutput,
		SqlQueries: label.SqlQueries,
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
	var setupOutput []entity.Setup
	err = json.Unmarshal([]byte(label.Setup), &setupOutput)
	if err != nil {
		return &entity.CreateDTO{}, err
	}
	result := &entity.CreateDTO{
		ID:         label.ID,
		Customer:   label.Customer,
		Model:      label.Model,
		PartNumber: label.PartNumber,
		Station:    label.Station,
		Dpi:        label.Dpi,
		Label:      label.Label,
		Setup:      setupOutput,
		SqlQueries: label.SqlQueries,
		Author:     label.Author,
		CreatedAt:  label.CreatedAt,
	}
	return result, nil
}

func (l *Labels) ListPaginate(limit, offset int) ([]*entity.CreateDTO, error) {
	arg := labels.ListPaginateParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	list, err := l.queries.ListPaginate(context.Background(), arg)
	if err != nil {
		return nil, err
	}
	result := make([]*entity.CreateDTO, len(list))
	for i, label := range list {
		var setupOutput []entity.Setup
		json.Unmarshal([]byte(label.Setup), &setupOutput)
		result[i] = &entity.CreateDTO{
			ID:         label.ID,
			Customer:   label.Customer,
			Model:      label.Model,
			PartNumber: label.PartNumber,
			Station:    label.Station,
			Dpi:        label.Dpi,
			Label:      label.Label,
			Setup:      setupOutput,
			SqlQueries: label.SqlQueries,
			Author:     label.Author,
			CreatedAt:  label.CreatedAt,
		}
	}
	return result, nil
}

func (l *Labels) ListByModel(model string) ([]*entity.CreateDTO, error) {
	list, err := l.queries.ListByModel(context.Background(), model)
	if err != nil {
		return nil, err
	}
	result := make([]*entity.CreateDTO, len(list))
	for i, label := range list {
		var setupOutput []entity.Setup
		json.Unmarshal([]byte(label.Setup), &setupOutput)
		result[i] = &entity.CreateDTO{
			ID:         label.ID,
			Customer:   label.Customer,
			Model:      label.Model,
			PartNumber: label.PartNumber,
			Station:    label.Station,
			Dpi:        label.Dpi,
			Label:      label.Label,
			Setup:      setupOutput,
			Author:     label.Author,
			SqlQueries: label.SqlQueries,
		}
	}
	return result, nil
}

func (l *Labels) ListByParts(partNumber string) ([]*entity.CreateDTO, error) {
	list, err := l.queries.ListByParts(context.Background(), partNumber)
	if err != nil {
		return nil, err
	}
	result := make([]*entity.CreateDTO, len(list))
	for i, label := range list {
		var setupOutput []entity.Setup
		json.Unmarshal([]byte(label.Setup), &setupOutput)
		result[i] = &entity.CreateDTO{
			ID:         label.ID,
			Customer:   label.Customer,
			Model:      label.Model,
			PartNumber: label.PartNumber,
			Station:    label.Station,
			Dpi:        label.Dpi,
			Label:      label.Label,
			Setup:      setupOutput,
			Author:     label.Author,
			SqlQueries: label.SqlQueries,
		}
	}
	return result, nil
}

func (l *Labels) ListZPLByPartsAndStationAndDpi(partNumber, station string, dpi int) ([]*entity.ZplDTO, error) {
	arg := labels.ListByPartsAndStationAndDpiParams{
		PartNumber: partNumber,
		Station:    station,
		Dpi:        int32(dpi),
	}
	list, err := l.queries.ListByPartsAndStationAndDpi(context.Background(), arg)
	if err != nil {
		return nil, err
	}
	result := make([]*entity.ZplDTO, len(list))
	for i, label := range list {
		var setupOutput []entity.Setup
		json.Unmarshal([]byte(label.Setup), &setupOutput)
		result[i] = &entity.ZplDTO{
			Label:      label.Label,
			SqlQueries: label.SqlQueries,
			Setup:      setupOutput,
		}
	}
	return result, nil
}

func (l *Labels) ListByModelAndStationAndDpi(model, station string, dpi int) ([]*entity.CreateDTO, error) {
	arg := labels.ListByModelAndStationAndDpiParams{
		Model:   model,
		Station: station,
		Dpi:     int32(dpi),
	}
	list, err := l.queries.ListByModelAndStationAndDpi(context.Background(), arg)
	if err != nil {
		return nil, err
	}
	result := make([]*entity.CreateDTO, len(list))
	for i, label := range list {
		var setupOutput []entity.Setup
		json.Unmarshal([]byte(label.Setup), &setupOutput)
		result[i] = &entity.CreateDTO{
			ID:         label.ID,
			Customer:   label.Customer,
			Model:      label.Model,
			PartNumber: label.PartNumber,
			Station:    label.Station,
			Dpi:        label.Dpi,
			Label:      label.Label,
			Setup:      setupOutput,
			Author:     label.Author,
			SqlQueries: label.SqlQueries,
		}
	}
	return result, nil
}

func (l *Labels) ListByPartsAndStationAndDpi(partNumber, station string, dpi int) ([]*entity.CreateDTO, error) {
	arg := labels.ListByPartsAndStationAndDpiParams{
		PartNumber: partNumber,
		Station:    station,
		Dpi:        int32(dpi),
	}
	list, err := l.queries.ListByPartsAndStationAndDpi(context.Background(), arg)
	if err != nil {
		return nil, err
	}
	result := make([]*entity.CreateDTO, len(list))
	for i, label := range list {
		var setupOutput []entity.Setup
		json.Unmarshal([]byte(label.Setup), &setupOutput)
		result[i] = &entity.CreateDTO{
			ID:         label.ID,
			Customer:   label.Customer,
			Model:      label.Model,
			PartNumber: label.PartNumber,
			Station:    label.Station,
			Dpi:        label.Dpi,
			Label:      label.Label,
			Setup:      setupOutput,
			Author:     label.Author,
			SqlQueries: label.SqlQueries,
		}
	}
	return result, nil
}

func (l *Labels) Update(dto *entity.UpdateDTO) (*entity.CreateDTO, error) {
	ctx := context.Background()
	setupInput, err := json.Marshal(dto.Setup)
	if err != nil {
		return &entity.CreateDTO{}, err
	}
	arg := labels.UpdateParams{
		ID:         dto.ID,
		Customer:   dto.Customer,
		Model:      dto.Model,
		PartNumber: dto.PartNumber,
		Station:    dto.Station,
		Dpi:        dto.Dpi,
		Label:      dto.Label,
		Setup:      string(setupInput),
		SqlQueries: dto.SqlQueries,
		Author:     dto.Author,
	}
	_, err = l.queries.Update(ctx, arg)
	if err != nil {
		return &entity.CreateDTO{}, err
	}
	label, err := l.queries.GetByID(ctx, dto.ID)
	if err != nil {
		return &entity.CreateDTO{}, err
	}
	var setupOutput []entity.Setup
	err = json.Unmarshal([]byte(label.Setup), &setupOutput)
	if err != nil {
		return &entity.CreateDTO{}, err
	}
	resultDTO := entity.CreateDTO{
		ID:         label.ID,
		Customer:   label.Customer,
		Model:      label.Model,
		PartNumber: label.PartNumber,
		Station:    label.Station,
		Dpi:        label.Dpi,
		Label:      label.Label,
		Setup:      setupOutput,
		SqlQueries: label.SqlQueries,
		Author:     label.Author,
		CreatedAt:  label.CreatedAt,
	}
	return &resultDTO, nil
}
