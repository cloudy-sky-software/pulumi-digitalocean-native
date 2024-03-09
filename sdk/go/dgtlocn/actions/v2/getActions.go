// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetActions(ctx *pulumi.Context, args *GetActionsArgs, opts ...pulumi.InvokeOption) (*GetActionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetActionsResult
	err := ctx.Invoke("digitalocean-native:actions/v2:getActions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type GetActionsArgs struct {
	// A unique numeric ID that can be used to identify and reference an action.
	ActionId string `pulumi:"actionId"`
}

type GetActionsResult struct {
	Items GetActionsProperties `pulumi:"items"`
}

// Defaults sets the appropriate defaults for GetActionsResult
func (val *GetActionsResult) Defaults() *GetActionsResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Items = *tmp.Items.Defaults()

	return &tmp
}

func GetActionsOutput(ctx *pulumi.Context, args GetActionsOutputArgs, opts ...pulumi.InvokeOption) GetActionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetActionsResult, error) {
			args := v.(GetActionsArgs)
			r, err := GetActions(ctx, &args, opts...)
			var s GetActionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetActionsResultOutput)
}

type GetActionsOutputArgs struct {
	// A unique numeric ID that can be used to identify and reference an action.
	ActionId pulumi.StringInput `pulumi:"actionId"`
}

func (GetActionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetActionsArgs)(nil)).Elem()
}

type GetActionsResultOutput struct{ *pulumi.OutputState }

func (GetActionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetActionsResult)(nil)).Elem()
}

func (o GetActionsResultOutput) ToGetActionsResultOutput() GetActionsResultOutput {
	return o
}

func (o GetActionsResultOutput) ToGetActionsResultOutputWithContext(ctx context.Context) GetActionsResultOutput {
	return o
}

func (o GetActionsResultOutput) Items() GetActionsPropertiesOutput {
	return o.ApplyT(func(v GetActionsResult) GetActionsProperties { return v.Items }).(GetActionsPropertiesOutput)
}

func init() {
	pulumi.RegisterOutputType(GetActionsResultOutput{})
}
