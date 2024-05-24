// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDatabasesMigrationStatu(ctx *pulumi.Context, args *GetDatabasesMigrationStatuArgs, opts ...pulumi.InvokeOption) (*GetDatabasesMigrationStatuResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDatabasesMigrationStatuResult
	err := ctx.Invoke("digitalocean-native:databases/v2:getDatabasesMigrationStatu", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDatabasesMigrationStatuArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid string `pulumi:"databaseClusterUuid"`
}

type GetDatabasesMigrationStatuResult struct {
	Items OnlineMigration `pulumi:"items"`
}

func GetDatabasesMigrationStatuOutput(ctx *pulumi.Context, args GetDatabasesMigrationStatuOutputArgs, opts ...pulumi.InvokeOption) GetDatabasesMigrationStatuResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDatabasesMigrationStatuResult, error) {
			args := v.(GetDatabasesMigrationStatuArgs)
			r, err := GetDatabasesMigrationStatu(ctx, &args, opts...)
			var s GetDatabasesMigrationStatuResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDatabasesMigrationStatuResultOutput)
}

type GetDatabasesMigrationStatuOutputArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid pulumi.StringInput `pulumi:"databaseClusterUuid"`
}

func (GetDatabasesMigrationStatuOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabasesMigrationStatuArgs)(nil)).Elem()
}

type GetDatabasesMigrationStatuResultOutput struct{ *pulumi.OutputState }

func (GetDatabasesMigrationStatuResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabasesMigrationStatuResult)(nil)).Elem()
}

func (o GetDatabasesMigrationStatuResultOutput) ToGetDatabasesMigrationStatuResultOutput() GetDatabasesMigrationStatuResultOutput {
	return o
}

func (o GetDatabasesMigrationStatuResultOutput) ToGetDatabasesMigrationStatuResultOutputWithContext(ctx context.Context) GetDatabasesMigrationStatuResultOutput {
	return o
}

func (o GetDatabasesMigrationStatuResultOutput) Items() OnlineMigrationOutput {
	return o.ApplyT(func(v GetDatabasesMigrationStatuResult) OnlineMigration { return v.Items }).(OnlineMigrationOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDatabasesMigrationStatuResultOutput{})
}
