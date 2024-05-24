// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetFirewall(ctx *pulumi.Context, args *GetFirewallArgs, opts ...pulumi.InvokeOption) (*GetFirewallResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetFirewallResult
	err := ctx.Invoke("digitalocean-native:firewalls/v2:getFirewall", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetFirewallArgs struct {
	// A unique ID that can be used to identify and reference a firewall.
	FirewallId string `pulumi:"firewallId"`
}

type GetFirewallResult struct {
	Items GetFirewallProperties `pulumi:"items"`
}

func GetFirewallOutput(ctx *pulumi.Context, args GetFirewallOutputArgs, opts ...pulumi.InvokeOption) GetFirewallResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFirewallResult, error) {
			args := v.(GetFirewallArgs)
			r, err := GetFirewall(ctx, &args, opts...)
			var s GetFirewallResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFirewallResultOutput)
}

type GetFirewallOutputArgs struct {
	// A unique ID that can be used to identify and reference a firewall.
	FirewallId pulumi.StringInput `pulumi:"firewallId"`
}

func (GetFirewallOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallArgs)(nil)).Elem()
}

type GetFirewallResultOutput struct{ *pulumi.OutputState }

func (GetFirewallResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallResult)(nil)).Elem()
}

func (o GetFirewallResultOutput) ToGetFirewallResultOutput() GetFirewallResultOutput {
	return o
}

func (o GetFirewallResultOutput) ToGetFirewallResultOutputWithContext(ctx context.Context) GetFirewallResultOutput {
	return o
}

func (o GetFirewallResultOutput) Items() GetFirewallPropertiesOutput {
	return o.ApplyT(func(v GetFirewallResult) GetFirewallProperties { return v.Items }).(GetFirewallPropertiesOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFirewallResultOutput{})
}