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
type DropletActionsPasswordReset struct {
	pulumi.CustomResourceState

	Action ActionPtrOutput `pulumi:"action"`
	// The type of action to initiate for the Droplet.
	Type TypeOutput `pulumi:"type"`
}

// NewDropletActionsPasswordReset registers a new resource with the given unique name, arguments, and options.
func NewDropletActionsPasswordReset(ctx *pulumi.Context,
	name string, args *DropletActionsPasswordResetArgs, opts ...pulumi.ResourceOption) (*DropletActionsPasswordReset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DropletActionsPasswordReset
	err := ctx.RegisterResource("digitalocean-native:droplets/v2:DropletActionsPasswordReset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDropletActionsPasswordReset gets an existing DropletActionsPasswordReset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDropletActionsPasswordReset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DropletActionsPasswordResetState, opts ...pulumi.ResourceOption) (*DropletActionsPasswordReset, error) {
	var resource DropletActionsPasswordReset
	err := ctx.ReadResource("digitalocean-native:droplets/v2:DropletActionsPasswordReset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DropletActionsPasswordReset resources.
type dropletActionsPasswordResetState struct {
}

type DropletActionsPasswordResetState struct {
}

func (DropletActionsPasswordResetState) ElementType() reflect.Type {
	return reflect.TypeOf((*dropletActionsPasswordResetState)(nil)).Elem()
}

type dropletActionsPasswordResetArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId *string `pulumi:"dropletId"`
	// The type of action to initiate for the Droplet.
	Type Type `pulumi:"type"`
}

// The set of arguments for constructing a DropletActionsPasswordReset resource.
type DropletActionsPasswordResetArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId pulumi.StringPtrInput
	// The type of action to initiate for the Droplet.
	Type TypeInput
}

func (DropletActionsPasswordResetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dropletActionsPasswordResetArgs)(nil)).Elem()
}

type DropletActionsPasswordResetInput interface {
	pulumi.Input

	ToDropletActionsPasswordResetOutput() DropletActionsPasswordResetOutput
	ToDropletActionsPasswordResetOutputWithContext(ctx context.Context) DropletActionsPasswordResetOutput
}

func (*DropletActionsPasswordReset) ElementType() reflect.Type {
	return reflect.TypeOf((**DropletActionsPasswordReset)(nil)).Elem()
}

func (i *DropletActionsPasswordReset) ToDropletActionsPasswordResetOutput() DropletActionsPasswordResetOutput {
	return i.ToDropletActionsPasswordResetOutputWithContext(context.Background())
}

func (i *DropletActionsPasswordReset) ToDropletActionsPasswordResetOutputWithContext(ctx context.Context) DropletActionsPasswordResetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DropletActionsPasswordResetOutput)
}

type DropletActionsPasswordResetOutput struct{ *pulumi.OutputState }

func (DropletActionsPasswordResetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DropletActionsPasswordReset)(nil)).Elem()
}

func (o DropletActionsPasswordResetOutput) ToDropletActionsPasswordResetOutput() DropletActionsPasswordResetOutput {
	return o
}

func (o DropletActionsPasswordResetOutput) ToDropletActionsPasswordResetOutputWithContext(ctx context.Context) DropletActionsPasswordResetOutput {
	return o
}

func (o DropletActionsPasswordResetOutput) Action() ActionPtrOutput {
	return o.ApplyT(func(v *DropletActionsPasswordReset) ActionPtrOutput { return v.Action }).(ActionPtrOutput)
}

// The type of action to initiate for the Droplet.
func (o DropletActionsPasswordResetOutput) Type() TypeOutput {
	return o.ApplyT(func(v *DropletActionsPasswordReset) TypeOutput { return v.Type }).(TypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DropletActionsPasswordResetInput)(nil)).Elem(), &DropletActionsPasswordReset{})
	pulumi.RegisterOutputType(DropletActionsPasswordResetOutput{})
}