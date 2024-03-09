// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListRegistryRepositories(ctx *pulumi.Context, args *ListRegistryRepositoriesArgs, opts ...pulumi.InvokeOption) (*ListRegistryRepositoriesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListRegistryRepositoriesResult
	err := ctx.Invoke("digitalocean-native:registry/v2:listRegistryRepositories", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRegistryRepositoriesArgs struct {
	// The name of a container registry.
	RegistryName string `pulumi:"registryName"`
}

type ListRegistryRepositoriesResult struct {
	Items ListRegistryRepositories `pulumi:"items"`
}

func ListRegistryRepositoriesOutput(ctx *pulumi.Context, args ListRegistryRepositoriesOutputArgs, opts ...pulumi.InvokeOption) ListRegistryRepositoriesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListRegistryRepositoriesResult, error) {
			args := v.(ListRegistryRepositoriesArgs)
			r, err := ListRegistryRepositories(ctx, &args, opts...)
			var s ListRegistryRepositoriesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListRegistryRepositoriesResultOutput)
}

type ListRegistryRepositoriesOutputArgs struct {
	// The name of a container registry.
	RegistryName pulumi.StringInput `pulumi:"registryName"`
}

func (ListRegistryRepositoriesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRegistryRepositoriesArgs)(nil)).Elem()
}

type ListRegistryRepositoriesResultOutput struct{ *pulumi.OutputState }

func (ListRegistryRepositoriesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRegistryRepositoriesResult)(nil)).Elem()
}

func (o ListRegistryRepositoriesResultOutput) ToListRegistryRepositoriesResultOutput() ListRegistryRepositoriesResultOutput {
	return o
}

func (o ListRegistryRepositoriesResultOutput) ToListRegistryRepositoriesResultOutputWithContext(ctx context.Context) ListRegistryRepositoriesResultOutput {
	return o
}

func (o ListRegistryRepositoriesResultOutput) Items() ListRegistryRepositoriesOutput {
	return o.ApplyT(func(v ListRegistryRepositoriesResult) ListRegistryRepositories { return v.Items }).(ListRegistryRepositoriesOutput)
}

func init() {
	pulumi.RegisterOutputType(ListRegistryRepositoriesResultOutput{})
}
