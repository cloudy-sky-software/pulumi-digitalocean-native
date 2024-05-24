// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetMonitoringDropletMemoryAvailableMetric(ctx *pulumi.Context, args *GetMonitoringDropletMemoryAvailableMetricArgs, opts ...pulumi.InvokeOption) (*GetMonitoringDropletMemoryAvailableMetricResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetMonitoringDropletMemoryAvailableMetricResult
	err := ctx.Invoke("digitalocean-native:monitoring/v2:getMonitoringDropletMemoryAvailableMetric", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetMonitoringDropletMemoryAvailableMetricArgs struct {
}

type GetMonitoringDropletMemoryAvailableMetricResult struct {
	Items Metrics `pulumi:"items"`
}

func GetMonitoringDropletMemoryAvailableMetricOutput(ctx *pulumi.Context, args GetMonitoringDropletMemoryAvailableMetricOutputArgs, opts ...pulumi.InvokeOption) GetMonitoringDropletMemoryAvailableMetricResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMonitoringDropletMemoryAvailableMetricResult, error) {
			args := v.(GetMonitoringDropletMemoryAvailableMetricArgs)
			r, err := GetMonitoringDropletMemoryAvailableMetric(ctx, &args, opts...)
			var s GetMonitoringDropletMemoryAvailableMetricResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMonitoringDropletMemoryAvailableMetricResultOutput)
}

type GetMonitoringDropletMemoryAvailableMetricOutputArgs struct {
}

func (GetMonitoringDropletMemoryAvailableMetricOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMonitoringDropletMemoryAvailableMetricArgs)(nil)).Elem()
}

type GetMonitoringDropletMemoryAvailableMetricResultOutput struct{ *pulumi.OutputState }

func (GetMonitoringDropletMemoryAvailableMetricResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMonitoringDropletMemoryAvailableMetricResult)(nil)).Elem()
}

func (o GetMonitoringDropletMemoryAvailableMetricResultOutput) ToGetMonitoringDropletMemoryAvailableMetricResultOutput() GetMonitoringDropletMemoryAvailableMetricResultOutput {
	return o
}

func (o GetMonitoringDropletMemoryAvailableMetricResultOutput) ToGetMonitoringDropletMemoryAvailableMetricResultOutputWithContext(ctx context.Context) GetMonitoringDropletMemoryAvailableMetricResultOutput {
	return o
}

func (o GetMonitoringDropletMemoryAvailableMetricResultOutput) Items() MetricsOutput {
	return o.ApplyT(func(v GetMonitoringDropletMemoryAvailableMetricResult) Metrics { return v.Items }).(MetricsOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMonitoringDropletMemoryAvailableMetricResultOutput{})
}