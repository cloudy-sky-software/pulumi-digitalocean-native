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

type GetSnapshotsProperties struct {
	Snapshot *Snapshots `pulumi:"snapshot"`
}

type GetSnapshotsPropertiesOutput struct{ *pulumi.OutputState }

func (GetSnapshotsPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSnapshotsProperties)(nil)).Elem()
}

func (o GetSnapshotsPropertiesOutput) ToGetSnapshotsPropertiesOutput() GetSnapshotsPropertiesOutput {
	return o
}

func (o GetSnapshotsPropertiesOutput) ToGetSnapshotsPropertiesOutputWithContext(ctx context.Context) GetSnapshotsPropertiesOutput {
	return o
}

func (o GetSnapshotsPropertiesOutput) Snapshot() SnapshotsPtrOutput {
	return o.ApplyT(func(v GetSnapshotsProperties) *Snapshots { return v.Snapshot }).(SnapshotsPtrOutput)
}

type ListSnapshots struct {
	Links     *PageLinks  `pulumi:"links"`
	Meta      MetaMeta    `pulumi:"meta"`
	Snapshots []Snapshots `pulumi:"snapshots"`
}

type ListSnapshotsOutput struct{ *pulumi.OutputState }

func (ListSnapshotsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSnapshots)(nil)).Elem()
}

func (o ListSnapshotsOutput) ToListSnapshotsOutput() ListSnapshotsOutput {
	return o
}

func (o ListSnapshotsOutput) ToListSnapshotsOutputWithContext(ctx context.Context) ListSnapshotsOutput {
	return o
}

func (o ListSnapshotsOutput) Links() PageLinksPtrOutput {
	return o.ApplyT(func(v ListSnapshots) *PageLinks { return v.Links }).(PageLinksPtrOutput)
}

func (o ListSnapshotsOutput) Meta() MetaMetaOutput {
	return o.ApplyT(func(v ListSnapshots) MetaMeta { return v.Meta }).(MetaMetaOutput)
}

func (o ListSnapshotsOutput) Snapshots() SnapshotsArrayOutput {
	return o.ApplyT(func(v ListSnapshots) []Snapshots { return v.Snapshots }).(SnapshotsArrayOutput)
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

type Snapshots struct {
	// A time value given in ISO8601 combined date and time format that represents when the snapshot was created.
	CreatedAt string `pulumi:"createdAt"`
	// The minimum size in GB required for a volume or Droplet to use this snapshot.
	MinDiskSize int `pulumi:"minDiskSize"`
	// A human-readable name for the snapshot.
	Name string `pulumi:"name"`
	// An array of the regions that the snapshot is available in. The regions are represented by their identifying slug values.
	Regions []string `pulumi:"regions"`
	// The unique identifier for the resource that the snapshot originated from.
	ResourceId string `pulumi:"resourceId"`
	// The type of resource that the snapshot originated from.
	ResourceType SnapshotsPropertiesResourceType `pulumi:"resourceType"`
	// The billable size of the snapshot in gigabytes.
	SizeGigabytes float64 `pulumi:"sizeGigabytes"`
	// An array of Tags the snapshot has been tagged with.
	Tags []string `pulumi:"tags"`
}

type SnapshotsOutput struct{ *pulumi.OutputState }

func (SnapshotsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Snapshots)(nil)).Elem()
}

func (o SnapshotsOutput) ToSnapshotsOutput() SnapshotsOutput {
	return o
}

func (o SnapshotsOutput) ToSnapshotsOutputWithContext(ctx context.Context) SnapshotsOutput {
	return o
}

// A time value given in ISO8601 combined date and time format that represents when the snapshot was created.
func (o SnapshotsOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v Snapshots) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The minimum size in GB required for a volume or Droplet to use this snapshot.
func (o SnapshotsOutput) MinDiskSize() pulumi.IntOutput {
	return o.ApplyT(func(v Snapshots) int { return v.MinDiskSize }).(pulumi.IntOutput)
}

// A human-readable name for the snapshot.
func (o SnapshotsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Snapshots) string { return v.Name }).(pulumi.StringOutput)
}

// An array of the regions that the snapshot is available in. The regions are represented by their identifying slug values.
func (o SnapshotsOutput) Regions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Snapshots) []string { return v.Regions }).(pulumi.StringArrayOutput)
}

// The unique identifier for the resource that the snapshot originated from.
func (o SnapshotsOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v Snapshots) string { return v.ResourceId }).(pulumi.StringOutput)
}

// The type of resource that the snapshot originated from.
func (o SnapshotsOutput) ResourceType() SnapshotsPropertiesResourceTypeOutput {
	return o.ApplyT(func(v Snapshots) SnapshotsPropertiesResourceType { return v.ResourceType }).(SnapshotsPropertiesResourceTypeOutput)
}

// The billable size of the snapshot in gigabytes.
func (o SnapshotsOutput) SizeGigabytes() pulumi.Float64Output {
	return o.ApplyT(func(v Snapshots) float64 { return v.SizeGigabytes }).(pulumi.Float64Output)
}

// An array of Tags the snapshot has been tagged with.
func (o SnapshotsOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Snapshots) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

type SnapshotsPtrOutput struct{ *pulumi.OutputState }

func (SnapshotsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshots)(nil)).Elem()
}

func (o SnapshotsPtrOutput) ToSnapshotsPtrOutput() SnapshotsPtrOutput {
	return o
}

func (o SnapshotsPtrOutput) ToSnapshotsPtrOutputWithContext(ctx context.Context) SnapshotsPtrOutput {
	return o
}

func (o SnapshotsPtrOutput) Elem() SnapshotsOutput {
	return o.ApplyT(func(v *Snapshots) Snapshots {
		if v != nil {
			return *v
		}
		var ret Snapshots
		return ret
	}).(SnapshotsOutput)
}

// A time value given in ISO8601 combined date and time format that represents when the snapshot was created.
func (o SnapshotsPtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snapshots) *string {
		if v == nil {
			return nil
		}
		return &v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

// The minimum size in GB required for a volume or Droplet to use this snapshot.
func (o SnapshotsPtrOutput) MinDiskSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Snapshots) *int {
		if v == nil {
			return nil
		}
		return &v.MinDiskSize
	}).(pulumi.IntPtrOutput)
}

// A human-readable name for the snapshot.
func (o SnapshotsPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snapshots) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

// An array of the regions that the snapshot is available in. The regions are represented by their identifying slug values.
func (o SnapshotsPtrOutput) Regions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Snapshots) []string {
		if v == nil {
			return nil
		}
		return v.Regions
	}).(pulumi.StringArrayOutput)
}

// The unique identifier for the resource that the snapshot originated from.
func (o SnapshotsPtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snapshots) *string {
		if v == nil {
			return nil
		}
		return &v.ResourceId
	}).(pulumi.StringPtrOutput)
}

// The type of resource that the snapshot originated from.
func (o SnapshotsPtrOutput) ResourceType() SnapshotsPropertiesResourceTypePtrOutput {
	return o.ApplyT(func(v *Snapshots) *SnapshotsPropertiesResourceType {
		if v == nil {
			return nil
		}
		return &v.ResourceType
	}).(SnapshotsPropertiesResourceTypePtrOutput)
}

// The billable size of the snapshot in gigabytes.
func (o SnapshotsPtrOutput) SizeGigabytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Snapshots) *float64 {
		if v == nil {
			return nil
		}
		return &v.SizeGigabytes
	}).(pulumi.Float64PtrOutput)
}

// An array of Tags the snapshot has been tagged with.
func (o SnapshotsPtrOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Snapshots) []string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringArrayOutput)
}

type SnapshotsArrayOutput struct{ *pulumi.OutputState }

func (SnapshotsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Snapshots)(nil)).Elem()
}

func (o SnapshotsArrayOutput) ToSnapshotsArrayOutput() SnapshotsArrayOutput {
	return o
}

func (o SnapshotsArrayOutput) ToSnapshotsArrayOutputWithContext(ctx context.Context) SnapshotsArrayOutput {
	return o
}

func (o SnapshotsArrayOutput) Index(i pulumi.IntInput) SnapshotsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Snapshots {
		return vs[0].([]Snapshots)[vs[1].(int)]
	}).(SnapshotsOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSnapshotsPropertiesOutput{})
	pulumi.RegisterOutputType(ListSnapshotsOutput{})
	pulumi.RegisterOutputType(MetaMetaOutput{})
	pulumi.RegisterOutputType(PageLinksOutput{})
	pulumi.RegisterOutputType(PageLinksPtrOutput{})
	pulumi.RegisterOutputType(PageLinksPagesPropertiesOutput{})
	pulumi.RegisterOutputType(PageLinksPagesPropertiesPtrOutput{})
	pulumi.RegisterOutputType(SnapshotsOutput{})
	pulumi.RegisterOutputType(SnapshotsPtrOutput{})
	pulumi.RegisterOutputType(SnapshotsArrayOutput{})
}
