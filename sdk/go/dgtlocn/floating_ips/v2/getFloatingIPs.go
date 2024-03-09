// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFloatingIPs(ctx *pulumi.Context, args *LookupFloatingIPsArgs, opts ...pulumi.InvokeOption) (*LookupFloatingIPsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFloatingIPsResult
	err := ctx.Invoke("digitalocean-native:floating_ips/v2:getFloatingIPs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupFloatingIPsArgs struct {
	// A floating IP address.
	FloatingIp string `pulumi:"floatingIp"`
}

type LookupFloatingIPsResult struct {
	Items GetFloatingIPsProperties `pulumi:"items"`
}

// Defaults sets the appropriate defaults for LookupFloatingIPsResult
func (val *LookupFloatingIPsResult) Defaults() *LookupFloatingIPsResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Items = *tmp.Items.Defaults()

	return &tmp
}

func LookupFloatingIPsOutput(ctx *pulumi.Context, args LookupFloatingIPsOutputArgs, opts ...pulumi.InvokeOption) LookupFloatingIPsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFloatingIPsResult, error) {
			args := v.(LookupFloatingIPsArgs)
			r, err := LookupFloatingIPs(ctx, &args, opts...)
			var s LookupFloatingIPsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFloatingIPsResultOutput)
}

type LookupFloatingIPsOutputArgs struct {
	// A floating IP address.
	FloatingIp pulumi.StringInput `pulumi:"floatingIp"`
}

func (LookupFloatingIPsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFloatingIPsArgs)(nil)).Elem()
}

type LookupFloatingIPsResultOutput struct{ *pulumi.OutputState }

func (LookupFloatingIPsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFloatingIPsResult)(nil)).Elem()
}

func (o LookupFloatingIPsResultOutput) ToLookupFloatingIPsResultOutput() LookupFloatingIPsResultOutput {
	return o
}

func (o LookupFloatingIPsResultOutput) ToLookupFloatingIPsResultOutputWithContext(ctx context.Context) LookupFloatingIPsResultOutput {
	return o
}

func (o LookupFloatingIPsResultOutput) Items() GetFloatingIPsPropertiesOutput {
	return o.ApplyT(func(v LookupFloatingIPsResult) GetFloatingIPsProperties { return v.Items }).(GetFloatingIPsPropertiesOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFloatingIPsResultOutput{})
}