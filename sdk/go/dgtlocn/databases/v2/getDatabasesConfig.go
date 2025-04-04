// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDatabasesConfig(ctx *pulumi.Context, args *GetDatabasesConfigArgs, opts ...pulumi.InvokeOption) (*GetDatabasesConfigResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDatabasesConfigResult
	err := ctx.Invoke("digitalocean-native:databases/v2:getDatabasesConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDatabasesConfigArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid string `pulumi:"databaseClusterUuid"`
}

type GetDatabasesConfigResult struct {
	Config interface{} `pulumi:"config"`
}

func GetDatabasesConfigOutput(ctx *pulumi.Context, args GetDatabasesConfigOutputArgs, opts ...pulumi.InvokeOption) GetDatabasesConfigResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetDatabasesConfigResultOutput, error) {
			args := v.(GetDatabasesConfigArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("digitalocean-native:databases/v2:getDatabasesConfig", args, GetDatabasesConfigResultOutput{}, options).(GetDatabasesConfigResultOutput), nil
		}).(GetDatabasesConfigResultOutput)
}

type GetDatabasesConfigOutputArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid pulumi.StringInput `pulumi:"databaseClusterUuid"`
}

func (GetDatabasesConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabasesConfigArgs)(nil)).Elem()
}

type GetDatabasesConfigResultOutput struct{ *pulumi.OutputState }

func (GetDatabasesConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabasesConfigResult)(nil)).Elem()
}

func (o GetDatabasesConfigResultOutput) ToGetDatabasesConfigResultOutput() GetDatabasesConfigResultOutput {
	return o
}

func (o GetDatabasesConfigResultOutput) ToGetDatabasesConfigResultOutputWithContext(ctx context.Context) GetDatabasesConfigResultOutput {
	return o
}

func (o GetDatabasesConfigResultOutput) Config() pulumi.AnyOutput {
	return o.ApplyT(func(v GetDatabasesConfigResult) interface{} { return v.Config }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDatabasesConfigResultOutput{})
}
