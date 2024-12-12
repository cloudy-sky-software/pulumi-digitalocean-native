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
	Namespaces []NamespaceInfo `pulumi:"namespaces"`
}

func ListFunctionsNamespacesOutput(ctx *pulumi.Context, args ListFunctionsNamespacesOutputArgs, opts ...pulumi.InvokeOption) ListFunctionsNamespacesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListFunctionsNamespacesResultOutput, error) {
			args := v.(ListFunctionsNamespacesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("digitalocean-native:functions/v2:listFunctionsNamespaces", args, ListFunctionsNamespacesResultOutput{}, options).(ListFunctionsNamespacesResultOutput), nil
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

func (o ListFunctionsNamespacesResultOutput) Namespaces() NamespaceInfoArrayOutput {
	return o.ApplyT(func(v ListFunctionsNamespacesResult) []NamespaceInfo { return v.Namespaces }).(NamespaceInfoArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListFunctionsNamespacesResultOutput{})
}
