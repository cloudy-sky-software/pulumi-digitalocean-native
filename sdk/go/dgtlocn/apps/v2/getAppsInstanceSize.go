// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetAppsInstanceSize(ctx *pulumi.Context, args *GetAppsInstanceSizeArgs, opts ...pulumi.InvokeOption) (*GetAppsInstanceSizeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAppsInstanceSizeResult
	err := ctx.Invoke("digitalocean-native:apps/v2:getAppsInstanceSize", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type GetAppsInstanceSizeArgs struct {
	// The slug of the instance size
	Slug string `pulumi:"slug"`
}

type GetAppsInstanceSizeResult struct {
	Items AppsGetInstanceSizeResponse `pulumi:"items"`
}

// Defaults sets the appropriate defaults for GetAppsInstanceSizeResult
func (val *GetAppsInstanceSizeResult) Defaults() *GetAppsInstanceSizeResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Items = *tmp.Items.Defaults()

	return &tmp
}

func GetAppsInstanceSizeOutput(ctx *pulumi.Context, args GetAppsInstanceSizeOutputArgs, opts ...pulumi.InvokeOption) GetAppsInstanceSizeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppsInstanceSizeResult, error) {
			args := v.(GetAppsInstanceSizeArgs)
			r, err := GetAppsInstanceSize(ctx, &args, opts...)
			var s GetAppsInstanceSizeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAppsInstanceSizeResultOutput)
}

type GetAppsInstanceSizeOutputArgs struct {
	// The slug of the instance size
	Slug pulumi.StringInput `pulumi:"slug"`
}

func (GetAppsInstanceSizeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppsInstanceSizeArgs)(nil)).Elem()
}

type GetAppsInstanceSizeResultOutput struct{ *pulumi.OutputState }

func (GetAppsInstanceSizeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppsInstanceSizeResult)(nil)).Elem()
}

func (o GetAppsInstanceSizeResultOutput) ToGetAppsInstanceSizeResultOutput() GetAppsInstanceSizeResultOutput {
	return o
}

func (o GetAppsInstanceSizeResultOutput) ToGetAppsInstanceSizeResultOutputWithContext(ctx context.Context) GetAppsInstanceSizeResultOutput {
	return o
}

func (o GetAppsInstanceSizeResultOutput) Items() AppsGetInstanceSizeResponseOutput {
	return o.ApplyT(func(v GetAppsInstanceSizeResult) AppsGetInstanceSizeResponse { return v.Items }).(AppsGetInstanceSizeResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppsInstanceSizeResultOutput{})
}
