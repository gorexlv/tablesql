package tablesql

import (
	"errors"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/xwb1989/sqlparser"
)

type Parser struct {
	client *tablestore.TableStoreClient

	tableNames []string
	tables     map[string]*TableBean
}

func (parser *Parser) Raw(sql string) (val interface{}, err error) {
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		return nil, err
	}

	switch stmt := stmt.(type) {
	case *sqlparser.Select:
	case *sqlparser.Insert:
	case *sqlparser.Delete:
	case *sqlparser.Update:
	case *sqlparser.DDL: // Create|Drop
		return parser.rawDBDDL(stmt)
	}

	return nil, errors.New("unsupport stmt")
}

func (parser *Parser) rawDBDDL(stmt *sqlparser.DDL) (val interface{}, err error) {
	switch action := stmt.Action; action {
	case sqlparser.CreateStr:
		parser.createTable(stmt)
	case sqlparser.DropStr:
	case sqlparser.AlterStr,
		sqlparser.RenameStr,
		sqlparser.TruncateStr,
		sqlparser.CreateVindexStr,
		sqlparser.AddColVindexStr,
		sqlparser.DropColVindexStr:
		return nil, errors.New("unsupport sql: " + action)
	}
	return nil, nil
}

func (parser *Parser) rawSelect(stmt *sqlparser.Select) (val interface{}, err error) {
	return nil, nil
}

func (parser *Parser) rawUpdate(stmt *sqlparser.Select) (val interface{}, err error) {
	return nil, nil
}

func (parser *Parser) rawInsert(stmt *sqlparser.Select) (val interface{}, err error) {
	return nil, nil
}

func (parser *Parser) rawDelete(stmt *sqlparser.Select) (val interface{}, err error) {
	return nil, nil
}

func (parser *Parser) ListTables() error {
	res, err := parser.client.ListTable()
	if err != nil {
		return err
	}

	parser.tableNames = res.TableNames
	return nil
}

func (parser *Parser) PreloadTable(name string) error {
	req := &tablestore.DescribeTableRequest{
		TableName: name,
	}
	res, err := parser.client.DescribeTable(req)
	if err != nil {
		return err
	}

	parser.tables[name] = &TableBean{
		TableMeta:          res.TableMeta,
		TableOption:        res.TableOption,
		ReservedThroughput: res.ReservedThroughput,
		StreamDetails:      res.StreamDetails,
	}
	return nil
}
