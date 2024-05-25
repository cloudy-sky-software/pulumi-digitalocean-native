// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetReservedIP(ctx *pulumi.Context, args *GetReservedIPArgs, opts ...pulumi.InvokeOption) (*GetReservedIPResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetReservedIPResult
	err := ctx.Invoke("digitalocean-native:reserved_ips/v2:getReservedIP", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type GetReservedIPArgs struct {
	// A reserved IP address.
	ReservedIp string `pulumi:"reservedIp"`
}

type GetReservedIPResult struct {
	Items GetReservedIPProperties `pulumi:"items"`
}

// Defaults sets the appropriate defaults for GetReservedIPResult
func (val *GetReservedIPResult) Defaults() *GetReservedIPResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Items = *tmp.Items.Defaults()

	return &tmp
}

func GetReservedIPOutput(ctx *pulumi.Context, args GetReservedIPOutputArgs, opts ...pulumi.InvokeOption) GetReservedIPResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetReservedIPResult, error) {
			args := v.(GetReservedIPArgs)
			r, err := GetReservedIP(ctx, &args, opts...)
			var s GetReservedIPResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetReservedIPResultOutput)
}

type GetReservedIPOutputArgs struct {
	// A reserved IP address.
	ReservedIp pulumi.StringInput `pulumi:"reservedIp"`
}

func (GetReservedIPOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetReservedIPArgs)(nil)).Elem()
}

type GetReservedIPResultOutput struct{ *pulumi.OutputState }

func (GetReservedIPResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetReservedIPResult)(nil)).Elem()
}

func (o GetReservedIPResultOutput) ToGetReservedIPResultOutput() GetReservedIPResultOutput {
	return o
}

func (o GetReservedIPResultOutput) ToGetReservedIPResultOutputWithContext(ctx context.Context) GetReservedIPResultOutput {
	return o
}

func (o GetReservedIPResultOutput) Items() GetReservedIPPropertiesOutput {
	return o.ApplyT(func(v GetReservedIPResult) GetReservedIPProperties { return v.Items }).(GetReservedIPPropertiesOutput)
}

func init() {
	pulumi.RegisterOutputType(GetReservedIPResultOutput{})
}
