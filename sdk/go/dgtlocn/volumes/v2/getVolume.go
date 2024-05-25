// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetVolume(ctx *pulumi.Context, args *GetVolumeArgs, opts ...pulumi.InvokeOption) (*GetVolumeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetVolumeResult
	err := ctx.Invoke("digitalocean-native:volumes/v2:getVolume", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetVolumeArgs struct {
	// The ID of the block storage volume.
	VolumeId string `pulumi:"volumeId"`
}

type GetVolumeResult struct {
	Items GetVolumeProperties `pulumi:"items"`
}

func GetVolumeOutput(ctx *pulumi.Context, args GetVolumeOutputArgs, opts ...pulumi.InvokeOption) GetVolumeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVolumeResult, error) {
			args := v.(GetVolumeArgs)
			r, err := GetVolume(ctx, &args, opts...)
			var s GetVolumeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVolumeResultOutput)
}

type GetVolumeOutputArgs struct {
	// The ID of the block storage volume.
	VolumeId pulumi.StringInput `pulumi:"volumeId"`
}

func (GetVolumeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVolumeArgs)(nil)).Elem()
}

type GetVolumeResultOutput struct{ *pulumi.OutputState }

func (GetVolumeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVolumeResult)(nil)).Elem()
}

func (o GetVolumeResultOutput) ToGetVolumeResultOutput() GetVolumeResultOutput {
	return o
}

func (o GetVolumeResultOutput) ToGetVolumeResultOutputWithContext(ctx context.Context) GetVolumeResultOutput {
	return o
}

func (o GetVolumeResultOutput) Items() GetVolumePropertiesOutput {
	return o.ApplyT(func(v GetVolumeResult) GetVolumeProperties { return v.Items }).(GetVolumePropertiesOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVolumeResultOutput{})
}
