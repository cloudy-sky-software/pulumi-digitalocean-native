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

// Specifies the action that will be taken on the Droplet.
type DropletActionsByTagPowerOn struct {
	pulumi.CustomResourceState

	Actions ActionArrayOutput `pulumi:"actions"`
	// The type of action to initiate for the Droplet.
	Type TypeOutput `pulumi:"type"`
}

// NewDropletActionsByTagPowerOn registers a new resource with the given unique name, arguments, and options.
func NewDropletActionsByTagPowerOn(ctx *pulumi.Context,
	name string, args *DropletActionsByTagPowerOnArgs, opts ...pulumi.ResourceOption) (*DropletActionsByTagPowerOn, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DropletActionsByTagPowerOn
	err := ctx.RegisterResource("digitalocean-native:droplets/v2:DropletActionsByTagPowerOn", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDropletActionsByTagPowerOn gets an existing DropletActionsByTagPowerOn resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDropletActionsByTagPowerOn(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DropletActionsByTagPowerOnState, opts ...pulumi.ResourceOption) (*DropletActionsByTagPowerOn, error) {
	var resource DropletActionsByTagPowerOn
	err := ctx.ReadResource("digitalocean-native:droplets/v2:DropletActionsByTagPowerOn", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DropletActionsByTagPowerOn resources.
type dropletActionsByTagPowerOnState struct {
}

type DropletActionsByTagPowerOnState struct {
}

func (DropletActionsByTagPowerOnState) ElementType() reflect.Type {
	return reflect.TypeOf((*dropletActionsByTagPowerOnState)(nil)).Elem()
}

type dropletActionsByTagPowerOnArgs struct {
	// The type of action to initiate for the Droplet.
	Type Type `pulumi:"type"`
}

// The set of arguments for constructing a DropletActionsByTagPowerOn resource.
type DropletActionsByTagPowerOnArgs struct {
	// The type of action to initiate for the Droplet.
	Type TypeInput
}

func (DropletActionsByTagPowerOnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dropletActionsByTagPowerOnArgs)(nil)).Elem()
}

type DropletActionsByTagPowerOnInput interface {
	pulumi.Input

	ToDropletActionsByTagPowerOnOutput() DropletActionsByTagPowerOnOutput
	ToDropletActionsByTagPowerOnOutputWithContext(ctx context.Context) DropletActionsByTagPowerOnOutput
}

func (*DropletActionsByTagPowerOn) ElementType() reflect.Type {
	return reflect.TypeOf((**DropletActionsByTagPowerOn)(nil)).Elem()
}

func (i *DropletActionsByTagPowerOn) ToDropletActionsByTagPowerOnOutput() DropletActionsByTagPowerOnOutput {
	return i.ToDropletActionsByTagPowerOnOutputWithContext(context.Background())
}

func (i *DropletActionsByTagPowerOn) ToDropletActionsByTagPowerOnOutputWithContext(ctx context.Context) DropletActionsByTagPowerOnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DropletActionsByTagPowerOnOutput)
}

type DropletActionsByTagPowerOnOutput struct{ *pulumi.OutputState }

func (DropletActionsByTagPowerOnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DropletActionsByTagPowerOn)(nil)).Elem()
}

func (o DropletActionsByTagPowerOnOutput) ToDropletActionsByTagPowerOnOutput() DropletActionsByTagPowerOnOutput {
	return o
}

func (o DropletActionsByTagPowerOnOutput) ToDropletActionsByTagPowerOnOutputWithContext(ctx context.Context) DropletActionsByTagPowerOnOutput {
	return o
}

func (o DropletActionsByTagPowerOnOutput) Actions() ActionArrayOutput {
	return o.ApplyT(func(v *DropletActionsByTagPowerOn) ActionArrayOutput { return v.Actions }).(ActionArrayOutput)
}

// The type of action to initiate for the Droplet.
func (o DropletActionsByTagPowerOnOutput) Type() TypeOutput {
	return o.ApplyT(func(v *DropletActionsByTagPowerOn) TypeOutput { return v.Type }).(TypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DropletActionsByTagPowerOnInput)(nil)).Elem(), &DropletActionsByTagPowerOn{})
	pulumi.RegisterOutputType(DropletActionsByTagPowerOnOutput{})
}