// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDatabasesFirewallRules(ctx *pulumi.Context, args *ListDatabasesFirewallRulesArgs, opts ...pulumi.InvokeOption) (*ListDatabasesFirewallRulesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListDatabasesFirewallRulesResult
	err := ctx.Invoke("digitalocean-native:databases/v2:listDatabasesFirewallRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDatabasesFirewallRulesArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid string `pulumi:"databaseClusterUuid"`
}

type ListDatabasesFirewallRulesResult struct {
	Items ListDatabasesFirewallRulesProperties `pulumi:"items"`
}

func ListDatabasesFirewallRulesOutput(ctx *pulumi.Context, args ListDatabasesFirewallRulesOutputArgs, opts ...pulumi.InvokeOption) ListDatabasesFirewallRulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDatabasesFirewallRulesResult, error) {
			args := v.(ListDatabasesFirewallRulesArgs)
			r, err := ListDatabasesFirewallRules(ctx, &args, opts...)
			var s ListDatabasesFirewallRulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDatabasesFirewallRulesResultOutput)
}

type ListDatabasesFirewallRulesOutputArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid pulumi.StringInput `pulumi:"databaseClusterUuid"`
}

func (ListDatabasesFirewallRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabasesFirewallRulesArgs)(nil)).Elem()
}

type ListDatabasesFirewallRulesResultOutput struct{ *pulumi.OutputState }

func (ListDatabasesFirewallRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabasesFirewallRulesResult)(nil)).Elem()
}

func (o ListDatabasesFirewallRulesResultOutput) ToListDatabasesFirewallRulesResultOutput() ListDatabasesFirewallRulesResultOutput {
	return o
}

func (o ListDatabasesFirewallRulesResultOutput) ToListDatabasesFirewallRulesResultOutputWithContext(ctx context.Context) ListDatabasesFirewallRulesResultOutput {
	return o
}

func (o ListDatabasesFirewallRulesResultOutput) Items() ListDatabasesFirewallRulesPropertiesOutput {
	return o.ApplyT(func(v ListDatabasesFirewallRulesResult) ListDatabasesFirewallRulesProperties { return v.Items }).(ListDatabasesFirewallRulesPropertiesOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDatabasesFirewallRulesResultOutput{})
}