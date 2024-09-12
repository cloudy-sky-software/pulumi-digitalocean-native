// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetMonitoringDropletBandwidthMetric(ctx *pulumi.Context, args *GetMonitoringDropletBandwidthMetricArgs, opts ...pulumi.InvokeOption) (*GetMonitoringDropletBandwidthMetricResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetMonitoringDropletBandwidthMetricResult
	err := ctx.Invoke("digitalocean-native:monitoring/v2:getMonitoringDropletBandwidthMetric", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetMonitoringDropletBandwidthMetricArgs struct {
}

type GetMonitoringDropletBandwidthMetricResult struct {
	Data   MetricsData   `pulumi:"data"`
	Status MetricsStatus `pulumi:"status"`
}

func GetMonitoringDropletBandwidthMetricOutput(ctx *pulumi.Context, args GetMonitoringDropletBandwidthMetricOutputArgs, opts ...pulumi.InvokeOption) GetMonitoringDropletBandwidthMetricResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMonitoringDropletBandwidthMetricResultOutput, error) {
			args := v.(GetMonitoringDropletBandwidthMetricArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetMonitoringDropletBandwidthMetricResult
			secret, err := ctx.InvokePackageRaw("digitalocean-native:monitoring/v2:getMonitoringDropletBandwidthMetric", args, &rv, "", opts...)
			if err != nil {
				return GetMonitoringDropletBandwidthMetricResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetMonitoringDropletBandwidthMetricResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetMonitoringDropletBandwidthMetricResultOutput), nil
			}
			return output, nil
		}).(GetMonitoringDropletBandwidthMetricResultOutput)
}

type GetMonitoringDropletBandwidthMetricOutputArgs struct {
}

func (GetMonitoringDropletBandwidthMetricOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMonitoringDropletBandwidthMetricArgs)(nil)).Elem()
}

type GetMonitoringDropletBandwidthMetricResultOutput struct{ *pulumi.OutputState }

func (GetMonitoringDropletBandwidthMetricResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMonitoringDropletBandwidthMetricResult)(nil)).Elem()
}

func (o GetMonitoringDropletBandwidthMetricResultOutput) ToGetMonitoringDropletBandwidthMetricResultOutput() GetMonitoringDropletBandwidthMetricResultOutput {
	return o
}

func (o GetMonitoringDropletBandwidthMetricResultOutput) ToGetMonitoringDropletBandwidthMetricResultOutputWithContext(ctx context.Context) GetMonitoringDropletBandwidthMetricResultOutput {
	return o
}

func (o GetMonitoringDropletBandwidthMetricResultOutput) Data() MetricsDataOutput {
	return o.ApplyT(func(v GetMonitoringDropletBandwidthMetricResult) MetricsData { return v.Data }).(MetricsDataOutput)
}

func (o GetMonitoringDropletBandwidthMetricResultOutput) Status() MetricsStatusOutput {
	return o.ApplyT(func(v GetMonitoringDropletBandwidthMetricResult) MetricsStatus { return v.Status }).(MetricsStatusOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMonitoringDropletBandwidthMetricResultOutput{})
}
