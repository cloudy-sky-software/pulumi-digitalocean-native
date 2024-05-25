// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListTags(ctx *pulumi.Context, args *ListTagsArgs, opts ...pulumi.InvokeOption) (*ListTagsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListTagsResult
	err := ctx.Invoke("digitalocean-native:tags/v2:listTags", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListTagsArgs struct {
}

type ListTagsResult struct {
	Items ListTagsItems `pulumi:"items"`
}

func ListTagsOutput(ctx *pulumi.Context, args ListTagsOutputArgs, opts ...pulumi.InvokeOption) ListTagsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListTagsResult, error) {
			args := v.(ListTagsArgs)
			r, err := ListTags(ctx, &args, opts...)
			var s ListTagsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListTagsResultOutput)
}

type ListTagsOutputArgs struct {
}

func (ListTagsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListTagsArgs)(nil)).Elem()
}

type ListTagsResultOutput struct{ *pulumi.OutputState }

func (ListTagsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListTagsResult)(nil)).Elem()
}

func (o ListTagsResultOutput) ToListTagsResultOutput() ListTagsResultOutput {
	return o
}

func (o ListTagsResultOutput) ToListTagsResultOutputWithContext(ctx context.Context) ListTagsResultOutput {
	return o
}

func (o ListTagsResultOutput) Items() ListTagsItemsOutput {
	return o.ApplyT(func(v ListTagsResult) ListTagsItems { return v.Items }).(ListTagsItemsOutput)
}

func init() {
	pulumi.RegisterOutputType(ListTagsResultOutput{})
}
