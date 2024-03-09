// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetMonitoringDropletCpuMetrics(ctx *pulumi.Context, args *GetMonitoringDropletCpuMetricsArgs, opts ...pulumi.InvokeOption) (*GetMonitoringDropletCpuMetricsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetMonitoringDropletCpuMetricsResult
	err := ctx.Invoke("digitalocean-native:monitoring/v2:getMonitoringDropletCpuMetrics", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetMonitoringDropletCpuMetricsArgs struct {
}

type GetMonitoringDropletCpuMetricsResult struct {
	Items Metrics `pulumi:"items"`
}

func GetMonitoringDropletCpuMetricsOutput(ctx *pulumi.Context, args GetMonitoringDropletCpuMetricsOutputArgs, opts ...pulumi.InvokeOption) GetMonitoringDropletCpuMetricsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMonitoringDropletCpuMetricsResult, error) {
			args := v.(GetMonitoringDropletCpuMetricsArgs)
			r, err := GetMonitoringDropletCpuMetrics(ctx, &args, opts...)
			var s GetMonitoringDropletCpuMetricsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMonitoringDropletCpuMetricsResultOutput)
}

type GetMonitoringDropletCpuMetricsOutputArgs struct {
}

func (GetMonitoringDropletCpuMetricsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMonitoringDropletCpuMetricsArgs)(nil)).Elem()
}

type GetMonitoringDropletCpuMetricsResultOutput struct{ *pulumi.OutputState }

func (GetMonitoringDropletCpuMetricsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMonitoringDropletCpuMetricsResult)(nil)).Elem()
}

func (o GetMonitoringDropletCpuMetricsResultOutput) ToGetMonitoringDropletCpuMetricsResultOutput() GetMonitoringDropletCpuMetricsResultOutput {
	return o
}

func (o GetMonitoringDropletCpuMetricsResultOutput) ToGetMonitoringDropletCpuMetricsResultOutputWithContext(ctx context.Context) GetMonitoringDropletCpuMetricsResultOutput {
	return o
}

func (o GetMonitoringDropletCpuMetricsResultOutput) Items() MetricsOutput {
	return o.ApplyT(func(v GetMonitoringDropletCpuMetricsResult) Metrics { return v.Items }).(MetricsOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMonitoringDropletCpuMetricsResultOutput{})
}