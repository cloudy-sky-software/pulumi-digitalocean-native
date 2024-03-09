// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListMonitoringAlertPolicy(ctx *pulumi.Context, args *ListMonitoringAlertPolicyArgs, opts ...pulumi.InvokeOption) (*ListMonitoringAlertPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListMonitoringAlertPolicyResult
	err := ctx.Invoke("digitalocean-native:monitoring/v2:listMonitoringAlertPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListMonitoringAlertPolicyArgs struct {
}

type ListMonitoringAlertPolicyResult struct {
	Items ListMonitoringAlertPolicy `pulumi:"items"`
}

func ListMonitoringAlertPolicyOutput(ctx *pulumi.Context, args ListMonitoringAlertPolicyOutputArgs, opts ...pulumi.InvokeOption) ListMonitoringAlertPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListMonitoringAlertPolicyResult, error) {
			args := v.(ListMonitoringAlertPolicyArgs)
			r, err := ListMonitoringAlertPolicy(ctx, &args, opts...)
			var s ListMonitoringAlertPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListMonitoringAlertPolicyResultOutput)
}

type ListMonitoringAlertPolicyOutputArgs struct {
}

func (ListMonitoringAlertPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMonitoringAlertPolicyArgs)(nil)).Elem()
}

type ListMonitoringAlertPolicyResultOutput struct{ *pulumi.OutputState }

func (ListMonitoringAlertPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMonitoringAlertPolicyResult)(nil)).Elem()
}

func (o ListMonitoringAlertPolicyResultOutput) ToListMonitoringAlertPolicyResultOutput() ListMonitoringAlertPolicyResultOutput {
	return o
}

func (o ListMonitoringAlertPolicyResultOutput) ToListMonitoringAlertPolicyResultOutputWithContext(ctx context.Context) ListMonitoringAlertPolicyResultOutput {
	return o
}

func (o ListMonitoringAlertPolicyResultOutput) Items() ListMonitoringAlertPolicyOutput {
	return o.ApplyT(func(v ListMonitoringAlertPolicyResult) ListMonitoringAlertPolicy { return v.Items }).(ListMonitoringAlertPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(ListMonitoringAlertPolicyResultOutput{})
}