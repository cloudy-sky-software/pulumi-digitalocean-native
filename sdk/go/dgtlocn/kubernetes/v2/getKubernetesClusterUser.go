// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetKubernetesClusterUser(ctx *pulumi.Context, args *GetKubernetesClusterUserArgs, opts ...pulumi.InvokeOption) (*GetKubernetesClusterUserResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetKubernetesClusterUserResult
	err := ctx.Invoke("digitalocean-native:kubernetes/v2:getKubernetesClusterUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetKubernetesClusterUserArgs struct {
	// A unique ID that can be used to reference a Kubernetes cluster.
	ClusterId string `pulumi:"clusterId"`
}

type GetKubernetesClusterUserResult struct {
	KubernetesClusterUser *UserKubernetesClusterUserProperties `pulumi:"kubernetesClusterUser"`
}

func GetKubernetesClusterUserOutput(ctx *pulumi.Context, args GetKubernetesClusterUserOutputArgs, opts ...pulumi.InvokeOption) GetKubernetesClusterUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetKubernetesClusterUserResult, error) {
			args := v.(GetKubernetesClusterUserArgs)
			r, err := GetKubernetesClusterUser(ctx, &args, opts...)
			var s GetKubernetesClusterUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetKubernetesClusterUserResultOutput)
}

type GetKubernetesClusterUserOutputArgs struct {
	// A unique ID that can be used to reference a Kubernetes cluster.
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
}

func (GetKubernetesClusterUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKubernetesClusterUserArgs)(nil)).Elem()
}

type GetKubernetesClusterUserResultOutput struct{ *pulumi.OutputState }

func (GetKubernetesClusterUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKubernetesClusterUserResult)(nil)).Elem()
}

func (o GetKubernetesClusterUserResultOutput) ToGetKubernetesClusterUserResultOutput() GetKubernetesClusterUserResultOutput {
	return o
}

func (o GetKubernetesClusterUserResultOutput) ToGetKubernetesClusterUserResultOutputWithContext(ctx context.Context) GetKubernetesClusterUserResultOutput {
	return o
}

func (o GetKubernetesClusterUserResultOutput) KubernetesClusterUser() UserKubernetesClusterUserPropertiesPtrOutput {
	return o.ApplyT(func(v GetKubernetesClusterUserResult) *UserKubernetesClusterUserProperties {
		return v.KubernetesClusterUser
	}).(UserKubernetesClusterUserPropertiesPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetKubernetesClusterUserResultOutput{})
}
