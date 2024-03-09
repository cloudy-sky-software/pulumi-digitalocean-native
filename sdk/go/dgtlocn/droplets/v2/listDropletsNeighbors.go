// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDropletsNeighbors(ctx *pulumi.Context, args *ListDropletsNeighborsArgs, opts ...pulumi.InvokeOption) (*ListDropletsNeighborsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListDropletsNeighborsResult
	err := ctx.Invoke("digitalocean-native:droplets/v2:listDropletsNeighbors", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDropletsNeighborsArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId string `pulumi:"dropletId"`
}

type ListDropletsNeighborsResult struct {
	Items ListDropletsNeighbors `pulumi:"items"`
}

func ListDropletsNeighborsOutput(ctx *pulumi.Context, args ListDropletsNeighborsOutputArgs, opts ...pulumi.InvokeOption) ListDropletsNeighborsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDropletsNeighborsResult, error) {
			args := v.(ListDropletsNeighborsArgs)
			r, err := ListDropletsNeighbors(ctx, &args, opts...)
			var s ListDropletsNeighborsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDropletsNeighborsResultOutput)
}

type ListDropletsNeighborsOutputArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId pulumi.StringInput `pulumi:"dropletId"`
}

func (ListDropletsNeighborsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDropletsNeighborsArgs)(nil)).Elem()
}

type ListDropletsNeighborsResultOutput struct{ *pulumi.OutputState }

func (ListDropletsNeighborsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDropletsNeighborsResult)(nil)).Elem()
}

func (o ListDropletsNeighborsResultOutput) ToListDropletsNeighborsResultOutput() ListDropletsNeighborsResultOutput {
	return o
}

func (o ListDropletsNeighborsResultOutput) ToListDropletsNeighborsResultOutputWithContext(ctx context.Context) ListDropletsNeighborsResultOutput {
	return o
}

func (o ListDropletsNeighborsResultOutput) Items() ListDropletsNeighborsOutput {
	return o.ApplyT(func(v ListDropletsNeighborsResult) ListDropletsNeighbors { return v.Items }).(ListDropletsNeighborsOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDropletsNeighborsResultOutput{})
}
