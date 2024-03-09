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

type ImageActionsConvert struct {
	pulumi.CustomResourceState

	// A time value given in ISO8601 combined date and time format that represents when the action was completed.
	CompletedAt pulumi.StringPtrOutput `pulumi:"completedAt"`
	Region      RegionPtrOutput        `pulumi:"region"`
	RegionSlug  RegionSlugPtrOutput    `pulumi:"regionSlug"`
	// A unique identifier for the resource that the action is associated with.
	ResourceId pulumi.IntPtrOutput `pulumi:"resourceId"`
	// The type of resource that the action is associated with.
	ResourceType pulumi.StringPtrOutput `pulumi:"resourceType"`
	// A time value given in ISO8601 combined date and time format that represents when the action was initiated.
	StartedAt pulumi.StringPtrOutput `pulumi:"startedAt"`
	// The current status of the action. This can be "in-progress", "completed", or "errored".
	Status StatusPtrOutput `pulumi:"status"`
	// This is the type of action that the object represents. For example, this could be "transfer" to represent the state of an image transfer action.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewImageActionsConvert registers a new resource with the given unique name, arguments, and options.
func NewImageActionsConvert(ctx *pulumi.Context,
	name string, args *ImageActionsConvertArgs, opts ...pulumi.ResourceOption) (*ImageActionsConvert, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ImageActionsConvert
	err := ctx.RegisterResource("digitalocean-native:images/v2:ImageActionsConvert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImageActionsConvert gets an existing ImageActionsConvert resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImageActionsConvert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageActionsConvertState, opts ...pulumi.ResourceOption) (*ImageActionsConvert, error) {
	var resource ImageActionsConvert
	err := ctx.ReadResource("digitalocean-native:images/v2:ImageActionsConvert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ImageActionsConvert resources.
type imageActionsConvertState struct {
}

type ImageActionsConvertState struct {
}

func (ImageActionsConvertState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageActionsConvertState)(nil)).Elem()
}

type imageActionsConvertArgs struct {
	// A unique number that can be used to identify and reference a specific image.
	ImageId *string `pulumi:"imageId"`
	// The action to be taken on the image. Can be either `convert` or `transfer`.
	Type Type `pulumi:"type"`
}

// The set of arguments for constructing a ImageActionsConvert resource.
type ImageActionsConvertArgs struct {
	// A unique number that can be used to identify and reference a specific image.
	ImageId pulumi.StringPtrInput
	// The action to be taken on the image. Can be either `convert` or `transfer`.
	Type TypeInput
}

func (ImageActionsConvertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageActionsConvertArgs)(nil)).Elem()
}

type ImageActionsConvertInput interface {
	pulumi.Input

	ToImageActionsConvertOutput() ImageActionsConvertOutput
	ToImageActionsConvertOutputWithContext(ctx context.Context) ImageActionsConvertOutput
}

func (*ImageActionsConvert) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageActionsConvert)(nil)).Elem()
}

func (i *ImageActionsConvert) ToImageActionsConvertOutput() ImageActionsConvertOutput {
	return i.ToImageActionsConvertOutputWithContext(context.Background())
}

func (i *ImageActionsConvert) ToImageActionsConvertOutputWithContext(ctx context.Context) ImageActionsConvertOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageActionsConvertOutput)
}

type ImageActionsConvertOutput struct{ *pulumi.OutputState }

func (ImageActionsConvertOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageActionsConvert)(nil)).Elem()
}

func (o ImageActionsConvertOutput) ToImageActionsConvertOutput() ImageActionsConvertOutput {
	return o
}

func (o ImageActionsConvertOutput) ToImageActionsConvertOutputWithContext(ctx context.Context) ImageActionsConvertOutput {
	return o
}

// A time value given in ISO8601 combined date and time format that represents when the action was completed.
func (o ImageActionsConvertOutput) CompletedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageActionsConvert) pulumi.StringPtrOutput { return v.CompletedAt }).(pulumi.StringPtrOutput)
}

func (o ImageActionsConvertOutput) Region() RegionPtrOutput {
	return o.ApplyT(func(v *ImageActionsConvert) RegionPtrOutput { return v.Region }).(RegionPtrOutput)
}

func (o ImageActionsConvertOutput) RegionSlug() RegionSlugPtrOutput {
	return o.ApplyT(func(v *ImageActionsConvert) RegionSlugPtrOutput { return v.RegionSlug }).(RegionSlugPtrOutput)
}

// A unique identifier for the resource that the action is associated with.
func (o ImageActionsConvertOutput) ResourceId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ImageActionsConvert) pulumi.IntPtrOutput { return v.ResourceId }).(pulumi.IntPtrOutput)
}

// The type of resource that the action is associated with.
func (o ImageActionsConvertOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageActionsConvert) pulumi.StringPtrOutput { return v.ResourceType }).(pulumi.StringPtrOutput)
}

// A time value given in ISO8601 combined date and time format that represents when the action was initiated.
func (o ImageActionsConvertOutput) StartedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageActionsConvert) pulumi.StringPtrOutput { return v.StartedAt }).(pulumi.StringPtrOutput)
}

// The current status of the action. This can be "in-progress", "completed", or "errored".
func (o ImageActionsConvertOutput) Status() StatusPtrOutput {
	return o.ApplyT(func(v *ImageActionsConvert) StatusPtrOutput { return v.Status }).(StatusPtrOutput)
}

// This is the type of action that the object represents. For example, this could be "transfer" to represent the state of an image transfer action.
func (o ImageActionsConvertOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageActionsConvert) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImageActionsConvertInput)(nil)).Elem(), &ImageActionsConvert{})
	pulumi.RegisterOutputType(ImageActionsConvertOutput{})
}
