// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDatabases(ctx *pulumi.Context, args *ListDatabasesArgs, opts ...pulumi.InvokeOption) (*ListDatabasesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListDatabasesResult
	err := ctx.Invoke("digitalocean-native:databases/v2:listDatabases", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDatabasesArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid string `pulumi:"databaseClusterUuid"`
}

type ListDatabasesResult struct {
	Dbs []DatabaseType `pulumi:"dbs"`
}

func ListDatabasesOutput(ctx *pulumi.Context, args ListDatabasesOutputArgs, opts ...pulumi.InvokeOption) ListDatabasesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDatabasesResultOutput, error) {
			args := v.(ListDatabasesArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv ListDatabasesResult
			secret, err := ctx.InvokePackageRaw("digitalocean-native:databases/v2:listDatabases", args, &rv, "", opts...)
			if err != nil {
				return ListDatabasesResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(ListDatabasesResultOutput)
			if secret {
				return pulumi.ToSecret(output).(ListDatabasesResultOutput), nil
			}
			return output, nil
		}).(ListDatabasesResultOutput)
}

type ListDatabasesOutputArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid pulumi.StringInput `pulumi:"databaseClusterUuid"`
}

func (ListDatabasesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabasesArgs)(nil)).Elem()
}

type ListDatabasesResultOutput struct{ *pulumi.OutputState }

func (ListDatabasesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabasesResult)(nil)).Elem()
}

func (o ListDatabasesResultOutput) ToListDatabasesResultOutput() ListDatabasesResultOutput {
	return o
}

func (o ListDatabasesResultOutput) ToListDatabasesResultOutputWithContext(ctx context.Context) ListDatabasesResultOutput {
	return o
}

func (o ListDatabasesResultOutput) Dbs() DatabaseTypeArrayOutput {
	return o.ApplyT(func(v ListDatabasesResult) []DatabaseType { return v.Dbs }).(DatabaseTypeArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDatabasesResultOutput{})
}
