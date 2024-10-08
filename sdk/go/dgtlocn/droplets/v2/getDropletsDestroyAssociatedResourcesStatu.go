// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDropletsDestroyAssociatedResourcesStatu(ctx *pulumi.Context, args *GetDropletsDestroyAssociatedResourcesStatuArgs, opts ...pulumi.InvokeOption) (*GetDropletsDestroyAssociatedResourcesStatuResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDropletsDestroyAssociatedResourcesStatuResult
	err := ctx.Invoke("digitalocean-native:droplets/v2:getDropletsDestroyAssociatedResourcesStatu", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDropletsDestroyAssociatedResourcesStatuArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId string `pulumi:"dropletId"`
}

// An objects containing information about a resources scheduled for deletion.
type GetDropletsDestroyAssociatedResourcesStatuResult struct {
	// A time value given in ISO8601 combined date and time format indicating when the requested action was completed.
	CompletedAt *string `pulumi:"completedAt"`
	// An object containing information about a resource scheduled for deletion.
	Droplet *DestroyedAssociatedResource `pulumi:"droplet"`
	// A count of the associated resources that failed to be destroyed, if any.
	Failures *int `pulumi:"failures"`
	// An object containing additional information about resource related to a Droplet requested to be destroyed.
	Resources *AssociatedResourceStatusResourcesProperties `pulumi:"resources"`
}

func GetDropletsDestroyAssociatedResourcesStatuOutput(ctx *pulumi.Context, args GetDropletsDestroyAssociatedResourcesStatuOutputArgs, opts ...pulumi.InvokeOption) GetDropletsDestroyAssociatedResourcesStatuResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDropletsDestroyAssociatedResourcesStatuResultOutput, error) {
			args := v.(GetDropletsDestroyAssociatedResourcesStatuArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetDropletsDestroyAssociatedResourcesStatuResult
			secret, err := ctx.InvokePackageRaw("digitalocean-native:droplets/v2:getDropletsDestroyAssociatedResourcesStatu", args, &rv, "", opts...)
			if err != nil {
				return GetDropletsDestroyAssociatedResourcesStatuResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetDropletsDestroyAssociatedResourcesStatuResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetDropletsDestroyAssociatedResourcesStatuResultOutput), nil
			}
			return output, nil
		}).(GetDropletsDestroyAssociatedResourcesStatuResultOutput)
}

type GetDropletsDestroyAssociatedResourcesStatuOutputArgs struct {
	// A unique identifier for a Droplet instance.
	DropletId pulumi.StringInput `pulumi:"dropletId"`
}

func (GetDropletsDestroyAssociatedResourcesStatuOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDropletsDestroyAssociatedResourcesStatuArgs)(nil)).Elem()
}

// An objects containing information about a resources scheduled for deletion.
type GetDropletsDestroyAssociatedResourcesStatuResultOutput struct{ *pulumi.OutputState }

func (GetDropletsDestroyAssociatedResourcesStatuResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDropletsDestroyAssociatedResourcesStatuResult)(nil)).Elem()
}

func (o GetDropletsDestroyAssociatedResourcesStatuResultOutput) ToGetDropletsDestroyAssociatedResourcesStatuResultOutput() GetDropletsDestroyAssociatedResourcesStatuResultOutput {
	return o
}

func (o GetDropletsDestroyAssociatedResourcesStatuResultOutput) ToGetDropletsDestroyAssociatedResourcesStatuResultOutputWithContext(ctx context.Context) GetDropletsDestroyAssociatedResourcesStatuResultOutput {
	return o
}

// A time value given in ISO8601 combined date and time format indicating when the requested action was completed.
func (o GetDropletsDestroyAssociatedResourcesStatuResultOutput) CompletedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDropletsDestroyAssociatedResourcesStatuResult) *string { return v.CompletedAt }).(pulumi.StringPtrOutput)
}

// An object containing information about a resource scheduled for deletion.
func (o GetDropletsDestroyAssociatedResourcesStatuResultOutput) Droplet() DestroyedAssociatedResourcePtrOutput {
	return o.ApplyT(func(v GetDropletsDestroyAssociatedResourcesStatuResult) *DestroyedAssociatedResource {
		return v.Droplet
	}).(DestroyedAssociatedResourcePtrOutput)
}

// A count of the associated resources that failed to be destroyed, if any.
func (o GetDropletsDestroyAssociatedResourcesStatuResultOutput) Failures() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetDropletsDestroyAssociatedResourcesStatuResult) *int { return v.Failures }).(pulumi.IntPtrOutput)
}

// An object containing additional information about resource related to a Droplet requested to be destroyed.
func (o GetDropletsDestroyAssociatedResourcesStatuResultOutput) Resources() AssociatedResourceStatusResourcesPropertiesPtrOutput {
	return o.ApplyT(func(v GetDropletsDestroyAssociatedResourcesStatuResult) *AssociatedResourceStatusResourcesProperties {
		return v.Resources
	}).(AssociatedResourceStatusResourcesPropertiesPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDropletsDestroyAssociatedResourcesStatuResultOutput{})
}
