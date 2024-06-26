// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListVolumes(ctx *pulumi.Context, args *ListVolumesArgs, opts ...pulumi.InvokeOption) (*ListVolumesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListVolumesResult
	err := ctx.Invoke("digitalocean-native:volumes/v2:listVolumes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListVolumesArgs struct {
}

type ListVolumesResult struct {
	Items ListVolumesItems `pulumi:"items"`
}

func ListVolumesOutput(ctx *pulumi.Context, args ListVolumesOutputArgs, opts ...pulumi.InvokeOption) ListVolumesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListVolumesResult, error) {
			args := v.(ListVolumesArgs)
			r, err := ListVolumes(ctx, &args, opts...)
			var s ListVolumesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListVolumesResultOutput)
}

type ListVolumesOutputArgs struct {
}

func (ListVolumesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVolumesArgs)(nil)).Elem()
}

type ListVolumesResultOutput struct{ *pulumi.OutputState }

func (ListVolumesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVolumesResult)(nil)).Elem()
}

func (o ListVolumesResultOutput) ToListVolumesResultOutput() ListVolumesResultOutput {
	return o
}

func (o ListVolumesResultOutput) ToListVolumesResultOutputWithContext(ctx context.Context) ListVolumesResultOutput {
	return o
}

func (o ListVolumesResultOutput) Items() ListVolumesItemsOutput {
	return o.ApplyT(func(v ListVolumesResult) ListVolumesItems { return v.Items }).(ListVolumesItemsOutput)
}

func init() {
	pulumi.RegisterOutputType(ListVolumesResultOutput{})
}
