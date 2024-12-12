// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFunctionsTrigger(ctx *pulumi.Context, args *LookupFunctionsTriggerArgs, opts ...pulumi.InvokeOption) (*LookupFunctionsTriggerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFunctionsTriggerResult
	err := ctx.Invoke("digitalocean-native:functions/v2:getFunctionsTrigger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFunctionsTriggerArgs struct {
	// The ID of the namespace to be managed.
	NamespaceId string `pulumi:"namespaceId"`
	// The name of the trigger to be managed.
	TriggerName string `pulumi:"triggerName"`
}

type LookupFunctionsTriggerResult struct {
	Trigger *TriggerInfo `pulumi:"trigger"`
}

func LookupFunctionsTriggerOutput(ctx *pulumi.Context, args LookupFunctionsTriggerOutputArgs, opts ...pulumi.InvokeOption) LookupFunctionsTriggerResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupFunctionsTriggerResultOutput, error) {
			args := v.(LookupFunctionsTriggerArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("digitalocean-native:functions/v2:getFunctionsTrigger", args, LookupFunctionsTriggerResultOutput{}, options).(LookupFunctionsTriggerResultOutput), nil
		}).(LookupFunctionsTriggerResultOutput)
}

type LookupFunctionsTriggerOutputArgs struct {
	// The ID of the namespace to be managed.
	NamespaceId pulumi.StringInput `pulumi:"namespaceId"`
	// The name of the trigger to be managed.
	TriggerName pulumi.StringInput `pulumi:"triggerName"`
}

func (LookupFunctionsTriggerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFunctionsTriggerArgs)(nil)).Elem()
}

type LookupFunctionsTriggerResultOutput struct{ *pulumi.OutputState }

func (LookupFunctionsTriggerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFunctionsTriggerResult)(nil)).Elem()
}

func (o LookupFunctionsTriggerResultOutput) ToLookupFunctionsTriggerResultOutput() LookupFunctionsTriggerResultOutput {
	return o
}

func (o LookupFunctionsTriggerResultOutput) ToLookupFunctionsTriggerResultOutputWithContext(ctx context.Context) LookupFunctionsTriggerResultOutput {
	return o
}

func (o LookupFunctionsTriggerResultOutput) Trigger() TriggerInfoPtrOutput {
	return o.ApplyT(func(v LookupFunctionsTriggerResult) *TriggerInfo { return v.Trigger }).(TriggerInfoPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFunctionsTriggerResultOutput{})
}
