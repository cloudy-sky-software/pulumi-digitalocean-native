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

type Resize struct {
	pulumi.CustomResourceState

	Action ActionPtrOutput `pulumi:"action"`
	// When `true`, the Droplet's disk will be resized in addition to its RAM and CPU. This is a permanent change and cannot be reversed as a Droplet's disk size cannot be decreased.
	Disk pulumi.BoolPtrOutput `pulumi:"disk"`
	// The slug identifier for the size to which you wish to resize the Droplet.
	Size pulumi.StringPtrOutput `pulumi:"size"`
	// The type of action to initiate for the Droplet.
	Type DropletActionTypePtrOutput `pulumi:"type"`
}

// NewResize registers a new resource with the given unique name, arguments, and options.
func NewResize(ctx *pulumi.Context,
	name string, args *ResizeArgs, opts ...pulumi.ResourceOption) (*Resize, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Resize
	err := ctx.RegisterResource("digitalocean-native:droplets/v2:Resize", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResize gets an existing Resize resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResize(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResizeState, opts ...pulumi.ResourceOption) (*Resize, error) {
	var resource Resize
	err := ctx.ReadResource("digitalocean-native:droplets/v2:Resize", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Resize resources.
type resizeState struct {
}

type ResizeState struct {
}

func (ResizeState) ElementType() reflect.Type {
	return reflect.TypeOf((*resizeState)(nil)).Elem()
}

type resizeArgs struct {
	// When `true`, the Droplet's disk will be resized in addition to its RAM and CPU. This is a permanent change and cannot be reversed as a Droplet's disk size cannot be decreased.
	Disk *bool `pulumi:"disk"`
	// A unique identifier for a Droplet instance.
	DropletId *string `pulumi:"dropletId"`
	// The slug identifier for the size to which you wish to resize the Droplet.
	Size *string `pulumi:"size"`
	// The type of action to initiate for the Droplet.
	Type DropletActionType `pulumi:"type"`
}

// The set of arguments for constructing a Resize resource.
type ResizeArgs struct {
	// When `true`, the Droplet's disk will be resized in addition to its RAM and CPU. This is a permanent change and cannot be reversed as a Droplet's disk size cannot be decreased.
	Disk pulumi.BoolPtrInput
	// A unique identifier for a Droplet instance.
	DropletId pulumi.StringPtrInput
	// The slug identifier for the size to which you wish to resize the Droplet.
	Size pulumi.StringPtrInput
	// The type of action to initiate for the Droplet.
	Type DropletActionTypeInput
}

func (ResizeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resizeArgs)(nil)).Elem()
}

type ResizeInput interface {
	pulumi.Input

	ToResizeOutput() ResizeOutput
	ToResizeOutputWithContext(ctx context.Context) ResizeOutput
}

func (*Resize) ElementType() reflect.Type {
	return reflect.TypeOf((**Resize)(nil)).Elem()
}

func (i *Resize) ToResizeOutput() ResizeOutput {
	return i.ToResizeOutputWithContext(context.Background())
}

func (i *Resize) ToResizeOutputWithContext(ctx context.Context) ResizeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResizeOutput)
}

type ResizeOutput struct{ *pulumi.OutputState }

func (ResizeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Resize)(nil)).Elem()
}

func (o ResizeOutput) ToResizeOutput() ResizeOutput {
	return o
}

func (o ResizeOutput) ToResizeOutputWithContext(ctx context.Context) ResizeOutput {
	return o
}

func (o ResizeOutput) Action() ActionPtrOutput {
	return o.ApplyT(func(v *Resize) ActionPtrOutput { return v.Action }).(ActionPtrOutput)
}

// When `true`, the Droplet's disk will be resized in addition to its RAM and CPU. This is a permanent change and cannot be reversed as a Droplet's disk size cannot be decreased.
func (o ResizeOutput) Disk() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Resize) pulumi.BoolPtrOutput { return v.Disk }).(pulumi.BoolPtrOutput)
}

// The slug identifier for the size to which you wish to resize the Droplet.
func (o ResizeOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Resize) pulumi.StringPtrOutput { return v.Size }).(pulumi.StringPtrOutput)
}

// The type of action to initiate for the Droplet.
func (o ResizeOutput) Type() DropletActionTypePtrOutput {
	return o.ApplyT(func(v *Resize) DropletActionTypePtrOutput { return v.Type }).(DropletActionTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResizeInput)(nil)).Elem(), &Resize{})
	pulumi.RegisterOutputType(ResizeOutput{})
}
