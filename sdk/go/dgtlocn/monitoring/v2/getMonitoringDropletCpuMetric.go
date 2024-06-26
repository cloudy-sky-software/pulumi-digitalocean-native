// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetMonitoringDropletCpuMetric(ctx *pulumi.Context, args *GetMonitoringDropletCpuMetricArgs, opts ...pulumi.InvokeOption) (*GetMonitoringDropletCpuMetricResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetMonitoringDropletCpuMetricResult
	err := ctx.Invoke("digitalocean-native:monitoring/v2:getMonitoringDropletCpuMetric", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetMonitoringDropletCpuMetricArgs struct {
}

type GetMonitoringDropletCpuMetricResult struct {
	Items Metrics `pulumi:"items"`
}

func GetMonitoringDropletCpuMetricOutput(ctx *pulumi.Context, args GetMonitoringDropletCpuMetricOutputArgs, opts ...pulumi.InvokeOption) GetMonitoringDropletCpuMetricResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMonitoringDropletCpuMetricResult, error) {
			args := v.(GetMonitoringDropletCpuMetricArgs)
			r, err := GetMonitoringDropletCpuMetric(ctx, &args, opts...)
			var s GetMonitoringDropletCpuMetricResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMonitoringDropletCpuMetricResultOutput)
}

type GetMonitoringDropletCpuMetricOutputArgs struct {
}

func (GetMonitoringDropletCpuMetricOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMonitoringDropletCpuMetricArgs)(nil)).Elem()
}

type GetMonitoringDropletCpuMetricResultOutput struct{ *pulumi.OutputState }

func (GetMonitoringDropletCpuMetricResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMonitoringDropletCpuMetricResult)(nil)).Elem()
}

func (o GetMonitoringDropletCpuMetricResultOutput) ToGetMonitoringDropletCpuMetricResultOutput() GetMonitoringDropletCpuMetricResultOutput {
	return o
}

func (o GetMonitoringDropletCpuMetricResultOutput) ToGetMonitoringDropletCpuMetricResultOutputWithContext(ctx context.Context) GetMonitoringDropletCpuMetricResultOutput {
	return o
}

func (o GetMonitoringDropletCpuMetricResultOutput) Items() MetricsOutput {
	return o.ApplyT(func(v GetMonitoringDropletCpuMetricResult) Metrics { return v.Items }).(MetricsOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMonitoringDropletCpuMetricResultOutput{})
}
