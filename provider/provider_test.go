package provider

import (
	"fmt"
	"github.com/selefra/selefra-provider-sdk/env"
	"github.com/selefra/selefra-provider-sdk/grpc/shard"
	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/storage/database_storage/postgresql_storage"
	"github.com/selefra/selefra-utils/pkg/json_util"
	"github.com/selefra/selefra-utils/pkg/pointer"
)

import (
	"context"

	"testing"
)

func TestProvider_PullTable(t *testing.T) {
	wk := "./"

	config := `foo-count: 1
bar-count: 1
sleep-seconds: 1`

	myProvider := GetProvider()

	Pull(myProvider, config, wk, "*")

}

func Test(t *testing.T) {

	myProvider := GetProvider()

	for _, table := range myProvider.TableList {

		fmt.Println(table.TableName)
	}

}

func Pull(myProvider *provider.Provider, config, workspace string, pullTables ...string) {

	diagnostics := schema.NewDiagnostics()

	initProviderRequest := &shard.ProviderInitRequest{

		Storage: &shard.Storage{

			Type: 0,

			StorageOptions: json_util.ToJsonBytes(postgresql_storage.NewPostgresqlStorageOptions(env.GetDatabaseDsn())),
		},

		Workspace:      &workspace,
		IsInstallInit:  pointer.TruePointer(),
		ProviderConfig: &config,
	}

	response, err := myProvider.Init(context.Background(), initProviderRequest)
	if err != nil {

		panic(diagnostics.AddFatal("init error: %s", err.Error()).ToString())

	}
	if diagnostics.AddDiagnostics(response.Diagnostics).HasError() {

		panic(diagnostics.ToString())

	}

	err = myProvider.PullTables(context.Background(), &shard.PullTablesRequest{
		Tables:        pullTables,
		MaxGoroutines: 1000,

		Timeout: 1000 * 60 * 60,
	}, shard.NewFakeProviderServerSender())

	if err != nil {
		panic(diagnostics.AddFatal("pull table error: %s", err.Error()).ToString())

	}

}
