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

type Rename struct {
	pulumi.CustomResourceState

	Action ActionPtrOutput `pulumi:"action"`
	// The new name for the Droplet.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The type of action to initiate for the Droplet.
	Type DropletActionTypePtrOutput `pulumi:"type"`
}

// NewRename registers a new resource with the given unique name, arguments, and options.
func NewRename(ctx *pulumi.Context,
	name string, args *RenameArgs, opts ...pulumi.ResourceOption) (*Rename, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Rename
	err := ctx.RegisterResource("digitalocean-native:droplets/v2:Rename", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRename gets an existing Rename resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRename(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RenameState, opts ...pulumi.ResourceOption) (*Rename, error) {
	var resource Rename
	err := ctx.ReadResource("digitalocean-native:droplets/v2:Rename", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rename resources.
type renameState struct {
}

type RenameState struct {
}

func (RenameState) ElementType() reflect.Type {
	return reflect.TypeOf((*renameState)(nil)).Elem()
}

type renameArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId *string `pulumi:"dropletId"`
	// The new name for the Droplet.
	Name *string `pulumi:"name"`
	// The type of action to initiate for the Droplet.
	Type DropletActionType `pulumi:"type"`
}

// The set of arguments for constructing a Rename resource.
type RenameArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId pulumi.StringPtrInput
	// The new name for the Droplet.
	Name pulumi.StringPtrInput
	// The type of action to initiate for the Droplet.
	Type DropletActionTypeInput
}

func (RenameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*renameArgs)(nil)).Elem()
}

type RenameInput interface {
	pulumi.Input

	ToRenameOutput() RenameOutput
	ToRenameOutputWithContext(ctx context.Context) RenameOutput
}

func (*Rename) ElementType() reflect.Type {
	return reflect.TypeOf((**Rename)(nil)).Elem()
}

func (i *Rename) ToRenameOutput() RenameOutput {
	return i.ToRenameOutputWithContext(context.Background())
}

func (i *Rename) ToRenameOutputWithContext(ctx context.Context) RenameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RenameOutput)
}

type RenameOutput struct{ *pulumi.OutputState }

func (RenameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Rename)(nil)).Elem()
}

func (o RenameOutput) ToRenameOutput() RenameOutput {
	return o
}

func (o RenameOutput) ToRenameOutputWithContext(ctx context.Context) RenameOutput {
	return o
}

func (o RenameOutput) Action() ActionPtrOutput {
	return o.ApplyT(func(v *Rename) ActionPtrOutput { return v.Action }).(ActionPtrOutput)
}

// The new name for the Droplet.
func (o RenameOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rename) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The type of action to initiate for the Droplet.
func (o RenameOutput) Type() DropletActionTypePtrOutput {
	return o.ApplyT(func(v *Rename) DropletActionTypePtrOutput { return v.Type }).(DropletActionTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RenameInput)(nil)).Elem(), &Rename{})
	pulumi.RegisterOutputType(RenameOutput{})
}