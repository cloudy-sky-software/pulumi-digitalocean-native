// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListUptimeChecks(ctx *pulumi.Context, args *ListUptimeChecksArgs, opts ...pulumi.InvokeOption) (*ListUptimeChecksResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListUptimeChecksResult
	err := ctx.Invoke("digitalocean-native:uptime/v2:listUptimeChecks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListUptimeChecksArgs struct {
}

type ListUptimeChecksResult struct {
	Items ListUptimeChecks `pulumi:"items"`
}

func ListUptimeChecksOutput(ctx *pulumi.Context, args ListUptimeChecksOutputArgs, opts ...pulumi.InvokeOption) ListUptimeChecksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListUptimeChecksResult, error) {
			args := v.(ListUptimeChecksArgs)
			r, err := ListUptimeChecks(ctx, &args, opts...)
			var s ListUptimeChecksResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListUptimeChecksResultOutput)
}

type ListUptimeChecksOutputArgs struct {
}

func (ListUptimeChecksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListUptimeChecksArgs)(nil)).Elem()
}

type ListUptimeChecksResultOutput struct{ *pulumi.OutputState }

func (ListUptimeChecksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListUptimeChecksResult)(nil)).Elem()
}

func (o ListUptimeChecksResultOutput) ToListUptimeChecksResultOutput() ListUptimeChecksResultOutput {
	return o
}

func (o ListUptimeChecksResultOutput) ToListUptimeChecksResultOutputWithContext(ctx context.Context) ListUptimeChecksResultOutput {
	return o
}

func (o ListUptimeChecksResultOutput) Items() ListUptimeChecksOutput {
	return o.ApplyT(func(v ListUptimeChecksResult) ListUptimeChecks { return v.Items }).(ListUptimeChecksOutput)
}

func init() {
	pulumi.RegisterOutputType(ListUptimeChecksResultOutput{})
}
