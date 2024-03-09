// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDomains(ctx *pulumi.Context, args *ListDomainsArgs, opts ...pulumi.InvokeOption) (*ListDomainsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListDomainsResult
	err := ctx.Invoke("digitalocean-native:domains/v2:listDomains", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDomainsArgs struct {
}

type ListDomainsResult struct {
	Items ListDomains `pulumi:"items"`
}

func ListDomainsOutput(ctx *pulumi.Context, args ListDomainsOutputArgs, opts ...pulumi.InvokeOption) ListDomainsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDomainsResult, error) {
			args := v.(ListDomainsArgs)
			r, err := ListDomains(ctx, &args, opts...)
			var s ListDomainsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDomainsResultOutput)
}

type ListDomainsOutputArgs struct {
}

func (ListDomainsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDomainsArgs)(nil)).Elem()
}

type ListDomainsResultOutput struct{ *pulumi.OutputState }

func (ListDomainsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDomainsResult)(nil)).Elem()
}

func (o ListDomainsResultOutput) ToListDomainsResultOutput() ListDomainsResultOutput {
	return o
}

func (o ListDomainsResultOutput) ToListDomainsResultOutputWithContext(ctx context.Context) ListDomainsResultOutput {
	return o
}

func (o ListDomainsResultOutput) Items() ListDomainsOutput {
	return o.ApplyT(func(v ListDomainsResult) ListDomains { return v.Items }).(ListDomainsOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDomainsResultOutput{})
}