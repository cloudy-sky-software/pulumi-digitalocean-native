// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetImage(ctx *pulumi.Context, args *GetImageArgs, opts ...pulumi.InvokeOption) (*GetImageResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetImageResult
	err := ctx.Invoke("digitalocean-native:images/v2:getImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetImageArgs struct {
	// A unique number (id) or string (slug) used to identify and reference a
	// specific image.
	//
	// **Public** images can be identified by image `id` or `slug`.
	//
	// **Private** images *must* be identified by image `id`.
	ImageId string `pulumi:"imageId"`
}

type GetImageResult struct {
	Image Image `pulumi:"image"`
}

func GetImageOutput(ctx *pulumi.Context, args GetImageOutputArgs, opts ...pulumi.InvokeOption) GetImageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetImageResult, error) {
			args := v.(GetImageArgs)
			r, err := GetImage(ctx, &args, opts...)
			var s GetImageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetImageResultOutput)
}

type GetImageOutputArgs struct {
	// A unique number (id) or string (slug) used to identify and reference a
	// specific image.
	//
	// **Public** images can be identified by image `id` or `slug`.
	//
	// **Private** images *must* be identified by image `id`.
	ImageId pulumi.StringInput `pulumi:"imageId"`
}

func (GetImageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetImageArgs)(nil)).Elem()
}

type GetImageResultOutput struct{ *pulumi.OutputState }

func (GetImageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetImageResult)(nil)).Elem()
}

func (o GetImageResultOutput) ToGetImageResultOutput() GetImageResultOutput {
	return o
}

func (o GetImageResultOutput) ToGetImageResultOutputWithContext(ctx context.Context) GetImageResultOutput {
	return o
}

func (o GetImageResultOutput) Image() ImageOutput {
	return o.ApplyT(func(v GetImageResult) Image { return v.Image }).(ImageOutput)
}

func init() {
	pulumi.RegisterOutputType(GetImageResultOutput{})
}
