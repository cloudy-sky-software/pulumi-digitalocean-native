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

type VolumeActionsByIdDetach struct {
	pulumi.CustomResourceState

	Action VolumeActionPtrOutput `pulumi:"action"`
	// The unique identifier for the Droplet the volume will be attached or detached from.
	DropletId pulumi.IntPtrOutput `pulumi:"dropletId"`
	// The slug identifier for the region where the resource will initially be  available.
	Region VolumeActionPostBaseRegionPtrOutput `pulumi:"region"`
	// The volume action to initiate.
	Type VolumeActionPostBaseTypePtrOutput `pulumi:"type"`
}

// NewVolumeActionsByIdDetach registers a new resource with the given unique name, arguments, and options.
func NewVolumeActionsByIdDetach(ctx *pulumi.Context,
	name string, args *VolumeActionsByIdDetachArgs, opts ...pulumi.ResourceOption) (*VolumeActionsByIdDetach, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DropletId == nil {
		return nil, errors.New("invalid value for required argument 'DropletId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VolumeActionsByIdDetach
	err := ctx.RegisterResource("digitalocean-native:volumes/v2:VolumeActionsByIdDetach", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolumeActionsByIdDetach gets an existing VolumeActionsByIdDetach resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolumeActionsByIdDetach(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeActionsByIdDetachState, opts ...pulumi.ResourceOption) (*VolumeActionsByIdDetach, error) {
	var resource VolumeActionsByIdDetach
	err := ctx.ReadResource("digitalocean-native:volumes/v2:VolumeActionsByIdDetach", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VolumeActionsByIdDetach resources.
type volumeActionsByIdDetachState struct {
}

type VolumeActionsByIdDetachState struct {
}

func (VolumeActionsByIdDetachState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeActionsByIdDetachState)(nil)).Elem()
}

type volumeActionsByIdDetachArgs struct {
	// The unique identifier for the Droplet the volume will be attached or detached from.
	DropletId int `pulumi:"dropletId"`
	// The slug identifier for the region where the resource will initially be  available.
	Region *VolumeActionPostBaseRegion `pulumi:"region"`
	// The volume action to initiate.
	Type VolumeActionPostBaseType `pulumi:"type"`
	// The ID of the block storage volume.
	VolumeId *string `pulumi:"volumeId"`
}

// The set of arguments for constructing a VolumeActionsByIdDetach resource.
type VolumeActionsByIdDetachArgs struct {
	// The unique identifier for the Droplet the volume will be attached or detached from.
	DropletId pulumi.IntInput
	// The slug identifier for the region where the resource will initially be  available.
	Region VolumeActionPostBaseRegionPtrInput
	// The volume action to initiate.
	Type VolumeActionPostBaseTypeInput
	// The ID of the block storage volume.
	VolumeId pulumi.StringPtrInput
}

func (VolumeActionsByIdDetachArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeActionsByIdDetachArgs)(nil)).Elem()
}

type VolumeActionsByIdDetachInput interface {
	pulumi.Input

	ToVolumeActionsByIdDetachOutput() VolumeActionsByIdDetachOutput
	ToVolumeActionsByIdDetachOutputWithContext(ctx context.Context) VolumeActionsByIdDetachOutput
}

func (*VolumeActionsByIdDetach) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeActionsByIdDetach)(nil)).Elem()
}

func (i *VolumeActionsByIdDetach) ToVolumeActionsByIdDetachOutput() VolumeActionsByIdDetachOutput {
	return i.ToVolumeActionsByIdDetachOutputWithContext(context.Background())
}

func (i *VolumeActionsByIdDetach) ToVolumeActionsByIdDetachOutputWithContext(ctx context.Context) VolumeActionsByIdDetachOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeActionsByIdDetachOutput)
}

type VolumeActionsByIdDetachOutput struct{ *pulumi.OutputState }

func (VolumeActionsByIdDetachOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeActionsByIdDetach)(nil)).Elem()
}

func (o VolumeActionsByIdDetachOutput) ToVolumeActionsByIdDetachOutput() VolumeActionsByIdDetachOutput {
	return o
}

func (o VolumeActionsByIdDetachOutput) ToVolumeActionsByIdDetachOutputWithContext(ctx context.Context) VolumeActionsByIdDetachOutput {
	return o
}

func (o VolumeActionsByIdDetachOutput) Action() VolumeActionPtrOutput {
	return o.ApplyT(func(v *VolumeActionsByIdDetach) VolumeActionPtrOutput { return v.Action }).(VolumeActionPtrOutput)
}

// The unique identifier for the Droplet the volume will be attached or detached from.
func (o VolumeActionsByIdDetachOutput) DropletId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VolumeActionsByIdDetach) pulumi.IntPtrOutput { return v.DropletId }).(pulumi.IntPtrOutput)
}

// The slug identifier for the region where the resource will initially be  available.
func (o VolumeActionsByIdDetachOutput) Region() VolumeActionPostBaseRegionPtrOutput {
	return o.ApplyT(func(v *VolumeActionsByIdDetach) VolumeActionPostBaseRegionPtrOutput { return v.Region }).(VolumeActionPostBaseRegionPtrOutput)
}

// The volume action to initiate.
func (o VolumeActionsByIdDetachOutput) Type() VolumeActionPostBaseTypePtrOutput {
	return o.ApplyT(func(v *VolumeActionsByIdDetach) VolumeActionPostBaseTypePtrOutput { return v.Type }).(VolumeActionPostBaseTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeActionsByIdDetachInput)(nil)).Elem(), &VolumeActionsByIdDetach{})
	pulumi.RegisterOutputType(VolumeActionsByIdDetachOutput{})
}