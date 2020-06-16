package tablesql_test

import (
	"testing"

	"github.com/gorexlv/tablesql"
	"github.com/stretchr/testify/assert"
)

func TestParse_CreateTable(t *testing.T) {
	parser := &tablesql.Parser{}

	resp, err := parser.Raw("create table aa(str varchar(20), inte int, primary key(str))")
	assert.Nil(t, err)
	t.Logf("rep: %+v", resp)
}
