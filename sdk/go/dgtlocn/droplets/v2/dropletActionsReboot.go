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
type DropletActionsReboot struct {
	pulumi.CustomResourceState

	Action ActionPtrOutput `pulumi:"action"`
	// The type of action to initiate for the Droplet.
	Type TypeOutput `pulumi:"type"`
}

// NewDropletActionsReboot registers a new resource with the given unique name, arguments, and options.
func NewDropletActionsReboot(ctx *pulumi.Context,
	name string, args *DropletActionsRebootArgs, opts ...pulumi.ResourceOption) (*DropletActionsReboot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DropletActionsReboot
	err := ctx.RegisterResource("digitalocean-native:droplets/v2:DropletActionsReboot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDropletActionsReboot gets an existing DropletActionsReboot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDropletActionsReboot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DropletActionsRebootState, opts ...pulumi.ResourceOption) (*DropletActionsReboot, error) {
	var resource DropletActionsReboot
	err := ctx.ReadResource("digitalocean-native:droplets/v2:DropletActionsReboot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DropletActionsReboot resources.
type dropletActionsRebootState struct {
}

type DropletActionsRebootState struct {
}

func (DropletActionsRebootState) ElementType() reflect.Type {
	return reflect.TypeOf((*dropletActionsRebootState)(nil)).Elem()
}

type dropletActionsRebootArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId *string `pulumi:"dropletId"`
	// The type of action to initiate for the Droplet.
	Type Type `pulumi:"type"`
}

// The set of arguments for constructing a DropletActionsReboot resource.
type DropletActionsRebootArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId pulumi.StringPtrInput
	// The type of action to initiate for the Droplet.
	Type TypeInput
}

func (DropletActionsRebootArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dropletActionsRebootArgs)(nil)).Elem()
}

type DropletActionsRebootInput interface {
	pulumi.Input

	ToDropletActionsRebootOutput() DropletActionsRebootOutput
	ToDropletActionsRebootOutputWithContext(ctx context.Context) DropletActionsRebootOutput
}

func (*DropletActionsReboot) ElementType() reflect.Type {
	return reflect.TypeOf((**DropletActionsReboot)(nil)).Elem()
}

func (i *DropletActionsReboot) ToDropletActionsRebootOutput() DropletActionsRebootOutput {
	return i.ToDropletActionsRebootOutputWithContext(context.Background())
}

func (i *DropletActionsReboot) ToDropletActionsRebootOutputWithContext(ctx context.Context) DropletActionsRebootOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DropletActionsRebootOutput)
}

type DropletActionsRebootOutput struct{ *pulumi.OutputState }

func (DropletActionsRebootOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DropletActionsReboot)(nil)).Elem()
}

func (o DropletActionsRebootOutput) ToDropletActionsRebootOutput() DropletActionsRebootOutput {
	return o
}

func (o DropletActionsRebootOutput) ToDropletActionsRebootOutputWithContext(ctx context.Context) DropletActionsRebootOutput {
	return o
}

func (o DropletActionsRebootOutput) Action() ActionPtrOutput {
	return o.ApplyT(func(v *DropletActionsReboot) ActionPtrOutput { return v.Action }).(ActionPtrOutput)
}

// The type of action to initiate for the Droplet.
func (o DropletActionsRebootOutput) Type() TypeOutput {
	return o.ApplyT(func(v *DropletActionsReboot) TypeOutput { return v.Type }).(TypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DropletActionsRebootInput)(nil)).Elem(), &DropletActionsReboot{})
	pulumi.RegisterOutputType(DropletActionsRebootOutput{})
}