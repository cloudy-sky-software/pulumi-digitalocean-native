// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListKubernetesNodePools(ctx *pulumi.Context, args *ListKubernetesNodePoolsArgs, opts ...pulumi.InvokeOption) (*ListKubernetesNodePoolsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListKubernetesNodePoolsResult
	err := ctx.Invoke("digitalocean-native:kubernetes/v2:listKubernetesNodePools", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListKubernetesNodePoolsArgs struct {
	// A unique ID that can be used to reference a Kubernetes cluster.
	ClusterId string `pulumi:"clusterId"`
}

type ListKubernetesNodePoolsResult struct {
	NodePools []KubernetesNodePoolType `pulumi:"nodePools"`
}

func ListKubernetesNodePoolsOutput(ctx *pulumi.Context, args ListKubernetesNodePoolsOutputArgs, opts ...pulumi.InvokeOption) ListKubernetesNodePoolsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListKubernetesNodePoolsResult, error) {
			args := v.(ListKubernetesNodePoolsArgs)
			r, err := ListKubernetesNodePools(ctx, &args, opts...)
			var s ListKubernetesNodePoolsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListKubernetesNodePoolsResultOutput)
}

type ListKubernetesNodePoolsOutputArgs struct {
	// A unique ID that can be used to reference a Kubernetes cluster.
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
}

func (ListKubernetesNodePoolsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListKubernetesNodePoolsArgs)(nil)).Elem()
}

type ListKubernetesNodePoolsResultOutput struct{ *pulumi.OutputState }

func (ListKubernetesNodePoolsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListKubernetesNodePoolsResult)(nil)).Elem()
}

func (o ListKubernetesNodePoolsResultOutput) ToListKubernetesNodePoolsResultOutput() ListKubernetesNodePoolsResultOutput {
	return o
}

func (o ListKubernetesNodePoolsResultOutput) ToListKubernetesNodePoolsResultOutputWithContext(ctx context.Context) ListKubernetesNodePoolsResultOutput {
	return o
}

func (o ListKubernetesNodePoolsResultOutput) NodePools() KubernetesNodePoolTypeArrayOutput {
	return o.ApplyT(func(v ListKubernetesNodePoolsResult) []KubernetesNodePoolType { return v.NodePools }).(KubernetesNodePoolTypeArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListKubernetesNodePoolsResultOutput{})
}
