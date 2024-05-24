// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetKubernetesCredential(ctx *pulumi.Context, args *GetKubernetesCredentialArgs, opts ...pulumi.InvokeOption) (*GetKubernetesCredentialResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetKubernetesCredentialResult
	err := ctx.Invoke("digitalocean-native:kubernetes/v2:getKubernetesCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetKubernetesCredentialArgs struct {
	// A unique ID that can be used to reference a Kubernetes cluster.
	ClusterId string `pulumi:"clusterId"`
}

type GetKubernetesCredentialResult struct {
	Items Credentials `pulumi:"items"`
}

func GetKubernetesCredentialOutput(ctx *pulumi.Context, args GetKubernetesCredentialOutputArgs, opts ...pulumi.InvokeOption) GetKubernetesCredentialResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetKubernetesCredentialResult, error) {
			args := v.(GetKubernetesCredentialArgs)
			r, err := GetKubernetesCredential(ctx, &args, opts...)
			var s GetKubernetesCredentialResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetKubernetesCredentialResultOutput)
}

type GetKubernetesCredentialOutputArgs struct {
	// A unique ID that can be used to reference a Kubernetes cluster.
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
}

func (GetKubernetesCredentialOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKubernetesCredentialArgs)(nil)).Elem()
}

type GetKubernetesCredentialResultOutput struct{ *pulumi.OutputState }

func (GetKubernetesCredentialResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKubernetesCredentialResult)(nil)).Elem()
}

func (o GetKubernetesCredentialResultOutput) ToGetKubernetesCredentialResultOutput() GetKubernetesCredentialResultOutput {
	return o
}

func (o GetKubernetesCredentialResultOutput) ToGetKubernetesCredentialResultOutputWithContext(ctx context.Context) GetKubernetesCredentialResultOutput {
	return o
}

func (o GetKubernetesCredentialResultOutput) Items() CredentialsOutput {
	return o.ApplyT(func(v GetKubernetesCredentialResult) Credentials { return v.Items }).(CredentialsOutput)
}

func init() {
	pulumi.RegisterOutputType(GetKubernetesCredentialResultOutput{})
}
