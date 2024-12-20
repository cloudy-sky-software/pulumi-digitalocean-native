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
	// A base64 encoding of bytes representing the certificate authority data for accessing the cluster.
	CertificateAuthorityData *string `pulumi:"certificateAuthorityData"`
	// A base64 encoding of bytes representing the x509 client
	// certificate data for access the cluster. This is only returned for clusters
	// without support for token-based authentication.
	//
	// Newly created Kubernetes clusters do not return credentials using
	// certificate-based authentication. For additional information,
	// [see here](https://www.digitalocean.com/docs/kubernetes/how-to/connect-to-cluster/#authenticate).
	ClientCertificateData *string `pulumi:"clientCertificateData"`
	// A base64 encoding of bytes representing the x509 client key
	// data for access the cluster. This is only returned for clusters without
	// support for token-based authentication.
	//
	// Newly created Kubernetes clusters do not return credentials using
	// certificate-based authentication. For additional information,
	// [see here](https://www.digitalocean.com/docs/kubernetes/how-to/connect-to-cluster/#authenticate).
	ClientKeyData *string `pulumi:"clientKeyData"`
	// A time value given in ISO8601 combined date and time format that represents when the access token expires.
	ExpiresAt *string `pulumi:"expiresAt"`
	// The URL used to access the cluster API server.
	Server *string `pulumi:"server"`
	// An access token used to authenticate with the cluster. This is only returned for clusters with support for token-based authentication.
	Token *string `pulumi:"token"`
}

func GetKubernetesCredentialOutput(ctx *pulumi.Context, args GetKubernetesCredentialOutputArgs, opts ...pulumi.InvokeOption) GetKubernetesCredentialResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetKubernetesCredentialResultOutput, error) {
			args := v.(GetKubernetesCredentialArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("digitalocean-native:kubernetes/v2:getKubernetesCredential", args, GetKubernetesCredentialResultOutput{}, options).(GetKubernetesCredentialResultOutput), nil
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

// A base64 encoding of bytes representing the certificate authority data for accessing the cluster.
func (o GetKubernetesCredentialResultOutput) CertificateAuthorityData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKubernetesCredentialResult) *string { return v.CertificateAuthorityData }).(pulumi.StringPtrOutput)
}

// A base64 encoding of bytes representing the x509 client
// certificate data for access the cluster. This is only returned for clusters
// without support for token-based authentication.
//
// Newly created Kubernetes clusters do not return credentials using
// certificate-based authentication. For additional information,
// [see here](https://www.digitalocean.com/docs/kubernetes/how-to/connect-to-cluster/#authenticate).
func (o GetKubernetesCredentialResultOutput) ClientCertificateData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKubernetesCredentialResult) *string { return v.ClientCertificateData }).(pulumi.StringPtrOutput)
}

// A base64 encoding of bytes representing the x509 client key
// data for access the cluster. This is only returned for clusters without
// support for token-based authentication.
//
// Newly created Kubernetes clusters do not return credentials using
// certificate-based authentication. For additional information,
// [see here](https://www.digitalocean.com/docs/kubernetes/how-to/connect-to-cluster/#authenticate).
func (o GetKubernetesCredentialResultOutput) ClientKeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKubernetesCredentialResult) *string { return v.ClientKeyData }).(pulumi.StringPtrOutput)
}

// A time value given in ISO8601 combined date and time format that represents when the access token expires.
func (o GetKubernetesCredentialResultOutput) ExpiresAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKubernetesCredentialResult) *string { return v.ExpiresAt }).(pulumi.StringPtrOutput)
}

// The URL used to access the cluster API server.
func (o GetKubernetesCredentialResultOutput) Server() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKubernetesCredentialResult) *string { return v.Server }).(pulumi.StringPtrOutput)
}

// An access token used to authenticate with the cluster. This is only returned for clusters with support for token-based authentication.
func (o GetKubernetesCredentialResultOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKubernetesCredentialResult) *string { return v.Token }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetKubernetesCredentialResultOutput{})
}
