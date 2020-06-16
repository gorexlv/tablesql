package tablesql

import "github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"

type TableBean struct {
	TableMeta          *tablestore.TableMeta
	TableOption        *tablestore.TableOption
	ReservedThroughput *tablestore.ReservedThroughput
	StreamDetails      *tablestore.StreamDetails
	IndexMetas         []*tablestore.IndexMeta
}
