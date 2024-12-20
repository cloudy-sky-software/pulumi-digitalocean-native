// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDropletsKernels(ctx *pulumi.Context, args *ListDropletsKernelsArgs, opts ...pulumi.InvokeOption) (*ListDropletsKernelsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListDropletsKernelsResult
	err := ctx.Invoke("digitalocean-native:droplets/v2:listDropletsKernels", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDropletsKernelsArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId string `pulumi:"dropletId"`
}

type ListDropletsKernelsResult struct {
	Kernels []Kernel   `pulumi:"kernels"`
	Links   *PageLinks `pulumi:"links"`
	Meta    MetaMeta   `pulumi:"meta"`
}

func ListDropletsKernelsOutput(ctx *pulumi.Context, args ListDropletsKernelsOutputArgs, opts ...pulumi.InvokeOption) ListDropletsKernelsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListDropletsKernelsResultOutput, error) {
			args := v.(ListDropletsKernelsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("digitalocean-native:droplets/v2:listDropletsKernels", args, ListDropletsKernelsResultOutput{}, options).(ListDropletsKernelsResultOutput), nil
		}).(ListDropletsKernelsResultOutput)
}

type ListDropletsKernelsOutputArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId pulumi.StringInput `pulumi:"dropletId"`
}

func (ListDropletsKernelsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDropletsKernelsArgs)(nil)).Elem()
}

type ListDropletsKernelsResultOutput struct{ *pulumi.OutputState }

func (ListDropletsKernelsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDropletsKernelsResult)(nil)).Elem()
}

func (o ListDropletsKernelsResultOutput) ToListDropletsKernelsResultOutput() ListDropletsKernelsResultOutput {
	return o
}

func (o ListDropletsKernelsResultOutput) ToListDropletsKernelsResultOutputWithContext(ctx context.Context) ListDropletsKernelsResultOutput {
	return o
}

func (o ListDropletsKernelsResultOutput) Kernels() KernelArrayOutput {
	return o.ApplyT(func(v ListDropletsKernelsResult) []Kernel { return v.Kernels }).(KernelArrayOutput)
}

func (o ListDropletsKernelsResultOutput) Links() PageLinksPtrOutput {
	return o.ApplyT(func(v ListDropletsKernelsResult) *PageLinks { return v.Links }).(PageLinksPtrOutput)
}

func (o ListDropletsKernelsResultOutput) Meta() MetaMetaOutput {
	return o.ApplyT(func(v ListDropletsKernelsResult) MetaMeta { return v.Meta }).(MetaMetaOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDropletsKernelsResultOutput{})
}
