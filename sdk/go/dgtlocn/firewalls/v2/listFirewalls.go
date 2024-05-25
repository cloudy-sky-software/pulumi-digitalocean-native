// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListFirewalls(ctx *pulumi.Context, args *ListFirewallsArgs, opts ...pulumi.InvokeOption) (*ListFirewallsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListFirewallsResult
	err := ctx.Invoke("digitalocean-native:firewalls/v2:listFirewalls", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListFirewallsArgs struct {
}

type ListFirewallsResult struct {
	Items ListFirewallsItems `pulumi:"items"`
}

func ListFirewallsOutput(ctx *pulumi.Context, args ListFirewallsOutputArgs, opts ...pulumi.InvokeOption) ListFirewallsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListFirewallsResult, error) {
			args := v.(ListFirewallsArgs)
			r, err := ListFirewalls(ctx, &args, opts...)
			var s ListFirewallsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListFirewallsResultOutput)
}

type ListFirewallsOutputArgs struct {
}

func (ListFirewallsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListFirewallsArgs)(nil)).Elem()
}

type ListFirewallsResultOutput struct{ *pulumi.OutputState }

func (ListFirewallsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListFirewallsResult)(nil)).Elem()
}

func (o ListFirewallsResultOutput) ToListFirewallsResultOutput() ListFirewallsResultOutput {
	return o
}

func (o ListFirewallsResultOutput) ToListFirewallsResultOutputWithContext(ctx context.Context) ListFirewallsResultOutput {
	return o
}

func (o ListFirewallsResultOutput) Items() ListFirewallsItemsOutput {
	return o.ApplyT(func(v ListFirewallsResult) ListFirewallsItems { return v.Items }).(ListFirewallsItemsOutput)
}

func init() {
	pulumi.RegisterOutputType(ListFirewallsResultOutput{})
}
