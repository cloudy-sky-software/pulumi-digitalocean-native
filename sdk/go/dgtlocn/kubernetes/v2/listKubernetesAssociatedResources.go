// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListKubernetesAssociatedResources(ctx *pulumi.Context, args *ListKubernetesAssociatedResourcesArgs, opts ...pulumi.InvokeOption) (*ListKubernetesAssociatedResourcesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListKubernetesAssociatedResourcesResult
	err := ctx.Invoke("digitalocean-native:kubernetes/v2:listKubernetesAssociatedResources", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListKubernetesAssociatedResourcesArgs struct {
	// A unique ID that can be used to reference a Kubernetes cluster.
	ClusterId string `pulumi:"clusterId"`
}

type ListKubernetesAssociatedResourcesResult struct {
	Items AssociatedKubernetesResources `pulumi:"items"`
}

func ListKubernetesAssociatedResourcesOutput(ctx *pulumi.Context, args ListKubernetesAssociatedResourcesOutputArgs, opts ...pulumi.InvokeOption) ListKubernetesAssociatedResourcesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListKubernetesAssociatedResourcesResult, error) {
			args := v.(ListKubernetesAssociatedResourcesArgs)
			r, err := ListKubernetesAssociatedResources(ctx, &args, opts...)
			var s ListKubernetesAssociatedResourcesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListKubernetesAssociatedResourcesResultOutput)
}

type ListKubernetesAssociatedResourcesOutputArgs struct {
	// A unique ID that can be used to reference a Kubernetes cluster.
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
}

func (ListKubernetesAssociatedResourcesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListKubernetesAssociatedResourcesArgs)(nil)).Elem()
}

type ListKubernetesAssociatedResourcesResultOutput struct{ *pulumi.OutputState }

func (ListKubernetesAssociatedResourcesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListKubernetesAssociatedResourcesResult)(nil)).Elem()
}

func (o ListKubernetesAssociatedResourcesResultOutput) ToListKubernetesAssociatedResourcesResultOutput() ListKubernetesAssociatedResourcesResultOutput {
	return o
}

func (o ListKubernetesAssociatedResourcesResultOutput) ToListKubernetesAssociatedResourcesResultOutputWithContext(ctx context.Context) ListKubernetesAssociatedResourcesResultOutput {
	return o
}

func (o ListKubernetesAssociatedResourcesResultOutput) Items() AssociatedKubernetesResourcesOutput {
	return o.ApplyT(func(v ListKubernetesAssociatedResourcesResult) AssociatedKubernetesResources { return v.Items }).(AssociatedKubernetesResourcesOutput)
}

func init() {
	pulumi.RegisterOutputType(ListKubernetesAssociatedResourcesResultOutput{})
}