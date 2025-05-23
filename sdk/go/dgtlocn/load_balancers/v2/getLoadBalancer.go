// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLoadBalancer(ctx *pulumi.Context, args *LookupLoadBalancerArgs, opts ...pulumi.InvokeOption) (*LookupLoadBalancerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLoadBalancerResult
	err := ctx.Invoke("digitalocean-native:load_balancers/v2:getLoadBalancer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupLoadBalancerArgs struct {
	// A unique identifier for a load balancer.
	LbId string `pulumi:"lbId"`
}

type LookupLoadBalancerResult struct {
	LoadBalancer *LoadBalancerType `pulumi:"loadBalancer"`
}

// Defaults sets the appropriate defaults for LookupLoadBalancerResult
func (val *LookupLoadBalancerResult) Defaults() *LookupLoadBalancerResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.LoadBalancer = tmp.LoadBalancer.Defaults()

	return &tmp
}
func LookupLoadBalancerOutput(ctx *pulumi.Context, args LookupLoadBalancerOutputArgs, opts ...pulumi.InvokeOption) LookupLoadBalancerResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupLoadBalancerResultOutput, error) {
			args := v.(LookupLoadBalancerArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("digitalocean-native:load_balancers/v2:getLoadBalancer", args, LookupLoadBalancerResultOutput{}, options).(LookupLoadBalancerResultOutput), nil
		}).(LookupLoadBalancerResultOutput)
}

type LookupLoadBalancerOutputArgs struct {
	// A unique identifier for a load balancer.
	LbId pulumi.StringInput `pulumi:"lbId"`
}

func (LookupLoadBalancerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoadBalancerArgs)(nil)).Elem()
}

type LookupLoadBalancerResultOutput struct{ *pulumi.OutputState }

func (LookupLoadBalancerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoadBalancerResult)(nil)).Elem()
}

func (o LookupLoadBalancerResultOutput) ToLookupLoadBalancerResultOutput() LookupLoadBalancerResultOutput {
	return o
}

func (o LookupLoadBalancerResultOutput) ToLookupLoadBalancerResultOutputWithContext(ctx context.Context) LookupLoadBalancerResultOutput {
	return o
}

func (o LookupLoadBalancerResultOutput) LoadBalancer() LoadBalancerTypePtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *LoadBalancerType { return v.LoadBalancer }).(LoadBalancerTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLoadBalancerResultOutput{})
}
