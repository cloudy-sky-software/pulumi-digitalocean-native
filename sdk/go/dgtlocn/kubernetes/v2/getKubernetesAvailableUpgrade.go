// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetKubernetesAvailableUpgrade(ctx *pulumi.Context, args *GetKubernetesAvailableUpgradeArgs, opts ...pulumi.InvokeOption) (*GetKubernetesAvailableUpgradeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetKubernetesAvailableUpgradeResult
	err := ctx.Invoke("digitalocean-native:kubernetes/v2:getKubernetesAvailableUpgrade", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetKubernetesAvailableUpgradeArgs struct {
	// A unique ID that can be used to reference a Kubernetes cluster.
	ClusterId string `pulumi:"clusterId"`
}

type GetKubernetesAvailableUpgradeResult struct {
	Items GetKubernetesAvailableUpgradeProperties `pulumi:"items"`
}

func GetKubernetesAvailableUpgradeOutput(ctx *pulumi.Context, args GetKubernetesAvailableUpgradeOutputArgs, opts ...pulumi.InvokeOption) GetKubernetesAvailableUpgradeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetKubernetesAvailableUpgradeResult, error) {
			args := v.(GetKubernetesAvailableUpgradeArgs)
			r, err := GetKubernetesAvailableUpgrade(ctx, &args, opts...)
			var s GetKubernetesAvailableUpgradeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetKubernetesAvailableUpgradeResultOutput)
}

type GetKubernetesAvailableUpgradeOutputArgs struct {
	// A unique ID that can be used to reference a Kubernetes cluster.
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
}

func (GetKubernetesAvailableUpgradeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKubernetesAvailableUpgradeArgs)(nil)).Elem()
}

type GetKubernetesAvailableUpgradeResultOutput struct{ *pulumi.OutputState }

func (GetKubernetesAvailableUpgradeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKubernetesAvailableUpgradeResult)(nil)).Elem()
}

func (o GetKubernetesAvailableUpgradeResultOutput) ToGetKubernetesAvailableUpgradeResultOutput() GetKubernetesAvailableUpgradeResultOutput {
	return o
}

func (o GetKubernetesAvailableUpgradeResultOutput) ToGetKubernetesAvailableUpgradeResultOutputWithContext(ctx context.Context) GetKubernetesAvailableUpgradeResultOutput {
	return o
}

func (o GetKubernetesAvailableUpgradeResultOutput) Items() GetKubernetesAvailableUpgradePropertiesOutput {
	return o.ApplyT(func(v GetKubernetesAvailableUpgradeResult) GetKubernetesAvailableUpgradeProperties { return v.Items }).(GetKubernetesAvailableUpgradePropertiesOutput)
}

func init() {
	pulumi.RegisterOutputType(GetKubernetesAvailableUpgradeResultOutput{})
}