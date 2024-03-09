// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type ListOneClicksProperties struct {
	_1Clicks []OneClicks `pulumi:"_1Clicks"`
}

type ListOneClicksPropertiesOutput struct{ *pulumi.OutputState }

func (ListOneClicksPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListOneClicksProperties)(nil)).Elem()
}

func (o ListOneClicksPropertiesOutput) ToListOneClicksPropertiesOutput() ListOneClicksPropertiesOutput {
	return o
}

func (o ListOneClicksPropertiesOutput) ToListOneClicksPropertiesOutputWithContext(ctx context.Context) ListOneClicksPropertiesOutput {
	return o
}

func (o ListOneClicksPropertiesOutput) _1Clicks() OneClicksArrayOutput {
	return o.ApplyT(func(v ListOneClicksProperties) []OneClicks { return v._1Clicks }).(OneClicksArrayOutput)
}

type OneClicks struct {
	// The slug identifier for the 1-Click application.
	Slug string `pulumi:"slug"`
	// The type of the 1-Click application.
	Type string `pulumi:"type"`
}

type OneClicksOutput struct{ *pulumi.OutputState }

func (OneClicksOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OneClicks)(nil)).Elem()
}

func (o OneClicksOutput) ToOneClicksOutput() OneClicksOutput {
	return o
}

func (o OneClicksOutput) ToOneClicksOutputWithContext(ctx context.Context) OneClicksOutput {
	return o
}

// The slug identifier for the 1-Click application.
func (o OneClicksOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v OneClicks) string { return v.Slug }).(pulumi.StringOutput)
}

// The type of the 1-Click application.
func (o OneClicksOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v OneClicks) string { return v.Type }).(pulumi.StringOutput)
}

type OneClicksArrayOutput struct{ *pulumi.OutputState }

func (OneClicksArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OneClicks)(nil)).Elem()
}

func (o OneClicksArrayOutput) ToOneClicksArrayOutput() OneClicksArrayOutput {
	return o
}

func (o OneClicksArrayOutput) ToOneClicksArrayOutputWithContext(ctx context.Context) OneClicksArrayOutput {
	return o
}

func (o OneClicksArrayOutput) Index(i pulumi.IntInput) OneClicksOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OneClicks {
		return vs[0].([]OneClicks)[vs[1].(int)]
	}).(OneClicksOutput)
}

func init() {
	pulumi.RegisterOutputType(ListOneClicksPropertiesOutput{})
	pulumi.RegisterOutputType(OneClicksOutput{})
	pulumi.RegisterOutputType(OneClicksArrayOutput{})
}
