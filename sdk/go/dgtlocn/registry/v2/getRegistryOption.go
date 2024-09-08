// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetRegistryOption(ctx *pulumi.Context, args *GetRegistryOptionArgs, opts ...pulumi.InvokeOption) (*GetRegistryOptionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRegistryOptionResult
	err := ctx.Invoke("digitalocean-native:registry/v2:getRegistryOption", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetRegistryOptionArgs struct {
}

type GetRegistryOptionResult struct {
	Options *GetRegistryOptionPropertiesOptionsProperties `pulumi:"options"`
}

func GetRegistryOptionOutput(ctx *pulumi.Context, args GetRegistryOptionOutputArgs, opts ...pulumi.InvokeOption) GetRegistryOptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRegistryOptionResult, error) {
			args := v.(GetRegistryOptionArgs)
			r, err := GetRegistryOption(ctx, &args, opts...)
			var s GetRegistryOptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRegistryOptionResultOutput)
}

type GetRegistryOptionOutputArgs struct {
}

func (GetRegistryOptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRegistryOptionArgs)(nil)).Elem()
}

type GetRegistryOptionResultOutput struct{ *pulumi.OutputState }

func (GetRegistryOptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRegistryOptionResult)(nil)).Elem()
}

func (o GetRegistryOptionResultOutput) ToGetRegistryOptionResultOutput() GetRegistryOptionResultOutput {
	return o
}

func (o GetRegistryOptionResultOutput) ToGetRegistryOptionResultOutputWithContext(ctx context.Context) GetRegistryOptionResultOutput {
	return o
}

func (o GetRegistryOptionResultOutput) Options() GetRegistryOptionPropertiesOptionsPropertiesPtrOutput {
	return o.ApplyT(func(v GetRegistryOptionResult) *GetRegistryOptionPropertiesOptionsProperties { return v.Options }).(GetRegistryOptionPropertiesOptionsPropertiesPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRegistryOptionResultOutput{})
}
