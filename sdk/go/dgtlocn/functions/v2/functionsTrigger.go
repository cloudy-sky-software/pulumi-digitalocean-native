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

type FunctionsTrigger struct {
	pulumi.CustomResourceState

	// Name of function(action) that exists in the given namespace.
	Function pulumi.StringOutput `pulumi:"function"`
	// Indicates weather the trigger is paused or unpaused.
	IsEnabled pulumi.BoolOutput `pulumi:"isEnabled"`
	// The trigger's unique name within the namespace.
	Name pulumi.StringOutput `pulumi:"name"`
	// Trigger details for SCHEDULED type, where body is optional.
	ScheduledDetails ScheduledDetailsOutput `pulumi:"scheduledDetails"`
	Trigger          TriggerInfoPtrOutput   `pulumi:"trigger"`
	// One of different type of triggers. Currently only SCHEDULED is supported.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFunctionsTrigger registers a new resource with the given unique name, arguments, and options.
func NewFunctionsTrigger(ctx *pulumi.Context,
	name string, args *FunctionsTriggerArgs, opts ...pulumi.ResourceOption) (*FunctionsTrigger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Function == nil {
		return nil, errors.New("invalid value for required argument 'Function'")
	}
	if args.IsEnabled == nil {
		return nil, errors.New("invalid value for required argument 'IsEnabled'")
	}
	if args.ScheduledDetails == nil {
		return nil, errors.New("invalid value for required argument 'ScheduledDetails'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FunctionsTrigger
	err := ctx.RegisterResource("digitalocean-native:functions/v2:FunctionsTrigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunctionsTrigger gets an existing FunctionsTrigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunctionsTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionsTriggerState, opts ...pulumi.ResourceOption) (*FunctionsTrigger, error) {
	var resource FunctionsTrigger
	err := ctx.ReadResource("digitalocean-native:functions/v2:FunctionsTrigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FunctionsTrigger resources.
type functionsTriggerState struct {
}

type FunctionsTriggerState struct {
}

func (FunctionsTriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionsTriggerState)(nil)).Elem()
}

type functionsTriggerArgs struct {
	// Name of function(action) that exists in the given namespace.
	Function string `pulumi:"function"`
	// Indicates weather the trigger is paused or unpaused.
	IsEnabled bool `pulumi:"isEnabled"`
	// The trigger's unique name within the namespace.
	Name *string `pulumi:"name"`
	// The ID of the namespace to be managed.
	NamespaceId *string `pulumi:"namespaceId"`
	// Trigger details for SCHEDULED type, where body is optional.
	ScheduledDetails ScheduledDetails `pulumi:"scheduledDetails"`
	// One of different type of triggers. Currently only SCHEDULED is supported.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a FunctionsTrigger resource.
type FunctionsTriggerArgs struct {
	// Name of function(action) that exists in the given namespace.
	Function pulumi.StringInput
	// Indicates weather the trigger is paused or unpaused.
	IsEnabled pulumi.BoolInput
	// The trigger's unique name within the namespace.
	Name pulumi.StringPtrInput
	// The ID of the namespace to be managed.
	NamespaceId pulumi.StringPtrInput
	// Trigger details for SCHEDULED type, where body is optional.
	ScheduledDetails ScheduledDetailsInput
	// One of different type of triggers. Currently only SCHEDULED is supported.
	Type pulumi.StringInput
}

func (FunctionsTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionsTriggerArgs)(nil)).Elem()
}

type FunctionsTriggerInput interface {
	pulumi.Input

	ToFunctionsTriggerOutput() FunctionsTriggerOutput
	ToFunctionsTriggerOutputWithContext(ctx context.Context) FunctionsTriggerOutput
}

func (*FunctionsTrigger) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionsTrigger)(nil)).Elem()
}

func (i *FunctionsTrigger) ToFunctionsTriggerOutput() FunctionsTriggerOutput {
	return i.ToFunctionsTriggerOutputWithContext(context.Background())
}

func (i *FunctionsTrigger) ToFunctionsTriggerOutputWithContext(ctx context.Context) FunctionsTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionsTriggerOutput)
}

type FunctionsTriggerOutput struct{ *pulumi.OutputState }

func (FunctionsTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionsTrigger)(nil)).Elem()
}

func (o FunctionsTriggerOutput) ToFunctionsTriggerOutput() FunctionsTriggerOutput {
	return o
}

func (o FunctionsTriggerOutput) ToFunctionsTriggerOutputWithContext(ctx context.Context) FunctionsTriggerOutput {
	return o
}

// Name of function(action) that exists in the given namespace.
func (o FunctionsTriggerOutput) Function() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionsTrigger) pulumi.StringOutput { return v.Function }).(pulumi.StringOutput)
}

// Indicates weather the trigger is paused or unpaused.
func (o FunctionsTriggerOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *FunctionsTrigger) pulumi.BoolOutput { return v.IsEnabled }).(pulumi.BoolOutput)
}

// The trigger's unique name within the namespace.
func (o FunctionsTriggerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionsTrigger) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Trigger details for SCHEDULED type, where body is optional.
func (o FunctionsTriggerOutput) ScheduledDetails() ScheduledDetailsOutput {
	return o.ApplyT(func(v *FunctionsTrigger) ScheduledDetailsOutput { return v.ScheduledDetails }).(ScheduledDetailsOutput)
}

func (o FunctionsTriggerOutput) Trigger() TriggerInfoPtrOutput {
	return o.ApplyT(func(v *FunctionsTrigger) TriggerInfoPtrOutput { return v.Trigger }).(TriggerInfoPtrOutput)
}

// One of different type of triggers. Currently only SCHEDULED is supported.
func (o FunctionsTriggerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionsTrigger) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionsTriggerInput)(nil)).Elem(), &FunctionsTrigger{})
	pulumi.RegisterOutputType(FunctionsTriggerOutput{})
}