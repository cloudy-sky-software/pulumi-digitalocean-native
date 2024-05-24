// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDroplet(ctx *pulumi.Context, args *GetDropletArgs, opts ...pulumi.InvokeOption) (*GetDropletResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDropletResult
	err := ctx.Invoke("digitalocean-native:droplets/v2:getDroplet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type GetDropletArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId string `pulumi:"dropletId"`
}

type GetDropletResult struct {
	Items GetDropletProperties `pulumi:"items"`
}

// Defaults sets the appropriate defaults for GetDropletResult
func (val *GetDropletResult) Defaults() *GetDropletResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Items = *tmp.Items.Defaults()

	return &tmp
}

func GetDropletOutput(ctx *pulumi.Context, args GetDropletOutputArgs, opts ...pulumi.InvokeOption) GetDropletResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDropletResult, error) {
			args := v.(GetDropletArgs)
			r, err := GetDroplet(ctx, &args, opts...)
			var s GetDropletResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDropletResultOutput)
}

type GetDropletOutputArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId pulumi.StringInput `pulumi:"dropletId"`
}

func (GetDropletOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDropletArgs)(nil)).Elem()
}

type GetDropletResultOutput struct{ *pulumi.OutputState }

func (GetDropletResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDropletResult)(nil)).Elem()
}

func (o GetDropletResultOutput) ToGetDropletResultOutput() GetDropletResultOutput {
	return o
}

func (o GetDropletResultOutput) ToGetDropletResultOutputWithContext(ctx context.Context) GetDropletResultOutput {
	return o
}

func (o GetDropletResultOutput) Items() GetDropletPropertiesOutput {
	return o.ApplyT(func(v GetDropletResult) GetDropletProperties { return v.Items }).(GetDropletPropertiesOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDropletResultOutput{})
}
