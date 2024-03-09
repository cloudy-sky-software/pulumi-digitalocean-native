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
type DropletActionsPowerOn struct {
	pulumi.CustomResourceState

	Action ActionPtrOutput `pulumi:"action"`
	// The type of action to initiate for the Droplet.
	Type TypeOutput `pulumi:"type"`
}

// NewDropletActionsPowerOn registers a new resource with the given unique name, arguments, and options.
func NewDropletActionsPowerOn(ctx *pulumi.Context,
	name string, args *DropletActionsPowerOnArgs, opts ...pulumi.ResourceOption) (*DropletActionsPowerOn, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DropletActionsPowerOn
	err := ctx.RegisterResource("digitalocean-native:droplets/v2:DropletActionsPowerOn", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDropletActionsPowerOn gets an existing DropletActionsPowerOn resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDropletActionsPowerOn(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DropletActionsPowerOnState, opts ...pulumi.ResourceOption) (*DropletActionsPowerOn, error) {
	var resource DropletActionsPowerOn
	err := ctx.ReadResource("digitalocean-native:droplets/v2:DropletActionsPowerOn", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DropletActionsPowerOn resources.
type dropletActionsPowerOnState struct {
}

type DropletActionsPowerOnState struct {
}

func (DropletActionsPowerOnState) ElementType() reflect.Type {
	return reflect.TypeOf((*dropletActionsPowerOnState)(nil)).Elem()
}

type dropletActionsPowerOnArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId *string `pulumi:"dropletId"`
	// The type of action to initiate for the Droplet.
	Type Type `pulumi:"type"`
}

// The set of arguments for constructing a DropletActionsPowerOn resource.
type DropletActionsPowerOnArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId pulumi.StringPtrInput
	// The type of action to initiate for the Droplet.
	Type TypeInput
}

func (DropletActionsPowerOnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dropletActionsPowerOnArgs)(nil)).Elem()
}

type DropletActionsPowerOnInput interface {
	pulumi.Input

	ToDropletActionsPowerOnOutput() DropletActionsPowerOnOutput
	ToDropletActionsPowerOnOutputWithContext(ctx context.Context) DropletActionsPowerOnOutput
}

func (*DropletActionsPowerOn) ElementType() reflect.Type {
	return reflect.TypeOf((**DropletActionsPowerOn)(nil)).Elem()
}

func (i *DropletActionsPowerOn) ToDropletActionsPowerOnOutput() DropletActionsPowerOnOutput {
	return i.ToDropletActionsPowerOnOutputWithContext(context.Background())
}

func (i *DropletActionsPowerOn) ToDropletActionsPowerOnOutputWithContext(ctx context.Context) DropletActionsPowerOnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DropletActionsPowerOnOutput)
}

type DropletActionsPowerOnOutput struct{ *pulumi.OutputState }

func (DropletActionsPowerOnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DropletActionsPowerOn)(nil)).Elem()
}

func (o DropletActionsPowerOnOutput) ToDropletActionsPowerOnOutput() DropletActionsPowerOnOutput {
	return o
}

func (o DropletActionsPowerOnOutput) ToDropletActionsPowerOnOutputWithContext(ctx context.Context) DropletActionsPowerOnOutput {
	return o
}

func (o DropletActionsPowerOnOutput) Action() ActionPtrOutput {
	return o.ApplyT(func(v *DropletActionsPowerOn) ActionPtrOutput { return v.Action }).(ActionPtrOutput)
}

// The type of action to initiate for the Droplet.
func (o DropletActionsPowerOnOutput) Type() TypeOutput {
	return o.ApplyT(func(v *DropletActionsPowerOn) TypeOutput { return v.Type }).(TypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DropletActionsPowerOnInput)(nil)).Elem(), &DropletActionsPowerOn{})
	pulumi.RegisterOutputType(DropletActionsPowerOnOutput{})
}