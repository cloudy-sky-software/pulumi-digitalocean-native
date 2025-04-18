// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetMonitoringDropletLoad5Metric(ctx *pulumi.Context, args *GetMonitoringDropletLoad5MetricArgs, opts ...pulumi.InvokeOption) (*GetMonitoringDropletLoad5MetricResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetMonitoringDropletLoad5MetricResult
	err := ctx.Invoke("digitalocean-native:monitoring/v2:getMonitoringDropletLoad5Metric", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetMonitoringDropletLoad5MetricArgs struct {
}

type GetMonitoringDropletLoad5MetricResult struct {
	Data   MetricsData   `pulumi:"data"`
	Status MetricsStatus `pulumi:"status"`
}

func GetMonitoringDropletLoad5MetricOutput(ctx *pulumi.Context, args GetMonitoringDropletLoad5MetricOutputArgs, opts ...pulumi.InvokeOption) GetMonitoringDropletLoad5MetricResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetMonitoringDropletLoad5MetricResultOutput, error) {
			args := v.(GetMonitoringDropletLoad5MetricArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("digitalocean-native:monitoring/v2:getMonitoringDropletLoad5Metric", args, GetMonitoringDropletLoad5MetricResultOutput{}, options).(GetMonitoringDropletLoad5MetricResultOutput), nil
		}).(GetMonitoringDropletLoad5MetricResultOutput)
}

type GetMonitoringDropletLoad5MetricOutputArgs struct {
}

func (GetMonitoringDropletLoad5MetricOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMonitoringDropletLoad5MetricArgs)(nil)).Elem()
}

type GetMonitoringDropletLoad5MetricResultOutput struct{ *pulumi.OutputState }

func (GetMonitoringDropletLoad5MetricResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMonitoringDropletLoad5MetricResult)(nil)).Elem()
}

func (o GetMonitoringDropletLoad5MetricResultOutput) ToGetMonitoringDropletLoad5MetricResultOutput() GetMonitoringDropletLoad5MetricResultOutput {
	return o
}

func (o GetMonitoringDropletLoad5MetricResultOutput) ToGetMonitoringDropletLoad5MetricResultOutputWithContext(ctx context.Context) GetMonitoringDropletLoad5MetricResultOutput {
	return o
}

func (o GetMonitoringDropletLoad5MetricResultOutput) Data() MetricsDataOutput {
	return o.ApplyT(func(v GetMonitoringDropletLoad5MetricResult) MetricsData { return v.Data }).(MetricsDataOutput)
}

func (o GetMonitoringDropletLoad5MetricResultOutput) Status() MetricsStatusOutput {
	return o.ApplyT(func(v GetMonitoringDropletLoad5MetricResult) MetricsStatus { return v.Status }).(MetricsStatusOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMonitoringDropletLoad5MetricResultOutput{})
}
