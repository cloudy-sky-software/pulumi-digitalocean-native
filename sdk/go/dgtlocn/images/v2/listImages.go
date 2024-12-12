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
	Images []Image    `pulumi:"images"`
	Links  *PageLinks `pulumi:"links"`
	Meta   MetaMeta   `pulumi:"meta"`
}

func ListImagesOutput(ctx *pulumi.Context, args ListImagesOutputArgs, opts ...pulumi.InvokeOption) ListImagesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListImagesResultOutput, error) {
			args := v.(ListImagesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("digitalocean-native:images/v2:listImages", args, ListImagesResultOutput{}, options).(ListImagesResultOutput), nil
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

func (o ListImagesResultOutput) Images() ImageArrayOutput {
	return o.ApplyT(func(v ListImagesResult) []Image { return v.Images }).(ImageArrayOutput)
}

func (o ListImagesResultOutput) Links() PageLinksPtrOutput {
	return o.ApplyT(func(v ListImagesResult) *PageLinks { return v.Links }).(PageLinksPtrOutput)
}

func (o ListImagesResultOutput) Meta() MetaMetaOutput {
	return o.ApplyT(func(v ListImagesResult) MetaMeta { return v.Meta }).(MetaMetaOutput)
}

func init() {
	pulumi.RegisterOutputType(ListImagesResultOutput{})
}
