package tables

import (
	"context"
	"fmt"
	"github.com/selefra/selefra-provider-mock/client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-utils/pkg/id_util"
	"time"
)

type TableMockBarGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableMockBarGenerator{}

func (x *TableMockBarGenerator) GetTableName() string {
	return "mock_bar"
}

func (x *TableMockBarGenerator) GetTableDescription() string {
	return ""
}

func (x *TableMockBarGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableMockBarGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableMockBarGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			client := taskClient.(*client.Client)
			for i := 0; i < client.Config.BarCount; i++ {
				bar := &Bar{
					ID:    "bar-" + id_util.RandomId(),
					Key:   fmt.Sprintf("bar-key-%d", i),
					Value: fmt.Sprintf("bar-value-%d", i),
				}
				resultChannel <- bar

				if client.Config.SleepSeconds > 0 {
					time.Sleep(time.Second * time.Duration(client.Config.SleepSeconds))
				}

			}
			return nil
		},
	}
}

func (x *TableMockBarGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

type Bar struct {
	ID    string
	Key   string
	Value string
}

func (x *TableMockBarGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("foo_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Key")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("value").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Value")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return time.Now(), nil
			})).Build(),
	}
}

func (x *TableMockBarGenerator) GetSubTables() []*schema.Table {
	return nil
}
