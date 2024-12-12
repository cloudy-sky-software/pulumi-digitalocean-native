// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFloatingIP(ctx *pulumi.Context, args *LookupFloatingIPArgs, opts ...pulumi.InvokeOption) (*LookupFloatingIPResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFloatingIPResult
	err := ctx.Invoke("digitalocean-native:floating_ips/v2:getFloatingIP", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupFloatingIPArgs struct {
	// A floating IP address.
	FloatingIp string `pulumi:"floatingIp"`
}

type LookupFloatingIPResult struct {
	FloatingIp *FloatingIpType `pulumi:"floatingIp"`
}

// Defaults sets the appropriate defaults for LookupFloatingIPResult
func (val *LookupFloatingIPResult) Defaults() *LookupFloatingIPResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.FloatingIp = tmp.FloatingIp.Defaults()

	return &tmp
}
func LookupFloatingIPOutput(ctx *pulumi.Context, args LookupFloatingIPOutputArgs, opts ...pulumi.InvokeOption) LookupFloatingIPResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupFloatingIPResultOutput, error) {
			args := v.(LookupFloatingIPArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("digitalocean-native:floating_ips/v2:getFloatingIP", args, LookupFloatingIPResultOutput{}, options).(LookupFloatingIPResultOutput), nil
		}).(LookupFloatingIPResultOutput)
}

type LookupFloatingIPOutputArgs struct {
	// A floating IP address.
	FloatingIp pulumi.StringInput `pulumi:"floatingIp"`
}

func (LookupFloatingIPOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFloatingIPArgs)(nil)).Elem()
}

type LookupFloatingIPResultOutput struct{ *pulumi.OutputState }

func (LookupFloatingIPResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFloatingIPResult)(nil)).Elem()
}

func (o LookupFloatingIPResultOutput) ToLookupFloatingIPResultOutput() LookupFloatingIPResultOutput {
	return o
}

func (o LookupFloatingIPResultOutput) ToLookupFloatingIPResultOutputWithContext(ctx context.Context) LookupFloatingIPResultOutput {
	return o
}

func (o LookupFloatingIPResultOutput) FloatingIp() FloatingIpTypePtrOutput {
	return o.ApplyT(func(v LookupFloatingIPResult) *FloatingIpType { return v.FloatingIp }).(FloatingIpTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFloatingIPResultOutput{})
}
