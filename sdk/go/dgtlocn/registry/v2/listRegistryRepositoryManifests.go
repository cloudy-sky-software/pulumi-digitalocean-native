// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListRegistryRepositoryManifests(ctx *pulumi.Context, args *ListRegistryRepositoryManifestsArgs, opts ...pulumi.InvokeOption) (*ListRegistryRepositoryManifestsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListRegistryRepositoryManifestsResult
	err := ctx.Invoke("digitalocean-native:registry/v2:listRegistryRepositoryManifests", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRegistryRepositoryManifestsArgs struct {
	// The name of a container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of a container registry repository. If the name contains `/` characters, they must be URL-encoded, e.g. `%2F`.
	RepositoryName string `pulumi:"repositoryName"`
}

type ListRegistryRepositoryManifestsResult struct {
	Links     *PageLinks           `pulumi:"links"`
	Manifests []RepositoryManifest `pulumi:"manifests"`
	Meta      MetaMeta             `pulumi:"meta"`
}

func ListRegistryRepositoryManifestsOutput(ctx *pulumi.Context, args ListRegistryRepositoryManifestsOutputArgs, opts ...pulumi.InvokeOption) ListRegistryRepositoryManifestsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListRegistryRepositoryManifestsResult, error) {
			args := v.(ListRegistryRepositoryManifestsArgs)
			r, err := ListRegistryRepositoryManifests(ctx, &args, opts...)
			var s ListRegistryRepositoryManifestsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListRegistryRepositoryManifestsResultOutput)
}

type ListRegistryRepositoryManifestsOutputArgs struct {
	// The name of a container registry.
	RegistryName pulumi.StringInput `pulumi:"registryName"`
	// The name of a container registry repository. If the name contains `/` characters, they must be URL-encoded, e.g. `%2F`.
	RepositoryName pulumi.StringInput `pulumi:"repositoryName"`
}

func (ListRegistryRepositoryManifestsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRegistryRepositoryManifestsArgs)(nil)).Elem()
}

type ListRegistryRepositoryManifestsResultOutput struct{ *pulumi.OutputState }

func (ListRegistryRepositoryManifestsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRegistryRepositoryManifestsResult)(nil)).Elem()
}

func (o ListRegistryRepositoryManifestsResultOutput) ToListRegistryRepositoryManifestsResultOutput() ListRegistryRepositoryManifestsResultOutput {
	return o
}

func (o ListRegistryRepositoryManifestsResultOutput) ToListRegistryRepositoryManifestsResultOutputWithContext(ctx context.Context) ListRegistryRepositoryManifestsResultOutput {
	return o
}

func (o ListRegistryRepositoryManifestsResultOutput) Links() PageLinksPtrOutput {
	return o.ApplyT(func(v ListRegistryRepositoryManifestsResult) *PageLinks { return v.Links }).(PageLinksPtrOutput)
}

func (o ListRegistryRepositoryManifestsResultOutput) Manifests() RepositoryManifestArrayOutput {
	return o.ApplyT(func(v ListRegistryRepositoryManifestsResult) []RepositoryManifest { return v.Manifests }).(RepositoryManifestArrayOutput)
}

func (o ListRegistryRepositoryManifestsResultOutput) Meta() MetaMetaOutput {
	return o.ApplyT(func(v ListRegistryRepositoryManifestsResult) MetaMeta { return v.Meta }).(MetaMetaOutput)
}

func init() {
	pulumi.RegisterOutputType(ListRegistryRepositoryManifestsResultOutput{})
}
