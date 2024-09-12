// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetAppsLog(ctx *pulumi.Context, args *GetAppsLogArgs, opts ...pulumi.InvokeOption) (*GetAppsLogResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAppsLogResult
	err := ctx.Invoke("digitalocean-native:apps/v2:getAppsLog", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetAppsLogArgs struct {
	// The app ID
	AppId string `pulumi:"appId"`
	// An optional component name. If set, logs will be limited to this component only.
	ComponentName string `pulumi:"componentName"`
	// The deployment ID
	DeploymentId string `pulumi:"deploymentId"`
}

type GetAppsLogResult struct {
	HistoricUrls []string `pulumi:"historicUrls"`
	// A URL of the real-time live logs. This URL may use either the `https://` or `wss://` protocols and will keep pushing live logs as they become available.
	LiveUrl *string `pulumi:"liveUrl"`
}

func GetAppsLogOutput(ctx *pulumi.Context, args GetAppsLogOutputArgs, opts ...pulumi.InvokeOption) GetAppsLogResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppsLogResultOutput, error) {
			args := v.(GetAppsLogArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetAppsLogResult
			secret, err := ctx.InvokePackageRaw("digitalocean-native:apps/v2:getAppsLog", args, &rv, "", opts...)
			if err != nil {
				return GetAppsLogResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetAppsLogResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetAppsLogResultOutput), nil
			}
			return output, nil
		}).(GetAppsLogResultOutput)
}

type GetAppsLogOutputArgs struct {
	// The app ID
	AppId pulumi.StringInput `pulumi:"appId"`
	// An optional component name. If set, logs will be limited to this component only.
	ComponentName pulumi.StringInput `pulumi:"componentName"`
	// The deployment ID
	DeploymentId pulumi.StringInput `pulumi:"deploymentId"`
}

func (GetAppsLogOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppsLogArgs)(nil)).Elem()
}

type GetAppsLogResultOutput struct{ *pulumi.OutputState }

func (GetAppsLogResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppsLogResult)(nil)).Elem()
}

func (o GetAppsLogResultOutput) ToGetAppsLogResultOutput() GetAppsLogResultOutput {
	return o
}

func (o GetAppsLogResultOutput) ToGetAppsLogResultOutputWithContext(ctx context.Context) GetAppsLogResultOutput {
	return o
}

func (o GetAppsLogResultOutput) HistoricUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAppsLogResult) []string { return v.HistoricUrls }).(pulumi.StringArrayOutput)
}

// A URL of the real-time live logs. This URL may use either the `https://` or `wss://` protocols and will keep pushing live logs as they become available.
func (o GetAppsLogResultOutput) LiveUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAppsLogResult) *string { return v.LiveUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppsLogResultOutput{})
}
