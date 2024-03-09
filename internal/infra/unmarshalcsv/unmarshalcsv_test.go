package unmarshalcsv_test

import (
	"testing"

	"github.com/i3onilha/MESEnterpriseSmart/internal/infra/unmarshalcsv"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshalCSVSemiColon(t *testing.T) {
	const txt = `palete;masterbox;serial_number;part_number
	LK24025CRR900014;LK24025CER700010;12459240000085;254034187_D01
	LK24025CRR900014;LK24025CER700010;12459240000086;254034187_D01
	LK24025CRR900014;LK24025CER700009;12459240000083;254034187_D01
	LK24025CRR900014;LK24025CER700009;12459240000084;254034187_D01`
	buf := []byte(txt)
	res, err := unmarshalcsv.UnmarshalCSV(buf, ';')
	assert.Nil(t, err)
	assert.Equal(t, 4, len(res))
}

func TestUnmarshalCSVComa(t *testing.T) {
	const txt = `palete,masterbox,serial_number,part_number
	LK24025CRR900014,LK24025CER700010,12459240000085,254034187_D01
	LK24025CRR900014,LK24025CER700010,12459240000086,254034187_D01
	LK24025CRR900014,LK24025CER700009,12459240000083,254034187_D01
	LK24025CRR900014,LK24025CER700009,12459240000084,254034187_D01`
	buf := []byte(txt)
	res, err := unmarshalcsv.UnmarshalCSV(buf, ',')
	assert.Nil(t, err)
	assert.Equal(t, 4, len(res))
}
