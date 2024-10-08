// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListAppsTiers(ctx *pulumi.Context, args *ListAppsTiersArgs, opts ...pulumi.InvokeOption) (*ListAppsTiersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListAppsTiersResult
	err := ctx.Invoke("digitalocean-native:apps/v2:listAppsTiers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAppsTiersArgs struct {
}

type ListAppsTiersResult struct {
	Tiers []AppsTier `pulumi:"tiers"`
}

func ListAppsTiersOutput(ctx *pulumi.Context, args ListAppsTiersOutputArgs, opts ...pulumi.InvokeOption) ListAppsTiersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListAppsTiersResultOutput, error) {
			args := v.(ListAppsTiersArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv ListAppsTiersResult
			secret, err := ctx.InvokePackageRaw("digitalocean-native:apps/v2:listAppsTiers", args, &rv, "", opts...)
			if err != nil {
				return ListAppsTiersResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(ListAppsTiersResultOutput)
			if secret {
				return pulumi.ToSecret(output).(ListAppsTiersResultOutput), nil
			}
			return output, nil
		}).(ListAppsTiersResultOutput)
}

type ListAppsTiersOutputArgs struct {
}

func (ListAppsTiersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAppsTiersArgs)(nil)).Elem()
}

type ListAppsTiersResultOutput struct{ *pulumi.OutputState }

func (ListAppsTiersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAppsTiersResult)(nil)).Elem()
}

func (o ListAppsTiersResultOutput) ToListAppsTiersResultOutput() ListAppsTiersResultOutput {
	return o
}

func (o ListAppsTiersResultOutput) ToListAppsTiersResultOutputWithContext(ctx context.Context) ListAppsTiersResultOutput {
	return o
}

func (o ListAppsTiersResultOutput) Tiers() AppsTierArrayOutput {
	return o.ApplyT(func(v ListAppsTiersResult) []AppsTier { return v.Tiers }).(AppsTierArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListAppsTiersResultOutput{})
}
