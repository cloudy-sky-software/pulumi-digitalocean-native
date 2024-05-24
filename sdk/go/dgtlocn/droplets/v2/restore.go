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

type Restore struct {
	pulumi.CustomResourceState

	Action ActionPtrOutput `pulumi:"action"`
	// The ID of a backup of the current Droplet instance to restore from.
	Image pulumi.IntPtrOutput `pulumi:"image"`
	// The type of action to initiate for the Droplet.
	Type RestoreDropletActionTypePtrOutput `pulumi:"type"`
}

// NewRestore registers a new resource with the given unique name, arguments, and options.
func NewRestore(ctx *pulumi.Context,
	name string, args *RestoreArgs, opts ...pulumi.ResourceOption) (*Restore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Restore
	err := ctx.RegisterResource("digitalocean-native:droplets/v2:Restore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRestore gets an existing Restore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRestore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RestoreState, opts ...pulumi.ResourceOption) (*Restore, error) {
	var resource Restore
	err := ctx.ReadResource("digitalocean-native:droplets/v2:Restore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Restore resources.
type restoreState struct {
}

type RestoreState struct {
}

func (RestoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*restoreState)(nil)).Elem()
}

type restoreArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId *string `pulumi:"dropletId"`
	// The ID of a backup of the current Droplet instance to restore from.
	Image *int `pulumi:"image"`
	// The type of action to initiate for the Droplet.
	Type RestoreDropletActionType `pulumi:"type"`
}

// The set of arguments for constructing a Restore resource.
type RestoreArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId pulumi.StringPtrInput
	// The ID of a backup of the current Droplet instance to restore from.
	Image pulumi.IntPtrInput
	// The type of action to initiate for the Droplet.
	Type RestoreDropletActionTypeInput
}

func (RestoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*restoreArgs)(nil)).Elem()
}

type RestoreInput interface {
	pulumi.Input

	ToRestoreOutput() RestoreOutput
	ToRestoreOutputWithContext(ctx context.Context) RestoreOutput
}

func (*Restore) ElementType() reflect.Type {
	return reflect.TypeOf((**Restore)(nil)).Elem()
}

func (i *Restore) ToRestoreOutput() RestoreOutput {
	return i.ToRestoreOutputWithContext(context.Background())
}

func (i *Restore) ToRestoreOutputWithContext(ctx context.Context) RestoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestoreOutput)
}

type RestoreOutput struct{ *pulumi.OutputState }

func (RestoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Restore)(nil)).Elem()
}

func (o RestoreOutput) ToRestoreOutput() RestoreOutput {
	return o
}

func (o RestoreOutput) ToRestoreOutputWithContext(ctx context.Context) RestoreOutput {
	return o
}

func (o RestoreOutput) Action() ActionPtrOutput {
	return o.ApplyT(func(v *Restore) ActionPtrOutput { return v.Action }).(ActionPtrOutput)
}

// The ID of a backup of the current Droplet instance to restore from.
func (o RestoreOutput) Image() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Restore) pulumi.IntPtrOutput { return v.Image }).(pulumi.IntPtrOutput)
}

// The type of action to initiate for the Droplet.
func (o RestoreOutput) Type() RestoreDropletActionTypePtrOutput {
	return o.ApplyT(func(v *Restore) RestoreDropletActionTypePtrOutput { return v.Type }).(RestoreDropletActionTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RestoreInput)(nil)).Elem(), &Restore{})
	pulumi.RegisterOutputType(RestoreOutput{})
}
