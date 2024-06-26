// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCdnEndpoint(ctx *pulumi.Context, args *LookupCdnEndpointArgs, opts ...pulumi.InvokeOption) (*LookupCdnEndpointResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCdnEndpointResult
	err := ctx.Invoke("digitalocean-native:cdn/v2:getCdnEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupCdnEndpointArgs struct {
	// A unique identifier for a CDN endpoint.
	CdnId string `pulumi:"cdnId"`
}

type LookupCdnEndpointResult struct {
	Items GetCdnEndpointProperties `pulumi:"items"`
}

// Defaults sets the appropriate defaults for LookupCdnEndpointResult
func (val *LookupCdnEndpointResult) Defaults() *LookupCdnEndpointResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Items = *tmp.Items.Defaults()

	return &tmp
}

func LookupCdnEndpointOutput(ctx *pulumi.Context, args LookupCdnEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupCdnEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCdnEndpointResult, error) {
			args := v.(LookupCdnEndpointArgs)
			r, err := LookupCdnEndpoint(ctx, &args, opts...)
			var s LookupCdnEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCdnEndpointResultOutput)
}

type LookupCdnEndpointOutputArgs struct {
	// A unique identifier for a CDN endpoint.
	CdnId pulumi.StringInput `pulumi:"cdnId"`
}

func (LookupCdnEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCdnEndpointArgs)(nil)).Elem()
}

type LookupCdnEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupCdnEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCdnEndpointResult)(nil)).Elem()
}

func (o LookupCdnEndpointResultOutput) ToLookupCdnEndpointResultOutput() LookupCdnEndpointResultOutput {
	return o
}

func (o LookupCdnEndpointResultOutput) ToLookupCdnEndpointResultOutputWithContext(ctx context.Context) LookupCdnEndpointResultOutput {
	return o
}

func (o LookupCdnEndpointResultOutput) Items() GetCdnEndpointPropertiesOutput {
	return o.ApplyT(func(v LookupCdnEndpointResult) GetCdnEndpointProperties { return v.Items }).(GetCdnEndpointPropertiesOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCdnEndpointResultOutput{})
}
