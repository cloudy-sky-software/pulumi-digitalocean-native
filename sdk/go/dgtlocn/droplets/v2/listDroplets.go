// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDroplets(ctx *pulumi.Context, args *ListDropletsArgs, opts ...pulumi.InvokeOption) (*ListDropletsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListDropletsResult
	err := ctx.Invoke("digitalocean-native:droplets/v2:listDroplets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDropletsArgs struct {
}

type ListDropletsResult struct {
	Items ListDropletsItems `pulumi:"items"`
}

func ListDropletsOutput(ctx *pulumi.Context, args ListDropletsOutputArgs, opts ...pulumi.InvokeOption) ListDropletsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDropletsResult, error) {
			args := v.(ListDropletsArgs)
			r, err := ListDroplets(ctx, &args, opts...)
			var s ListDropletsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDropletsResultOutput)
}

type ListDropletsOutputArgs struct {
}

func (ListDropletsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDropletsArgs)(nil)).Elem()
}

type ListDropletsResultOutput struct{ *pulumi.OutputState }

func (ListDropletsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDropletsResult)(nil)).Elem()
}

func (o ListDropletsResultOutput) ToListDropletsResultOutput() ListDropletsResultOutput {
	return o
}

func (o ListDropletsResultOutput) ToListDropletsResultOutputWithContext(ctx context.Context) ListDropletsResultOutput {
	return o
}

func (o ListDropletsResultOutput) Items() ListDropletsItemsOutput {
	return o.ApplyT(func(v ListDropletsResult) ListDropletsItems { return v.Items }).(ListDropletsItemsOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDropletsResultOutput{})
}
