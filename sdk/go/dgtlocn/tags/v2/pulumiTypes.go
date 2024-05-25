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

type GetTagProperties struct {
	// A tag is a label that can be applied to a resource (currently Droplets, Images, Volumes, Volume Snapshots, and Database clusters) in order to better organize or facilitate the lookups and actions on it.
	// Tags have two attributes: a user defined `name` attribute and an embedded `resources` attribute with information about resources that have been tagged.
	Tag *TagsType `pulumi:"tag"`
}

type GetTagPropertiesOutput struct{ *pulumi.OutputState }

func (GetTagPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTagProperties)(nil)).Elem()
}

func (o GetTagPropertiesOutput) ToGetTagPropertiesOutput() GetTagPropertiesOutput {
	return o
}

func (o GetTagPropertiesOutput) ToGetTagPropertiesOutputWithContext(ctx context.Context) GetTagPropertiesOutput {
	return o
}

// A tag is a label that can be applied to a resource (currently Droplets, Images, Volumes, Volume Snapshots, and Database clusters) in order to better organize or facilitate the lookups and actions on it.
// Tags have two attributes: a user defined `name` attribute and an embedded `resources` attribute with information about resources that have been tagged.
func (o GetTagPropertiesOutput) Tag() TagsTypePtrOutput {
	return o.ApplyT(func(v GetTagProperties) *TagsType { return v.Tag }).(TagsTypePtrOutput)
}

type ListTagsItems struct {
	Links *PageLinks `pulumi:"links"`
	Meta  MetaMeta   `pulumi:"meta"`
	Tags  []TagsType `pulumi:"tags"`
}

type ListTagsItemsOutput struct{ *pulumi.OutputState }

func (ListTagsItemsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListTagsItems)(nil)).Elem()
}

func (o ListTagsItemsOutput) ToListTagsItemsOutput() ListTagsItemsOutput {
	return o
}

func (o ListTagsItemsOutput) ToListTagsItemsOutputWithContext(ctx context.Context) ListTagsItemsOutput {
	return o
}

func (o ListTagsItemsOutput) Links() PageLinksPtrOutput {
	return o.ApplyT(func(v ListTagsItems) *PageLinks { return v.Links }).(PageLinksPtrOutput)
}

func (o ListTagsItemsOutput) Meta() MetaMetaOutput {
	return o.ApplyT(func(v ListTagsItems) MetaMeta { return v.Meta }).(MetaMetaOutput)
}

func (o ListTagsItemsOutput) Tags() TagsTypeArrayOutput {
	return o.ApplyT(func(v ListTagsItems) []TagsType { return v.Tags }).(TagsTypeArrayOutput)
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

// An embedded object containing key value pairs of resource type and resource statistics. It also includes a count of the total number of resources tagged with the current tag as well as a `last_tagged_uri` attribute set to the last resource tagged with the current tag.
type Resources struct {
	// The number of tagged objects for this type of resource.
	Count *int `pulumi:"count"`
	// The URI for the last tagged object for this type of resource.
	LastTaggedUri *string `pulumi:"lastTaggedUri"`
}

// An embedded object containing key value pairs of resource type and resource statistics. It also includes a count of the total number of resources tagged with the current tag as well as a `last_tagged_uri` attribute set to the last resource tagged with the current tag.
type ResourcesOutput struct{ *pulumi.OutputState }

func (ResourcesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Resources)(nil)).Elem()
}

func (o ResourcesOutput) ToResourcesOutput() ResourcesOutput {
	return o
}

func (o ResourcesOutput) ToResourcesOutputWithContext(ctx context.Context) ResourcesOutput {
	return o
}

// The number of tagged objects for this type of resource.
func (o ResourcesOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Resources) *int { return v.Count }).(pulumi.IntPtrOutput)
}

// The URI for the last tagged object for this type of resource.
func (o ResourcesOutput) LastTaggedUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Resources) *string { return v.LastTaggedUri }).(pulumi.StringPtrOutput)
}

type ResourcesPtrOutput struct{ *pulumi.OutputState }

func (ResourcesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Resources)(nil)).Elem()
}

func (o ResourcesPtrOutput) ToResourcesPtrOutput() ResourcesPtrOutput {
	return o
}

func (o ResourcesPtrOutput) ToResourcesPtrOutputWithContext(ctx context.Context) ResourcesPtrOutput {
	return o
}

func (o ResourcesPtrOutput) Elem() ResourcesOutput {
	return o.ApplyT(func(v *Resources) Resources {
		if v != nil {
			return *v
		}
		var ret Resources
		return ret
	}).(ResourcesOutput)
}

// The number of tagged objects for this type of resource.
func (o ResourcesPtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Resources) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

// The URI for the last tagged object for this type of resource.
func (o ResourcesPtrOutput) LastTaggedUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Resources) *string {
		if v == nil {
			return nil
		}
		return v.LastTaggedUri
	}).(pulumi.StringPtrOutput)
}

type ResourcesItemProperties struct {
	// The identifier of a resource.
	ResourceId *string `pulumi:"resourceId"`
	// The type of the resource.
	ResourceType *ResourcesItemPropertiesResourceType `pulumi:"resourceType"`
}

// ResourcesItemPropertiesInput is an input type that accepts ResourcesItemPropertiesArgs and ResourcesItemPropertiesOutput values.
// You can construct a concrete instance of `ResourcesItemPropertiesInput` via:
//
//	ResourcesItemPropertiesArgs{...}
type ResourcesItemPropertiesInput interface {
	pulumi.Input

	ToResourcesItemPropertiesOutput() ResourcesItemPropertiesOutput
	ToResourcesItemPropertiesOutputWithContext(context.Context) ResourcesItemPropertiesOutput
}

type ResourcesItemPropertiesArgs struct {
	// The identifier of a resource.
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
	// The type of the resource.
	ResourceType ResourcesItemPropertiesResourceTypePtrInput `pulumi:"resourceType"`
}

func (ResourcesItemPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourcesItemProperties)(nil)).Elem()
}

func (i ResourcesItemPropertiesArgs) ToResourcesItemPropertiesOutput() ResourcesItemPropertiesOutput {
	return i.ToResourcesItemPropertiesOutputWithContext(context.Background())
}

func (i ResourcesItemPropertiesArgs) ToResourcesItemPropertiesOutputWithContext(ctx context.Context) ResourcesItemPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcesItemPropertiesOutput)
}

// ResourcesItemPropertiesArrayInput is an input type that accepts ResourcesItemPropertiesArray and ResourcesItemPropertiesArrayOutput values.
// You can construct a concrete instance of `ResourcesItemPropertiesArrayInput` via:
//
//	ResourcesItemPropertiesArray{ ResourcesItemPropertiesArgs{...} }
type ResourcesItemPropertiesArrayInput interface {
	pulumi.Input

	ToResourcesItemPropertiesArrayOutput() ResourcesItemPropertiesArrayOutput
	ToResourcesItemPropertiesArrayOutputWithContext(context.Context) ResourcesItemPropertiesArrayOutput
}

type ResourcesItemPropertiesArray []ResourcesItemPropertiesInput

func (ResourcesItemPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourcesItemProperties)(nil)).Elem()
}

func (i ResourcesItemPropertiesArray) ToResourcesItemPropertiesArrayOutput() ResourcesItemPropertiesArrayOutput {
	return i.ToResourcesItemPropertiesArrayOutputWithContext(context.Background())
}

func (i ResourcesItemPropertiesArray) ToResourcesItemPropertiesArrayOutputWithContext(ctx context.Context) ResourcesItemPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcesItemPropertiesArrayOutput)
}

type ResourcesItemPropertiesOutput struct{ *pulumi.OutputState }

func (ResourcesItemPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourcesItemProperties)(nil)).Elem()
}

func (o ResourcesItemPropertiesOutput) ToResourcesItemPropertiesOutput() ResourcesItemPropertiesOutput {
	return o
}

func (o ResourcesItemPropertiesOutput) ToResourcesItemPropertiesOutputWithContext(ctx context.Context) ResourcesItemPropertiesOutput {
	return o
}

// The identifier of a resource.
func (o ResourcesItemPropertiesOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourcesItemProperties) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// The type of the resource.
func (o ResourcesItemPropertiesOutput) ResourceType() ResourcesItemPropertiesResourceTypePtrOutput {
	return o.ApplyT(func(v ResourcesItemProperties) *ResourcesItemPropertiesResourceType { return v.ResourceType }).(ResourcesItemPropertiesResourceTypePtrOutput)
}

type ResourcesItemPropertiesArrayOutput struct{ *pulumi.OutputState }

func (ResourcesItemPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourcesItemProperties)(nil)).Elem()
}

func (o ResourcesItemPropertiesArrayOutput) ToResourcesItemPropertiesArrayOutput() ResourcesItemPropertiesArrayOutput {
	return o
}

func (o ResourcesItemPropertiesArrayOutput) ToResourcesItemPropertiesArrayOutputWithContext(ctx context.Context) ResourcesItemPropertiesArrayOutput {
	return o
}

func (o ResourcesItemPropertiesArrayOutput) Index(i pulumi.IntInput) ResourcesItemPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourcesItemProperties {
		return vs[0].([]ResourcesItemProperties)[vs[1].(int)]
	}).(ResourcesItemPropertiesOutput)
}

// A tag is a label that can be applied to a resource (currently Droplets, Images, Volumes, Volume Snapshots, and Database clusters) in order to better organize or facilitate the lookups and actions on it.
// Tags have two attributes: a user defined `name` attribute and an embedded `resources` attribute with information about resources that have been tagged.
type TagsType struct {
	// The name of the tag. Tags may contain letters, numbers, colons, dashes, and underscores.
	// There is a limit of 255 characters per tag.
	//
	// **Note:** Tag names are case stable, which means the capitalization you use when you first create a tag is canonical.
	//
	// When working with tags in the API, you must use the tag's canonical capitalization. For example, if you create a tag named "PROD", the URL to add that tag to a resource would be `https://api.digitalocean.com/v2/tags/PROD/resources` (not `/v2/tags/prod/resources`).
	//
	// Tagged resources in the control panel will always display the canonical capitalization. For example, if you create a tag named "PROD", you can tag resources in the control panel by entering "prod". The tag will still display with its canonical capitalization, "PROD".
	Name *string `pulumi:"name"`
	// An embedded object containing key value pairs of resource type and resource statistics. It also includes a count of the total number of resources tagged with the current tag as well as a `last_tagged_uri` attribute set to the last resource tagged with the current tag.
	Resources *TagsResources `pulumi:"resources"`
}

// A tag is a label that can be applied to a resource (currently Droplets, Images, Volumes, Volume Snapshots, and Database clusters) in order to better organize or facilitate the lookups and actions on it.
// Tags have two attributes: a user defined `name` attribute and an embedded `resources` attribute with information about resources that have been tagged.
type TagsTypeOutput struct{ *pulumi.OutputState }

func (TagsTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagsType)(nil)).Elem()
}

func (o TagsTypeOutput) ToTagsTypeOutput() TagsTypeOutput {
	return o
}

func (o TagsTypeOutput) ToTagsTypeOutputWithContext(ctx context.Context) TagsTypeOutput {
	return o
}

// The name of the tag. Tags may contain letters, numbers, colons, dashes, and underscores.
// There is a limit of 255 characters per tag.
//
// **Note:** Tag names are case stable, which means the capitalization you use when you first create a tag is canonical.
//
// When working with tags in the API, you must use the tag's canonical capitalization. For example, if you create a tag named "PROD", the URL to add that tag to a resource would be `https://api.digitalocean.com/v2/tags/PROD/resources` (not `/v2/tags/prod/resources`).
//
// Tagged resources in the control panel will always display the canonical capitalization. For example, if you create a tag named "PROD", you can tag resources in the control panel by entering "prod". The tag will still display with its canonical capitalization, "PROD".
func (o TagsTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TagsType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// An embedded object containing key value pairs of resource type and resource statistics. It also includes a count of the total number of resources tagged with the current tag as well as a `last_tagged_uri` attribute set to the last resource tagged with the current tag.
func (o TagsTypeOutput) Resources() TagsResourcesPtrOutput {
	return o.ApplyT(func(v TagsType) *TagsResources { return v.Resources }).(TagsResourcesPtrOutput)
}

type TagsTypePtrOutput struct{ *pulumi.OutputState }

func (TagsTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagsType)(nil)).Elem()
}

func (o TagsTypePtrOutput) ToTagsTypePtrOutput() TagsTypePtrOutput {
	return o
}

func (o TagsTypePtrOutput) ToTagsTypePtrOutputWithContext(ctx context.Context) TagsTypePtrOutput {
	return o
}

func (o TagsTypePtrOutput) Elem() TagsTypeOutput {
	return o.ApplyT(func(v *TagsType) TagsType {
		if v != nil {
			return *v
		}
		var ret TagsType
		return ret
	}).(TagsTypeOutput)
}

// The name of the tag. Tags may contain letters, numbers, colons, dashes, and underscores.
// There is a limit of 255 characters per tag.
//
// **Note:** Tag names are case stable, which means the capitalization you use when you first create a tag is canonical.
//
// When working with tags in the API, you must use the tag's canonical capitalization. For example, if you create a tag named "PROD", the URL to add that tag to a resource would be `https://api.digitalocean.com/v2/tags/PROD/resources` (not `/v2/tags/prod/resources`).
//
// Tagged resources in the control panel will always display the canonical capitalization. For example, if you create a tag named "PROD", you can tag resources in the control panel by entering "prod". The tag will still display with its canonical capitalization, "PROD".
func (o TagsTypePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TagsType) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

// An embedded object containing key value pairs of resource type and resource statistics. It also includes a count of the total number of resources tagged with the current tag as well as a `last_tagged_uri` attribute set to the last resource tagged with the current tag.
func (o TagsTypePtrOutput) Resources() TagsResourcesPtrOutput {
	return o.ApplyT(func(v *TagsType) *TagsResources {
		if v == nil {
			return nil
		}
		return v.Resources
	}).(TagsResourcesPtrOutput)
}

type TagsTypeArrayOutput struct{ *pulumi.OutputState }

func (TagsTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TagsType)(nil)).Elem()
}

func (o TagsTypeArrayOutput) ToTagsTypeArrayOutput() TagsTypeArrayOutput {
	return o
}

func (o TagsTypeArrayOutput) ToTagsTypeArrayOutputWithContext(ctx context.Context) TagsTypeArrayOutput {
	return o
}

func (o TagsTypeArrayOutput) Index(i pulumi.IntInput) TagsTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TagsType {
		return vs[0].([]TagsType)[vs[1].(int)]
	}).(TagsTypeOutput)
}

// An embedded object containing key value pairs of resource type and resource statistics. It also includes a count of the total number of resources tagged with the current tag as well as a `last_tagged_uri` attribute set to the last resource tagged with the current tag.
type TagsResources struct {
	// The number of tagged objects for this type of resource.
	Count *int `pulumi:"count"`
	// The URI for the last tagged object for this type of resource.
	LastTaggedUri *string `pulumi:"lastTaggedUri"`
}

// An embedded object containing key value pairs of resource type and resource statistics. It also includes a count of the total number of resources tagged with the current tag as well as a `last_tagged_uri` attribute set to the last resource tagged with the current tag.
type TagsResourcesOutput struct{ *pulumi.OutputState }

func (TagsResourcesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagsResources)(nil)).Elem()
}

func (o TagsResourcesOutput) ToTagsResourcesOutput() TagsResourcesOutput {
	return o
}

func (o TagsResourcesOutput) ToTagsResourcesOutputWithContext(ctx context.Context) TagsResourcesOutput {
	return o
}

// The number of tagged objects for this type of resource.
func (o TagsResourcesOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v TagsResources) *int { return v.Count }).(pulumi.IntPtrOutput)
}

// The URI for the last tagged object for this type of resource.
func (o TagsResourcesOutput) LastTaggedUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TagsResources) *string { return v.LastTaggedUri }).(pulumi.StringPtrOutput)
}

type TagsResourcesPtrOutput struct{ *pulumi.OutputState }

func (TagsResourcesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagsResources)(nil)).Elem()
}

func (o TagsResourcesPtrOutput) ToTagsResourcesPtrOutput() TagsResourcesPtrOutput {
	return o
}

func (o TagsResourcesPtrOutput) ToTagsResourcesPtrOutputWithContext(ctx context.Context) TagsResourcesPtrOutput {
	return o
}

func (o TagsResourcesPtrOutput) Elem() TagsResourcesOutput {
	return o.ApplyT(func(v *TagsResources) TagsResources {
		if v != nil {
			return *v
		}
		var ret TagsResources
		return ret
	}).(TagsResourcesOutput)
}

// The number of tagged objects for this type of resource.
func (o TagsResourcesPtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TagsResources) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

// The URI for the last tagged object for this type of resource.
func (o TagsResourcesPtrOutput) LastTaggedUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TagsResources) *string {
		if v == nil {
			return nil
		}
		return v.LastTaggedUri
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourcesItemPropertiesInput)(nil)).Elem(), ResourcesItemPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourcesItemPropertiesArrayInput)(nil)).Elem(), ResourcesItemPropertiesArray{})
	pulumi.RegisterOutputType(GetTagPropertiesOutput{})
	pulumi.RegisterOutputType(ListTagsItemsOutput{})
	pulumi.RegisterOutputType(MetaMetaOutput{})
	pulumi.RegisterOutputType(PageLinksOutput{})
	pulumi.RegisterOutputType(PageLinksPtrOutput{})
	pulumi.RegisterOutputType(PageLinksPagesPropertiesOutput{})
	pulumi.RegisterOutputType(PageLinksPagesPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ResourcesOutput{})
	pulumi.RegisterOutputType(ResourcesPtrOutput{})
	pulumi.RegisterOutputType(ResourcesItemPropertiesOutput{})
	pulumi.RegisterOutputType(ResourcesItemPropertiesArrayOutput{})
	pulumi.RegisterOutputType(TagsTypeOutput{})
	pulumi.RegisterOutputType(TagsTypePtrOutput{})
	pulumi.RegisterOutputType(TagsTypeArrayOutput{})
	pulumi.RegisterOutputType(TagsResourcesOutput{})
	pulumi.RegisterOutputType(TagsResourcesPtrOutput{})
}
