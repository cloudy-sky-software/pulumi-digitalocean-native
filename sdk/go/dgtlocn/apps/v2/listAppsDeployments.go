// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListAppsDeployments(ctx *pulumi.Context, args *ListAppsDeploymentsArgs, opts ...pulumi.InvokeOption) (*ListAppsDeploymentsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListAppsDeploymentsResult
	err := ctx.Invoke("digitalocean-native:apps/v2:listAppsDeployments", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAppsDeploymentsArgs struct {
	// The app ID
	AppId string `pulumi:"appId"`
}

type ListAppsDeploymentsResult struct {
	Deployments []AppsDeploymentType `pulumi:"deployments"`
	Links       *PageLinks           `pulumi:"links"`
	Meta        MetaMeta             `pulumi:"meta"`
}

func ListAppsDeploymentsOutput(ctx *pulumi.Context, args ListAppsDeploymentsOutputArgs, opts ...pulumi.InvokeOption) ListAppsDeploymentsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListAppsDeploymentsResultOutput, error) {
			args := v.(ListAppsDeploymentsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("digitalocean-native:apps/v2:listAppsDeployments", args, ListAppsDeploymentsResultOutput{}, options).(ListAppsDeploymentsResultOutput), nil
		}).(ListAppsDeploymentsResultOutput)
}

type ListAppsDeploymentsOutputArgs struct {
	// The app ID
	AppId pulumi.StringInput `pulumi:"appId"`
}

func (ListAppsDeploymentsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAppsDeploymentsArgs)(nil)).Elem()
}

type ListAppsDeploymentsResultOutput struct{ *pulumi.OutputState }

func (ListAppsDeploymentsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAppsDeploymentsResult)(nil)).Elem()
}

func (o ListAppsDeploymentsResultOutput) ToListAppsDeploymentsResultOutput() ListAppsDeploymentsResultOutput {
	return o
}

func (o ListAppsDeploymentsResultOutput) ToListAppsDeploymentsResultOutputWithContext(ctx context.Context) ListAppsDeploymentsResultOutput {
	return o
}

func (o ListAppsDeploymentsResultOutput) Deployments() AppsDeploymentTypeArrayOutput {
	return o.ApplyT(func(v ListAppsDeploymentsResult) []AppsDeploymentType { return v.Deployments }).(AppsDeploymentTypeArrayOutput)
}

func (o ListAppsDeploymentsResultOutput) Links() PageLinksPtrOutput {
	return o.ApplyT(func(v ListAppsDeploymentsResult) *PageLinks { return v.Links }).(PageLinksPtrOutput)
}

func (o ListAppsDeploymentsResultOutput) Meta() MetaMetaOutput {
	return o.ApplyT(func(v ListAppsDeploymentsResult) MetaMeta { return v.Meta }).(MetaMetaOutput)
}

func init() {
	pulumi.RegisterOutputType(ListAppsDeploymentsResultOutput{})
}
