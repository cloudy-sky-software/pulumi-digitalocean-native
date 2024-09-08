// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMonitoringAlertPolicy(ctx *pulumi.Context, args *LookupMonitoringAlertPolicyArgs, opts ...pulumi.InvokeOption) (*LookupMonitoringAlertPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupMonitoringAlertPolicyResult
	err := ctx.Invoke("digitalocean-native:monitoring/v2:getMonitoringAlertPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMonitoringAlertPolicyArgs struct {
	// A unique identifier for an alert policy.
	AlertUuid string `pulumi:"alertUuid"`
}

type LookupMonitoringAlertPolicyResult struct {
	Policy *AlertPolicy `pulumi:"policy"`
}

func LookupMonitoringAlertPolicyOutput(ctx *pulumi.Context, args LookupMonitoringAlertPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupMonitoringAlertPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMonitoringAlertPolicyResult, error) {
			args := v.(LookupMonitoringAlertPolicyArgs)
			r, err := LookupMonitoringAlertPolicy(ctx, &args, opts...)
			var s LookupMonitoringAlertPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMonitoringAlertPolicyResultOutput)
}

type LookupMonitoringAlertPolicyOutputArgs struct {
	// A unique identifier for an alert policy.
	AlertUuid pulumi.StringInput `pulumi:"alertUuid"`
}

func (LookupMonitoringAlertPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMonitoringAlertPolicyArgs)(nil)).Elem()
}

type LookupMonitoringAlertPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupMonitoringAlertPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMonitoringAlertPolicyResult)(nil)).Elem()
}

func (o LookupMonitoringAlertPolicyResultOutput) ToLookupMonitoringAlertPolicyResultOutput() LookupMonitoringAlertPolicyResultOutput {
	return o
}

func (o LookupMonitoringAlertPolicyResultOutput) ToLookupMonitoringAlertPolicyResultOutputWithContext(ctx context.Context) LookupMonitoringAlertPolicyResultOutput {
	return o
}

func (o LookupMonitoringAlertPolicyResultOutput) Policy() AlertPolicyPtrOutput {
	return o.ApplyT(func(v LookupMonitoringAlertPolicyResult) *AlertPolicy { return v.Policy }).(AlertPolicyPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMonitoringAlertPolicyResultOutput{})
}
