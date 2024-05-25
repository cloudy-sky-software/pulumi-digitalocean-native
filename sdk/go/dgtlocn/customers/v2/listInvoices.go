// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListInvoices(ctx *pulumi.Context, args *ListInvoicesArgs, opts ...pulumi.InvokeOption) (*ListInvoicesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListInvoicesResult
	err := ctx.Invoke("digitalocean-native:customers/v2:listInvoices", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListInvoicesArgs struct {
}

type ListInvoicesResult struct {
	Items ListInvoicesItems `pulumi:"items"`
}

func ListInvoicesOutput(ctx *pulumi.Context, args ListInvoicesOutputArgs, opts ...pulumi.InvokeOption) ListInvoicesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListInvoicesResult, error) {
			args := v.(ListInvoicesArgs)
			r, err := ListInvoices(ctx, &args, opts...)
			var s ListInvoicesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListInvoicesResultOutput)
}

type ListInvoicesOutputArgs struct {
}

func (ListInvoicesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListInvoicesArgs)(nil)).Elem()
}

type ListInvoicesResultOutput struct{ *pulumi.OutputState }

func (ListInvoicesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListInvoicesResult)(nil)).Elem()
}

func (o ListInvoicesResultOutput) ToListInvoicesResultOutput() ListInvoicesResultOutput {
	return o
}

func (o ListInvoicesResultOutput) ToListInvoicesResultOutputWithContext(ctx context.Context) ListInvoicesResultOutput {
	return o
}

func (o ListInvoicesResultOutput) Items() ListInvoicesItemsOutput {
	return o.ApplyT(func(v ListInvoicesResult) ListInvoicesItems { return v.Items }).(ListInvoicesItemsOutput)
}

func init() {
	pulumi.RegisterOutputType(ListInvoicesResultOutput{})
}
