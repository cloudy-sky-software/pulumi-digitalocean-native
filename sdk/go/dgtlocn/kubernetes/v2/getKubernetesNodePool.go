// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKubernetesNodePool(ctx *pulumi.Context, args *LookupKubernetesNodePoolArgs, opts ...pulumi.InvokeOption) (*LookupKubernetesNodePoolResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupKubernetesNodePoolResult
	err := ctx.Invoke("digitalocean-native:kubernetes/v2:getKubernetesNodePool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKubernetesNodePoolArgs struct {
	// A unique ID that can be used to reference a Kubernetes cluster.
	ClusterId string `pulumi:"clusterId"`
	// A unique ID that can be used to reference a Kubernetes node pool.
	NodePoolId string `pulumi:"nodePoolId"`
}

type LookupKubernetesNodePoolResult struct {
	NodePool *KubernetesNodePoolType `pulumi:"nodePool"`
}

func LookupKubernetesNodePoolOutput(ctx *pulumi.Context, args LookupKubernetesNodePoolOutputArgs, opts ...pulumi.InvokeOption) LookupKubernetesNodePoolResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupKubernetesNodePoolResultOutput, error) {
			args := v.(LookupKubernetesNodePoolArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("digitalocean-native:kubernetes/v2:getKubernetesNodePool", args, LookupKubernetesNodePoolResultOutput{}, options).(LookupKubernetesNodePoolResultOutput), nil
		}).(LookupKubernetesNodePoolResultOutput)
}

type LookupKubernetesNodePoolOutputArgs struct {
	// A unique ID that can be used to reference a Kubernetes cluster.
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// A unique ID that can be used to reference a Kubernetes node pool.
	NodePoolId pulumi.StringInput `pulumi:"nodePoolId"`
}

func (LookupKubernetesNodePoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKubernetesNodePoolArgs)(nil)).Elem()
}

type LookupKubernetesNodePoolResultOutput struct{ *pulumi.OutputState }

func (LookupKubernetesNodePoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKubernetesNodePoolResult)(nil)).Elem()
}

func (o LookupKubernetesNodePoolResultOutput) ToLookupKubernetesNodePoolResultOutput() LookupKubernetesNodePoolResultOutput {
	return o
}

func (o LookupKubernetesNodePoolResultOutput) ToLookupKubernetesNodePoolResultOutputWithContext(ctx context.Context) LookupKubernetesNodePoolResultOutput {
	return o
}

func (o LookupKubernetesNodePoolResultOutput) NodePool() KubernetesNodePoolTypePtrOutput {
	return o.ApplyT(func(v LookupKubernetesNodePoolResult) *KubernetesNodePoolType { return v.NodePool }).(KubernetesNodePoolTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKubernetesNodePoolResultOutput{})
}
