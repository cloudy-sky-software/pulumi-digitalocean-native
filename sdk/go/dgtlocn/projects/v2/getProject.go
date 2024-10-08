// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProjectResult
	err := ctx.Invoke("digitalocean-native:projects/v2:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProjectArgs struct {
	// A unique identifier for a project.
	ProjectId string `pulumi:"projectId"`
}

type LookupProjectResult struct {
	Project *ProjectType `pulumi:"project"`
}

func LookupProjectOutput(ctx *pulumi.Context, args LookupProjectOutputArgs, opts ...pulumi.InvokeOption) LookupProjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectResultOutput, error) {
			args := v.(LookupProjectArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupProjectResult
			secret, err := ctx.InvokePackageRaw("digitalocean-native:projects/v2:getProject", args, &rv, "", opts...)
			if err != nil {
				return LookupProjectResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupProjectResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupProjectResultOutput), nil
			}
			return output, nil
		}).(LookupProjectResultOutput)
}

type LookupProjectOutputArgs struct {
	// A unique identifier for a project.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (LookupProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectArgs)(nil)).Elem()
}

type LookupProjectResultOutput struct{ *pulumi.OutputState }

func (LookupProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectResult)(nil)).Elem()
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutput() LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutputWithContext(ctx context.Context) LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) Project() ProjectTypePtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *ProjectType { return v.Project }).(ProjectTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectResultOutput{})
}
