// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListProjectsResources(ctx *pulumi.Context, args *ListProjectsResourcesArgs, opts ...pulumi.InvokeOption) (*ListProjectsResourcesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListProjectsResourcesResult
	err := ctx.Invoke("digitalocean-native:projects/v2:listProjectsResources", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListProjectsResourcesArgs struct {
	// A unique identifier for a project.
	ProjectId string `pulumi:"projectId"`
}

type ListProjectsResourcesResult struct {
	Items ListProjectsResourcesItems `pulumi:"items"`
}

func ListProjectsResourcesOutput(ctx *pulumi.Context, args ListProjectsResourcesOutputArgs, opts ...pulumi.InvokeOption) ListProjectsResourcesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListProjectsResourcesResult, error) {
			args := v.(ListProjectsResourcesArgs)
			r, err := ListProjectsResources(ctx, &args, opts...)
			var s ListProjectsResourcesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListProjectsResourcesResultOutput)
}

type ListProjectsResourcesOutputArgs struct {
	// A unique identifier for a project.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (ListProjectsResourcesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListProjectsResourcesArgs)(nil)).Elem()
}

type ListProjectsResourcesResultOutput struct{ *pulumi.OutputState }

func (ListProjectsResourcesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListProjectsResourcesResult)(nil)).Elem()
}

func (o ListProjectsResourcesResultOutput) ToListProjectsResourcesResultOutput() ListProjectsResourcesResultOutput {
	return o
}

func (o ListProjectsResourcesResultOutput) ToListProjectsResourcesResultOutputWithContext(ctx context.Context) ListProjectsResourcesResultOutput {
	return o
}

func (o ListProjectsResourcesResultOutput) Items() ListProjectsResourcesItemsOutput {
	return o.ApplyT(func(v ListProjectsResourcesResult) ListProjectsResourcesItems { return v.Items }).(ListProjectsResourcesItemsOutput)
}

func init() {
	pulumi.RegisterOutputType(ListProjectsResourcesResultOutput{})
}
