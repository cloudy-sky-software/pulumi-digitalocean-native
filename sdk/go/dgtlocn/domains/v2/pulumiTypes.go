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

type Domain struct {
	// This optional attribute may contain an IP address. When provided, an A record will be automatically created pointing to the apex domain.
	IpAddress *string `pulumi:"ipAddress"`
	// The name of the domain itself. This should follow the standard domain format of domain.TLD. For instance, `example.com` is a valid domain name.
	Name *string `pulumi:"name"`
	// This value is the time to live for the records on this domain, in seconds. This defines the time frame that clients can cache queried information before a refresh should be requested.
	Ttl *int `pulumi:"ttl"`
	// This attribute contains the complete contents of the zone file for the selected domain. Individual domain record resources should be used to get more granular control over records. However, this attribute can also be used to get information about the SOA record, which is created automatically and is not accessible as an individual record resource.
	ZoneFile *string `pulumi:"zoneFile"`
}

type DomainOutput struct{ *pulumi.OutputState }

func (DomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Domain)(nil)).Elem()
}

func (o DomainOutput) ToDomainOutput() DomainOutput {
	return o
}

func (o DomainOutput) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return o
}

// This optional attribute may contain an IP address. When provided, an A record will be automatically created pointing to the apex domain.
func (o DomainOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Domain) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

// The name of the domain itself. This should follow the standard domain format of domain.TLD. For instance, `example.com` is a valid domain name.
func (o DomainOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Domain) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// This value is the time to live for the records on this domain, in seconds. This defines the time frame that clients can cache queried information before a refresh should be requested.
func (o DomainOutput) Ttl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Domain) *int { return v.Ttl }).(pulumi.IntPtrOutput)
}

// This attribute contains the complete contents of the zone file for the selected domain. Individual domain record resources should be used to get more granular control over records. However, this attribute can also be used to get information about the SOA record, which is created automatically and is not accessible as an individual record resource.
func (o DomainOutput) ZoneFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Domain) *string { return v.ZoneFile }).(pulumi.StringPtrOutput)
}

type DomainPtrOutput struct{ *pulumi.OutputState }

func (DomainPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (o DomainPtrOutput) ToDomainPtrOutput() DomainPtrOutput {
	return o
}

func (o DomainPtrOutput) ToDomainPtrOutputWithContext(ctx context.Context) DomainPtrOutput {
	return o
}

func (o DomainPtrOutput) Elem() DomainOutput {
	return o.ApplyT(func(v *Domain) Domain {
		if v != nil {
			return *v
		}
		var ret Domain
		return ret
	}).(DomainOutput)
}

// This optional attribute may contain an IP address. When provided, an A record will be automatically created pointing to the apex domain.
func (o DomainPtrOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) *string {
		if v == nil {
			return nil
		}
		return v.IpAddress
	}).(pulumi.StringPtrOutput)
}

// The name of the domain itself. This should follow the standard domain format of domain.TLD. For instance, `example.com` is a valid domain name.
func (o DomainPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

// This value is the time to live for the records on this domain, in seconds. This defines the time frame that clients can cache queried information before a refresh should be requested.
func (o DomainPtrOutput) Ttl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Domain) *int {
		if v == nil {
			return nil
		}
		return v.Ttl
	}).(pulumi.IntPtrOutput)
}

// This attribute contains the complete contents of the zone file for the selected domain. Individual domain record resources should be used to get more granular control over records. However, this attribute can also be used to get information about the SOA record, which is created automatically and is not accessible as an individual record resource.
func (o DomainPtrOutput) ZoneFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) *string {
		if v == nil {
			return nil
		}
		return v.ZoneFile
	}).(pulumi.StringPtrOutput)
}

type DomainArrayOutput struct{ *pulumi.OutputState }

func (DomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Domain)(nil)).Elem()
}

func (o DomainArrayOutput) ToDomainArrayOutput() DomainArrayOutput {
	return o
}

func (o DomainArrayOutput) ToDomainArrayOutputWithContext(ctx context.Context) DomainArrayOutput {
	return o
}

func (o DomainArrayOutput) Index(i pulumi.IntInput) DomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Domain {
		return vs[0].([]Domain)[vs[1].(int)]
	}).(DomainOutput)
}

type DomainRecord struct {
	// Variable data depending on record type. For example, the "data" value for an A record would be the IPv4 address to which the domain will be mapped. For a CAA record, it would contain the domain name of the CA being granted permission to issue certificates.
	Data *string `pulumi:"data"`
	// An unsigned integer between 0-255 used for CAA records.
	Flags *int `pulumi:"flags"`
	// A unique identifier for each domain record.
	Id *int `pulumi:"id"`
	// The host name, alias, or service being defined by the record.
	Name *string `pulumi:"name"`
	// The port for SRV records.
	Port *int `pulumi:"port"`
	// The priority for SRV and MX records.
	Priority *int `pulumi:"priority"`
	// The parameter tag for CAA records. Valid values are "issue", "issuewild", or "iodef"
	Tag *string `pulumi:"tag"`
	// This value is the time to live for the record, in seconds. This defines the time frame that clients can cache queried information before a refresh should be requested.
	Ttl *int `pulumi:"ttl"`
	// The type of the DNS record. For example: A, CNAME, TXT, ...
	Type string `pulumi:"type"`
	// The weight for SRV records.
	Weight *int `pulumi:"weight"`
}

type DomainRecordOutput struct{ *pulumi.OutputState }

func (DomainRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainRecord)(nil)).Elem()
}

func (o DomainRecordOutput) ToDomainRecordOutput() DomainRecordOutput {
	return o
}

func (o DomainRecordOutput) ToDomainRecordOutputWithContext(ctx context.Context) DomainRecordOutput {
	return o
}

// Variable data depending on record type. For example, the "data" value for an A record would be the IPv4 address to which the domain will be mapped. For a CAA record, it would contain the domain name of the CA being granted permission to issue certificates.
func (o DomainRecordOutput) Data() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainRecord) *string { return v.Data }).(pulumi.StringPtrOutput)
}

// An unsigned integer between 0-255 used for CAA records.
func (o DomainRecordOutput) Flags() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DomainRecord) *int { return v.Flags }).(pulumi.IntPtrOutput)
}

// A unique identifier for each domain record.
func (o DomainRecordOutput) Id() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DomainRecord) *int { return v.Id }).(pulumi.IntPtrOutput)
}

// The host name, alias, or service being defined by the record.
func (o DomainRecordOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainRecord) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The port for SRV records.
func (o DomainRecordOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DomainRecord) *int { return v.Port }).(pulumi.IntPtrOutput)
}

// The priority for SRV and MX records.
func (o DomainRecordOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DomainRecord) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

// The parameter tag for CAA records. Valid values are "issue", "issuewild", or "iodef"
func (o DomainRecordOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainRecord) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

// This value is the time to live for the record, in seconds. This defines the time frame that clients can cache queried information before a refresh should be requested.
func (o DomainRecordOutput) Ttl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DomainRecord) *int { return v.Ttl }).(pulumi.IntPtrOutput)
}

// The type of the DNS record. For example: A, CNAME, TXT, ...
func (o DomainRecordOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v DomainRecord) string { return v.Type }).(pulumi.StringOutput)
}

// The weight for SRV records.
func (o DomainRecordOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DomainRecord) *int { return v.Weight }).(pulumi.IntPtrOutput)
}

type DomainRecordPtrOutput struct{ *pulumi.OutputState }

func (DomainRecordPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainRecord)(nil)).Elem()
}

func (o DomainRecordPtrOutput) ToDomainRecordPtrOutput() DomainRecordPtrOutput {
	return o
}

func (o DomainRecordPtrOutput) ToDomainRecordPtrOutputWithContext(ctx context.Context) DomainRecordPtrOutput {
	return o
}

func (o DomainRecordPtrOutput) Elem() DomainRecordOutput {
	return o.ApplyT(func(v *DomainRecord) DomainRecord {
		if v != nil {
			return *v
		}
		var ret DomainRecord
		return ret
	}).(DomainRecordOutput)
}

// Variable data depending on record type. For example, the "data" value for an A record would be the IPv4 address to which the domain will be mapped. For a CAA record, it would contain the domain name of the CA being granted permission to issue certificates.
func (o DomainRecordPtrOutput) Data() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainRecord) *string {
		if v == nil {
			return nil
		}
		return v.Data
	}).(pulumi.StringPtrOutput)
}

// An unsigned integer between 0-255 used for CAA records.
func (o DomainRecordPtrOutput) Flags() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DomainRecord) *int {
		if v == nil {
			return nil
		}
		return v.Flags
	}).(pulumi.IntPtrOutput)
}

// A unique identifier for each domain record.
func (o DomainRecordPtrOutput) Id() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DomainRecord) *int {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.IntPtrOutput)
}

// The host name, alias, or service being defined by the record.
func (o DomainRecordPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainRecord) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

// The port for SRV records.
func (o DomainRecordPtrOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DomainRecord) *int {
		if v == nil {
			return nil
		}
		return v.Port
	}).(pulumi.IntPtrOutput)
}

// The priority for SRV and MX records.
func (o DomainRecordPtrOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DomainRecord) *int {
		if v == nil {
			return nil
		}
		return v.Priority
	}).(pulumi.IntPtrOutput)
}

// The parameter tag for CAA records. Valid values are "issue", "issuewild", or "iodef"
func (o DomainRecordPtrOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainRecord) *string {
		if v == nil {
			return nil
		}
		return v.Tag
	}).(pulumi.StringPtrOutput)
}

// This value is the time to live for the record, in seconds. This defines the time frame that clients can cache queried information before a refresh should be requested.
func (o DomainRecordPtrOutput) Ttl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DomainRecord) *int {
		if v == nil {
			return nil
		}
		return v.Ttl
	}).(pulumi.IntPtrOutput)
}

// The type of the DNS record. For example: A, CNAME, TXT, ...
func (o DomainRecordPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainRecord) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

// The weight for SRV records.
func (o DomainRecordPtrOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DomainRecord) *int {
		if v == nil {
			return nil
		}
		return v.Weight
	}).(pulumi.IntPtrOutput)
}

type DomainRecordArrayOutput struct{ *pulumi.OutputState }

func (DomainRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DomainRecord)(nil)).Elem()
}

func (o DomainRecordArrayOutput) ToDomainRecordArrayOutput() DomainRecordArrayOutput {
	return o
}

func (o DomainRecordArrayOutput) ToDomainRecordArrayOutputWithContext(ctx context.Context) DomainRecordArrayOutput {
	return o
}

func (o DomainRecordArrayOutput) Index(i pulumi.IntInput) DomainRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DomainRecord {
		return vs[0].([]DomainRecord)[vs[1].(int)]
	}).(DomainRecordOutput)
}

type GetDomainsProperties struct {
	Domain *Domain `pulumi:"domain"`
}

type GetDomainsPropertiesOutput struct{ *pulumi.OutputState }

func (GetDomainsPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainsProperties)(nil)).Elem()
}

func (o GetDomainsPropertiesOutput) ToGetDomainsPropertiesOutput() GetDomainsPropertiesOutput {
	return o
}

func (o GetDomainsPropertiesOutput) ToGetDomainsPropertiesOutputWithContext(ctx context.Context) GetDomainsPropertiesOutput {
	return o
}

func (o GetDomainsPropertiesOutput) Domain() DomainPtrOutput {
	return o.ApplyT(func(v GetDomainsProperties) *Domain { return v.Domain }).(DomainPtrOutput)
}

type GetDomainsRecordProperties struct {
	DomainRecord *DomainRecord `pulumi:"domainRecord"`
}

type GetDomainsRecordPropertiesOutput struct{ *pulumi.OutputState }

func (GetDomainsRecordPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainsRecordProperties)(nil)).Elem()
}

func (o GetDomainsRecordPropertiesOutput) ToGetDomainsRecordPropertiesOutput() GetDomainsRecordPropertiesOutput {
	return o
}

func (o GetDomainsRecordPropertiesOutput) ToGetDomainsRecordPropertiesOutputWithContext(ctx context.Context) GetDomainsRecordPropertiesOutput {
	return o
}

func (o GetDomainsRecordPropertiesOutput) DomainRecord() DomainRecordPtrOutput {
	return o.ApplyT(func(v GetDomainsRecordProperties) *DomainRecord { return v.DomainRecord }).(DomainRecordPtrOutput)
}

type ListDomains struct {
	// Array of volumes.
	Domains []Domain   `pulumi:"domains"`
	Links   *PageLinks `pulumi:"links"`
	Meta    MetaMeta   `pulumi:"meta"`
}

type ListDomainsOutput struct{ *pulumi.OutputState }

func (ListDomainsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDomains)(nil)).Elem()
}

func (o ListDomainsOutput) ToListDomainsOutput() ListDomainsOutput {
	return o
}

func (o ListDomainsOutput) ToListDomainsOutputWithContext(ctx context.Context) ListDomainsOutput {
	return o
}

// Array of volumes.
func (o ListDomainsOutput) Domains() DomainArrayOutput {
	return o.ApplyT(func(v ListDomains) []Domain { return v.Domains }).(DomainArrayOutput)
}

func (o ListDomainsOutput) Links() PageLinksPtrOutput {
	return o.ApplyT(func(v ListDomains) *PageLinks { return v.Links }).(PageLinksPtrOutput)
}

func (o ListDomainsOutput) Meta() MetaMetaOutput {
	return o.ApplyT(func(v ListDomains) MetaMeta { return v.Meta }).(MetaMetaOutput)
}

type ListDomainsRecords struct {
	DomainRecords []DomainRecord `pulumi:"domainRecords"`
	Links         *PageLinks     `pulumi:"links"`
	Meta          MetaMeta       `pulumi:"meta"`
}

type ListDomainsRecordsOutput struct{ *pulumi.OutputState }

func (ListDomainsRecordsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDomainsRecords)(nil)).Elem()
}

func (o ListDomainsRecordsOutput) ToListDomainsRecordsOutput() ListDomainsRecordsOutput {
	return o
}

func (o ListDomainsRecordsOutput) ToListDomainsRecordsOutputWithContext(ctx context.Context) ListDomainsRecordsOutput {
	return o
}

func (o ListDomainsRecordsOutput) DomainRecords() DomainRecordArrayOutput {
	return o.ApplyT(func(v ListDomainsRecords) []DomainRecord { return v.DomainRecords }).(DomainRecordArrayOutput)
}

func (o ListDomainsRecordsOutput) Links() PageLinksPtrOutput {
	return o.ApplyT(func(v ListDomainsRecords) *PageLinks { return v.Links }).(PageLinksPtrOutput)
}

func (o ListDomainsRecordsOutput) Meta() MetaMetaOutput {
	return o.ApplyT(func(v ListDomainsRecords) MetaMeta { return v.Meta }).(MetaMetaOutput)
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
	pulumi.RegisterOutputType(DomainOutput{})
	pulumi.RegisterOutputType(DomainPtrOutput{})
	pulumi.RegisterOutputType(DomainArrayOutput{})
	pulumi.RegisterOutputType(DomainRecordOutput{})
	pulumi.RegisterOutputType(DomainRecordPtrOutput{})
	pulumi.RegisterOutputType(DomainRecordArrayOutput{})
	pulumi.RegisterOutputType(GetDomainsPropertiesOutput{})
	pulumi.RegisterOutputType(GetDomainsRecordPropertiesOutput{})
	pulumi.RegisterOutputType(ListDomainsOutput{})
	pulumi.RegisterOutputType(ListDomainsRecordsOutput{})
	pulumi.RegisterOutputType(MetaMetaOutput{})
	pulumi.RegisterOutputType(PageLinksOutput{})
	pulumi.RegisterOutputType(PageLinksPtrOutput{})
	pulumi.RegisterOutputType(PageLinksPagesPropertiesOutput{})
	pulumi.RegisterOutputType(PageLinksPagesPropertiesPtrOutput{})
}