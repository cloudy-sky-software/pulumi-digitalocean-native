// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSnapshots(ctx *pulumi.Context, args *ListSnapshotsArgs, opts ...pulumi.InvokeOption) (*ListSnapshotsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListSnapshotsResult
	err := ctx.Invoke("digitalocean-native:snapshots/v2:listSnapshots", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSnapshotsArgs struct {
}

type ListSnapshotsResult struct {
	Items ListSnapshots `pulumi:"items"`
}

func ListSnapshotsOutput(ctx *pulumi.Context, args ListSnapshotsOutputArgs, opts ...pulumi.InvokeOption) ListSnapshotsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSnapshotsResult, error) {
			args := v.(ListSnapshotsArgs)
			r, err := ListSnapshots(ctx, &args, opts...)
			var s ListSnapshotsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSnapshotsResultOutput)
}

type ListSnapshotsOutputArgs struct {
}

func (ListSnapshotsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSnapshotsArgs)(nil)).Elem()
}

type ListSnapshotsResultOutput struct{ *pulumi.OutputState }

func (ListSnapshotsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSnapshotsResult)(nil)).Elem()
}

func (o ListSnapshotsResultOutput) ToListSnapshotsResultOutput() ListSnapshotsResultOutput {
	return o
}

func (o ListSnapshotsResultOutput) ToListSnapshotsResultOutputWithContext(ctx context.Context) ListSnapshotsResultOutput {
	return o
}

func (o ListSnapshotsResultOutput) Items() ListSnapshotsOutput {
	return o.ApplyT(func(v ListSnapshotsResult) ListSnapshots { return v.Items }).(ListSnapshotsOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSnapshotsResultOutput{})
}