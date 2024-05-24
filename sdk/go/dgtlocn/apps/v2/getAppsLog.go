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
	Items AppsGetLogsResponse `pulumi:"items"`
}

func GetAppsLogOutput(ctx *pulumi.Context, args GetAppsLogOutputArgs, opts ...pulumi.InvokeOption) GetAppsLogResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppsLogResult, error) {
			args := v.(GetAppsLogArgs)
			r, err := GetAppsLog(ctx, &args, opts...)
			var s GetAppsLogResult
			if r != nil {
				s = *r
			}
			return s, err
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

func (o GetAppsLogResultOutput) Items() AppsGetLogsResponseOutput {
	return o.ApplyT(func(v GetAppsLogResult) AppsGetLogsResponse { return v.Items }).(AppsGetLogsResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppsLogResultOutput{})
}