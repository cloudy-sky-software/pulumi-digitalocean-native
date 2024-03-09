// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDatabasesMigrationStatus(ctx *pulumi.Context, args *GetDatabasesMigrationStatusArgs, opts ...pulumi.InvokeOption) (*GetDatabasesMigrationStatusResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDatabasesMigrationStatusResult
	err := ctx.Invoke("digitalocean-native:databases/v2:getDatabasesMigrationStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDatabasesMigrationStatusArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid string `pulumi:"databaseClusterUuid"`
}

type GetDatabasesMigrationStatusResult struct {
	Items OnlineMigration `pulumi:"items"`
}

func GetDatabasesMigrationStatusOutput(ctx *pulumi.Context, args GetDatabasesMigrationStatusOutputArgs, opts ...pulumi.InvokeOption) GetDatabasesMigrationStatusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDatabasesMigrationStatusResult, error) {
			args := v.(GetDatabasesMigrationStatusArgs)
			r, err := GetDatabasesMigrationStatus(ctx, &args, opts...)
			var s GetDatabasesMigrationStatusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDatabasesMigrationStatusResultOutput)
}

type GetDatabasesMigrationStatusOutputArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid pulumi.StringInput `pulumi:"databaseClusterUuid"`
}

func (GetDatabasesMigrationStatusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabasesMigrationStatusArgs)(nil)).Elem()
}

type GetDatabasesMigrationStatusResultOutput struct{ *pulumi.OutputState }

func (GetDatabasesMigrationStatusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabasesMigrationStatusResult)(nil)).Elem()
}

func (o GetDatabasesMigrationStatusResultOutput) ToGetDatabasesMigrationStatusResultOutput() GetDatabasesMigrationStatusResultOutput {
	return o
}

func (o GetDatabasesMigrationStatusResultOutput) ToGetDatabasesMigrationStatusResultOutputWithContext(ctx context.Context) GetDatabasesMigrationStatusResultOutput {
	return o
}

func (o GetDatabasesMigrationStatusResultOutput) Items() OnlineMigrationOutput {
	return o.ApplyT(func(v GetDatabasesMigrationStatusResult) OnlineMigration { return v.Items }).(OnlineMigrationOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDatabasesMigrationStatusResultOutput{})
}