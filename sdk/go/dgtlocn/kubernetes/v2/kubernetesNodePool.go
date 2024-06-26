// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type KubernetesNodePool struct {
	pulumi.CustomResourceState

	// A boolean value indicating whether auto-scaling is enabled for this node pool.
	AutoScale pulumi.BoolPtrOutput `pulumi:"autoScale"`
	// The number of Droplet instances in the node pool.
	Count pulumi.IntOutput `pulumi:"count"`
	// An object of key/value mappings specifying labels to apply to all nodes in a pool. Labels will automatically be applied to all existing nodes and any subsequent nodes added to the pool. Note that when a label is removed, it is not deleted from the nodes in the pool.
	Labels pulumi.AnyOutput `pulumi:"labels"`
	// The maximum number of nodes that this node pool can be auto-scaled to. The value will be `0` if `auto_scale` is set to `false`.
	MaxNodes pulumi.IntPtrOutput `pulumi:"maxNodes"`
	// The minimum number of nodes that this node pool can be auto-scaled to. The value will be `0` if `auto_scale` is set to `false`.
	MinNodes pulumi.IntPtrOutput `pulumi:"minNodes"`
	// A human-readable name for the node pool.
	Name     pulumi.StringOutput             `pulumi:"name"`
	NodePool KubernetesNodePoolTypePtrOutput `pulumi:"nodePool"`
	// An object specifying the details of a specific worker node in a node pool.
	Nodes NodeArrayOutput `pulumi:"nodes"`
	// The slug identifier for the type of Droplet used as workers in the node pool.
	Size pulumi.StringOutput `pulumi:"size"`
	// An array containing the tags applied to the node pool. All node pools are automatically tagged `k8s`, `k8s-worker`, and `k8s:$K8S_CLUSTER_ID`.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// An array of taints to apply to all nodes in a pool. Taints will automatically be applied to all existing nodes and any subsequent nodes added to the pool. When a taint is removed, it is deleted from all nodes in the pool.
	Taints KubernetesNodePoolTaintArrayOutput `pulumi:"taints"`
}

// NewKubernetesNodePool registers a new resource with the given unique name, arguments, and options.
func NewKubernetesNodePool(ctx *pulumi.Context,
	name string, args *KubernetesNodePoolArgs, opts ...pulumi.ResourceOption) (*KubernetesNodePool, error) {
	if args == nil {
		args = &KubernetesNodePoolArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource KubernetesNodePool
	err := ctx.RegisterResource("digitalocean-native:kubernetes/v2:KubernetesNodePool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKubernetesNodePool gets an existing KubernetesNodePool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKubernetesNodePool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KubernetesNodePoolState, opts ...pulumi.ResourceOption) (*KubernetesNodePool, error) {
	var resource KubernetesNodePool
	err := ctx.ReadResource("digitalocean-native:kubernetes/v2:KubernetesNodePool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KubernetesNodePool resources.
type kubernetesNodePoolState struct {
}

type KubernetesNodePoolState struct {
}

func (KubernetesNodePoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesNodePoolState)(nil)).Elem()
}

type kubernetesNodePoolArgs struct {
	// A boolean value indicating whether auto-scaling is enabled for this node pool.
	AutoScale *bool `pulumi:"autoScale"`
	// A unique ID that can be used to reference a Kubernetes cluster.
	ClusterId *string `pulumi:"clusterId"`
	// The number of Droplet instances in the node pool.
	Count *int `pulumi:"count"`
	// An object of key/value mappings specifying labels to apply to all nodes in a pool. Labels will automatically be applied to all existing nodes and any subsequent nodes added to the pool. Note that when a label is removed, it is not deleted from the nodes in the pool.
	Labels interface{} `pulumi:"labels"`
	// The maximum number of nodes that this node pool can be auto-scaled to. The value will be `0` if `auto_scale` is set to `false`.
	MaxNodes *int `pulumi:"maxNodes"`
	// The minimum number of nodes that this node pool can be auto-scaled to. The value will be `0` if `auto_scale` is set to `false`.
	MinNodes *int `pulumi:"minNodes"`
	// A human-readable name for the node pool.
	Name *string `pulumi:"name"`
	// An object specifying the details of a specific worker node in a node pool.
	Nodes []Node `pulumi:"nodes"`
	// The slug identifier for the type of Droplet used as workers in the node pool.
	Size *string `pulumi:"size"`
	// An array containing the tags applied to the node pool. All node pools are automatically tagged `k8s`, `k8s-worker`, and `k8s:$K8S_CLUSTER_ID`.
	Tags []string `pulumi:"tags"`
	// An array of taints to apply to all nodes in a pool. Taints will automatically be applied to all existing nodes and any subsequent nodes added to the pool. When a taint is removed, it is deleted from all nodes in the pool.
	Taints []KubernetesNodePoolTaint `pulumi:"taints"`
}

// The set of arguments for constructing a KubernetesNodePool resource.
type KubernetesNodePoolArgs struct {
	// A boolean value indicating whether auto-scaling is enabled for this node pool.
	AutoScale pulumi.BoolPtrInput
	// A unique ID that can be used to reference a Kubernetes cluster.
	ClusterId pulumi.StringPtrInput
	// The number of Droplet instances in the node pool.
	Count pulumi.IntPtrInput
	// An object of key/value mappings specifying labels to apply to all nodes in a pool. Labels will automatically be applied to all existing nodes and any subsequent nodes added to the pool. Note that when a label is removed, it is not deleted from the nodes in the pool.
	Labels pulumi.Input
	// The maximum number of nodes that this node pool can be auto-scaled to. The value will be `0` if `auto_scale` is set to `false`.
	MaxNodes pulumi.IntPtrInput
	// The minimum number of nodes that this node pool can be auto-scaled to. The value will be `0` if `auto_scale` is set to `false`.
	MinNodes pulumi.IntPtrInput
	// A human-readable name for the node pool.
	Name pulumi.StringPtrInput
	// An object specifying the details of a specific worker node in a node pool.
	Nodes NodeArrayInput
	// The slug identifier for the type of Droplet used as workers in the node pool.
	Size pulumi.StringPtrInput
	// An array containing the tags applied to the node pool. All node pools are automatically tagged `k8s`, `k8s-worker`, and `k8s:$K8S_CLUSTER_ID`.
	Tags pulumi.StringArrayInput
	// An array of taints to apply to all nodes in a pool. Taints will automatically be applied to all existing nodes and any subsequent nodes added to the pool. When a taint is removed, it is deleted from all nodes in the pool.
	Taints KubernetesNodePoolTaintArrayInput
}

func (KubernetesNodePoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesNodePoolArgs)(nil)).Elem()
}

type KubernetesNodePoolInput interface {
	pulumi.Input

	ToKubernetesNodePoolOutput() KubernetesNodePoolOutput
	ToKubernetesNodePoolOutputWithContext(ctx context.Context) KubernetesNodePoolOutput
}

func (*KubernetesNodePool) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesNodePool)(nil)).Elem()
}

func (i *KubernetesNodePool) ToKubernetesNodePoolOutput() KubernetesNodePoolOutput {
	return i.ToKubernetesNodePoolOutputWithContext(context.Background())
}

func (i *KubernetesNodePool) ToKubernetesNodePoolOutputWithContext(ctx context.Context) KubernetesNodePoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesNodePoolOutput)
}

type KubernetesNodePoolOutput struct{ *pulumi.OutputState }

func (KubernetesNodePoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesNodePool)(nil)).Elem()
}

func (o KubernetesNodePoolOutput) ToKubernetesNodePoolOutput() KubernetesNodePoolOutput {
	return o
}

func (o KubernetesNodePoolOutput) ToKubernetesNodePoolOutputWithContext(ctx context.Context) KubernetesNodePoolOutput {
	return o
}

// A boolean value indicating whether auto-scaling is enabled for this node pool.
func (o KubernetesNodePoolOutput) AutoScale() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KubernetesNodePool) pulumi.BoolPtrOutput { return v.AutoScale }).(pulumi.BoolPtrOutput)
}

// The number of Droplet instances in the node pool.
func (o KubernetesNodePoolOutput) Count() pulumi.IntOutput {
	return o.ApplyT(func(v *KubernetesNodePool) pulumi.IntOutput { return v.Count }).(pulumi.IntOutput)
}

// An object of key/value mappings specifying labels to apply to all nodes in a pool. Labels will automatically be applied to all existing nodes and any subsequent nodes added to the pool. Note that when a label is removed, it is not deleted from the nodes in the pool.
func (o KubernetesNodePoolOutput) Labels() pulumi.AnyOutput {
	return o.ApplyT(func(v *KubernetesNodePool) pulumi.AnyOutput { return v.Labels }).(pulumi.AnyOutput)
}

// The maximum number of nodes that this node pool can be auto-scaled to. The value will be `0` if `auto_scale` is set to `false`.
func (o KubernetesNodePoolOutput) MaxNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *KubernetesNodePool) pulumi.IntPtrOutput { return v.MaxNodes }).(pulumi.IntPtrOutput)
}

// The minimum number of nodes that this node pool can be auto-scaled to. The value will be `0` if `auto_scale` is set to `false`.
func (o KubernetesNodePoolOutput) MinNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *KubernetesNodePool) pulumi.IntPtrOutput { return v.MinNodes }).(pulumi.IntPtrOutput)
}

// A human-readable name for the node pool.
func (o KubernetesNodePoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesNodePool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o KubernetesNodePoolOutput) NodePool() KubernetesNodePoolTypePtrOutput {
	return o.ApplyT(func(v *KubernetesNodePool) KubernetesNodePoolTypePtrOutput { return v.NodePool }).(KubernetesNodePoolTypePtrOutput)
}

// An object specifying the details of a specific worker node in a node pool.
func (o KubernetesNodePoolOutput) Nodes() NodeArrayOutput {
	return o.ApplyT(func(v *KubernetesNodePool) NodeArrayOutput { return v.Nodes }).(NodeArrayOutput)
}

// The slug identifier for the type of Droplet used as workers in the node pool.
func (o KubernetesNodePoolOutput) Size() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesNodePool) pulumi.StringOutput { return v.Size }).(pulumi.StringOutput)
}

// An array containing the tags applied to the node pool. All node pools are automatically tagged `k8s`, `k8s-worker`, and `k8s:$K8S_CLUSTER_ID`.
func (o KubernetesNodePoolOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *KubernetesNodePool) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// An array of taints to apply to all nodes in a pool. Taints will automatically be applied to all existing nodes and any subsequent nodes added to the pool. When a taint is removed, it is deleted from all nodes in the pool.
func (o KubernetesNodePoolOutput) Taints() KubernetesNodePoolTaintArrayOutput {
	return o.ApplyT(func(v *KubernetesNodePool) KubernetesNodePoolTaintArrayOutput { return v.Taints }).(KubernetesNodePoolTaintArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesNodePoolInput)(nil)).Elem(), &KubernetesNodePool{})
	pulumi.RegisterOutputType(KubernetesNodePoolOutput{})
}
