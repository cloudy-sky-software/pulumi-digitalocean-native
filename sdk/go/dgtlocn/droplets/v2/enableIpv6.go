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
type EnableIpv6 struct {
	pulumi.CustomResourceState

	Action ActionPtrOutput `pulumi:"action"`
	// The type of action to initiate for the Droplet.
	Type TypeOutput `pulumi:"type"`
}

// NewEnableIpv6 registers a new resource with the given unique name, arguments, and options.
func NewEnableIpv6(ctx *pulumi.Context,
	name string, args *EnableIpv6Args, opts ...pulumi.ResourceOption) (*EnableIpv6, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EnableIpv6
	err := ctx.RegisterResource("digitalocean-native:droplets/v2:EnableIpv6", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnableIpv6 gets an existing EnableIpv6 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnableIpv6(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnableIpv6State, opts ...pulumi.ResourceOption) (*EnableIpv6, error) {
	var resource EnableIpv6
	err := ctx.ReadResource("digitalocean-native:droplets/v2:EnableIpv6", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnableIpv6 resources.
type enableIpv6State struct {
}

type EnableIpv6State struct {
}

func (EnableIpv6State) ElementType() reflect.Type {
	return reflect.TypeOf((*enableIpv6State)(nil)).Elem()
}

type enableIpv6Args struct {
	// A unique identifier for a Droplet instance.
	DropletId *string `pulumi:"dropletId"`
	// The type of action to initiate for the Droplet.
	Type Type `pulumi:"type"`
}

// The set of arguments for constructing a EnableIpv6 resource.
type EnableIpv6Args struct {
	// A unique identifier for a Droplet instance.
	DropletId pulumi.StringPtrInput
	// The type of action to initiate for the Droplet.
	Type TypeInput
}

func (EnableIpv6Args) ElementType() reflect.Type {
	return reflect.TypeOf((*enableIpv6Args)(nil)).Elem()
}

type EnableIpv6Input interface {
	pulumi.Input

	ToEnableIpv6Output() EnableIpv6Output
	ToEnableIpv6OutputWithContext(ctx context.Context) EnableIpv6Output
}

func (*EnableIpv6) ElementType() reflect.Type {
	return reflect.TypeOf((**EnableIpv6)(nil)).Elem()
}

func (i *EnableIpv6) ToEnableIpv6Output() EnableIpv6Output {
	return i.ToEnableIpv6OutputWithContext(context.Background())
}

func (i *EnableIpv6) ToEnableIpv6OutputWithContext(ctx context.Context) EnableIpv6Output {
	return pulumi.ToOutputWithContext(ctx, i).(EnableIpv6Output)
}

type EnableIpv6Output struct{ *pulumi.OutputState }

func (EnableIpv6Output) ElementType() reflect.Type {
	return reflect.TypeOf((**EnableIpv6)(nil)).Elem()
}

func (o EnableIpv6Output) ToEnableIpv6Output() EnableIpv6Output {
	return o
}

func (o EnableIpv6Output) ToEnableIpv6OutputWithContext(ctx context.Context) EnableIpv6Output {
	return o
}

func (o EnableIpv6Output) Action() ActionPtrOutput {
	return o.ApplyT(func(v *EnableIpv6) ActionPtrOutput { return v.Action }).(ActionPtrOutput)
}

// The type of action to initiate for the Droplet.
func (o EnableIpv6Output) Type() TypeOutput {
	return o.ApplyT(func(v *EnableIpv6) TypeOutput { return v.Type }).(TypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnableIpv6Input)(nil)).Elem(), &EnableIpv6{})
	pulumi.RegisterOutputType(EnableIpv6Output{})
}