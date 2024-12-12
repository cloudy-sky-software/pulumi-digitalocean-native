// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSizes(ctx *pulumi.Context, args *ListSizesArgs, opts ...pulumi.InvokeOption) (*ListSizesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListSizesResult
	err := ctx.Invoke("digitalocean-native:sizes/v2:listSizes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSizesArgs struct {
}

type ListSizesResult struct {
	Links *PageLinks `pulumi:"links"`
	Meta  MetaMeta   `pulumi:"meta"`
	Sizes []Size     `pulumi:"sizes"`
}

func ListSizesOutput(ctx *pulumi.Context, args ListSizesOutputArgs, opts ...pulumi.InvokeOption) ListSizesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListSizesResultOutput, error) {
			args := v.(ListSizesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("digitalocean-native:sizes/v2:listSizes", args, ListSizesResultOutput{}, options).(ListSizesResultOutput), nil
		}).(ListSizesResultOutput)
}

type ListSizesOutputArgs struct {
}

func (ListSizesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSizesArgs)(nil)).Elem()
}

type ListSizesResultOutput struct{ *pulumi.OutputState }

func (ListSizesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSizesResult)(nil)).Elem()
}

func (o ListSizesResultOutput) ToListSizesResultOutput() ListSizesResultOutput {
	return o
}

func (o ListSizesResultOutput) ToListSizesResultOutputWithContext(ctx context.Context) ListSizesResultOutput {
	return o
}

func (o ListSizesResultOutput) Links() PageLinksPtrOutput {
	return o.ApplyT(func(v ListSizesResult) *PageLinks { return v.Links }).(PageLinksPtrOutput)
}

func (o ListSizesResultOutput) Meta() MetaMetaOutput {
	return o.ApplyT(func(v ListSizesResult) MetaMeta { return v.Meta }).(MetaMetaOutput)
}

func (o ListSizesResultOutput) Sizes() SizeArrayOutput {
	return o.ApplyT(func(v ListSizesResult) []Size { return v.Sizes }).(SizeArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSizesResultOutput{})
}
