// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type KubernetesRegistry struct {
	pulumi.CustomResourceState

	// An array containing the UUIDs of Kubernetes clusters.
	ClusterUuids pulumi.StringArrayOutput `pulumi:"clusterUuids"`
}

// NewKubernetesRegistry registers a new resource with the given unique name, arguments, and options.
func NewKubernetesRegistry(ctx *pulumi.Context,
	name string, args *KubernetesRegistryArgs, opts ...pulumi.ResourceOption) (*KubernetesRegistry, error) {
	if args == nil {
		args = &KubernetesRegistryArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource KubernetesRegistry
	err := ctx.RegisterResource("digitalocean-native:kubernetes/v2:KubernetesRegistry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKubernetesRegistry gets an existing KubernetesRegistry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKubernetesRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KubernetesRegistryState, opts ...pulumi.ResourceOption) (*KubernetesRegistry, error) {
	var resource KubernetesRegistry
	err := ctx.ReadResource("digitalocean-native:kubernetes/v2:KubernetesRegistry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KubernetesRegistry resources.
type kubernetesRegistryState struct {
}

type KubernetesRegistryState struct {
}

func (KubernetesRegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesRegistryState)(nil)).Elem()
}

type kubernetesRegistryArgs struct {
	// An array containing the UUIDs of Kubernetes clusters.
	ClusterUuids []string `pulumi:"clusterUuids"`
}

// The set of arguments for constructing a KubernetesRegistry resource.
type KubernetesRegistryArgs struct {
	// An array containing the UUIDs of Kubernetes clusters.
	ClusterUuids pulumi.StringArrayInput
}

func (KubernetesRegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesRegistryArgs)(nil)).Elem()
}

type KubernetesRegistryInput interface {
	pulumi.Input

	ToKubernetesRegistryOutput() KubernetesRegistryOutput
	ToKubernetesRegistryOutputWithContext(ctx context.Context) KubernetesRegistryOutput
}

func (*KubernetesRegistry) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesRegistry)(nil)).Elem()
}

func (i *KubernetesRegistry) ToKubernetesRegistryOutput() KubernetesRegistryOutput {
	return i.ToKubernetesRegistryOutputWithContext(context.Background())
}

func (i *KubernetesRegistry) ToKubernetesRegistryOutputWithContext(ctx context.Context) KubernetesRegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesRegistryOutput)
}

type KubernetesRegistryOutput struct{ *pulumi.OutputState }

func (KubernetesRegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesRegistry)(nil)).Elem()
}

func (o KubernetesRegistryOutput) ToKubernetesRegistryOutput() KubernetesRegistryOutput {
	return o
}

func (o KubernetesRegistryOutput) ToKubernetesRegistryOutputWithContext(ctx context.Context) KubernetesRegistryOutput {
	return o
}

// An array containing the UUIDs of Kubernetes clusters.
func (o KubernetesRegistryOutput) ClusterUuids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *KubernetesRegistry) pulumi.StringArrayOutput { return v.ClusterUuids }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesRegistryInput)(nil)).Elem(), &KubernetesRegistry{})
	pulumi.RegisterOutputType(KubernetesRegistryOutput{})
}