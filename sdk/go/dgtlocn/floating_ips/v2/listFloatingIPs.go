// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListFloatingIPs(ctx *pulumi.Context, args *ListFloatingIPsArgs, opts ...pulumi.InvokeOption) (*ListFloatingIPsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListFloatingIPsResult
	err := ctx.Invoke("digitalocean-native:floating_ips/v2:listFloatingIPs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListFloatingIPsArgs struct {
}

type ListFloatingIPsResult struct {
	Items ListFloatingIPs `pulumi:"items"`
}

func ListFloatingIPsOutput(ctx *pulumi.Context, args ListFloatingIPsOutputArgs, opts ...pulumi.InvokeOption) ListFloatingIPsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListFloatingIPsResult, error) {
			args := v.(ListFloatingIPsArgs)
			r, err := ListFloatingIPs(ctx, &args, opts...)
			var s ListFloatingIPsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListFloatingIPsResultOutput)
}

type ListFloatingIPsOutputArgs struct {
}

func (ListFloatingIPsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListFloatingIPsArgs)(nil)).Elem()
}

type ListFloatingIPsResultOutput struct{ *pulumi.OutputState }

func (ListFloatingIPsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListFloatingIPsResult)(nil)).Elem()
}

func (o ListFloatingIPsResultOutput) ToListFloatingIPsResultOutput() ListFloatingIPsResultOutput {
	return o
}

func (o ListFloatingIPsResultOutput) ToListFloatingIPsResultOutputWithContext(ctx context.Context) ListFloatingIPsResultOutput {
	return o
}

func (o ListFloatingIPsResultOutput) Items() ListFloatingIPsOutput {
	return o.ApplyT(func(v ListFloatingIPsResult) ListFloatingIPs { return v.Items }).(ListFloatingIPsOutput)
}

func init() {
	pulumi.RegisterOutputType(ListFloatingIPsResultOutput{})
}