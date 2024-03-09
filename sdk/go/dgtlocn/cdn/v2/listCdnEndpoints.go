// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListCdnEndpoints(ctx *pulumi.Context, args *ListCdnEndpointsArgs, opts ...pulumi.InvokeOption) (*ListCdnEndpointsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListCdnEndpointsResult
	err := ctx.Invoke("digitalocean-native:cdn/v2:listCdnEndpoints", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListCdnEndpointsArgs struct {
}

type ListCdnEndpointsResult struct {
	Items ListCdnEndpoints `pulumi:"items"`
}

func ListCdnEndpointsOutput(ctx *pulumi.Context, args ListCdnEndpointsOutputArgs, opts ...pulumi.InvokeOption) ListCdnEndpointsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListCdnEndpointsResult, error) {
			args := v.(ListCdnEndpointsArgs)
			r, err := ListCdnEndpoints(ctx, &args, opts...)
			var s ListCdnEndpointsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListCdnEndpointsResultOutput)
}

type ListCdnEndpointsOutputArgs struct {
}

func (ListCdnEndpointsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListCdnEndpointsArgs)(nil)).Elem()
}

type ListCdnEndpointsResultOutput struct{ *pulumi.OutputState }

func (ListCdnEndpointsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListCdnEndpointsResult)(nil)).Elem()
}

func (o ListCdnEndpointsResultOutput) ToListCdnEndpointsResultOutput() ListCdnEndpointsResultOutput {
	return o
}

func (o ListCdnEndpointsResultOutput) ToListCdnEndpointsResultOutputWithContext(ctx context.Context) ListCdnEndpointsResultOutput {
	return o
}

func (o ListCdnEndpointsResultOutput) Items() ListCdnEndpointsOutput {
	return o.ApplyT(func(v ListCdnEndpointsResult) ListCdnEndpoints { return v.Items }).(ListCdnEndpointsOutput)
}

func init() {
	pulumi.RegisterOutputType(ListCdnEndpointsResultOutput{})
}