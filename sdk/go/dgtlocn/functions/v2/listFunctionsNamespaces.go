// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListFunctionsNamespaces(ctx *pulumi.Context, args *ListFunctionsNamespacesArgs, opts ...pulumi.InvokeOption) (*ListFunctionsNamespacesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListFunctionsNamespacesResult
	err := ctx.Invoke("digitalocean-native:functions/v2:listFunctionsNamespaces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListFunctionsNamespacesArgs struct {
}

type ListFunctionsNamespacesResult struct {
	Items ListFunctionsNamespacesItems `pulumi:"items"`
}

func ListFunctionsNamespacesOutput(ctx *pulumi.Context, args ListFunctionsNamespacesOutputArgs, opts ...pulumi.InvokeOption) ListFunctionsNamespacesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListFunctionsNamespacesResult, error) {
			args := v.(ListFunctionsNamespacesArgs)
			r, err := ListFunctionsNamespaces(ctx, &args, opts...)
			var s ListFunctionsNamespacesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListFunctionsNamespacesResultOutput)
}

type ListFunctionsNamespacesOutputArgs struct {
}

func (ListFunctionsNamespacesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListFunctionsNamespacesArgs)(nil)).Elem()
}

type ListFunctionsNamespacesResultOutput struct{ *pulumi.OutputState }

func (ListFunctionsNamespacesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListFunctionsNamespacesResult)(nil)).Elem()
}

func (o ListFunctionsNamespacesResultOutput) ToListFunctionsNamespacesResultOutput() ListFunctionsNamespacesResultOutput {
	return o
}

func (o ListFunctionsNamespacesResultOutput) ToListFunctionsNamespacesResultOutputWithContext(ctx context.Context) ListFunctionsNamespacesResultOutput {
	return o
}

func (o ListFunctionsNamespacesResultOutput) Items() ListFunctionsNamespacesItemsOutput {
	return o.ApplyT(func(v ListFunctionsNamespacesResult) ListFunctionsNamespacesItems { return v.Items }).(ListFunctionsNamespacesItemsOutput)
}

func init() {
	pulumi.RegisterOutputType(ListFunctionsNamespacesResultOutput{})
}
