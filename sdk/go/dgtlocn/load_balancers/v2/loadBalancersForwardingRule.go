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

type LoadBalancersForwardingRule struct {
	pulumi.CustomResourceState

	ForwardingRules ForwardingRuleArrayOutput `pulumi:"forwardingRules"`
}

// NewLoadBalancersForwardingRule registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancersForwardingRule(ctx *pulumi.Context,
	name string, args *LoadBalancersForwardingRuleArgs, opts ...pulumi.ResourceOption) (*LoadBalancersForwardingRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ForwardingRules == nil {
		return nil, errors.New("invalid value for required argument 'ForwardingRules'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LoadBalancersForwardingRule
	err := ctx.RegisterResource("digitalocean-native:load_balancers/v2:LoadBalancersForwardingRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadBalancersForwardingRule gets an existing LoadBalancersForwardingRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancersForwardingRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadBalancersForwardingRuleState, opts ...pulumi.ResourceOption) (*LoadBalancersForwardingRule, error) {
	var resource LoadBalancersForwardingRule
	err := ctx.ReadResource("digitalocean-native:load_balancers/v2:LoadBalancersForwardingRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoadBalancersForwardingRule resources.
type loadBalancersForwardingRuleState struct {
}

type LoadBalancersForwardingRuleState struct {
}

func (LoadBalancersForwardingRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancersForwardingRuleState)(nil)).Elem()
}

type loadBalancersForwardingRuleArgs struct {
	ForwardingRules []ForwardingRule `pulumi:"forwardingRules"`
	// A unique identifier for a load balancer.
	LbId *string `pulumi:"lbId"`
}

// The set of arguments for constructing a LoadBalancersForwardingRule resource.
type LoadBalancersForwardingRuleArgs struct {
	ForwardingRules ForwardingRuleArrayInput
	// A unique identifier for a load balancer.
	LbId pulumi.StringPtrInput
}

func (LoadBalancersForwardingRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancersForwardingRuleArgs)(nil)).Elem()
}

type LoadBalancersForwardingRuleInput interface {
	pulumi.Input

	ToLoadBalancersForwardingRuleOutput() LoadBalancersForwardingRuleOutput
	ToLoadBalancersForwardingRuleOutputWithContext(ctx context.Context) LoadBalancersForwardingRuleOutput
}

func (*LoadBalancersForwardingRule) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancersForwardingRule)(nil)).Elem()
}

func (i *LoadBalancersForwardingRule) ToLoadBalancersForwardingRuleOutput() LoadBalancersForwardingRuleOutput {
	return i.ToLoadBalancersForwardingRuleOutputWithContext(context.Background())
}

func (i *LoadBalancersForwardingRule) ToLoadBalancersForwardingRuleOutputWithContext(ctx context.Context) LoadBalancersForwardingRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancersForwardingRuleOutput)
}

type LoadBalancersForwardingRuleOutput struct{ *pulumi.OutputState }

func (LoadBalancersForwardingRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancersForwardingRule)(nil)).Elem()
}

func (o LoadBalancersForwardingRuleOutput) ToLoadBalancersForwardingRuleOutput() LoadBalancersForwardingRuleOutput {
	return o
}

func (o LoadBalancersForwardingRuleOutput) ToLoadBalancersForwardingRuleOutputWithContext(ctx context.Context) LoadBalancersForwardingRuleOutput {
	return o
}

func (o LoadBalancersForwardingRuleOutput) ForwardingRules() ForwardingRuleArrayOutput {
	return o.ApplyT(func(v *LoadBalancersForwardingRule) ForwardingRuleArrayOutput { return v.ForwardingRules }).(ForwardingRuleArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancersForwardingRuleInput)(nil)).Elem(), &LoadBalancersForwardingRule{})
	pulumi.RegisterOutputType(LoadBalancersForwardingRuleOutput{})
}
