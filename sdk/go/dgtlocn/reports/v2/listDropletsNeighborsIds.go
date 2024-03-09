// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDropletsNeighborsIds(ctx *pulumi.Context, args *ListDropletsNeighborsIdsArgs, opts ...pulumi.InvokeOption) (*ListDropletsNeighborsIdsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListDropletsNeighborsIdsResult
	err := ctx.Invoke("digitalocean-native:reports/v2:listDropletsNeighborsIds", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDropletsNeighborsIdsArgs struct {
}

type ListDropletsNeighborsIdsResult struct {
	Items NeighborIds `pulumi:"items"`
}

func ListDropletsNeighborsIdsOutput(ctx *pulumi.Context, args ListDropletsNeighborsIdsOutputArgs, opts ...pulumi.InvokeOption) ListDropletsNeighborsIdsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDropletsNeighborsIdsResult, error) {
			args := v.(ListDropletsNeighborsIdsArgs)
			r, err := ListDropletsNeighborsIds(ctx, &args, opts...)
			var s ListDropletsNeighborsIdsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDropletsNeighborsIdsResultOutput)
}

type ListDropletsNeighborsIdsOutputArgs struct {
}

func (ListDropletsNeighborsIdsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDropletsNeighborsIdsArgs)(nil)).Elem()
}

type ListDropletsNeighborsIdsResultOutput struct{ *pulumi.OutputState }

func (ListDropletsNeighborsIdsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDropletsNeighborsIdsResult)(nil)).Elem()
}

func (o ListDropletsNeighborsIdsResultOutput) ToListDropletsNeighborsIdsResultOutput() ListDropletsNeighborsIdsResultOutput {
	return o
}

func (o ListDropletsNeighborsIdsResultOutput) ToListDropletsNeighborsIdsResultOutputWithContext(ctx context.Context) ListDropletsNeighborsIdsResultOutput {
	return o
}

func (o ListDropletsNeighborsIdsResultOutput) Items() NeighborIdsOutput {
	return o.ApplyT(func(v ListDropletsNeighborsIdsResult) NeighborIds { return v.Items }).(NeighborIdsOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDropletsNeighborsIdsResultOutput{})
}