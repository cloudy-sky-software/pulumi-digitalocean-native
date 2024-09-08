// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetAppsTier(ctx *pulumi.Context, args *GetAppsTierArgs, opts ...pulumi.InvokeOption) (*GetAppsTierResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAppsTierResult
	err := ctx.Invoke("digitalocean-native:apps/v2:getAppsTier", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetAppsTierArgs struct {
	// The slug of the tier
	Slug string `pulumi:"slug"`
}

type GetAppsTierResult struct {
	Tier *AppsTier `pulumi:"tier"`
}

func GetAppsTierOutput(ctx *pulumi.Context, args GetAppsTierOutputArgs, opts ...pulumi.InvokeOption) GetAppsTierResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppsTierResult, error) {
			args := v.(GetAppsTierArgs)
			r, err := GetAppsTier(ctx, &args, opts...)
			var s GetAppsTierResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAppsTierResultOutput)
}

type GetAppsTierOutputArgs struct {
	// The slug of the tier
	Slug pulumi.StringInput `pulumi:"slug"`
}

func (GetAppsTierOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppsTierArgs)(nil)).Elem()
}

type GetAppsTierResultOutput struct{ *pulumi.OutputState }

func (GetAppsTierResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppsTierResult)(nil)).Elem()
}

func (o GetAppsTierResultOutput) ToGetAppsTierResultOutput() GetAppsTierResultOutput {
	return o
}

func (o GetAppsTierResultOutput) ToGetAppsTierResultOutputWithContext(ctx context.Context) GetAppsTierResultOutput {
	return o
}

func (o GetAppsTierResultOutput) Tier() AppsTierPtrOutput {
	return o.ApplyT(func(v GetAppsTierResult) *AppsTier { return v.Tier }).(AppsTierPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppsTierResultOutput{})
}
