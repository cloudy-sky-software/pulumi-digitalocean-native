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

type LoadBalancer struct {
	pulumi.CustomResourceState

	// This field has been deprecated. You can no longer specify an algorithm for load balancers.
	Algorithm LoadBalancerBaseAlgorithmPtrOutput `pulumi:"algorithm"`
	// A time value given in ISO8601 combined date and time format that represents when the load balancer was created.
	CreatedAt pulumi.StringPtrOutput `pulumi:"createdAt"`
	// A boolean value indicating whether to disable automatic DNS record creation for Let's Encrypt certificates that are added to the load balancer.
	DisableLetsEncryptDnsRecords pulumi.BoolPtrOutput `pulumi:"disableLetsEncryptDnsRecords"`
	// A boolean value indicating whether HTTP keepalive connections are maintained to target Droplets.
	EnableBackendKeepalive pulumi.BoolPtrOutput `pulumi:"enableBackendKeepalive"`
	// A boolean value indicating whether PROXY Protocol is in use.
	EnableProxyProtocol pulumi.BoolPtrOutput `pulumi:"enableProxyProtocol"`
	// An object specifying allow and deny rules to control traffic to the load balancer.
	Firewall LbFirewallPtrOutput `pulumi:"firewall"`
	// An array of objects specifying the forwarding rules for a load balancer.
	ForwardingRules ForwardingRuleArrayOutput `pulumi:"forwardingRules"`
	// An object specifying health check settings for the load balancer.
	HealthCheck HealthCheckPtrOutput `pulumi:"healthCheck"`
	// An integer value which configures the idle timeout for HTTP requests to the target droplets.
	HttpIdleTimeoutSeconds pulumi.IntPtrOutput `pulumi:"httpIdleTimeoutSeconds"`
	// An attribute containing the public-facing IP address of the load balancer.
	Ip           pulumi.StringPtrOutput    `pulumi:"ip"`
	LoadBalancer LoadBalancerTypePtrOutput `pulumi:"loadBalancer"`
	// A human-readable name for a load balancer instance.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The ID of the project that the load balancer is associated with. If no ID is provided at creation, the load balancer associates with the user's default project. If an invalid project ID is provided, the load balancer will not be created.
	ProjectId pulumi.StringPtrOutput `pulumi:"projectId"`
	// A boolean value indicating whether HTTP requests to the load balancer on port 80 will be redirected to HTTPS on port 443.
	RedirectHttpToHttps pulumi.BoolPtrOutput `pulumi:"redirectHttpToHttps"`
	// The slug identifier for the region where the resource will initially be  available.
	Region LoadBalancerPropertiesRegionEnumPtrOutput `pulumi:"region"`
	// This field has been replaced by the `size_unit` field for all regions except in AMS2, NYC2, and SFO1. Each available load balancer size now equates to the load balancer having a set number of nodes.
	// * `lb-small` = 1 node
	// * `lb-medium` = 3 nodes
	// * `lb-large` = 6 nodes
	//
	// You can resize load balancers after creation up to once per hour. You cannot resize a load balancer within the first hour of its creation.
	Size LoadBalancerBaseSizePtrOutput `pulumi:"size"`
	// How many nodes the load balancer contains. Each additional node increases the load balancer's ability to manage more connections. Load balancers can be scaled up or down, and you can change the number of nodes after creation up to once per hour. This field is currently not available in the AMS2, NYC2, or SFO1 regions. Use the `size` field to scale load balancers that reside in these regions.
	SizeUnit pulumi.IntPtrOutput `pulumi:"sizeUnit"`
	// A status string indicating the current state of the load balancer. This can be `new`, `active`, or `errored`.
	Status LoadBalancerBaseStatusPtrOutput `pulumi:"status"`
	// An object specifying sticky sessions settings for the load balancer.
	StickySessions StickySessionsPtrOutput `pulumi:"stickySessions"`
	// A string specifying the UUID of the VPC to which the load balancer is assigned.
	VpcUuid pulumi.StringPtrOutput `pulumi:"vpcUuid"`
}

// NewLoadBalancer registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancer(ctx *pulumi.Context,
	name string, args *LoadBalancerArgs, opts ...pulumi.ResourceOption) (*LoadBalancer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ForwardingRules == nil {
		return nil, errors.New("invalid value for required argument 'ForwardingRules'")
	}
	if args.Algorithm == nil {
		args.Algorithm = LoadBalancerBaseAlgorithm("round_robin")
	}
	if args.DisableLetsEncryptDnsRecords == nil {
		args.DisableLetsEncryptDnsRecords = pulumi.BoolPtr(false)
	}
	if args.EnableBackendKeepalive == nil {
		args.EnableBackendKeepalive = pulumi.BoolPtr(false)
	}
	if args.EnableProxyProtocol == nil {
		args.EnableProxyProtocol = pulumi.BoolPtr(false)
	}
	if args.HealthCheck != nil {
		args.HealthCheck = args.HealthCheck.ToHealthCheckPtrOutput().ApplyT(func(v *HealthCheck) *HealthCheck { return v.Defaults() }).(HealthCheckPtrOutput)
	}
	if args.HttpIdleTimeoutSeconds == nil {
		args.HttpIdleTimeoutSeconds = pulumi.IntPtr(60)
	}
	if args.RedirectHttpToHttps == nil {
		args.RedirectHttpToHttps = pulumi.BoolPtr(false)
	}
	if args.Size == nil {
		args.Size = LoadBalancerBaseSize("lb-small")
	}
	if args.SizeUnit == nil {
		args.SizeUnit = pulumi.IntPtr(1)
	}
	if args.StickySessions != nil {
		args.StickySessions = args.StickySessions.ToStickySessionsPtrOutput().ApplyT(func(v *StickySessions) *StickySessions { return v.Defaults() }).(StickySessionsPtrOutput)
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LoadBalancer
	err := ctx.RegisterResource("digitalocean-native:load_balancers/v2:LoadBalancer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadBalancer gets an existing LoadBalancer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadBalancerState, opts ...pulumi.ResourceOption) (*LoadBalancer, error) {
	var resource LoadBalancer
	err := ctx.ReadResource("digitalocean-native:load_balancers/v2:LoadBalancer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoadBalancer resources.
type loadBalancerState struct {
}

type LoadBalancerState struct {
}

func (LoadBalancerState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerState)(nil)).Elem()
}

type loadBalancerArgs struct {
	// This field has been deprecated. You can no longer specify an algorithm for load balancers.
	Algorithm *LoadBalancerBaseAlgorithm `pulumi:"algorithm"`
	// A time value given in ISO8601 combined date and time format that represents when the load balancer was created.
	CreatedAt *string `pulumi:"createdAt"`
	// A boolean value indicating whether to disable automatic DNS record creation for Let's Encrypt certificates that are added to the load balancer.
	DisableLetsEncryptDnsRecords *bool `pulumi:"disableLetsEncryptDnsRecords"`
	// A boolean value indicating whether HTTP keepalive connections are maintained to target Droplets.
	EnableBackendKeepalive *bool `pulumi:"enableBackendKeepalive"`
	// A boolean value indicating whether PROXY Protocol is in use.
	EnableProxyProtocol *bool `pulumi:"enableProxyProtocol"`
	// An object specifying allow and deny rules to control traffic to the load balancer.
	Firewall *LbFirewall `pulumi:"firewall"`
	// An array of objects specifying the forwarding rules for a load balancer.
	ForwardingRules []ForwardingRule `pulumi:"forwardingRules"`
	// An object specifying health check settings for the load balancer.
	HealthCheck *HealthCheck `pulumi:"healthCheck"`
	// An integer value which configures the idle timeout for HTTP requests to the target droplets.
	HttpIdleTimeoutSeconds *int `pulumi:"httpIdleTimeoutSeconds"`
	// An attribute containing the public-facing IP address of the load balancer.
	Ip *string `pulumi:"ip"`
	// A human-readable name for a load balancer instance.
	Name *string `pulumi:"name"`
	// The ID of the project that the load balancer is associated with. If no ID is provided at creation, the load balancer associates with the user's default project. If an invalid project ID is provided, the load balancer will not be created.
	ProjectId *string `pulumi:"projectId"`
	// A boolean value indicating whether HTTP requests to the load balancer on port 80 will be redirected to HTTPS on port 443.
	RedirectHttpToHttps *bool `pulumi:"redirectHttpToHttps"`
	// The slug identifier for the region where the resource will initially be  available.
	Region *LoadBalancerPropertiesRegionEnum `pulumi:"region"`
	// This field has been replaced by the `size_unit` field for all regions except in AMS2, NYC2, and SFO1. Each available load balancer size now equates to the load balancer having a set number of nodes.
	// * `lb-small` = 1 node
	// * `lb-medium` = 3 nodes
	// * `lb-large` = 6 nodes
	//
	// You can resize load balancers after creation up to once per hour. You cannot resize a load balancer within the first hour of its creation.
	Size *LoadBalancerBaseSize `pulumi:"size"`
	// How many nodes the load balancer contains. Each additional node increases the load balancer's ability to manage more connections. Load balancers can be scaled up or down, and you can change the number of nodes after creation up to once per hour. This field is currently not available in the AMS2, NYC2, or SFO1 regions. Use the `size` field to scale load balancers that reside in these regions.
	SizeUnit *int `pulumi:"sizeUnit"`
	// A status string indicating the current state of the load balancer. This can be `new`, `active`, or `errored`.
	Status *LoadBalancerBaseStatus `pulumi:"status"`
	// An object specifying sticky sessions settings for the load balancer.
	StickySessions *StickySessions `pulumi:"stickySessions"`
	// A string specifying the UUID of the VPC to which the load balancer is assigned.
	VpcUuid *string `pulumi:"vpcUuid"`
}

// The set of arguments for constructing a LoadBalancer resource.
type LoadBalancerArgs struct {
	// This field has been deprecated. You can no longer specify an algorithm for load balancers.
	Algorithm LoadBalancerBaseAlgorithmPtrInput
	// A time value given in ISO8601 combined date and time format that represents when the load balancer was created.
	CreatedAt pulumi.StringPtrInput
	// A boolean value indicating whether to disable automatic DNS record creation for Let's Encrypt certificates that are added to the load balancer.
	DisableLetsEncryptDnsRecords pulumi.BoolPtrInput
	// A boolean value indicating whether HTTP keepalive connections are maintained to target Droplets.
	EnableBackendKeepalive pulumi.BoolPtrInput
	// A boolean value indicating whether PROXY Protocol is in use.
	EnableProxyProtocol pulumi.BoolPtrInput
	// An object specifying allow and deny rules to control traffic to the load balancer.
	Firewall LbFirewallPtrInput
	// An array of objects specifying the forwarding rules for a load balancer.
	ForwardingRules ForwardingRuleArrayInput
	// An object specifying health check settings for the load balancer.
	HealthCheck HealthCheckPtrInput
	// An integer value which configures the idle timeout for HTTP requests to the target droplets.
	HttpIdleTimeoutSeconds pulumi.IntPtrInput
	// An attribute containing the public-facing IP address of the load balancer.
	Ip pulumi.StringPtrInput
	// A human-readable name for a load balancer instance.
	Name pulumi.StringPtrInput
	// The ID of the project that the load balancer is associated with. If no ID is provided at creation, the load balancer associates with the user's default project. If an invalid project ID is provided, the load balancer will not be created.
	ProjectId pulumi.StringPtrInput
	// A boolean value indicating whether HTTP requests to the load balancer on port 80 will be redirected to HTTPS on port 443.
	RedirectHttpToHttps pulumi.BoolPtrInput
	// The slug identifier for the region where the resource will initially be  available.
	Region LoadBalancerPropertiesRegionEnumPtrInput
	// This field has been replaced by the `size_unit` field for all regions except in AMS2, NYC2, and SFO1. Each available load balancer size now equates to the load balancer having a set number of nodes.
	// * `lb-small` = 1 node
	// * `lb-medium` = 3 nodes
	// * `lb-large` = 6 nodes
	//
	// You can resize load balancers after creation up to once per hour. You cannot resize a load balancer within the first hour of its creation.
	Size LoadBalancerBaseSizePtrInput
	// How many nodes the load balancer contains. Each additional node increases the load balancer's ability to manage more connections. Load balancers can be scaled up or down, and you can change the number of nodes after creation up to once per hour. This field is currently not available in the AMS2, NYC2, or SFO1 regions. Use the `size` field to scale load balancers that reside in these regions.
	SizeUnit pulumi.IntPtrInput
	// A status string indicating the current state of the load balancer. This can be `new`, `active`, or `errored`.
	Status LoadBalancerBaseStatusPtrInput
	// An object specifying sticky sessions settings for the load balancer.
	StickySessions StickySessionsPtrInput
	// A string specifying the UUID of the VPC to which the load balancer is assigned.
	VpcUuid pulumi.StringPtrInput
}

func (LoadBalancerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerArgs)(nil)).Elem()
}

type LoadBalancerInput interface {
	pulumi.Input

	ToLoadBalancerOutput() LoadBalancerOutput
	ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput
}

func (*LoadBalancer) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancer)(nil)).Elem()
}

func (i *LoadBalancer) ToLoadBalancerOutput() LoadBalancerOutput {
	return i.ToLoadBalancerOutputWithContext(context.Background())
}

func (i *LoadBalancer) ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerOutput)
}

type LoadBalancerOutput struct{ *pulumi.OutputState }

func (LoadBalancerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancer)(nil)).Elem()
}

func (o LoadBalancerOutput) ToLoadBalancerOutput() LoadBalancerOutput {
	return o
}

func (o LoadBalancerOutput) ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput {
	return o
}

// This field has been deprecated. You can no longer specify an algorithm for load balancers.
func (o LoadBalancerOutput) Algorithm() LoadBalancerBaseAlgorithmPtrOutput {
	return o.ApplyT(func(v *LoadBalancer) LoadBalancerBaseAlgorithmPtrOutput { return v.Algorithm }).(LoadBalancerBaseAlgorithmPtrOutput)
}

// A time value given in ISO8601 combined date and time format that represents when the load balancer was created.
func (o LoadBalancerOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// A boolean value indicating whether to disable automatic DNS record creation for Let's Encrypt certificates that are added to the load balancer.
func (o LoadBalancerOutput) DisableLetsEncryptDnsRecords() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.BoolPtrOutput { return v.DisableLetsEncryptDnsRecords }).(pulumi.BoolPtrOutput)
}

// A boolean value indicating whether HTTP keepalive connections are maintained to target Droplets.
func (o LoadBalancerOutput) EnableBackendKeepalive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.BoolPtrOutput { return v.EnableBackendKeepalive }).(pulumi.BoolPtrOutput)
}

// A boolean value indicating whether PROXY Protocol is in use.
func (o LoadBalancerOutput) EnableProxyProtocol() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.BoolPtrOutput { return v.EnableProxyProtocol }).(pulumi.BoolPtrOutput)
}

// An object specifying allow and deny rules to control traffic to the load balancer.
func (o LoadBalancerOutput) Firewall() LbFirewallPtrOutput {
	return o.ApplyT(func(v *LoadBalancer) LbFirewallPtrOutput { return v.Firewall }).(LbFirewallPtrOutput)
}

// An array of objects specifying the forwarding rules for a load balancer.
func (o LoadBalancerOutput) ForwardingRules() ForwardingRuleArrayOutput {
	return o.ApplyT(func(v *LoadBalancer) ForwardingRuleArrayOutput { return v.ForwardingRules }).(ForwardingRuleArrayOutput)
}

// An object specifying health check settings for the load balancer.
func (o LoadBalancerOutput) HealthCheck() HealthCheckPtrOutput {
	return o.ApplyT(func(v *LoadBalancer) HealthCheckPtrOutput { return v.HealthCheck }).(HealthCheckPtrOutput)
}

// An integer value which configures the idle timeout for HTTP requests to the target droplets.
func (o LoadBalancerOutput) HttpIdleTimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.IntPtrOutput { return v.HttpIdleTimeoutSeconds }).(pulumi.IntPtrOutput)
}

// An attribute containing the public-facing IP address of the load balancer.
func (o LoadBalancerOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringPtrOutput { return v.Ip }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerOutput) LoadBalancer() LoadBalancerTypePtrOutput {
	return o.ApplyT(func(v *LoadBalancer) LoadBalancerTypePtrOutput { return v.LoadBalancer }).(LoadBalancerTypePtrOutput)
}

// A human-readable name for a load balancer instance.
func (o LoadBalancerOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The ID of the project that the load balancer is associated with. If no ID is provided at creation, the load balancer associates with the user's default project. If an invalid project ID is provided, the load balancer will not be created.
func (o LoadBalancerOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringPtrOutput { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// A boolean value indicating whether HTTP requests to the load balancer on port 80 will be redirected to HTTPS on port 443.
func (o LoadBalancerOutput) RedirectHttpToHttps() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.BoolPtrOutput { return v.RedirectHttpToHttps }).(pulumi.BoolPtrOutput)
}

// The slug identifier for the region where the resource will initially be  available.
func (o LoadBalancerOutput) Region() LoadBalancerPropertiesRegionEnumPtrOutput {
	return o.ApplyT(func(v *LoadBalancer) LoadBalancerPropertiesRegionEnumPtrOutput { return v.Region }).(LoadBalancerPropertiesRegionEnumPtrOutput)
}

// This field has been replaced by the `size_unit` field for all regions except in AMS2, NYC2, and SFO1. Each available load balancer size now equates to the load balancer having a set number of nodes.
// * `lb-small` = 1 node
// * `lb-medium` = 3 nodes
// * `lb-large` = 6 nodes
//
// You can resize load balancers after creation up to once per hour. You cannot resize a load balancer within the first hour of its creation.
func (o LoadBalancerOutput) Size() LoadBalancerBaseSizePtrOutput {
	return o.ApplyT(func(v *LoadBalancer) LoadBalancerBaseSizePtrOutput { return v.Size }).(LoadBalancerBaseSizePtrOutput)
}

// How many nodes the load balancer contains. Each additional node increases the load balancer's ability to manage more connections. Load balancers can be scaled up or down, and you can change the number of nodes after creation up to once per hour. This field is currently not available in the AMS2, NYC2, or SFO1 regions. Use the `size` field to scale load balancers that reside in these regions.
func (o LoadBalancerOutput) SizeUnit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.IntPtrOutput { return v.SizeUnit }).(pulumi.IntPtrOutput)
}

// A status string indicating the current state of the load balancer. This can be `new`, `active`, or `errored`.
func (o LoadBalancerOutput) Status() LoadBalancerBaseStatusPtrOutput {
	return o.ApplyT(func(v *LoadBalancer) LoadBalancerBaseStatusPtrOutput { return v.Status }).(LoadBalancerBaseStatusPtrOutput)
}

// An object specifying sticky sessions settings for the load balancer.
func (o LoadBalancerOutput) StickySessions() StickySessionsPtrOutput {
	return o.ApplyT(func(v *LoadBalancer) StickySessionsPtrOutput { return v.StickySessions }).(StickySessionsPtrOutput)
}

// A string specifying the UUID of the VPC to which the load balancer is assigned.
func (o LoadBalancerOutput) VpcUuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringPtrOutput { return v.VpcUuid }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerInput)(nil)).Elem(), &LoadBalancer{})
	pulumi.RegisterOutputType(LoadBalancerOutput{})
}