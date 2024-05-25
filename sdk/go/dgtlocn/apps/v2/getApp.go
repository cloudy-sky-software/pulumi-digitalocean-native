// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetApp(ctx *pulumi.Context, args *GetAppArgs, opts ...pulumi.InvokeOption) (*GetAppResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAppResult
	err := ctx.Invoke("digitalocean-native:apps/v2:getApp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type GetAppArgs struct {
	// The ID of the app
	Id string `pulumi:"id"`
}

type GetAppResult struct {
	Items AppResponse `pulumi:"items"`
}

// Defaults sets the appropriate defaults for GetAppResult
func (val *GetAppResult) Defaults() *GetAppResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Items = *tmp.Items.Defaults()

	return &tmp
}

func GetAppOutput(ctx *pulumi.Context, args GetAppOutputArgs, opts ...pulumi.InvokeOption) GetAppResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppResult, error) {
			args := v.(GetAppArgs)
			r, err := GetApp(ctx, &args, opts...)
			var s GetAppResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAppResultOutput)
}

type GetAppOutputArgs struct {
	// The ID of the app
	Id pulumi.StringInput `pulumi:"id"`
}

func (GetAppOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppArgs)(nil)).Elem()
}

type GetAppResultOutput struct{ *pulumi.OutputState }

func (GetAppResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppResult)(nil)).Elem()
}

func (o GetAppResultOutput) ToGetAppResultOutput() GetAppResultOutput {
	return o
}

func (o GetAppResultOutput) ToGetAppResultOutputWithContext(ctx context.Context) GetAppResultOutput {
	return o
}

func (o GetAppResultOutput) Items() AppResponseOutput {
	return o.ApplyT(func(v GetAppResult) AppResponse { return v.Items }).(AppResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppResultOutput{})
}
