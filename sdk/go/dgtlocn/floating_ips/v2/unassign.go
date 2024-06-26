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

type Unassign struct {
	pulumi.CustomResourceState

	Action ActionPtrOutput `pulumi:"action"`
	// The type of action to initiate for the floating IP.
	Type FloatingIPsActionTypePtrOutput `pulumi:"type"`
}

// NewUnassign registers a new resource with the given unique name, arguments, and options.
func NewUnassign(ctx *pulumi.Context,
	name string, args *UnassignArgs, opts ...pulumi.ResourceOption) (*Unassign, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Unassign
	err := ctx.RegisterResource("digitalocean-native:floating_ips/v2:Unassign", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUnassign gets an existing Unassign resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUnassign(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UnassignState, opts ...pulumi.ResourceOption) (*Unassign, error) {
	var resource Unassign
	err := ctx.ReadResource("digitalocean-native:floating_ips/v2:Unassign", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Unassign resources.
type unassignState struct {
}

type UnassignState struct {
}

func (UnassignState) ElementType() reflect.Type {
	return reflect.TypeOf((*unassignState)(nil)).Elem()
}

type unassignArgs struct {
	// A floating IP address.
	FloatingIp *string `pulumi:"floatingIp"`
	// The type of action to initiate for the floating IP.
	Type FloatingIPsActionType `pulumi:"type"`
}

// The set of arguments for constructing a Unassign resource.
type UnassignArgs struct {
	// A floating IP address.
	FloatingIp pulumi.StringPtrInput
	// The type of action to initiate for the floating IP.
	Type FloatingIPsActionTypeInput
}

func (UnassignArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*unassignArgs)(nil)).Elem()
}

type UnassignInput interface {
	pulumi.Input

	ToUnassignOutput() UnassignOutput
	ToUnassignOutputWithContext(ctx context.Context) UnassignOutput
}

func (*Unassign) ElementType() reflect.Type {
	return reflect.TypeOf((**Unassign)(nil)).Elem()
}

func (i *Unassign) ToUnassignOutput() UnassignOutput {
	return i.ToUnassignOutputWithContext(context.Background())
}

func (i *Unassign) ToUnassignOutputWithContext(ctx context.Context) UnassignOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UnassignOutput)
}

type UnassignOutput struct{ *pulumi.OutputState }

func (UnassignOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Unassign)(nil)).Elem()
}

func (o UnassignOutput) ToUnassignOutput() UnassignOutput {
	return o
}

func (o UnassignOutput) ToUnassignOutputWithContext(ctx context.Context) UnassignOutput {
	return o
}

func (o UnassignOutput) Action() ActionPtrOutput {
	return o.ApplyT(func(v *Unassign) ActionPtrOutput { return v.Action }).(ActionPtrOutput)
}

// The type of action to initiate for the floating IP.
func (o UnassignOutput) Type() FloatingIPsActionTypePtrOutput {
	return o.ApplyT(func(v *Unassign) FloatingIPsActionTypePtrOutput { return v.Type }).(FloatingIPsActionTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UnassignInput)(nil)).Elem(), &Unassign{})
	pulumi.RegisterOutputType(UnassignOutput{})
}
