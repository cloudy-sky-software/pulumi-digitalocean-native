// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDatabasesOptions(ctx *pulumi.Context, args *ListDatabasesOptionsArgs, opts ...pulumi.InvokeOption) (*ListDatabasesOptionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListDatabasesOptionsResult
	err := ctx.Invoke("digitalocean-native:databases/v2:listDatabasesOptions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDatabasesOptionsArgs struct {
}

type ListDatabasesOptionsResult struct {
	Options             *OptionsOptionsProperties             `pulumi:"options"`
	VersionAvailability *OptionsVersionAvailabilityProperties `pulumi:"versionAvailability"`
}

func ListDatabasesOptionsOutput(ctx *pulumi.Context, args ListDatabasesOptionsOutputArgs, opts ...pulumi.InvokeOption) ListDatabasesOptionsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListDatabasesOptionsResultOutput, error) {
			args := v.(ListDatabasesOptionsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("digitalocean-native:databases/v2:listDatabasesOptions", args, ListDatabasesOptionsResultOutput{}, options).(ListDatabasesOptionsResultOutput), nil
		}).(ListDatabasesOptionsResultOutput)
}

type ListDatabasesOptionsOutputArgs struct {
}

func (ListDatabasesOptionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabasesOptionsArgs)(nil)).Elem()
}

type ListDatabasesOptionsResultOutput struct{ *pulumi.OutputState }

func (ListDatabasesOptionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabasesOptionsResult)(nil)).Elem()
}

func (o ListDatabasesOptionsResultOutput) ToListDatabasesOptionsResultOutput() ListDatabasesOptionsResultOutput {
	return o
}

func (o ListDatabasesOptionsResultOutput) ToListDatabasesOptionsResultOutputWithContext(ctx context.Context) ListDatabasesOptionsResultOutput {
	return o
}

func (o ListDatabasesOptionsResultOutput) Options() OptionsOptionsPropertiesPtrOutput {
	return o.ApplyT(func(v ListDatabasesOptionsResult) *OptionsOptionsProperties { return v.Options }).(OptionsOptionsPropertiesPtrOutput)
}

func (o ListDatabasesOptionsResultOutput) VersionAvailability() OptionsVersionAvailabilityPropertiesPtrOutput {
	return o.ApplyT(func(v ListDatabasesOptionsResult) *OptionsVersionAvailabilityProperties { return v.VersionAvailability }).(OptionsVersionAvailabilityPropertiesPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDatabasesOptionsResultOutput{})
}
