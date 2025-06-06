// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListFunctionsTriggers(ctx *pulumi.Context, args *ListFunctionsTriggersArgs, opts ...pulumi.InvokeOption) (*ListFunctionsTriggersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListFunctionsTriggersResult
	err := ctx.Invoke("digitalocean-native:functions/v2:listFunctionsTriggers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListFunctionsTriggersArgs struct {
	// The ID of the namespace to be managed.
	NamespaceId string `pulumi:"namespaceId"`
}

type ListFunctionsTriggersResult struct {
	Triggers []TriggerInfo `pulumi:"triggers"`
}

func ListFunctionsTriggersOutput(ctx *pulumi.Context, args ListFunctionsTriggersOutputArgs, opts ...pulumi.InvokeOption) ListFunctionsTriggersResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListFunctionsTriggersResultOutput, error) {
			args := v.(ListFunctionsTriggersArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("digitalocean-native:functions/v2:listFunctionsTriggers", args, ListFunctionsTriggersResultOutput{}, options).(ListFunctionsTriggersResultOutput), nil
		}).(ListFunctionsTriggersResultOutput)
}

type ListFunctionsTriggersOutputArgs struct {
	// The ID of the namespace to be managed.
	NamespaceId pulumi.StringInput `pulumi:"namespaceId"`
}

func (ListFunctionsTriggersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListFunctionsTriggersArgs)(nil)).Elem()
}

type ListFunctionsTriggersResultOutput struct{ *pulumi.OutputState }

func (ListFunctionsTriggersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListFunctionsTriggersResult)(nil)).Elem()
}

func (o ListFunctionsTriggersResultOutput) ToListFunctionsTriggersResultOutput() ListFunctionsTriggersResultOutput {
	return o
}

func (o ListFunctionsTriggersResultOutput) ToListFunctionsTriggersResultOutputWithContext(ctx context.Context) ListFunctionsTriggersResultOutput {
	return o
}

func (o ListFunctionsTriggersResultOutput) Triggers() TriggerInfoArrayOutput {
	return o.ApplyT(func(v ListFunctionsTriggersResult) []TriggerInfo { return v.Triggers }).(TriggerInfoArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListFunctionsTriggersResultOutput{})
}
