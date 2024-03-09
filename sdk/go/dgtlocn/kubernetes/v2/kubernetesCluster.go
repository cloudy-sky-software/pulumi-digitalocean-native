// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"errors"
	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type KubernetesCluster struct {
	pulumi.CustomResourceState

	// A boolean value indicating whether the cluster will be automatically upgraded to new patch releases during its maintenance window.
	AutoUpgrade pulumi.BoolPtrOutput `pulumi:"autoUpgrade"`
	// The range of IP addresses in the overlay network of the Kubernetes cluster in CIDR notation.
	ClusterSubnet pulumi.StringPtrOutput `pulumi:"clusterSubnet"`
	// A time value given in ISO8601 combined date and time format that represents when the Kubernetes cluster was created.
	CreatedAt pulumi.StringPtrOutput `pulumi:"createdAt"`
	// The base URL of the API server on the Kubernetes master node.
	Endpoint pulumi.StringPtrOutput `pulumi:"endpoint"`
	// A boolean value indicating whether the control plane is run in a highly available configuration in the cluster. Highly available control planes incur less downtime. The property cannot be disabled.
	Ha pulumi.BoolPtrOutput `pulumi:"ha"`
	// The public IPv4 address of the Kubernetes master node. This will not be set if high availability is configured on the cluster (v1.21+)
	Ipv4              pulumi.StringPtrOutput `pulumi:"ipv4"`
	KubernetesCluster ClusterPtrOutput       `pulumi:"kubernetesCluster"`
	// An object specifying the maintenance window policy for the Kubernetes cluster.
	MaintenancePolicy MaintenancePolicyPtrOutput `pulumi:"maintenancePolicy"`
	// A human-readable name for a Kubernetes cluster.
	Name pulumi.StringOutput `pulumi:"name"`
	// An object specifying the details of the worker nodes available to the Kubernetes cluster.
	NodePools KubernetesNodePoolTypeArrayOutput `pulumi:"nodePools"`
	// The slug identifier for the region where the Kubernetes cluster is located.
	Region pulumi.StringOutput `pulumi:"region"`
	// A read-only boolean value indicating if a container registry is integrated with the cluster.
	RegistryEnabled pulumi.BoolPtrOutput `pulumi:"registryEnabled"`
	// The range of assignable IP addresses for services running in the Kubernetes cluster in CIDR notation.
	ServiceSubnet pulumi.StringPtrOutput `pulumi:"serviceSubnet"`
	// An object containing a `state` attribute whose value is set to a string indicating the current status of the cluster.
	Status StatusPropertiesPtrOutput `pulumi:"status"`
	// A boolean value indicating whether surge upgrade is enabled/disabled for the cluster. Surge upgrade makes cluster upgrades fast and reliable by bringing up new nodes before destroying the outdated nodes.
	SurgeUpgrade pulumi.BoolPtrOutput `pulumi:"surgeUpgrade"`
	// An array of tags applied to the Kubernetes cluster. All clusters are automatically tagged `k8s` and `k8s:$K8S_CLUSTER_ID`.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// A time value given in ISO8601 combined date and time format that represents when the Kubernetes cluster was last updated.
	UpdatedAt pulumi.StringPtrOutput `pulumi:"updatedAt"`
	// The slug identifier for the version of Kubernetes used for the cluster. If set to a minor version (e.g. "1.14"), the latest version within it will be used (e.g. "1.14.6-do.1"); if set to "latest", the latest published version will be used. See the `/v2/kubernetes/options` endpoint to find all currently available versions.
	Version pulumi.StringOutput `pulumi:"version"`
	// A string specifying the UUID of the VPC to which the Kubernetes cluster is assigned.
	VpcUuid pulumi.StringPtrOutput `pulumi:"vpcUuid"`
}

// NewKubernetesCluster registers a new resource with the given unique name, arguments, and options.
func NewKubernetesCluster(ctx *pulumi.Context,
	name string, args *KubernetesClusterArgs, opts ...pulumi.ResourceOption) (*KubernetesCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NodePools == nil {
		return nil, errors.New("invalid value for required argument 'NodePools'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	if args.AutoUpgrade == nil {
		args.AutoUpgrade = pulumi.BoolPtr(false)
	}
	if args.Ha == nil {
		args.Ha = pulumi.BoolPtr(false)
	}
	if args.SurgeUpgrade == nil {
		args.SurgeUpgrade = pulumi.BoolPtr(false)
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource KubernetesCluster
	err := ctx.RegisterResource("digitalocean-native:kubernetes/v2:KubernetesCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKubernetesCluster gets an existing KubernetesCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKubernetesCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KubernetesClusterState, opts ...pulumi.ResourceOption) (*KubernetesCluster, error) {
	var resource KubernetesCluster
	err := ctx.ReadResource("digitalocean-native:kubernetes/v2:KubernetesCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KubernetesCluster resources.
type kubernetesClusterState struct {
}

type KubernetesClusterState struct {
}

func (KubernetesClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesClusterState)(nil)).Elem()
}

type kubernetesClusterArgs struct {
	// A boolean value indicating whether the cluster will be automatically upgraded to new patch releases during its maintenance window.
	AutoUpgrade *bool `pulumi:"autoUpgrade"`
	// A boolean value indicating whether the control plane is run in a highly available configuration in the cluster. Highly available control planes incur less downtime. The property cannot be disabled.
	Ha *bool `pulumi:"ha"`
	// An object specifying the maintenance window policy for the Kubernetes cluster.
	MaintenancePolicy *MaintenancePolicy `pulumi:"maintenancePolicy"`
	// A human-readable name for a Kubernetes cluster.
	Name *string `pulumi:"name"`
	// An object specifying the details of the worker nodes available to the Kubernetes cluster.
	NodePools []KubernetesNodePoolType `pulumi:"nodePools"`
	// The slug identifier for the region where the Kubernetes cluster is located.
	Region string `pulumi:"region"`
	// A boolean value indicating whether surge upgrade is enabled/disabled for the cluster. Surge upgrade makes cluster upgrades fast and reliable by bringing up new nodes before destroying the outdated nodes.
	SurgeUpgrade *bool `pulumi:"surgeUpgrade"`
	// An array of tags applied to the Kubernetes cluster. All clusters are automatically tagged `k8s` and `k8s:$K8S_CLUSTER_ID`.
	Tags []string `pulumi:"tags"`
	// The slug identifier for the version of Kubernetes used for the cluster. If set to a minor version (e.g. "1.14"), the latest version within it will be used (e.g. "1.14.6-do.1"); if set to "latest", the latest published version will be used. See the `/v2/kubernetes/options` endpoint to find all currently available versions.
	Version string `pulumi:"version"`
	// A string specifying the UUID of the VPC to which the Kubernetes cluster is assigned.
	VpcUuid *string `pulumi:"vpcUuid"`
}

// The set of arguments for constructing a KubernetesCluster resource.
type KubernetesClusterArgs struct {
	// A boolean value indicating whether the cluster will be automatically upgraded to new patch releases during its maintenance window.
	AutoUpgrade pulumi.BoolPtrInput
	// A boolean value indicating whether the control plane is run in a highly available configuration in the cluster. Highly available control planes incur less downtime. The property cannot be disabled.
	Ha pulumi.BoolPtrInput
	// An object specifying the maintenance window policy for the Kubernetes cluster.
	MaintenancePolicy MaintenancePolicyPtrInput
	// A human-readable name for a Kubernetes cluster.
	Name pulumi.StringPtrInput
	// An object specifying the details of the worker nodes available to the Kubernetes cluster.
	NodePools KubernetesNodePoolTypeArrayInput
	// The slug identifier for the region where the Kubernetes cluster is located.
	Region pulumi.StringInput
	// A boolean value indicating whether surge upgrade is enabled/disabled for the cluster. Surge upgrade makes cluster upgrades fast and reliable by bringing up new nodes before destroying the outdated nodes.
	SurgeUpgrade pulumi.BoolPtrInput
	// An array of tags applied to the Kubernetes cluster. All clusters are automatically tagged `k8s` and `k8s:$K8S_CLUSTER_ID`.
	Tags pulumi.StringArrayInput
	// The slug identifier for the version of Kubernetes used for the cluster. If set to a minor version (e.g. "1.14"), the latest version within it will be used (e.g. "1.14.6-do.1"); if set to "latest", the latest published version will be used. See the `/v2/kubernetes/options` endpoint to find all currently available versions.
	Version pulumi.StringInput
	// A string specifying the UUID of the VPC to which the Kubernetes cluster is assigned.
	VpcUuid pulumi.StringPtrInput
}

func (KubernetesClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesClusterArgs)(nil)).Elem()
}

type KubernetesClusterInput interface {
	pulumi.Input

	ToKubernetesClusterOutput() KubernetesClusterOutput
	ToKubernetesClusterOutputWithContext(ctx context.Context) KubernetesClusterOutput
}

func (*KubernetesCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesCluster)(nil)).Elem()
}

func (i *KubernetesCluster) ToKubernetesClusterOutput() KubernetesClusterOutput {
	return i.ToKubernetesClusterOutputWithContext(context.Background())
}

func (i *KubernetesCluster) ToKubernetesClusterOutputWithContext(ctx context.Context) KubernetesClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesClusterOutput)
}

type KubernetesClusterOutput struct{ *pulumi.OutputState }

func (KubernetesClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesCluster)(nil)).Elem()
}

func (o KubernetesClusterOutput) ToKubernetesClusterOutput() KubernetesClusterOutput {
	return o
}

func (o KubernetesClusterOutput) ToKubernetesClusterOutputWithContext(ctx context.Context) KubernetesClusterOutput {
	return o
}

// A boolean value indicating whether the cluster will be automatically upgraded to new patch releases during its maintenance window.
func (o KubernetesClusterOutput) AutoUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KubernetesCluster) pulumi.BoolPtrOutput { return v.AutoUpgrade }).(pulumi.BoolPtrOutput)
}

// The range of IP addresses in the overlay network of the Kubernetes cluster in CIDR notation.
func (o KubernetesClusterOutput) ClusterSubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesCluster) pulumi.StringPtrOutput { return v.ClusterSubnet }).(pulumi.StringPtrOutput)
}

// A time value given in ISO8601 combined date and time format that represents when the Kubernetes cluster was created.
func (o KubernetesClusterOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesCluster) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The base URL of the API server on the Kubernetes master node.
func (o KubernetesClusterOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesCluster) pulumi.StringPtrOutput { return v.Endpoint }).(pulumi.StringPtrOutput)
}

// A boolean value indicating whether the control plane is run in a highly available configuration in the cluster. Highly available control planes incur less downtime. The property cannot be disabled.
func (o KubernetesClusterOutput) Ha() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KubernetesCluster) pulumi.BoolPtrOutput { return v.Ha }).(pulumi.BoolPtrOutput)
}

// The public IPv4 address of the Kubernetes master node. This will not be set if high availability is configured on the cluster (v1.21+)
func (o KubernetesClusterOutput) Ipv4() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesCluster) pulumi.StringPtrOutput { return v.Ipv4 }).(pulumi.StringPtrOutput)
}

func (o KubernetesClusterOutput) KubernetesCluster() ClusterPtrOutput {
	return o.ApplyT(func(v *KubernetesCluster) ClusterPtrOutput { return v.KubernetesCluster }).(ClusterPtrOutput)
}

// An object specifying the maintenance window policy for the Kubernetes cluster.
func (o KubernetesClusterOutput) MaintenancePolicy() MaintenancePolicyPtrOutput {
	return o.ApplyT(func(v *KubernetesCluster) MaintenancePolicyPtrOutput { return v.MaintenancePolicy }).(MaintenancePolicyPtrOutput)
}

// A human-readable name for a Kubernetes cluster.
func (o KubernetesClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// An object specifying the details of the worker nodes available to the Kubernetes cluster.
func (o KubernetesClusterOutput) NodePools() KubernetesNodePoolTypeArrayOutput {
	return o.ApplyT(func(v *KubernetesCluster) KubernetesNodePoolTypeArrayOutput { return v.NodePools }).(KubernetesNodePoolTypeArrayOutput)
}

// The slug identifier for the region where the Kubernetes cluster is located.
func (o KubernetesClusterOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesCluster) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// A read-only boolean value indicating if a container registry is integrated with the cluster.
func (o KubernetesClusterOutput) RegistryEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KubernetesCluster) pulumi.BoolPtrOutput { return v.RegistryEnabled }).(pulumi.BoolPtrOutput)
}

// The range of assignable IP addresses for services running in the Kubernetes cluster in CIDR notation.
func (o KubernetesClusterOutput) ServiceSubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesCluster) pulumi.StringPtrOutput { return v.ServiceSubnet }).(pulumi.StringPtrOutput)
}

// An object containing a `state` attribute whose value is set to a string indicating the current status of the cluster.
func (o KubernetesClusterOutput) Status() StatusPropertiesPtrOutput {
	return o.ApplyT(func(v *KubernetesCluster) StatusPropertiesPtrOutput { return v.Status }).(StatusPropertiesPtrOutput)
}

// A boolean value indicating whether surge upgrade is enabled/disabled for the cluster. Surge upgrade makes cluster upgrades fast and reliable by bringing up new nodes before destroying the outdated nodes.
func (o KubernetesClusterOutput) SurgeUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KubernetesCluster) pulumi.BoolPtrOutput { return v.SurgeUpgrade }).(pulumi.BoolPtrOutput)
}

// An array of tags applied to the Kubernetes cluster. All clusters are automatically tagged `k8s` and `k8s:$K8S_CLUSTER_ID`.
func (o KubernetesClusterOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *KubernetesCluster) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// A time value given in ISO8601 combined date and time format that represents when the Kubernetes cluster was last updated.
func (o KubernetesClusterOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesCluster) pulumi.StringPtrOutput { return v.UpdatedAt }).(pulumi.StringPtrOutput)
}

// The slug identifier for the version of Kubernetes used for the cluster. If set to a minor version (e.g. "1.14"), the latest version within it will be used (e.g. "1.14.6-do.1"); if set to "latest", the latest published version will be used. See the `/v2/kubernetes/options` endpoint to find all currently available versions.
func (o KubernetesClusterOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesCluster) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

// A string specifying the UUID of the VPC to which the Kubernetes cluster is assigned.
func (o KubernetesClusterOutput) VpcUuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesCluster) pulumi.StringPtrOutput { return v.VpcUuid }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesClusterInput)(nil)).Elem(), &KubernetesCluster{})
	pulumi.RegisterOutputType(KubernetesClusterOutput{})
}