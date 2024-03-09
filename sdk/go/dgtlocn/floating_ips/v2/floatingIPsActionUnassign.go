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

type FloatingIPsActionUnassign struct {
	pulumi.CustomResourceState

	Action ActionPtrOutput `pulumi:"action"`
	// The type of action to initiate for the floating IP.
	Type FloatingIPsActionTypePtrOutput `pulumi:"type"`
}

// NewFloatingIPsActionUnassign registers a new resource with the given unique name, arguments, and options.
func NewFloatingIPsActionUnassign(ctx *pulumi.Context,
	name string, args *FloatingIPsActionUnassignArgs, opts ...pulumi.ResourceOption) (*FloatingIPsActionUnassign, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FloatingIPsActionUnassign
	err := ctx.RegisterResource("digitalocean-native:floating_ips/v2:FloatingIPsActionUnassign", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFloatingIPsActionUnassign gets an existing FloatingIPsActionUnassign resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFloatingIPsActionUnassign(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FloatingIPsActionUnassignState, opts ...pulumi.ResourceOption) (*FloatingIPsActionUnassign, error) {
	var resource FloatingIPsActionUnassign
	err := ctx.ReadResource("digitalocean-native:floating_ips/v2:FloatingIPsActionUnassign", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FloatingIPsActionUnassign resources.
type floatingIPsActionUnassignState struct {
}

type FloatingIPsActionUnassignState struct {
}

func (FloatingIPsActionUnassignState) ElementType() reflect.Type {
	return reflect.TypeOf((*floatingIPsActionUnassignState)(nil)).Elem()
}

type floatingIPsActionUnassignArgs struct {
	// A floating IP address.
	FloatingIp *string `pulumi:"floatingIp"`
	// The type of action to initiate for the floating IP.
	Type FloatingIPsActionType `pulumi:"type"`
}

// The set of arguments for constructing a FloatingIPsActionUnassign resource.
type FloatingIPsActionUnassignArgs struct {
	// A floating IP address.
	FloatingIp pulumi.StringPtrInput
	// The type of action to initiate for the floating IP.
	Type FloatingIPsActionTypeInput
}

func (FloatingIPsActionUnassignArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*floatingIPsActionUnassignArgs)(nil)).Elem()
}

type FloatingIPsActionUnassignInput interface {
	pulumi.Input

	ToFloatingIPsActionUnassignOutput() FloatingIPsActionUnassignOutput
	ToFloatingIPsActionUnassignOutputWithContext(ctx context.Context) FloatingIPsActionUnassignOutput
}

func (*FloatingIPsActionUnassign) ElementType() reflect.Type {
	return reflect.TypeOf((**FloatingIPsActionUnassign)(nil)).Elem()
}

func (i *FloatingIPsActionUnassign) ToFloatingIPsActionUnassignOutput() FloatingIPsActionUnassignOutput {
	return i.ToFloatingIPsActionUnassignOutputWithContext(context.Background())
}

func (i *FloatingIPsActionUnassign) ToFloatingIPsActionUnassignOutputWithContext(ctx context.Context) FloatingIPsActionUnassignOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FloatingIPsActionUnassignOutput)
}

type FloatingIPsActionUnassignOutput struct{ *pulumi.OutputState }

func (FloatingIPsActionUnassignOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FloatingIPsActionUnassign)(nil)).Elem()
}

func (o FloatingIPsActionUnassignOutput) ToFloatingIPsActionUnassignOutput() FloatingIPsActionUnassignOutput {
	return o
}

func (o FloatingIPsActionUnassignOutput) ToFloatingIPsActionUnassignOutputWithContext(ctx context.Context) FloatingIPsActionUnassignOutput {
	return o
}

func (o FloatingIPsActionUnassignOutput) Action() ActionPtrOutput {
	return o.ApplyT(func(v *FloatingIPsActionUnassign) ActionPtrOutput { return v.Action }).(ActionPtrOutput)
}

// The type of action to initiate for the floating IP.
func (o FloatingIPsActionUnassignOutput) Type() FloatingIPsActionTypePtrOutput {
	return o.ApplyT(func(v *FloatingIPsActionUnassign) FloatingIPsActionTypePtrOutput { return v.Type }).(FloatingIPsActionTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FloatingIPsActionUnassignInput)(nil)).Elem(), &FloatingIPsActionUnassign{})
	pulumi.RegisterOutputType(FloatingIPsActionUnassignOutput{})
}