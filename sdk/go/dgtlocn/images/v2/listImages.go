// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListImages(ctx *pulumi.Context, args *ListImagesArgs, opts ...pulumi.InvokeOption) (*ListImagesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListImagesResult
	err := ctx.Invoke("digitalocean-native:images/v2:listImages", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListImagesArgs struct {
}

type ListImagesResult struct {
	Items ListImages `pulumi:"items"`
}

func ListImagesOutput(ctx *pulumi.Context, args ListImagesOutputArgs, opts ...pulumi.InvokeOption) ListImagesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListImagesResult, error) {
			args := v.(ListImagesArgs)
			r, err := ListImages(ctx, &args, opts...)
			var s ListImagesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListImagesResultOutput)
}

type ListImagesOutputArgs struct {
}

func (ListImagesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListImagesArgs)(nil)).Elem()
}

type ListImagesResultOutput struct{ *pulumi.OutputState }

func (ListImagesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListImagesResult)(nil)).Elem()
}

func (o ListImagesResultOutput) ToListImagesResultOutput() ListImagesResultOutput {
	return o
}

func (o ListImagesResultOutput) ToListImagesResultOutputWithContext(ctx context.Context) ListImagesResultOutput {
	return o
}

func (o ListImagesResultOutput) Items() ListImagesOutput {
	return o.ApplyT(func(v ListImagesResult) ListImages { return v.Items }).(ListImagesOutput)
}

func init() {
	pulumi.RegisterOutputType(ListImagesResultOutput{})
}