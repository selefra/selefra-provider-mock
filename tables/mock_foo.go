package tables

import (
	"context"
	"fmt"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-utils/pkg/id_util"
)

type TableMockFooGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableMockFooGenerator{}

func (x *TableMockFooGenerator) GetTableName() string {
	return "mock_foo"
}

func (x *TableMockFooGenerator) GetTableDescription() string {
	return ""
}

func (x *TableMockFooGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableMockFooGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableMockFooGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			for i := 0; i < 100; i++ {
				foo := &Foo{
					ID:    "foo-" + id_util.RandomId(),
					Key:   fmt.Sprintf("foo-key-%d", i),
					Value: fmt.Sprintf("foo-value-%d", i),
				}
				resultChannel <- foo
			}
			return nil
		},
	}
}

func (x *TableMockFooGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

type Foo struct {
	ID    string
	Key   string
	Value string
}

func (x *TableMockFooGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Key")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("value").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Value")).Build(),
	}
}

func (x *TableMockFooGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableMockBarGenerator{}),
	}
}