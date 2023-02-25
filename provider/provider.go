package provider

import (
	"context"
	client2 "github.com/selefra/selefra-provider-mock/client"
	"github.com/selefra/selefra-provider-mock/tables"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"

	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/spf13/viper"
)

var Version = "v0.0.1"

func GetProvider() *provider.Provider {
	return &provider.Provider{
		Name:    "mock",
		Version: Version,
		TableList: []*schema.Table{
			table_schema_generator.GenTableSchema(&tables.TableMockFooGenerator{}),
		},
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {
				client, diagnostics := client2.NewClient(config)
				return []any{client}, diagnostics
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `#foo-count: 3
#bar-count: 2
#sleep-seconds: 0`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				_, diagnostics := client2.NewClient(config)
				return diagnostics
			},
		},
		TransformerMeta:   schema.TransformerMeta{},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{},
	}
}
