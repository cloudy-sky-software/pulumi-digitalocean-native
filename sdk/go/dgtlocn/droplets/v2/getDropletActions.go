// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDropletActions(ctx *pulumi.Context, args *GetDropletActionsArgs, opts ...pulumi.InvokeOption) (*GetDropletActionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDropletActionsResult
	err := ctx.Invoke("digitalocean-native:droplets/v2:getDropletActions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type GetDropletActionsArgs struct {
	// A unique numeric ID that can be used to identify and reference an action.
	ActionId string `pulumi:"actionId"`
	// A unique identifier for a Droplet instance.
	DropletId string `pulumi:"dropletId"`
}

type GetDropletActionsResult struct {
	Items GetDropletActionsProperties `pulumi:"items"`
}

// Defaults sets the appropriate defaults for GetDropletActionsResult
func (val *GetDropletActionsResult) Defaults() *GetDropletActionsResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Items = *tmp.Items.Defaults()

	return &tmp
}

func GetDropletActionsOutput(ctx *pulumi.Context, args GetDropletActionsOutputArgs, opts ...pulumi.InvokeOption) GetDropletActionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDropletActionsResult, error) {
			args := v.(GetDropletActionsArgs)
			r, err := GetDropletActions(ctx, &args, opts...)
			var s GetDropletActionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDropletActionsResultOutput)
}

type GetDropletActionsOutputArgs struct {
	// A unique numeric ID that can be used to identify and reference an action.
	ActionId pulumi.StringInput `pulumi:"actionId"`
	// A unique identifier for a Droplet instance.
	DropletId pulumi.StringInput `pulumi:"dropletId"`
}

func (GetDropletActionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDropletActionsArgs)(nil)).Elem()
}

type GetDropletActionsResultOutput struct{ *pulumi.OutputState }

func (GetDropletActionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDropletActionsResult)(nil)).Elem()
}

func (o GetDropletActionsResultOutput) ToGetDropletActionsResultOutput() GetDropletActionsResultOutput {
	return o
}

func (o GetDropletActionsResultOutput) ToGetDropletActionsResultOutputWithContext(ctx context.Context) GetDropletActionsResultOutput {
	return o
}

func (o GetDropletActionsResultOutput) Items() GetDropletActionsPropertiesOutput {
	return o.ApplyT(func(v GetDropletActionsResult) GetDropletActionsProperties { return v.Items }).(GetDropletActionsPropertiesOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDropletActionsResultOutput{})
}