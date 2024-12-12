// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListVolumeActions(ctx *pulumi.Context, args *ListVolumeActionsArgs, opts ...pulumi.InvokeOption) (*ListVolumeActionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListVolumeActionsResult
	err := ctx.Invoke("digitalocean-native:volumes/v2:listVolumeActions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListVolumeActionsArgs struct {
	// The ID of the block storage volume.
	VolumeId string `pulumi:"volumeId"`
}

type ListVolumeActionsResult struct {
	Actions []VolumeAction `pulumi:"actions"`
	Links   *PageLinks     `pulumi:"links"`
	Meta    MetaMeta       `pulumi:"meta"`
}

func ListVolumeActionsOutput(ctx *pulumi.Context, args ListVolumeActionsOutputArgs, opts ...pulumi.InvokeOption) ListVolumeActionsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListVolumeActionsResultOutput, error) {
			args := v.(ListVolumeActionsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("digitalocean-native:volumes/v2:listVolumeActions", args, ListVolumeActionsResultOutput{}, options).(ListVolumeActionsResultOutput), nil
		}).(ListVolumeActionsResultOutput)
}

type ListVolumeActionsOutputArgs struct {
	// The ID of the block storage volume.
	VolumeId pulumi.StringInput `pulumi:"volumeId"`
}

func (ListVolumeActionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVolumeActionsArgs)(nil)).Elem()
}

type ListVolumeActionsResultOutput struct{ *pulumi.OutputState }

func (ListVolumeActionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVolumeActionsResult)(nil)).Elem()
}

func (o ListVolumeActionsResultOutput) ToListVolumeActionsResultOutput() ListVolumeActionsResultOutput {
	return o
}

func (o ListVolumeActionsResultOutput) ToListVolumeActionsResultOutputWithContext(ctx context.Context) ListVolumeActionsResultOutput {
	return o
}

func (o ListVolumeActionsResultOutput) Actions() VolumeActionArrayOutput {
	return o.ApplyT(func(v ListVolumeActionsResult) []VolumeAction { return v.Actions }).(VolumeActionArrayOutput)
}

func (o ListVolumeActionsResultOutput) Links() PageLinksPtrOutput {
	return o.ApplyT(func(v ListVolumeActionsResult) *PageLinks { return v.Links }).(PageLinksPtrOutput)
}

func (o ListVolumeActionsResultOutput) Meta() MetaMetaOutput {
	return o.ApplyT(func(v ListVolumeActionsResult) MetaMeta { return v.Meta }).(MetaMetaOutput)
}

func init() {
	pulumi.RegisterOutputType(ListVolumeActionsResultOutput{})
}
