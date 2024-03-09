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

type OneClicksInstallKubernetes struct {
	pulumi.CustomResourceState

	// An array of 1-Click Application slugs to be installed to the Kubernetes cluster.
	AddonSlugs pulumi.StringArrayOutput `pulumi:"addonSlugs"`
	// A unique ID for the Kubernetes cluster to which the 1-Click Applications will be installed.
	ClusterUuid pulumi.StringOutput `pulumi:"clusterUuid"`
	// A message about the result of the request.
	Message pulumi.StringPtrOutput `pulumi:"message"`
}

// NewOneClicksInstallKubernetes registers a new resource with the given unique name, arguments, and options.
func NewOneClicksInstallKubernetes(ctx *pulumi.Context,
	name string, args *OneClicksInstallKubernetesArgs, opts ...pulumi.ResourceOption) (*OneClicksInstallKubernetes, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AddonSlugs == nil {
		return nil, errors.New("invalid value for required argument 'AddonSlugs'")
	}
	if args.ClusterUuid == nil {
		return nil, errors.New("invalid value for required argument 'ClusterUuid'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OneClicksInstallKubernetes
	err := ctx.RegisterResource("digitalocean-native:1-clicks/v2:OneClicksInstallKubernetes", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOneClicksInstallKubernetes gets an existing OneClicksInstallKubernetes resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOneClicksInstallKubernetes(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OneClicksInstallKubernetesState, opts ...pulumi.ResourceOption) (*OneClicksInstallKubernetes, error) {
	var resource OneClicksInstallKubernetes
	err := ctx.ReadResource("digitalocean-native:1-clicks/v2:OneClicksInstallKubernetes", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OneClicksInstallKubernetes resources.
type oneClicksInstallKubernetesState struct {
}

type OneClicksInstallKubernetesState struct {
}

func (OneClicksInstallKubernetesState) ElementType() reflect.Type {
	return reflect.TypeOf((*oneClicksInstallKubernetesState)(nil)).Elem()
}

type oneClicksInstallKubernetesArgs struct {
	// An array of 1-Click Application slugs to be installed to the Kubernetes cluster.
	AddonSlugs []string `pulumi:"addonSlugs"`
	// A unique ID for the Kubernetes cluster to which the 1-Click Applications will be installed.
	ClusterUuid string `pulumi:"clusterUuid"`
}

// The set of arguments for constructing a OneClicksInstallKubernetes resource.
type OneClicksInstallKubernetesArgs struct {
	// An array of 1-Click Application slugs to be installed to the Kubernetes cluster.
	AddonSlugs pulumi.StringArrayInput
	// A unique ID for the Kubernetes cluster to which the 1-Click Applications will be installed.
	ClusterUuid pulumi.StringInput
}

func (OneClicksInstallKubernetesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*oneClicksInstallKubernetesArgs)(nil)).Elem()
}

type OneClicksInstallKubernetesInput interface {
	pulumi.Input

	ToOneClicksInstallKubernetesOutput() OneClicksInstallKubernetesOutput
	ToOneClicksInstallKubernetesOutputWithContext(ctx context.Context) OneClicksInstallKubernetesOutput
}

func (*OneClicksInstallKubernetes) ElementType() reflect.Type {
	return reflect.TypeOf((**OneClicksInstallKubernetes)(nil)).Elem()
}

func (i *OneClicksInstallKubernetes) ToOneClicksInstallKubernetesOutput() OneClicksInstallKubernetesOutput {
	return i.ToOneClicksInstallKubernetesOutputWithContext(context.Background())
}

func (i *OneClicksInstallKubernetes) ToOneClicksInstallKubernetesOutputWithContext(ctx context.Context) OneClicksInstallKubernetesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OneClicksInstallKubernetesOutput)
}

type OneClicksInstallKubernetesOutput struct{ *pulumi.OutputState }

func (OneClicksInstallKubernetesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OneClicksInstallKubernetes)(nil)).Elem()
}

func (o OneClicksInstallKubernetesOutput) ToOneClicksInstallKubernetesOutput() OneClicksInstallKubernetesOutput {
	return o
}

func (o OneClicksInstallKubernetesOutput) ToOneClicksInstallKubernetesOutputWithContext(ctx context.Context) OneClicksInstallKubernetesOutput {
	return o
}

// An array of 1-Click Application slugs to be installed to the Kubernetes cluster.
func (o OneClicksInstallKubernetesOutput) AddonSlugs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OneClicksInstallKubernetes) pulumi.StringArrayOutput { return v.AddonSlugs }).(pulumi.StringArrayOutput)
}

// A unique ID for the Kubernetes cluster to which the 1-Click Applications will be installed.
func (o OneClicksInstallKubernetesOutput) ClusterUuid() pulumi.StringOutput {
	return o.ApplyT(func(v *OneClicksInstallKubernetes) pulumi.StringOutput { return v.ClusterUuid }).(pulumi.StringOutput)
}

// A message about the result of the request.
func (o OneClicksInstallKubernetesOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OneClicksInstallKubernetes) pulumi.StringPtrOutput { return v.Message }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OneClicksInstallKubernetesInput)(nil)).Elem(), &OneClicksInstallKubernetes{})
	pulumi.RegisterOutputType(OneClicksInstallKubernetesOutput{})
}
