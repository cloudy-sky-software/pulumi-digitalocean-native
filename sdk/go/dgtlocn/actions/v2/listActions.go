// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListActions(ctx *pulumi.Context, args *ListActionsArgs, opts ...pulumi.InvokeOption) (*ListActionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListActionsResult
	err := ctx.Invoke("digitalocean-native:actions/v2:listActions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListActionsArgs struct {
}

type ListActionsResult struct {
	Actions []Action   `pulumi:"actions"`
	Links   *PageLinks `pulumi:"links"`
	Meta    MetaMeta   `pulumi:"meta"`
}

func ListActionsOutput(ctx *pulumi.Context, args ListActionsOutputArgs, opts ...pulumi.InvokeOption) ListActionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListActionsResult, error) {
			args := v.(ListActionsArgs)
			r, err := ListActions(ctx, &args, opts...)
			var s ListActionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListActionsResultOutput)
}

type ListActionsOutputArgs struct {
}

func (ListActionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListActionsArgs)(nil)).Elem()
}

type ListActionsResultOutput struct{ *pulumi.OutputState }

func (ListActionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListActionsResult)(nil)).Elem()
}

func (o ListActionsResultOutput) ToListActionsResultOutput() ListActionsResultOutput {
	return o
}

func (o ListActionsResultOutput) ToListActionsResultOutputWithContext(ctx context.Context) ListActionsResultOutput {
	return o
}

func (o ListActionsResultOutput) Actions() ActionArrayOutput {
	return o.ApplyT(func(v ListActionsResult) []Action { return v.Actions }).(ActionArrayOutput)
}

func (o ListActionsResultOutput) Links() PageLinksPtrOutput {
	return o.ApplyT(func(v ListActionsResult) *PageLinks { return v.Links }).(PageLinksPtrOutput)
}

func (o ListActionsResultOutput) Meta() MetaMetaOutput {
	return o.ApplyT(func(v ListActionsResult) MetaMeta { return v.Meta }).(MetaMetaOutput)
}

func init() {
	pulumi.RegisterOutputType(ListActionsResultOutput{})
}
