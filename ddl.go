package tablesql

import (
	"fmt"

	"github.com/xwb1989/sqlparser"
)

func (parser *Parser) createTable(stmt *sqlparser.DDL) (val interface{}, err error) {
	// meta := &tablestore.TableMeta{
	// 	TableName: stmt.Table.Name.String(),
	// }
	for _, column := range stmt.TableSpec.Columns {
		fmt.Printf("column => %+v\n", column)
	}
	for _, index := range stmt.TableSpec.Indexes {
		fmt.Printf("index => %+v, %d\n", index.Info, len(index.Columns))
		for _, column := range index.Columns {
			fmt.Printf("index column => %+v\n", column)
		}
	}
	fmt.Printf("table spec: %+v\n", stmt)
	fmt.Printf("stmt.TableSpec => %+v\n", stmt.TableSpec)
	fmt.Printf("stmt.TableSpec => %+v\n", stmt.TableSpec.Indexes)
	fmt.Printf("stmt.VindexSpec => %+v\n", stmt.VindexSpec)
	return nil, nil
}

func (parser *Parser) dropTable(stmt *sqlparser.DDL) (val interface{}, err error) {
	return nil, nil
}
