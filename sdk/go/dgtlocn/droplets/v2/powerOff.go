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
type PowerOff struct {
	pulumi.CustomResourceState

	Action ActionPtrOutput `pulumi:"action"`
	// The type of action to initiate for the Droplet.
	Type PowerOffTypeOutput `pulumi:"type"`
}

// NewPowerOff registers a new resource with the given unique name, arguments, and options.
func NewPowerOff(ctx *pulumi.Context,
	name string, args *PowerOffArgs, opts ...pulumi.ResourceOption) (*PowerOff, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PowerOff
	err := ctx.RegisterResource("digitalocean-native:droplets/v2:PowerOff", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPowerOff gets an existing PowerOff resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPowerOff(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PowerOffState, opts ...pulumi.ResourceOption) (*PowerOff, error) {
	var resource PowerOff
	err := ctx.ReadResource("digitalocean-native:droplets/v2:PowerOff", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PowerOff resources.
type powerOffState struct {
}

type PowerOffState struct {
}

func (PowerOffState) ElementType() reflect.Type {
	return reflect.TypeOf((*powerOffState)(nil)).Elem()
}

type powerOffArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId *string `pulumi:"dropletId"`
	// The type of action to initiate for the Droplet.
	Type PowerOffType `pulumi:"type"`
}

// The set of arguments for constructing a PowerOff resource.
type PowerOffArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId pulumi.StringPtrInput
	// The type of action to initiate for the Droplet.
	Type PowerOffTypeInput
}

func (PowerOffArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*powerOffArgs)(nil)).Elem()
}

type PowerOffInput interface {
	pulumi.Input

	ToPowerOffOutput() PowerOffOutput
	ToPowerOffOutputWithContext(ctx context.Context) PowerOffOutput
}

func (*PowerOff) ElementType() reflect.Type {
	return reflect.TypeOf((**PowerOff)(nil)).Elem()
}

func (i *PowerOff) ToPowerOffOutput() PowerOffOutput {
	return i.ToPowerOffOutputWithContext(context.Background())
}

func (i *PowerOff) ToPowerOffOutputWithContext(ctx context.Context) PowerOffOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PowerOffOutput)
}

type PowerOffOutput struct{ *pulumi.OutputState }

func (PowerOffOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PowerOff)(nil)).Elem()
}

func (o PowerOffOutput) ToPowerOffOutput() PowerOffOutput {
	return o
}

func (o PowerOffOutput) ToPowerOffOutputWithContext(ctx context.Context) PowerOffOutput {
	return o
}

func (o PowerOffOutput) Action() ActionPtrOutput {
	return o.ApplyT(func(v *PowerOff) ActionPtrOutput { return v.Action }).(ActionPtrOutput)
}

// The type of action to initiate for the Droplet.
func (o PowerOffOutput) Type() PowerOffTypeOutput {
	return o.ApplyT(func(v *PowerOff) PowerOffTypeOutput { return v.Type }).(PowerOffTypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PowerOffInput)(nil)).Elem(), &PowerOff{})
	pulumi.RegisterOutputType(PowerOffOutput{})
}
