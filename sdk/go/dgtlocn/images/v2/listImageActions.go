// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListImageActions(ctx *pulumi.Context, args *ListImageActionsArgs, opts ...pulumi.InvokeOption) (*ListImageActionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListImageActionsResult
	err := ctx.Invoke("digitalocean-native:images/v2:listImageActions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListImageActionsArgs struct {
	// A unique number that can be used to identify and reference a specific image.
	ImageId string `pulumi:"imageId"`
}

type ListImageActionsResult struct {
	Actions []Action   `pulumi:"actions"`
	Links   *PageLinks `pulumi:"links"`
	Meta    MetaMeta   `pulumi:"meta"`
}

func ListImageActionsOutput(ctx *pulumi.Context, args ListImageActionsOutputArgs, opts ...pulumi.InvokeOption) ListImageActionsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListImageActionsResultOutput, error) {
			args := v.(ListImageActionsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("digitalocean-native:images/v2:listImageActions", args, ListImageActionsResultOutput{}, options).(ListImageActionsResultOutput), nil
		}).(ListImageActionsResultOutput)
}

type ListImageActionsOutputArgs struct {
	// A unique number that can be used to identify and reference a specific image.
	ImageId pulumi.StringInput `pulumi:"imageId"`
}

func (ListImageActionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListImageActionsArgs)(nil)).Elem()
}

type ListImageActionsResultOutput struct{ *pulumi.OutputState }

func (ListImageActionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListImageActionsResult)(nil)).Elem()
}

func (o ListImageActionsResultOutput) ToListImageActionsResultOutput() ListImageActionsResultOutput {
	return o
}

func (o ListImageActionsResultOutput) ToListImageActionsResultOutputWithContext(ctx context.Context) ListImageActionsResultOutput {
	return o
}

func (o ListImageActionsResultOutput) Actions() ActionArrayOutput {
	return o.ApplyT(func(v ListImageActionsResult) []Action { return v.Actions }).(ActionArrayOutput)
}

func (o ListImageActionsResultOutput) Links() PageLinksPtrOutput {
	return o.ApplyT(func(v ListImageActionsResult) *PageLinks { return v.Links }).(PageLinksPtrOutput)
}

func (o ListImageActionsResultOutput) Meta() MetaMetaOutput {
	return o.ApplyT(func(v ListImageActionsResult) MetaMeta { return v.Meta }).(MetaMetaOutput)
}

func init() {
	pulumi.RegisterOutputType(ListImageActionsResultOutput{})
}
