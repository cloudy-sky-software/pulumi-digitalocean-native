// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDroplet(ctx *pulumi.Context, args *LookupDropletArgs, opts ...pulumi.InvokeOption) (*LookupDropletResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDropletResult
	err := ctx.Invoke("digitalocean-native:droplets/v2:getDroplet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDropletArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId string `pulumi:"dropletId"`
}

type LookupDropletResult struct {
	Items GetDropletProperties `pulumi:"items"`
}

// Defaults sets the appropriate defaults for LookupDropletResult
func (val *LookupDropletResult) Defaults() *LookupDropletResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Items = *tmp.Items.Defaults()

	return &tmp
}

func LookupDropletOutput(ctx *pulumi.Context, args LookupDropletOutputArgs, opts ...pulumi.InvokeOption) LookupDropletResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDropletResult, error) {
			args := v.(LookupDropletArgs)
			r, err := LookupDroplet(ctx, &args, opts...)
			var s LookupDropletResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDropletResultOutput)
}

type LookupDropletOutputArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId pulumi.StringInput `pulumi:"dropletId"`
}

func (LookupDropletOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDropletArgs)(nil)).Elem()
}

type LookupDropletResultOutput struct{ *pulumi.OutputState }

func (LookupDropletResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDropletResult)(nil)).Elem()
}

func (o LookupDropletResultOutput) ToLookupDropletResultOutput() LookupDropletResultOutput {
	return o
}

func (o LookupDropletResultOutput) ToLookupDropletResultOutputWithContext(ctx context.Context) LookupDropletResultOutput {
	return o
}

func (o LookupDropletResultOutput) Items() GetDropletPropertiesOutput {
	return o.ApplyT(func(v LookupDropletResult) GetDropletProperties { return v.Items }).(GetDropletPropertiesOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDropletResultOutput{})
}
