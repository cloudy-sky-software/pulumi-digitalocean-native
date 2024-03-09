// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReservedIPs(ctx *pulumi.Context, args *LookupReservedIPsArgs, opts ...pulumi.InvokeOption) (*LookupReservedIPsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupReservedIPsResult
	err := ctx.Invoke("digitalocean-native:reserved_ips/v2:getReservedIPs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupReservedIPsArgs struct {
	// A reserved IP address.
	ReservedIp string `pulumi:"reservedIp"`
}

type LookupReservedIPsResult struct {
	Items GetReservedIPsProperties `pulumi:"items"`
}

// Defaults sets the appropriate defaults for LookupReservedIPsResult
func (val *LookupReservedIPsResult) Defaults() *LookupReservedIPsResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Items = *tmp.Items.Defaults()

	return &tmp
}

func LookupReservedIPsOutput(ctx *pulumi.Context, args LookupReservedIPsOutputArgs, opts ...pulumi.InvokeOption) LookupReservedIPsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReservedIPsResult, error) {
			args := v.(LookupReservedIPsArgs)
			r, err := LookupReservedIPs(ctx, &args, opts...)
			var s LookupReservedIPsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReservedIPsResultOutput)
}

type LookupReservedIPsOutputArgs struct {
	// A reserved IP address.
	ReservedIp pulumi.StringInput `pulumi:"reservedIp"`
}

func (LookupReservedIPsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReservedIPsArgs)(nil)).Elem()
}

type LookupReservedIPsResultOutput struct{ *pulumi.OutputState }

func (LookupReservedIPsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReservedIPsResult)(nil)).Elem()
}

func (o LookupReservedIPsResultOutput) ToLookupReservedIPsResultOutput() LookupReservedIPsResultOutput {
	return o
}

func (o LookupReservedIPsResultOutput) ToLookupReservedIPsResultOutputWithContext(ctx context.Context) LookupReservedIPsResultOutput {
	return o
}

func (o LookupReservedIPsResultOutput) Items() GetReservedIPsPropertiesOutput {
	return o.ApplyT(func(v LookupReservedIPsResult) GetReservedIPsProperties { return v.Items }).(GetReservedIPsPropertiesOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReservedIPsResultOutput{})
}
