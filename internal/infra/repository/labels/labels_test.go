package labels_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/i3onilha/MESEnterpriseSmart/internal/infra/mysql"
	"github.com/i3onilha/MESEnterpriseSmart/internal/infra/repository/labels"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	queries, err := mysql.New("")
	assert.Nil(t, err)
	label := labels.New(queries)
	assert.NotNil(t, label)
}
