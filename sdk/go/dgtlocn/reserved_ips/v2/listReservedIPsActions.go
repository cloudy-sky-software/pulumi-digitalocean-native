// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListReservedIPsActions(ctx *pulumi.Context, args *ListReservedIPsActionsArgs, opts ...pulumi.InvokeOption) (*ListReservedIPsActionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListReservedIPsActionsResult
	err := ctx.Invoke("digitalocean-native:reserved_ips/v2:listReservedIPsActions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListReservedIPsActionsArgs struct {
	// A reserved IP address.
	ReservedIp string `pulumi:"reservedIp"`
}

type ListReservedIPsActionsResult struct {
	Items ListReservedIPsActionsItems `pulumi:"items"`
}

func ListReservedIPsActionsOutput(ctx *pulumi.Context, args ListReservedIPsActionsOutputArgs, opts ...pulumi.InvokeOption) ListReservedIPsActionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListReservedIPsActionsResult, error) {
			args := v.(ListReservedIPsActionsArgs)
			r, err := ListReservedIPsActions(ctx, &args, opts...)
			var s ListReservedIPsActionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListReservedIPsActionsResultOutput)
}

type ListReservedIPsActionsOutputArgs struct {
	// A reserved IP address.
	ReservedIp pulumi.StringInput `pulumi:"reservedIp"`
}

func (ListReservedIPsActionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListReservedIPsActionsArgs)(nil)).Elem()
}

type ListReservedIPsActionsResultOutput struct{ *pulumi.OutputState }

func (ListReservedIPsActionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListReservedIPsActionsResult)(nil)).Elem()
}

func (o ListReservedIPsActionsResultOutput) ToListReservedIPsActionsResultOutput() ListReservedIPsActionsResultOutput {
	return o
}

func (o ListReservedIPsActionsResultOutput) ToListReservedIPsActionsResultOutputWithContext(ctx context.Context) ListReservedIPsActionsResultOutput {
	return o
}

func (o ListReservedIPsActionsResultOutput) Items() ListReservedIPsActionsItemsOutput {
	return o.ApplyT(func(v ListReservedIPsActionsResult) ListReservedIPsActionsItems { return v.Items }).(ListReservedIPsActionsItemsOutput)
}

func init() {
	pulumi.RegisterOutputType(ListReservedIPsActionsResultOutput{})
}
