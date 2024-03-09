// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetBalance(ctx *pulumi.Context, args *GetBalanceArgs, opts ...pulumi.InvokeOption) (*GetBalanceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetBalanceResult
	err := ctx.Invoke("digitalocean-native:customers/v2:getBalance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetBalanceArgs struct {
}

type GetBalanceResult struct {
	Items Balance `pulumi:"items"`
}

func GetBalanceOutput(ctx *pulumi.Context, args GetBalanceOutputArgs, opts ...pulumi.InvokeOption) GetBalanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBalanceResult, error) {
			args := v.(GetBalanceArgs)
			r, err := GetBalance(ctx, &args, opts...)
			var s GetBalanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBalanceResultOutput)
}

type GetBalanceOutputArgs struct {
}

func (GetBalanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBalanceArgs)(nil)).Elem()
}

type GetBalanceResultOutput struct{ *pulumi.OutputState }

func (GetBalanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBalanceResult)(nil)).Elem()
}

func (o GetBalanceResultOutput) ToGetBalanceResultOutput() GetBalanceResultOutput {
	return o
}

func (o GetBalanceResultOutput) ToGetBalanceResultOutputWithContext(ctx context.Context) GetBalanceResultOutput {
	return o
}

func (o GetBalanceResultOutput) Items() BalanceOutput {
	return o.ApplyT(func(v GetBalanceResult) Balance { return v.Items }).(BalanceOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBalanceResultOutput{})
}
