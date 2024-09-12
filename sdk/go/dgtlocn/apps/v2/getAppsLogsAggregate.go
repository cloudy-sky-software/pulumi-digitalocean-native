// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetAppsLogsAggregate(ctx *pulumi.Context, args *GetAppsLogsAggregateArgs, opts ...pulumi.InvokeOption) (*GetAppsLogsAggregateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAppsLogsAggregateResult
	err := ctx.Invoke("digitalocean-native:apps/v2:getAppsLogsAggregate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetAppsLogsAggregateArgs struct {
	// The app ID
	AppId string `pulumi:"appId"`
	// The deployment ID
	DeploymentId string `pulumi:"deploymentId"`
}

type GetAppsLogsAggregateResult struct {
	HistoricUrls []string `pulumi:"historicUrls"`
	// A URL of the real-time live logs. This URL may use either the `https://` or `wss://` protocols and will keep pushing live logs as they become available.
	LiveUrl *string `pulumi:"liveUrl"`
}

func GetAppsLogsAggregateOutput(ctx *pulumi.Context, args GetAppsLogsAggregateOutputArgs, opts ...pulumi.InvokeOption) GetAppsLogsAggregateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppsLogsAggregateResultOutput, error) {
			args := v.(GetAppsLogsAggregateArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetAppsLogsAggregateResult
			secret, err := ctx.InvokePackageRaw("digitalocean-native:apps/v2:getAppsLogsAggregate", args, &rv, "", opts...)
			if err != nil {
				return GetAppsLogsAggregateResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetAppsLogsAggregateResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetAppsLogsAggregateResultOutput), nil
			}
			return output, nil
		}).(GetAppsLogsAggregateResultOutput)
}

type GetAppsLogsAggregateOutputArgs struct {
	// The app ID
	AppId pulumi.StringInput `pulumi:"appId"`
	// The deployment ID
	DeploymentId pulumi.StringInput `pulumi:"deploymentId"`
}

func (GetAppsLogsAggregateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppsLogsAggregateArgs)(nil)).Elem()
}

type GetAppsLogsAggregateResultOutput struct{ *pulumi.OutputState }

func (GetAppsLogsAggregateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppsLogsAggregateResult)(nil)).Elem()
}

func (o GetAppsLogsAggregateResultOutput) ToGetAppsLogsAggregateResultOutput() GetAppsLogsAggregateResultOutput {
	return o
}

func (o GetAppsLogsAggregateResultOutput) ToGetAppsLogsAggregateResultOutputWithContext(ctx context.Context) GetAppsLogsAggregateResultOutput {
	return o
}

func (o GetAppsLogsAggregateResultOutput) HistoricUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAppsLogsAggregateResult) []string { return v.HistoricUrls }).(pulumi.StringArrayOutput)
}

// A URL of the real-time live logs. This URL may use either the `https://` or `wss://` protocols and will keep pushing live logs as they become available.
func (o GetAppsLogsAggregateResultOutput) LiveUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAppsLogsAggregateResult) *string { return v.LiveUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppsLogsAggregateResultOutput{})
}
