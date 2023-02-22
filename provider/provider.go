package provider

import (
	"context"
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
				return nil, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `# Optional. by default assumes all regions
# regions:
#   - us-east-1
#   - us-west-2`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				return nil
			},
		},
		TransformerMeta:   schema.TransformerMeta{},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{},
	}
}
