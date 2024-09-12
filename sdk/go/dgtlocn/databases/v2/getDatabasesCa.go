// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDatabasesCa(ctx *pulumi.Context, args *GetDatabasesCaArgs, opts ...pulumi.InvokeOption) (*GetDatabasesCaResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDatabasesCaResult
	err := ctx.Invoke("digitalocean-native:databases/v2:getDatabasesCa", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDatabasesCaArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid string `pulumi:"databaseClusterUuid"`
}

type GetDatabasesCaResult struct {
	Ca Ca `pulumi:"ca"`
}

func GetDatabasesCaOutput(ctx *pulumi.Context, args GetDatabasesCaOutputArgs, opts ...pulumi.InvokeOption) GetDatabasesCaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDatabasesCaResultOutput, error) {
			args := v.(GetDatabasesCaArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetDatabasesCaResult
			secret, err := ctx.InvokePackageRaw("digitalocean-native:databases/v2:getDatabasesCa", args, &rv, "", opts...)
			if err != nil {
				return GetDatabasesCaResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetDatabasesCaResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetDatabasesCaResultOutput), nil
			}
			return output, nil
		}).(GetDatabasesCaResultOutput)
}

type GetDatabasesCaOutputArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid pulumi.StringInput `pulumi:"databaseClusterUuid"`
}

func (GetDatabasesCaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabasesCaArgs)(nil)).Elem()
}

type GetDatabasesCaResultOutput struct{ *pulumi.OutputState }

func (GetDatabasesCaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabasesCaResult)(nil)).Elem()
}

func (o GetDatabasesCaResultOutput) ToGetDatabasesCaResultOutput() GetDatabasesCaResultOutput {
	return o
}

func (o GetDatabasesCaResultOutput) ToGetDatabasesCaResultOutputWithContext(ctx context.Context) GetDatabasesCaResultOutput {
	return o
}

func (o GetDatabasesCaResultOutput) Ca() CaOutput {
	return o.ApplyT(func(v GetDatabasesCaResult) Ca { return v.Ca }).(CaOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDatabasesCaResultOutput{})
}
