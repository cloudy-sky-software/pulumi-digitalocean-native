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
type PasswordReset struct {
	pulumi.CustomResourceState

	Action ActionPtrOutput `pulumi:"action"`
	// The type of action to initiate for the Droplet.
	Type TypeOutput `pulumi:"type"`
}

// NewPasswordReset registers a new resource with the given unique name, arguments, and options.
func NewPasswordReset(ctx *pulumi.Context,
	name string, args *PasswordResetArgs, opts ...pulumi.ResourceOption) (*PasswordReset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PasswordReset
	err := ctx.RegisterResource("digitalocean-native:droplets/v2:PasswordReset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPasswordReset gets an existing PasswordReset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPasswordReset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PasswordResetState, opts ...pulumi.ResourceOption) (*PasswordReset, error) {
	var resource PasswordReset
	err := ctx.ReadResource("digitalocean-native:droplets/v2:PasswordReset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PasswordReset resources.
type passwordResetState struct {
}

type PasswordResetState struct {
}

func (PasswordResetState) ElementType() reflect.Type {
	return reflect.TypeOf((*passwordResetState)(nil)).Elem()
}

type passwordResetArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId *string `pulumi:"dropletId"`
	// The type of action to initiate for the Droplet.
	Type Type `pulumi:"type"`
}

// The set of arguments for constructing a PasswordReset resource.
type PasswordResetArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId pulumi.StringPtrInput
	// The type of action to initiate for the Droplet.
	Type TypeInput
}

func (PasswordResetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*passwordResetArgs)(nil)).Elem()
}

type PasswordResetInput interface {
	pulumi.Input

	ToPasswordResetOutput() PasswordResetOutput
	ToPasswordResetOutputWithContext(ctx context.Context) PasswordResetOutput
}

func (*PasswordReset) ElementType() reflect.Type {
	return reflect.TypeOf((**PasswordReset)(nil)).Elem()
}

func (i *PasswordReset) ToPasswordResetOutput() PasswordResetOutput {
	return i.ToPasswordResetOutputWithContext(context.Background())
}

func (i *PasswordReset) ToPasswordResetOutputWithContext(ctx context.Context) PasswordResetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PasswordResetOutput)
}

type PasswordResetOutput struct{ *pulumi.OutputState }

func (PasswordResetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PasswordReset)(nil)).Elem()
}

func (o PasswordResetOutput) ToPasswordResetOutput() PasswordResetOutput {
	return o
}

func (o PasswordResetOutput) ToPasswordResetOutputWithContext(ctx context.Context) PasswordResetOutput {
	return o
}

func (o PasswordResetOutput) Action() ActionPtrOutput {
	return o.ApplyT(func(v *PasswordReset) ActionPtrOutput { return v.Action }).(ActionPtrOutput)
}

// The type of action to initiate for the Droplet.
func (o PasswordResetOutput) Type() TypeOutput {
	return o.ApplyT(func(v *PasswordReset) TypeOutput { return v.Type }).(TypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PasswordResetInput)(nil)).Elem(), &PasswordReset{})
	pulumi.RegisterOutputType(PasswordResetOutput{})
}