// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListOneClicks(ctx *pulumi.Context, args *ListOneClicksArgs, opts ...pulumi.InvokeOption) (*ListOneClicksResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListOneClicksResult
	err := ctx.Invoke("digitalocean-native:1-clicks/v2:listOneClicks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListOneClicksArgs struct {
}

type ListOneClicksResult struct {
	_1Clicks []OneClicks `pulumi:"_1Clicks"`
}

func ListOneClicksOutput(ctx *pulumi.Context, args ListOneClicksOutputArgs, opts ...pulumi.InvokeOption) ListOneClicksResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListOneClicksResultOutput, error) {
			args := v.(ListOneClicksArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("digitalocean-native:1-clicks/v2:listOneClicks", args, ListOneClicksResultOutput{}, options).(ListOneClicksResultOutput), nil
		}).(ListOneClicksResultOutput)
}

type ListOneClicksOutputArgs struct {
}

func (ListOneClicksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListOneClicksArgs)(nil)).Elem()
}

type ListOneClicksResultOutput struct{ *pulumi.OutputState }

func (ListOneClicksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListOneClicksResult)(nil)).Elem()
}

func (o ListOneClicksResultOutput) ToListOneClicksResultOutput() ListOneClicksResultOutput {
	return o
}

func (o ListOneClicksResultOutput) ToListOneClicksResultOutputWithContext(ctx context.Context) ListOneClicksResultOutput {
	return o
}

func (o ListOneClicksResultOutput) _1Clicks() OneClicksArrayOutput {
	return o.ApplyT(func(v ListOneClicksResult) []OneClicks { return v._1Clicks }).(OneClicksArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListOneClicksResultOutput{})
}
