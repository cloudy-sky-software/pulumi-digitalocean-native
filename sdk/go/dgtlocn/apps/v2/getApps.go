// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApps(ctx *pulumi.Context, args *LookupAppsArgs, opts ...pulumi.InvokeOption) (*LookupAppsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAppsResult
	err := ctx.Invoke("digitalocean-native:apps/v2:getApps", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAppsArgs struct {
	// The ID of the app
	Id string `pulumi:"id"`
}

type LookupAppsResult struct {
	Items AppResponse `pulumi:"items"`
}

// Defaults sets the appropriate defaults for LookupAppsResult
func (val *LookupAppsResult) Defaults() *LookupAppsResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Items = *tmp.Items.Defaults()

	return &tmp
}

func LookupAppsOutput(ctx *pulumi.Context, args LookupAppsOutputArgs, opts ...pulumi.InvokeOption) LookupAppsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAppsResult, error) {
			args := v.(LookupAppsArgs)
			r, err := LookupApps(ctx, &args, opts...)
			var s LookupAppsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAppsResultOutput)
}

type LookupAppsOutputArgs struct {
	// The ID of the app
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupAppsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppsArgs)(nil)).Elem()
}

type LookupAppsResultOutput struct{ *pulumi.OutputState }

func (LookupAppsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppsResult)(nil)).Elem()
}

func (o LookupAppsResultOutput) ToLookupAppsResultOutput() LookupAppsResultOutput {
	return o
}

func (o LookupAppsResultOutput) ToLookupAppsResultOutputWithContext(ctx context.Context) LookupAppsResultOutput {
	return o
}

func (o LookupAppsResultOutput) Items() AppResponseOutput {
	return o.ApplyT(func(v LookupAppsResult) AppResponse { return v.Items }).(AppResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppsResultOutput{})
}