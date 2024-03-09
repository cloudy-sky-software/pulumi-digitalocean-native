// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLoadBalancers(ctx *pulumi.Context, args *LookupLoadBalancersArgs, opts ...pulumi.InvokeOption) (*LookupLoadBalancersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLoadBalancersResult
	err := ctx.Invoke("digitalocean-native:load_balancers/v2:getLoadBalancers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupLoadBalancersArgs struct {
	// A unique identifier for a load balancer.
	LbId string `pulumi:"lbId"`
}

type LookupLoadBalancersResult struct {
	Items GetLoadBalancersProperties `pulumi:"items"`
}

// Defaults sets the appropriate defaults for LookupLoadBalancersResult
func (val *LookupLoadBalancersResult) Defaults() *LookupLoadBalancersResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Items = *tmp.Items.Defaults()

	return &tmp
}

func LookupLoadBalancersOutput(ctx *pulumi.Context, args LookupLoadBalancersOutputArgs, opts ...pulumi.InvokeOption) LookupLoadBalancersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLoadBalancersResult, error) {
			args := v.(LookupLoadBalancersArgs)
			r, err := LookupLoadBalancers(ctx, &args, opts...)
			var s LookupLoadBalancersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLoadBalancersResultOutput)
}

type LookupLoadBalancersOutputArgs struct {
	// A unique identifier for a load balancer.
	LbId pulumi.StringInput `pulumi:"lbId"`
}

func (LookupLoadBalancersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoadBalancersArgs)(nil)).Elem()
}

type LookupLoadBalancersResultOutput struct{ *pulumi.OutputState }

func (LookupLoadBalancersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoadBalancersResult)(nil)).Elem()
}

func (o LookupLoadBalancersResultOutput) ToLookupLoadBalancersResultOutput() LookupLoadBalancersResultOutput {
	return o
}

func (o LookupLoadBalancersResultOutput) ToLookupLoadBalancersResultOutputWithContext(ctx context.Context) LookupLoadBalancersResultOutput {
	return o
}

func (o LookupLoadBalancersResultOutput) Items() GetLoadBalancersPropertiesOutput {
	return o.ApplyT(func(v LookupLoadBalancersResult) GetLoadBalancersProperties { return v.Items }).(GetLoadBalancersPropertiesOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLoadBalancersResultOutput{})
}
