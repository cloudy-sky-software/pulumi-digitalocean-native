// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDropletsBackups(ctx *pulumi.Context, args *ListDropletsBackupsArgs, opts ...pulumi.InvokeOption) (*ListDropletsBackupsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListDropletsBackupsResult
	err := ctx.Invoke("digitalocean-native:droplets/v2:listDropletsBackups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDropletsBackupsArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId string `pulumi:"dropletId"`
}

type ListDropletsBackupsResult struct {
	Items ListDropletsBackupsItems `pulumi:"items"`
}

func ListDropletsBackupsOutput(ctx *pulumi.Context, args ListDropletsBackupsOutputArgs, opts ...pulumi.InvokeOption) ListDropletsBackupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDropletsBackupsResult, error) {
			args := v.(ListDropletsBackupsArgs)
			r, err := ListDropletsBackups(ctx, &args, opts...)
			var s ListDropletsBackupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDropletsBackupsResultOutput)
}

type ListDropletsBackupsOutputArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId pulumi.StringInput `pulumi:"dropletId"`
}

func (ListDropletsBackupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDropletsBackupsArgs)(nil)).Elem()
}

type ListDropletsBackupsResultOutput struct{ *pulumi.OutputState }

func (ListDropletsBackupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDropletsBackupsResult)(nil)).Elem()
}

func (o ListDropletsBackupsResultOutput) ToListDropletsBackupsResultOutput() ListDropletsBackupsResultOutput {
	return o
}

func (o ListDropletsBackupsResultOutput) ToListDropletsBackupsResultOutputWithContext(ctx context.Context) ListDropletsBackupsResultOutput {
	return o
}

func (o ListDropletsBackupsResultOutput) Items() ListDropletsBackupsItemsOutput {
	return o.ApplyT(func(v ListDropletsBackupsResult) ListDropletsBackupsItems { return v.Items }).(ListDropletsBackupsItemsOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDropletsBackupsResultOutput{})
}
