// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetAppsMetricsBandwidthDaily(ctx *pulumi.Context, args *GetAppsMetricsBandwidthDailyArgs, opts ...pulumi.InvokeOption) (*GetAppsMetricsBandwidthDailyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAppsMetricsBandwidthDailyResult
	err := ctx.Invoke("digitalocean-native:apps/v2:getAppsMetricsBandwidthDaily", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetAppsMetricsBandwidthDailyArgs struct {
	// The app ID
	AppId string `pulumi:"appId"`
}

type GetAppsMetricsBandwidthDailyResult struct {
	Items AppMetricsBandwidthUsage `pulumi:"items"`
}

func GetAppsMetricsBandwidthDailyOutput(ctx *pulumi.Context, args GetAppsMetricsBandwidthDailyOutputArgs, opts ...pulumi.InvokeOption) GetAppsMetricsBandwidthDailyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppsMetricsBandwidthDailyResult, error) {
			args := v.(GetAppsMetricsBandwidthDailyArgs)
			r, err := GetAppsMetricsBandwidthDaily(ctx, &args, opts...)
			var s GetAppsMetricsBandwidthDailyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAppsMetricsBandwidthDailyResultOutput)
}

type GetAppsMetricsBandwidthDailyOutputArgs struct {
	// The app ID
	AppId pulumi.StringInput `pulumi:"appId"`
}

func (GetAppsMetricsBandwidthDailyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppsMetricsBandwidthDailyArgs)(nil)).Elem()
}

type GetAppsMetricsBandwidthDailyResultOutput struct{ *pulumi.OutputState }

func (GetAppsMetricsBandwidthDailyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppsMetricsBandwidthDailyResult)(nil)).Elem()
}

func (o GetAppsMetricsBandwidthDailyResultOutput) ToGetAppsMetricsBandwidthDailyResultOutput() GetAppsMetricsBandwidthDailyResultOutput {
	return o
}

func (o GetAppsMetricsBandwidthDailyResultOutput) ToGetAppsMetricsBandwidthDailyResultOutputWithContext(ctx context.Context) GetAppsMetricsBandwidthDailyResultOutput {
	return o
}

func (o GetAppsMetricsBandwidthDailyResultOutput) Items() AppMetricsBandwidthUsageOutput {
	return o.ApplyT(func(v GetAppsMetricsBandwidthDailyResult) AppMetricsBandwidthUsage { return v.Items }).(AppMetricsBandwidthUsageOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppsMetricsBandwidthDailyResultOutput{})
}