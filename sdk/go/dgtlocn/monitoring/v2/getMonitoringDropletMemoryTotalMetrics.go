// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetMonitoringDropletMemoryTotalMetrics(ctx *pulumi.Context, args *GetMonitoringDropletMemoryTotalMetricsArgs, opts ...pulumi.InvokeOption) (*GetMonitoringDropletMemoryTotalMetricsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetMonitoringDropletMemoryTotalMetricsResult
	err := ctx.Invoke("digitalocean-native:monitoring/v2:getMonitoringDropletMemoryTotalMetrics", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetMonitoringDropletMemoryTotalMetricsArgs struct {
}

type GetMonitoringDropletMemoryTotalMetricsResult struct {
	Items Metrics `pulumi:"items"`
}

func GetMonitoringDropletMemoryTotalMetricsOutput(ctx *pulumi.Context, args GetMonitoringDropletMemoryTotalMetricsOutputArgs, opts ...pulumi.InvokeOption) GetMonitoringDropletMemoryTotalMetricsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMonitoringDropletMemoryTotalMetricsResult, error) {
			args := v.(GetMonitoringDropletMemoryTotalMetricsArgs)
			r, err := GetMonitoringDropletMemoryTotalMetrics(ctx, &args, opts...)
			var s GetMonitoringDropletMemoryTotalMetricsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMonitoringDropletMemoryTotalMetricsResultOutput)
}

type GetMonitoringDropletMemoryTotalMetricsOutputArgs struct {
}

func (GetMonitoringDropletMemoryTotalMetricsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMonitoringDropletMemoryTotalMetricsArgs)(nil)).Elem()
}

type GetMonitoringDropletMemoryTotalMetricsResultOutput struct{ *pulumi.OutputState }

func (GetMonitoringDropletMemoryTotalMetricsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMonitoringDropletMemoryTotalMetricsResult)(nil)).Elem()
}

func (o GetMonitoringDropletMemoryTotalMetricsResultOutput) ToGetMonitoringDropletMemoryTotalMetricsResultOutput() GetMonitoringDropletMemoryTotalMetricsResultOutput {
	return o
}

func (o GetMonitoringDropletMemoryTotalMetricsResultOutput) ToGetMonitoringDropletMemoryTotalMetricsResultOutputWithContext(ctx context.Context) GetMonitoringDropletMemoryTotalMetricsResultOutput {
	return o
}

func (o GetMonitoringDropletMemoryTotalMetricsResultOutput) Items() MetricsOutput {
	return o.ApplyT(func(v GetMonitoringDropletMemoryTotalMetricsResult) Metrics { return v.Items }).(MetricsOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMonitoringDropletMemoryTotalMetricsResultOutput{})
}
