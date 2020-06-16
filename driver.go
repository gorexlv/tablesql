package tablesql

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"net/url"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/xwb1989/sqlparser"
)

type TableStoreDriver struct{}

// accessId:accessKey@endpoint/instanceName?
// tablestoreConn
// "https://web-feed-dev.cn-beijing.ots.aliyuncs.com/instanceName?accessId=&accessKey="
func (TableStoreDriver) Open(dsn string) (driver.Conn, error) {
	u, err := url.Parse(dsn)
	if err != nil {
		return nil, err
	}

	client := tablestore.NewClientWithConfig(u.Host, u.Path, u.Query().Get("accessId"), u.Query().Get("accessKey"))
	rep, err := client.ListTable()
	if err != nil {
		return nil, err
	}
	tables := rep.TableNames
	schemes := make(map[string]*TableBean, len(tables))
	for _, table := range tables {
		res, err := client.DescribeTable(&tablestore.DescribeTableRequest{TableName: table})
		if err != nil {
			return nil, err
		}
		schemes[table] = &TableBean{
			TableMeta:          res.TableMeta,
			TableOption:        res.TableOption,
			ReservedThroughput: res.ReservedThroughput,
			StreamDetails:      res.StreamDetails,
		}
	}

	conn := &tablestoreConn{
		client:       client,
		tables:       tables,
		tableSchemes: schemes,
	}

	return conn, err
}

type tablestoreConn struct {
	client       *tablestore.TableStoreClient
	tables       []string
	tableSchemes map[string]*TableBean
}

// Prepare returns a prepared statement, bound to this connection.
func (tc *tablestoreConn) Prepare(query string) (driver.Stmt, error) {
	ts := &tablestoreStmt{
		tablestoreConn: tc,
	}

	return ts, nil
}

// Close invalidates and potentially stops any current
// prepared statements and transactions, marking this
// connection as no longer in use.
//
// Because the sql package maintains a free pool of
// connections and only calls Close when there's a surplus of
// idle connections, it shouldn't be necessary for drivers to
// do their own connection caching.
func (tablestoreConn) Close() error {
	panic("not implemented") // TODO: Implement
}

// Begin starts and returns a new transaction.
//
// Deprecated: Drivers should implement ConnBeginTx instead (or additionally).
func (tablestoreConn) Begin() (driver.Tx, error) {
	panic("not implemented") // TODO: Implement
}

type tablestoreStmt struct {
	*tablestoreConn
}

// Close closes the statement.
//
// As of Go 1.1, a Stmt will not be closed if it's in use
// by any queries.
func (tablestoreStmt) Close() error {
	panic("not implemented") // TODO: Implement
}

// NumInput returns the number of placeholder parameters.
//
// If NumInput returns >= 0, the sql package will sanity check
// argument counts from callers and return errors to the caller
// before the statement's Exec or Query methods are called.
//
// NumInput may also return -1, if the driver doesn't know
// its number of placeholders. In that case, the sql package
// will not sanity check Exec or Query argument counts.
func (tablestoreStmt) NumInput() int {
	panic("not implemented") // TODO: Implement
}

// Exec executes a query that doesn't return rows, such
// as an INSERT or UPDATE.
//
// Deprecated: Drivers should implement StmtExecContext instead (or additionally).
func (tablestoreStmt) Exec(args []driver.Value) (driver.Result, error) {
	stmt, err := sqlparser.Parse(query)
	if err != nil {
		return nil, err
	}

	switch stmt := stmt.(type) {
	case *sqlparser.Select:
	case *sqlparser.Insert:
	case *sqlparser.Delete:
	case *sqlparser.Update:
	case *sqlparser.DDL: // Create|Drop
		// return parser.rawDBDDL(stmt)
	}

	return nil, errors.New("unsupport stmt")
}

// Query executes a query that may return rows, such as a
// SELECT.
//
// Deprecated: Drivers should implement StmtQueryContext instead (or additionally).
func (tablestoreStmt) Query(args []driver.Value) (driver.Rows, error) {
	panic("not implemented") // TODO: Implement
}

type tablestoreResult struct{}

// LastInsertId returns the database's auto-generated ID
// after, for example, an INSERT into a table with primary
// key.
func (tablestoreResult) LastInsertId() (int64, error) {
	panic("not implemented") // TODO: Implement
}

// RowsAffected returns the number of rows affected by the
// query.
func (tablestoreResult) RowsAffected() (int64, error) {
	panic("not implemented") // TODO: Implement
}

type tablestoreRows struct{}

// Columns returns the names of the columns. The number of
// columns of the result is inferred from the length of the
// slice. If a particular column name isn't known, an empty
// string should be returned for that entry.
func (tablestoreRows) Columns() []string {
	panic("not implemented") // TODO: Implement
}

// Close closes the rows iterator.
func (tablestoreRows) Close() error {
	panic("not implemented") // TODO: Implement
}

// Next is called to populate the next row of data into
// the provided slice. The provided slice will be the same
// size as the Columns() are wide.
//
// Next should return io.EOF when there are no more rows.
//
// The dest should not be written to outside of Next. Care
// should be taken when closing Rows not to modify
// a buffer held in dest.
func (tablestoreRows) Next(dest []driver.Value) error {
	panic("not implemented") // TODO: Implement
}

type tablestoreValue struct{}

func init() {
	sql.Register("tablestore", &TableStoreDriver{})
}
