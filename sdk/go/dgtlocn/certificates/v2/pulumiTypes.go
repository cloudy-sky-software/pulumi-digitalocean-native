// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type Certificate struct {
	// A unique identifier generated from the SHA-1 fingerprint of the certificate.
	_sha1Fingerprint *string `pulumi:"_sha1Fingerprint"`
	// A time value given in ISO8601 combined date and time format that represents when the certificate was created.
	CreatedAt *string `pulumi:"createdAt"`
	// An array of fully qualified domain names (FQDNs) for which the certificate was issued.
	DnsNames []string `pulumi:"dnsNames"`
	// A unique ID that can be used to identify and reference a certificate.
	Id *string `pulumi:"id"`
	// A unique human-readable name referring to a certificate.
	Name *string `pulumi:"name"`
	// A time value given in ISO8601 combined date and time format that represents the certificate's expiration date.
	NotAfter *string `pulumi:"notAfter"`
	// A string representing the current state of the certificate. It may be `pending`, `verified`, or `error`.
	State *CertificateState `pulumi:"state"`
	// A string representing the type of the certificate. The value will be `custom` for a user-uploaded certificate or `lets_encrypt` for one automatically generated with Let's Encrypt.
	Type *CertificateType `pulumi:"type"`
}

type CertificateOutput struct{ *pulumi.OutputState }

func (CertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Certificate)(nil)).Elem()
}

func (o CertificateOutput) ToCertificateOutput() CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return o
}

// A unique identifier generated from the SHA-1 fingerprint of the certificate.
func (o CertificateOutput) _sha1Fingerprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Certificate) *string { return v._sha1Fingerprint }).(pulumi.StringPtrOutput)
}

// A time value given in ISO8601 combined date and time format that represents when the certificate was created.
func (o CertificateOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Certificate) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// An array of fully qualified domain names (FQDNs) for which the certificate was issued.
func (o CertificateOutput) DnsNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Certificate) []string { return v.DnsNames }).(pulumi.StringArrayOutput)
}

// A unique ID that can be used to identify and reference a certificate.
func (o CertificateOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Certificate) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// A unique human-readable name referring to a certificate.
func (o CertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Certificate) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A time value given in ISO8601 combined date and time format that represents the certificate's expiration date.
func (o CertificateOutput) NotAfter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Certificate) *string { return v.NotAfter }).(pulumi.StringPtrOutput)
}

// A string representing the current state of the certificate. It may be `pending`, `verified`, or `error`.
func (o CertificateOutput) State() CertificateStatePtrOutput {
	return o.ApplyT(func(v Certificate) *CertificateState { return v.State }).(CertificateStatePtrOutput)
}

// A string representing the type of the certificate. The value will be `custom` for a user-uploaded certificate or `lets_encrypt` for one automatically generated with Let's Encrypt.
func (o CertificateOutput) Type() CertificateTypePtrOutput {
	return o.ApplyT(func(v Certificate) *CertificateType { return v.Type }).(CertificateTypePtrOutput)
}

type CertificatePtrOutput struct{ *pulumi.OutputState }

func (CertificatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (o CertificatePtrOutput) ToCertificatePtrOutput() CertificatePtrOutput {
	return o
}

func (o CertificatePtrOutput) ToCertificatePtrOutputWithContext(ctx context.Context) CertificatePtrOutput {
	return o
}

func (o CertificatePtrOutput) Elem() CertificateOutput {
	return o.ApplyT(func(v *Certificate) Certificate {
		if v != nil {
			return *v
		}
		var ret Certificate
		return ret
	}).(CertificateOutput)
}

// A unique identifier generated from the SHA-1 fingerprint of the certificate.
func (o CertificatePtrOutput) _sha1Fingerprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) *string {
		if v == nil {
			return nil
		}
		return v._sha1Fingerprint
	}).(pulumi.StringPtrOutput)
}

// A time value given in ISO8601 combined date and time format that represents when the certificate was created.
func (o CertificatePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

// An array of fully qualified domain names (FQDNs) for which the certificate was issued.
func (o CertificatePtrOutput) DnsNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Certificate) []string {
		if v == nil {
			return nil
		}
		return v.DnsNames
	}).(pulumi.StringArrayOutput)
}

// A unique ID that can be used to identify and reference a certificate.
func (o CertificatePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

// A unique human-readable name referring to a certificate.
func (o CertificatePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

// A time value given in ISO8601 combined date and time format that represents the certificate's expiration date.
func (o CertificatePtrOutput) NotAfter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) *string {
		if v == nil {
			return nil
		}
		return v.NotAfter
	}).(pulumi.StringPtrOutput)
}

// A string representing the current state of the certificate. It may be `pending`, `verified`, or `error`.
func (o CertificatePtrOutput) State() CertificateStatePtrOutput {
	return o.ApplyT(func(v *Certificate) *CertificateState {
		if v == nil {
			return nil
		}
		return v.State
	}).(CertificateStatePtrOutput)
}

// A string representing the type of the certificate. The value will be `custom` for a user-uploaded certificate or `lets_encrypt` for one automatically generated with Let's Encrypt.
func (o CertificatePtrOutput) Type() CertificateTypePtrOutput {
	return o.ApplyT(func(v *Certificate) *CertificateType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(CertificateTypePtrOutput)
}

type CertificateArrayOutput struct{ *pulumi.OutputState }

func (CertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Certificate)(nil)).Elem()
}

func (o CertificateArrayOutput) ToCertificateArrayOutput() CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) ToCertificateArrayOutputWithContext(ctx context.Context) CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) Index(i pulumi.IntInput) CertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Certificate {
		return vs[0].([]Certificate)[vs[1].(int)]
	}).(CertificateOutput)
}

type GetCertificateProperties struct {
	Certificate *Certificate `pulumi:"certificate"`
}

type GetCertificatePropertiesOutput struct{ *pulumi.OutputState }

func (GetCertificatePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCertificateProperties)(nil)).Elem()
}

func (o GetCertificatePropertiesOutput) ToGetCertificatePropertiesOutput() GetCertificatePropertiesOutput {
	return o
}

func (o GetCertificatePropertiesOutput) ToGetCertificatePropertiesOutputWithContext(ctx context.Context) GetCertificatePropertiesOutput {
	return o
}

func (o GetCertificatePropertiesOutput) Certificate() CertificatePtrOutput {
	return o.ApplyT(func(v GetCertificateProperties) *Certificate { return v.Certificate }).(CertificatePtrOutput)
}

type ListCertificatesItems struct {
	Certificates []Certificate `pulumi:"certificates"`
	Links        *PageLinks    `pulumi:"links"`
	Meta         MetaMeta      `pulumi:"meta"`
}

type ListCertificatesItemsOutput struct{ *pulumi.OutputState }

func (ListCertificatesItemsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListCertificatesItems)(nil)).Elem()
}

func (o ListCertificatesItemsOutput) ToListCertificatesItemsOutput() ListCertificatesItemsOutput {
	return o
}

func (o ListCertificatesItemsOutput) ToListCertificatesItemsOutputWithContext(ctx context.Context) ListCertificatesItemsOutput {
	return o
}

func (o ListCertificatesItemsOutput) Certificates() CertificateArrayOutput {
	return o.ApplyT(func(v ListCertificatesItems) []Certificate { return v.Certificates }).(CertificateArrayOutput)
}

func (o ListCertificatesItemsOutput) Links() PageLinksPtrOutput {
	return o.ApplyT(func(v ListCertificatesItems) *PageLinks { return v.Links }).(PageLinksPtrOutput)
}

func (o ListCertificatesItemsOutput) Meta() MetaMetaOutput {
	return o.ApplyT(func(v ListCertificatesItems) MetaMeta { return v.Meta }).(MetaMetaOutput)
}

type MetaMeta struct {
	// Number of objects returned by the request.
	Total *int `pulumi:"total"`
}

type MetaMetaOutput struct{ *pulumi.OutputState }

func (MetaMetaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetaMeta)(nil)).Elem()
}

func (o MetaMetaOutput) ToMetaMetaOutput() MetaMetaOutput {
	return o
}

func (o MetaMetaOutput) ToMetaMetaOutputWithContext(ctx context.Context) MetaMetaOutput {
	return o
}

// Number of objects returned by the request.
func (o MetaMetaOutput) Total() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MetaMeta) *int { return v.Total }).(pulumi.IntPtrOutput)
}

type PageLinks struct {
	Pages *PageLinksPagesProperties `pulumi:"pages"`
}

type PageLinksOutput struct{ *pulumi.OutputState }

func (PageLinksOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PageLinks)(nil)).Elem()
}

func (o PageLinksOutput) ToPageLinksOutput() PageLinksOutput {
	return o
}

func (o PageLinksOutput) ToPageLinksOutputWithContext(ctx context.Context) PageLinksOutput {
	return o
}

func (o PageLinksOutput) Pages() PageLinksPagesPropertiesPtrOutput {
	return o.ApplyT(func(v PageLinks) *PageLinksPagesProperties { return v.Pages }).(PageLinksPagesPropertiesPtrOutput)
}

type PageLinksPtrOutput struct{ *pulumi.OutputState }

func (PageLinksPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PageLinks)(nil)).Elem()
}

func (o PageLinksPtrOutput) ToPageLinksPtrOutput() PageLinksPtrOutput {
	return o
}

func (o PageLinksPtrOutput) ToPageLinksPtrOutputWithContext(ctx context.Context) PageLinksPtrOutput {
	return o
}

func (o PageLinksPtrOutput) Elem() PageLinksOutput {
	return o.ApplyT(func(v *PageLinks) PageLinks {
		if v != nil {
			return *v
		}
		var ret PageLinks
		return ret
	}).(PageLinksOutput)
}

func (o PageLinksPtrOutput) Pages() PageLinksPagesPropertiesPtrOutput {
	return o.ApplyT(func(v *PageLinks) *PageLinksPagesProperties {
		if v == nil {
			return nil
		}
		return v.Pages
	}).(PageLinksPagesPropertiesPtrOutput)
}

type PageLinksPagesProperties struct {
	First *string `pulumi:"first"`
	Last  *string `pulumi:"last"`
	Next  *string `pulumi:"next"`
	Prev  *string `pulumi:"prev"`
}

type PageLinksPagesPropertiesOutput struct{ *pulumi.OutputState }

func (PageLinksPagesPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PageLinksPagesProperties)(nil)).Elem()
}

func (o PageLinksPagesPropertiesOutput) ToPageLinksPagesPropertiesOutput() PageLinksPagesPropertiesOutput {
	return o
}

func (o PageLinksPagesPropertiesOutput) ToPageLinksPagesPropertiesOutputWithContext(ctx context.Context) PageLinksPagesPropertiesOutput {
	return o
}

func (o PageLinksPagesPropertiesOutput) First() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PageLinksPagesProperties) *string { return v.First }).(pulumi.StringPtrOutput)
}

func (o PageLinksPagesPropertiesOutput) Last() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PageLinksPagesProperties) *string { return v.Last }).(pulumi.StringPtrOutput)
}

func (o PageLinksPagesPropertiesOutput) Next() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PageLinksPagesProperties) *string { return v.Next }).(pulumi.StringPtrOutput)
}

func (o PageLinksPagesPropertiesOutput) Prev() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PageLinksPagesProperties) *string { return v.Prev }).(pulumi.StringPtrOutput)
}

type PageLinksPagesPropertiesPtrOutput struct{ *pulumi.OutputState }

func (PageLinksPagesPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PageLinksPagesProperties)(nil)).Elem()
}

func (o PageLinksPagesPropertiesPtrOutput) ToPageLinksPagesPropertiesPtrOutput() PageLinksPagesPropertiesPtrOutput {
	return o
}

func (o PageLinksPagesPropertiesPtrOutput) ToPageLinksPagesPropertiesPtrOutputWithContext(ctx context.Context) PageLinksPagesPropertiesPtrOutput {
	return o
}

func (o PageLinksPagesPropertiesPtrOutput) Elem() PageLinksPagesPropertiesOutput {
	return o.ApplyT(func(v *PageLinksPagesProperties) PageLinksPagesProperties {
		if v != nil {
			return *v
		}
		var ret PageLinksPagesProperties
		return ret
	}).(PageLinksPagesPropertiesOutput)
}

func (o PageLinksPagesPropertiesPtrOutput) First() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PageLinksPagesProperties) *string {
		if v == nil {
			return nil
		}
		return v.First
	}).(pulumi.StringPtrOutput)
}

func (o PageLinksPagesPropertiesPtrOutput) Last() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PageLinksPagesProperties) *string {
		if v == nil {
			return nil
		}
		return v.Last
	}).(pulumi.StringPtrOutput)
}

func (o PageLinksPagesPropertiesPtrOutput) Next() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PageLinksPagesProperties) *string {
		if v == nil {
			return nil
		}
		return v.Next
	}).(pulumi.StringPtrOutput)
}

func (o PageLinksPagesPropertiesPtrOutput) Prev() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PageLinksPagesProperties) *string {
		if v == nil {
			return nil
		}
		return v.Prev
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CertificateOutput{})
	pulumi.RegisterOutputType(CertificatePtrOutput{})
	pulumi.RegisterOutputType(CertificateArrayOutput{})
	pulumi.RegisterOutputType(GetCertificatePropertiesOutput{})
	pulumi.RegisterOutputType(ListCertificatesItemsOutput{})
	pulumi.RegisterOutputType(MetaMetaOutput{})
	pulumi.RegisterOutputType(PageLinksOutput{})
	pulumi.RegisterOutputType(PageLinksPtrOutput{})
	pulumi.RegisterOutputType(PageLinksPagesPropertiesOutput{})
	pulumi.RegisterOutputType(PageLinksPagesPropertiesPtrOutput{})
}
