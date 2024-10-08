// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFirewall(ctx *pulumi.Context, args *LookupFirewallArgs, opts ...pulumi.InvokeOption) (*LookupFirewallResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFirewallResult
	err := ctx.Invoke("digitalocean-native:firewalls/v2:getFirewall", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFirewallArgs struct {
	// A unique ID that can be used to identify and reference a firewall.
	FirewallId string `pulumi:"firewallId"`
}

type LookupFirewallResult struct {
	Firewall *FirewallType `pulumi:"firewall"`
}

func LookupFirewallOutput(ctx *pulumi.Context, args LookupFirewallOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallResultOutput, error) {
			args := v.(LookupFirewallArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupFirewallResult
			secret, err := ctx.InvokePackageRaw("digitalocean-native:firewalls/v2:getFirewall", args, &rv, "", opts...)
			if err != nil {
				return LookupFirewallResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupFirewallResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupFirewallResultOutput), nil
			}
			return output, nil
		}).(LookupFirewallResultOutput)
}

type LookupFirewallOutputArgs struct {
	// A unique ID that can be used to identify and reference a firewall.
	FirewallId pulumi.StringInput `pulumi:"firewallId"`
}

func (LookupFirewallOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallArgs)(nil)).Elem()
}

type LookupFirewallResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallResult)(nil)).Elem()
}

func (o LookupFirewallResultOutput) ToLookupFirewallResultOutput() LookupFirewallResultOutput {
	return o
}

func (o LookupFirewallResultOutput) ToLookupFirewallResultOutputWithContext(ctx context.Context) LookupFirewallResultOutput {
	return o
}

func (o LookupFirewallResultOutput) Firewall() FirewallTypePtrOutput {
	return o.ApplyT(func(v LookupFirewallResult) *FirewallType { return v.Firewall }).(FirewallTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallResultOutput{})
}
