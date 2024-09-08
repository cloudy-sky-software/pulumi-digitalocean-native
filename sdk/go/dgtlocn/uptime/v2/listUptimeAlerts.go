// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListUptimeAlerts(ctx *pulumi.Context, args *ListUptimeAlertsArgs, opts ...pulumi.InvokeOption) (*ListUptimeAlertsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListUptimeAlertsResult
	err := ctx.Invoke("digitalocean-native:uptime/v2:listUptimeAlerts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListUptimeAlertsArgs struct {
	// A unique identifier for a check.
	CheckId string `pulumi:"checkId"`
}

type ListUptimeAlertsResult struct {
	Alerts []Alert    `pulumi:"alerts"`
	Links  *PageLinks `pulumi:"links"`
	Meta   MetaMeta   `pulumi:"meta"`
}

func ListUptimeAlertsOutput(ctx *pulumi.Context, args ListUptimeAlertsOutputArgs, opts ...pulumi.InvokeOption) ListUptimeAlertsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListUptimeAlertsResult, error) {
			args := v.(ListUptimeAlertsArgs)
			r, err := ListUptimeAlerts(ctx, &args, opts...)
			var s ListUptimeAlertsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListUptimeAlertsResultOutput)
}

type ListUptimeAlertsOutputArgs struct {
	// A unique identifier for a check.
	CheckId pulumi.StringInput `pulumi:"checkId"`
}

func (ListUptimeAlertsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListUptimeAlertsArgs)(nil)).Elem()
}

type ListUptimeAlertsResultOutput struct{ *pulumi.OutputState }

func (ListUptimeAlertsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListUptimeAlertsResult)(nil)).Elem()
}

func (o ListUptimeAlertsResultOutput) ToListUptimeAlertsResultOutput() ListUptimeAlertsResultOutput {
	return o
}

func (o ListUptimeAlertsResultOutput) ToListUptimeAlertsResultOutputWithContext(ctx context.Context) ListUptimeAlertsResultOutput {
	return o
}

func (o ListUptimeAlertsResultOutput) Alerts() AlertArrayOutput {
	return o.ApplyT(func(v ListUptimeAlertsResult) []Alert { return v.Alerts }).(AlertArrayOutput)
}

func (o ListUptimeAlertsResultOutput) Links() PageLinksPtrOutput {
	return o.ApplyT(func(v ListUptimeAlertsResult) *PageLinks { return v.Links }).(PageLinksPtrOutput)
}

func (o ListUptimeAlertsResultOutput) Meta() MetaMetaOutput {
	return o.ApplyT(func(v ListUptimeAlertsResult) MetaMeta { return v.Meta }).(MetaMetaOutput)
}

func init() {
	pulumi.RegisterOutputType(ListUptimeAlertsResultOutput{})
}
