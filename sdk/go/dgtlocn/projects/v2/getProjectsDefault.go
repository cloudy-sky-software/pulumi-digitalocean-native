// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProjectsDefault(ctx *pulumi.Context, args *LookupProjectsDefaultArgs, opts ...pulumi.InvokeOption) (*LookupProjectsDefaultResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProjectsDefaultResult
	err := ctx.Invoke("digitalocean-native:projects/v2:getProjectsDefault", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProjectsDefaultArgs struct {
}

type LookupProjectsDefaultResult struct {
	Items GetProjectsDefaultProperties `pulumi:"items"`
}

func LookupProjectsDefaultOutput(ctx *pulumi.Context, args LookupProjectsDefaultOutputArgs, opts ...pulumi.InvokeOption) LookupProjectsDefaultResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectsDefaultResult, error) {
			args := v.(LookupProjectsDefaultArgs)
			r, err := LookupProjectsDefault(ctx, &args, opts...)
			var s LookupProjectsDefaultResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProjectsDefaultResultOutput)
}

type LookupProjectsDefaultOutputArgs struct {
}

func (LookupProjectsDefaultOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectsDefaultArgs)(nil)).Elem()
}

type LookupProjectsDefaultResultOutput struct{ *pulumi.OutputState }

func (LookupProjectsDefaultResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectsDefaultResult)(nil)).Elem()
}

func (o LookupProjectsDefaultResultOutput) ToLookupProjectsDefaultResultOutput() LookupProjectsDefaultResultOutput {
	return o
}

func (o LookupProjectsDefaultResultOutput) ToLookupProjectsDefaultResultOutputWithContext(ctx context.Context) LookupProjectsDefaultResultOutput {
	return o
}

func (o LookupProjectsDefaultResultOutput) Items() GetProjectsDefaultPropertiesOutput {
	return o.ApplyT(func(v LookupProjectsDefaultResult) GetProjectsDefaultProperties { return v.Items }).(GetProjectsDefaultPropertiesOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectsDefaultResultOutput{})
}
