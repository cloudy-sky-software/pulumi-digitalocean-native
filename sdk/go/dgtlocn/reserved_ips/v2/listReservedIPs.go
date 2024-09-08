// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListReservedIPs(ctx *pulumi.Context, args *ListReservedIPsArgs, opts ...pulumi.InvokeOption) (*ListReservedIPsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListReservedIPsResult
	err := ctx.Invoke("digitalocean-native:reserved_ips/v2:listReservedIPs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListReservedIPsArgs struct {
}

type ListReservedIPsResult struct {
	Links       *PageLinks       `pulumi:"links"`
	Meta        MetaMeta         `pulumi:"meta"`
	ReservedIps []ReservedIpType `pulumi:"reservedIps"`
}

func ListReservedIPsOutput(ctx *pulumi.Context, args ListReservedIPsOutputArgs, opts ...pulumi.InvokeOption) ListReservedIPsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListReservedIPsResult, error) {
			args := v.(ListReservedIPsArgs)
			r, err := ListReservedIPs(ctx, &args, opts...)
			var s ListReservedIPsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListReservedIPsResultOutput)
}

type ListReservedIPsOutputArgs struct {
}

func (ListReservedIPsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListReservedIPsArgs)(nil)).Elem()
}

type ListReservedIPsResultOutput struct{ *pulumi.OutputState }

func (ListReservedIPsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListReservedIPsResult)(nil)).Elem()
}

func (o ListReservedIPsResultOutput) ToListReservedIPsResultOutput() ListReservedIPsResultOutput {
	return o
}

func (o ListReservedIPsResultOutput) ToListReservedIPsResultOutputWithContext(ctx context.Context) ListReservedIPsResultOutput {
	return o
}

func (o ListReservedIPsResultOutput) Links() PageLinksPtrOutput {
	return o.ApplyT(func(v ListReservedIPsResult) *PageLinks { return v.Links }).(PageLinksPtrOutput)
}

func (o ListReservedIPsResultOutput) Meta() MetaMetaOutput {
	return o.ApplyT(func(v ListReservedIPsResult) MetaMeta { return v.Meta }).(MetaMetaOutput)
}

func (o ListReservedIPsResultOutput) ReservedIps() ReservedIpTypeArrayOutput {
	return o.ApplyT(func(v ListReservedIPsResult) []ReservedIpType { return v.ReservedIps }).(ReservedIpTypeArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListReservedIPsResultOutput{})
}
