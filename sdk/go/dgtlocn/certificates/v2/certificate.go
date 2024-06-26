// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"errors"
	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Certificate struct {
	pulumi.CustomResourceState

	Certificate CertificateTypePtrOutput `pulumi:"certificate"`
	// The full PEM-formatted trust chain between the certificate authority's certificate and your domain's SSL certificate.
	CertificateChain pulumi.StringPtrOutput `pulumi:"certificateChain"`
	// An array of fully qualified domain names (FQDNs) for which the certificate was issued. A certificate covering all subdomains can be issued using a wildcard (e.g. `*.example.com`).
	DnsNames pulumi.StringArrayOutput `pulumi:"dnsNames"`
	// The contents of a PEM-formatted public SSL certificate.
	LeafCertificate pulumi.StringPtrOutput `pulumi:"leafCertificate"`
	// A unique human-readable name referring to a certificate.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The contents of a PEM-formatted private-key corresponding to the SSL certificate.
	PrivateKey pulumi.StringPtrOutput `pulumi:"privateKey"`
	// A string representing the type of the certificate. The value will be `custom` for a user-uploaded certificate or `lets_encrypt` for one automatically generated with Let's Encrypt.
	Type CertificateCreateBaseTypePtrOutput `pulumi:"type"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DnsNames == nil {
		return nil, errors.New("invalid value for required argument 'DnsNames'")
	}
	if args.LeafCertificate == nil {
		return nil, errors.New("invalid value for required argument 'LeafCertificate'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.PrivateKey == nil {
		return nil, errors.New("invalid value for required argument 'PrivateKey'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Certificate
	err := ctx.RegisterResource("digitalocean-native:certificates/v2:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificate gets an existing Certificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("digitalocean-native:certificates/v2:Certificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Certificate resources.
type certificateState struct {
}

type CertificateState struct {
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	// The full PEM-formatted trust chain between the certificate authority's certificate and your domain's SSL certificate.
	CertificateChain *string `pulumi:"certificateChain"`
	// An array of fully qualified domain names (FQDNs) for which the certificate was issued. A certificate covering all subdomains can be issued using a wildcard (e.g. `*.example.com`).
	DnsNames []string `pulumi:"dnsNames"`
	// The contents of a PEM-formatted public SSL certificate.
	LeafCertificate string `pulumi:"leafCertificate"`
	// A unique human-readable name referring to a certificate.
	Name string `pulumi:"name"`
	// The contents of a PEM-formatted private-key corresponding to the SSL certificate.
	PrivateKey string `pulumi:"privateKey"`
	// A string representing the type of the certificate. The value will be `custom` for a user-uploaded certificate or `lets_encrypt` for one automatically generated with Let's Encrypt.
	Type *CertificateCreateBaseType `pulumi:"type"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// The full PEM-formatted trust chain between the certificate authority's certificate and your domain's SSL certificate.
	CertificateChain pulumi.StringPtrInput
	// An array of fully qualified domain names (FQDNs) for which the certificate was issued. A certificate covering all subdomains can be issued using a wildcard (e.g. `*.example.com`).
	DnsNames pulumi.StringArrayInput
	// The contents of a PEM-formatted public SSL certificate.
	LeafCertificate pulumi.StringInput
	// A unique human-readable name referring to a certificate.
	Name pulumi.StringInput
	// The contents of a PEM-formatted private-key corresponding to the SSL certificate.
	PrivateKey pulumi.StringInput
	// A string representing the type of the certificate. The value will be `custom` for a user-uploaded certificate or `lets_encrypt` for one automatically generated with Let's Encrypt.
	Type CertificateCreateBaseTypePtrInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}

type CertificateInput interface {
	pulumi.Input

	ToCertificateOutput() CertificateOutput
	ToCertificateOutputWithContext(ctx context.Context) CertificateOutput
}

func (*Certificate) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (i *Certificate) ToCertificateOutput() CertificateOutput {
	return i.ToCertificateOutputWithContext(context.Background())
}

func (i *Certificate) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOutput)
}

type CertificateOutput struct{ *pulumi.OutputState }

func (CertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (o CertificateOutput) ToCertificateOutput() CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return o
}

func (o CertificateOutput) Certificate() CertificateTypePtrOutput {
	return o.ApplyT(func(v *Certificate) CertificateTypePtrOutput { return v.Certificate }).(CertificateTypePtrOutput)
}

// The full PEM-formatted trust chain between the certificate authority's certificate and your domain's SSL certificate.
func (o CertificateOutput) CertificateChain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.CertificateChain }).(pulumi.StringPtrOutput)
}

// An array of fully qualified domain names (FQDNs) for which the certificate was issued. A certificate covering all subdomains can be issued using a wildcard (e.g. `*.example.com`).
func (o CertificateOutput) DnsNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringArrayOutput { return v.DnsNames }).(pulumi.StringArrayOutput)
}

// The contents of a PEM-formatted public SSL certificate.
func (o CertificateOutput) LeafCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.LeafCertificate }).(pulumi.StringPtrOutput)
}

// A unique human-readable name referring to a certificate.
func (o CertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The contents of a PEM-formatted private-key corresponding to the SSL certificate.
func (o CertificateOutput) PrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.PrivateKey }).(pulumi.StringPtrOutput)
}

// A string representing the type of the certificate. The value will be `custom` for a user-uploaded certificate or `lets_encrypt` for one automatically generated with Let's Encrypt.
func (o CertificateOutput) Type() CertificateCreateBaseTypePtrOutput {
	return o.ApplyT(func(v *Certificate) CertificateCreateBaseTypePtrOutput { return v.Type }).(CertificateCreateBaseTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateInput)(nil)).Elem(), &Certificate{})
	pulumi.RegisterOutputType(CertificateOutput{})
}
