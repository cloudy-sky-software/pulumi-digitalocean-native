// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetMonitoringDropletLoad15Metric(ctx *pulumi.Context, args *GetMonitoringDropletLoad15MetricArgs, opts ...pulumi.InvokeOption) (*GetMonitoringDropletLoad15MetricResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetMonitoringDropletLoad15MetricResult
	err := ctx.Invoke("digitalocean-native:monitoring/v2:getMonitoringDropletLoad15Metric", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetMonitoringDropletLoad15MetricArgs struct {
}

type GetMonitoringDropletLoad15MetricResult struct {
	Data   MetricsData   `pulumi:"data"`
	Status MetricsStatus `pulumi:"status"`
}

func GetMonitoringDropletLoad15MetricOutput(ctx *pulumi.Context, args GetMonitoringDropletLoad15MetricOutputArgs, opts ...pulumi.InvokeOption) GetMonitoringDropletLoad15MetricResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetMonitoringDropletLoad15MetricResultOutput, error) {
			args := v.(GetMonitoringDropletLoad15MetricArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("digitalocean-native:monitoring/v2:getMonitoringDropletLoad15Metric", args, GetMonitoringDropletLoad15MetricResultOutput{}, options).(GetMonitoringDropletLoad15MetricResultOutput), nil
		}).(GetMonitoringDropletLoad15MetricResultOutput)
}

type GetMonitoringDropletLoad15MetricOutputArgs struct {
}

func (GetMonitoringDropletLoad15MetricOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMonitoringDropletLoad15MetricArgs)(nil)).Elem()
}

type GetMonitoringDropletLoad15MetricResultOutput struct{ *pulumi.OutputState }

func (GetMonitoringDropletLoad15MetricResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMonitoringDropletLoad15MetricResult)(nil)).Elem()
}

func (o GetMonitoringDropletLoad15MetricResultOutput) ToGetMonitoringDropletLoad15MetricResultOutput() GetMonitoringDropletLoad15MetricResultOutput {
	return o
}

func (o GetMonitoringDropletLoad15MetricResultOutput) ToGetMonitoringDropletLoad15MetricResultOutputWithContext(ctx context.Context) GetMonitoringDropletLoad15MetricResultOutput {
	return o
}

func (o GetMonitoringDropletLoad15MetricResultOutput) Data() MetricsDataOutput {
	return o.ApplyT(func(v GetMonitoringDropletLoad15MetricResult) MetricsData { return v.Data }).(MetricsDataOutput)
}

func (o GetMonitoringDropletLoad15MetricResultOutput) Status() MetricsStatusOutput {
	return o.ApplyT(func(v GetMonitoringDropletLoad15MetricResult) MetricsStatus { return v.Status }).(MetricsStatusOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMonitoringDropletLoad15MetricResultOutput{})
}
