// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListRegistryRepositoriesV2(ctx *pulumi.Context, args *ListRegistryRepositoriesV2Args, opts ...pulumi.InvokeOption) (*ListRegistryRepositoriesV2Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListRegistryRepositoriesV2Result
	err := ctx.Invoke("digitalocean-native:registry/v2:listRegistryRepositoriesV2", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRegistryRepositoriesV2Args struct {
	// The name of a container registry.
	RegistryName string `pulumi:"registryName"`
}

type ListRegistryRepositoriesV2Result struct {
	Links        *PageLinks     `pulumi:"links"`
	Meta         MetaMeta       `pulumi:"meta"`
	Repositories []RepositoryV2 `pulumi:"repositories"`
}

func ListRegistryRepositoriesV2Output(ctx *pulumi.Context, args ListRegistryRepositoriesV2OutputArgs, opts ...pulumi.InvokeOption) ListRegistryRepositoriesV2ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListRegistryRepositoriesV2Result, error) {
			args := v.(ListRegistryRepositoriesV2Args)
			r, err := ListRegistryRepositoriesV2(ctx, &args, opts...)
			var s ListRegistryRepositoriesV2Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListRegistryRepositoriesV2ResultOutput)
}

type ListRegistryRepositoriesV2OutputArgs struct {
	// The name of a container registry.
	RegistryName pulumi.StringInput `pulumi:"registryName"`
}

func (ListRegistryRepositoriesV2OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRegistryRepositoriesV2Args)(nil)).Elem()
}

type ListRegistryRepositoriesV2ResultOutput struct{ *pulumi.OutputState }

func (ListRegistryRepositoriesV2ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRegistryRepositoriesV2Result)(nil)).Elem()
}

func (o ListRegistryRepositoriesV2ResultOutput) ToListRegistryRepositoriesV2ResultOutput() ListRegistryRepositoriesV2ResultOutput {
	return o
}

func (o ListRegistryRepositoriesV2ResultOutput) ToListRegistryRepositoriesV2ResultOutputWithContext(ctx context.Context) ListRegistryRepositoriesV2ResultOutput {
	return o
}

func (o ListRegistryRepositoriesV2ResultOutput) Links() PageLinksPtrOutput {
	return o.ApplyT(func(v ListRegistryRepositoriesV2Result) *PageLinks { return v.Links }).(PageLinksPtrOutput)
}

func (o ListRegistryRepositoriesV2ResultOutput) Meta() MetaMetaOutput {
	return o.ApplyT(func(v ListRegistryRepositoriesV2Result) MetaMeta { return v.Meta }).(MetaMetaOutput)
}

func (o ListRegistryRepositoriesV2ResultOutput) Repositories() RepositoryV2ArrayOutput {
	return o.ApplyT(func(v ListRegistryRepositoriesV2Result) []RepositoryV2 { return v.Repositories }).(RepositoryV2ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListRegistryRepositoriesV2ResultOutput{})
}
