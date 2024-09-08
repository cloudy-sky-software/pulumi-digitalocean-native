// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKubernetesCluster(ctx *pulumi.Context, args *LookupKubernetesClusterArgs, opts ...pulumi.InvokeOption) (*LookupKubernetesClusterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupKubernetesClusterResult
	err := ctx.Invoke("digitalocean-native:kubernetes/v2:getKubernetesCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupKubernetesClusterArgs struct {
	// A unique ID that can be used to reference a Kubernetes cluster.
	ClusterId string `pulumi:"clusterId"`
}

type LookupKubernetesClusterResult struct {
	KubernetesCluster *Cluster `pulumi:"kubernetesCluster"`
}

// Defaults sets the appropriate defaults for LookupKubernetesClusterResult
func (val *LookupKubernetesClusterResult) Defaults() *LookupKubernetesClusterResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.KubernetesCluster = tmp.KubernetesCluster.Defaults()

	return &tmp
}

func LookupKubernetesClusterOutput(ctx *pulumi.Context, args LookupKubernetesClusterOutputArgs, opts ...pulumi.InvokeOption) LookupKubernetesClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKubernetesClusterResult, error) {
			args := v.(LookupKubernetesClusterArgs)
			r, err := LookupKubernetesCluster(ctx, &args, opts...)
			var s LookupKubernetesClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKubernetesClusterResultOutput)
}

type LookupKubernetesClusterOutputArgs struct {
	// A unique ID that can be used to reference a Kubernetes cluster.
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
}

func (LookupKubernetesClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKubernetesClusterArgs)(nil)).Elem()
}

type LookupKubernetesClusterResultOutput struct{ *pulumi.OutputState }

func (LookupKubernetesClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKubernetesClusterResult)(nil)).Elem()
}

func (o LookupKubernetesClusterResultOutput) ToLookupKubernetesClusterResultOutput() LookupKubernetesClusterResultOutput {
	return o
}

func (o LookupKubernetesClusterResultOutput) ToLookupKubernetesClusterResultOutputWithContext(ctx context.Context) LookupKubernetesClusterResultOutput {
	return o
}

func (o LookupKubernetesClusterResultOutput) KubernetesCluster() ClusterPtrOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) *Cluster { return v.KubernetesCluster }).(ClusterPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKubernetesClusterResultOutput{})
}
