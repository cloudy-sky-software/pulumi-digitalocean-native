// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDropletsAssociatedResources(ctx *pulumi.Context, args *ListDropletsAssociatedResourcesArgs, opts ...pulumi.InvokeOption) (*ListDropletsAssociatedResourcesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListDropletsAssociatedResourcesResult
	err := ctx.Invoke("digitalocean-native:droplets/v2:listDropletsAssociatedResources", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDropletsAssociatedResourcesArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId string `pulumi:"dropletId"`
}

type ListDropletsAssociatedResourcesResult struct {
	FloatingIps     []AssociatedResource `pulumi:"floatingIps"`
	ReservedIps     []AssociatedResource `pulumi:"reservedIps"`
	Snapshots       []AssociatedResource `pulumi:"snapshots"`
	VolumeSnapshots []AssociatedResource `pulumi:"volumeSnapshots"`
	Volumes         []AssociatedResource `pulumi:"volumes"`
}

func ListDropletsAssociatedResourcesOutput(ctx *pulumi.Context, args ListDropletsAssociatedResourcesOutputArgs, opts ...pulumi.InvokeOption) ListDropletsAssociatedResourcesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDropletsAssociatedResourcesResult, error) {
			args := v.(ListDropletsAssociatedResourcesArgs)
			r, err := ListDropletsAssociatedResources(ctx, &args, opts...)
			var s ListDropletsAssociatedResourcesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDropletsAssociatedResourcesResultOutput)
}

type ListDropletsAssociatedResourcesOutputArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId pulumi.StringInput `pulumi:"dropletId"`
}

func (ListDropletsAssociatedResourcesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDropletsAssociatedResourcesArgs)(nil)).Elem()
}

type ListDropletsAssociatedResourcesResultOutput struct{ *pulumi.OutputState }

func (ListDropletsAssociatedResourcesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDropletsAssociatedResourcesResult)(nil)).Elem()
}

func (o ListDropletsAssociatedResourcesResultOutput) ToListDropletsAssociatedResourcesResultOutput() ListDropletsAssociatedResourcesResultOutput {
	return o
}

func (o ListDropletsAssociatedResourcesResultOutput) ToListDropletsAssociatedResourcesResultOutputWithContext(ctx context.Context) ListDropletsAssociatedResourcesResultOutput {
	return o
}

func (o ListDropletsAssociatedResourcesResultOutput) FloatingIps() AssociatedResourceArrayOutput {
	return o.ApplyT(func(v ListDropletsAssociatedResourcesResult) []AssociatedResource { return v.FloatingIps }).(AssociatedResourceArrayOutput)
}

func (o ListDropletsAssociatedResourcesResultOutput) ReservedIps() AssociatedResourceArrayOutput {
	return o.ApplyT(func(v ListDropletsAssociatedResourcesResult) []AssociatedResource { return v.ReservedIps }).(AssociatedResourceArrayOutput)
}

func (o ListDropletsAssociatedResourcesResultOutput) Snapshots() AssociatedResourceArrayOutput {
	return o.ApplyT(func(v ListDropletsAssociatedResourcesResult) []AssociatedResource { return v.Snapshots }).(AssociatedResourceArrayOutput)
}

func (o ListDropletsAssociatedResourcesResultOutput) VolumeSnapshots() AssociatedResourceArrayOutput {
	return o.ApplyT(func(v ListDropletsAssociatedResourcesResult) []AssociatedResource { return v.VolumeSnapshots }).(AssociatedResourceArrayOutput)
}

func (o ListDropletsAssociatedResourcesResultOutput) Volumes() AssociatedResourceArrayOutput {
	return o.ApplyT(func(v ListDropletsAssociatedResourcesResult) []AssociatedResource { return v.Volumes }).(AssociatedResourceArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDropletsAssociatedResourcesResultOutput{})
}
