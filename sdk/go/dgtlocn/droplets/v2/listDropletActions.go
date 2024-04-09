// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDropletActions(ctx *pulumi.Context, args *ListDropletActionsArgs, opts ...pulumi.InvokeOption) (*ListDropletActionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListDropletActionsResult
	err := ctx.Invoke("digitalocean-native:droplets/v2:listDropletActions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDropletActionsArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId string `pulumi:"dropletId"`
}

type ListDropletActionsResult struct {
	Items ListDropletActionsItems `pulumi:"items"`
}

func ListDropletActionsOutput(ctx *pulumi.Context, args ListDropletActionsOutputArgs, opts ...pulumi.InvokeOption) ListDropletActionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDropletActionsResult, error) {
			args := v.(ListDropletActionsArgs)
			r, err := ListDropletActions(ctx, &args, opts...)
			var s ListDropletActionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDropletActionsResultOutput)
}

type ListDropletActionsOutputArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId pulumi.StringInput `pulumi:"dropletId"`
}

func (ListDropletActionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDropletActionsArgs)(nil)).Elem()
}

type ListDropletActionsResultOutput struct{ *pulumi.OutputState }

func (ListDropletActionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDropletActionsResult)(nil)).Elem()
}

func (o ListDropletActionsResultOutput) ToListDropletActionsResultOutput() ListDropletActionsResultOutput {
	return o
}

func (o ListDropletActionsResultOutput) ToListDropletActionsResultOutputWithContext(ctx context.Context) ListDropletActionsResultOutput {
	return o
}

func (o ListDropletActionsResultOutput) Items() ListDropletActionsItemsOutput {
	return o.ApplyT(func(v ListDropletActionsResult) ListDropletActionsItems { return v.Items }).(ListDropletActionsItemsOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDropletActionsResultOutput{})
}
