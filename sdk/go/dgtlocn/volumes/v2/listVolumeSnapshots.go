// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListVolumeSnapshots(ctx *pulumi.Context, args *ListVolumeSnapshotsArgs, opts ...pulumi.InvokeOption) (*ListVolumeSnapshotsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListVolumeSnapshotsResult
	err := ctx.Invoke("digitalocean-native:volumes/v2:listVolumeSnapshots", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListVolumeSnapshotsArgs struct {
	// The ID of the block storage volume.
	VolumeId string `pulumi:"volumeId"`
}

type ListVolumeSnapshotsResult struct {
	Items ListVolumeSnapshotsItems `pulumi:"items"`
}

func ListVolumeSnapshotsOutput(ctx *pulumi.Context, args ListVolumeSnapshotsOutputArgs, opts ...pulumi.InvokeOption) ListVolumeSnapshotsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListVolumeSnapshotsResult, error) {
			args := v.(ListVolumeSnapshotsArgs)
			r, err := ListVolumeSnapshots(ctx, &args, opts...)
			var s ListVolumeSnapshotsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListVolumeSnapshotsResultOutput)
}

type ListVolumeSnapshotsOutputArgs struct {
	// The ID of the block storage volume.
	VolumeId pulumi.StringInput `pulumi:"volumeId"`
}

func (ListVolumeSnapshotsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVolumeSnapshotsArgs)(nil)).Elem()
}

type ListVolumeSnapshotsResultOutput struct{ *pulumi.OutputState }

func (ListVolumeSnapshotsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVolumeSnapshotsResult)(nil)).Elem()
}

func (o ListVolumeSnapshotsResultOutput) ToListVolumeSnapshotsResultOutput() ListVolumeSnapshotsResultOutput {
	return o
}

func (o ListVolumeSnapshotsResultOutput) ToListVolumeSnapshotsResultOutputWithContext(ctx context.Context) ListVolumeSnapshotsResultOutput {
	return o
}

func (o ListVolumeSnapshotsResultOutput) Items() ListVolumeSnapshotsItemsOutput {
	return o.ApplyT(func(v ListVolumeSnapshotsResult) ListVolumeSnapshotsItems { return v.Items }).(ListVolumeSnapshotsItemsOutput)
}

func init() {
	pulumi.RegisterOutputType(ListVolumeSnapshotsResultOutput{})
}
