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

type GetProjectProperties struct {
	Project *Project `pulumi:"project"`
}

type GetProjectPropertiesOutput struct{ *pulumi.OutputState }

func (GetProjectPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectProperties)(nil)).Elem()
}

func (o GetProjectPropertiesOutput) ToGetProjectPropertiesOutput() GetProjectPropertiesOutput {
	return o
}

func (o GetProjectPropertiesOutput) ToGetProjectPropertiesOutputWithContext(ctx context.Context) GetProjectPropertiesOutput {
	return o
}

func (o GetProjectPropertiesOutput) Project() ProjectPtrOutput {
	return o.ApplyT(func(v GetProjectProperties) *Project { return v.Project }).(ProjectPtrOutput)
}

type GetProjectsDefaultProperties struct {
	Project *Project `pulumi:"project"`
}

type GetProjectsDefaultPropertiesOutput struct{ *pulumi.OutputState }

func (GetProjectsDefaultPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectsDefaultProperties)(nil)).Elem()
}

func (o GetProjectsDefaultPropertiesOutput) ToGetProjectsDefaultPropertiesOutput() GetProjectsDefaultPropertiesOutput {
	return o
}

func (o GetProjectsDefaultPropertiesOutput) ToGetProjectsDefaultPropertiesOutputWithContext(ctx context.Context) GetProjectsDefaultPropertiesOutput {
	return o
}

func (o GetProjectsDefaultPropertiesOutput) Project() ProjectPtrOutput {
	return o.ApplyT(func(v GetProjectsDefaultProperties) *Project { return v.Project }).(ProjectPtrOutput)
}

type ListProjectsItems struct {
	Links    *PageLinks `pulumi:"links"`
	Meta     MetaMeta   `pulumi:"meta"`
	Projects []Project  `pulumi:"projects"`
}

type ListProjectsItemsOutput struct{ *pulumi.OutputState }

func (ListProjectsItemsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListProjectsItems)(nil)).Elem()
}

func (o ListProjectsItemsOutput) ToListProjectsItemsOutput() ListProjectsItemsOutput {
	return o
}

func (o ListProjectsItemsOutput) ToListProjectsItemsOutputWithContext(ctx context.Context) ListProjectsItemsOutput {
	return o
}

func (o ListProjectsItemsOutput) Links() PageLinksPtrOutput {
	return o.ApplyT(func(v ListProjectsItems) *PageLinks { return v.Links }).(PageLinksPtrOutput)
}

func (o ListProjectsItemsOutput) Meta() MetaMetaOutput {
	return o.ApplyT(func(v ListProjectsItems) MetaMeta { return v.Meta }).(MetaMetaOutput)
}

func (o ListProjectsItemsOutput) Projects() ProjectArrayOutput {
	return o.ApplyT(func(v ListProjectsItems) []Project { return v.Projects }).(ProjectArrayOutput)
}

type ListProjectsResourcesDefaultItems struct {
	Links     *PageLinks `pulumi:"links"`
	Meta      MetaMeta   `pulumi:"meta"`
	Resources []Resource `pulumi:"resources"`
}

type ListProjectsResourcesDefaultItemsOutput struct{ *pulumi.OutputState }

func (ListProjectsResourcesDefaultItemsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListProjectsResourcesDefaultItems)(nil)).Elem()
}

func (o ListProjectsResourcesDefaultItemsOutput) ToListProjectsResourcesDefaultItemsOutput() ListProjectsResourcesDefaultItemsOutput {
	return o
}

func (o ListProjectsResourcesDefaultItemsOutput) ToListProjectsResourcesDefaultItemsOutputWithContext(ctx context.Context) ListProjectsResourcesDefaultItemsOutput {
	return o
}

func (o ListProjectsResourcesDefaultItemsOutput) Links() PageLinksPtrOutput {
	return o.ApplyT(func(v ListProjectsResourcesDefaultItems) *PageLinks { return v.Links }).(PageLinksPtrOutput)
}

func (o ListProjectsResourcesDefaultItemsOutput) Meta() MetaMetaOutput {
	return o.ApplyT(func(v ListProjectsResourcesDefaultItems) MetaMeta { return v.Meta }).(MetaMetaOutput)
}

func (o ListProjectsResourcesDefaultItemsOutput) Resources() ResourceArrayOutput {
	return o.ApplyT(func(v ListProjectsResourcesDefaultItems) []Resource { return v.Resources }).(ResourceArrayOutput)
}

type ListProjectsResourcesItems struct {
	Links     *PageLinks `pulumi:"links"`
	Meta      MetaMeta   `pulumi:"meta"`
	Resources []Resource `pulumi:"resources"`
}

type ListProjectsResourcesItemsOutput struct{ *pulumi.OutputState }

func (ListProjectsResourcesItemsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListProjectsResourcesItems)(nil)).Elem()
}

func (o ListProjectsResourcesItemsOutput) ToListProjectsResourcesItemsOutput() ListProjectsResourcesItemsOutput {
	return o
}

func (o ListProjectsResourcesItemsOutput) ToListProjectsResourcesItemsOutputWithContext(ctx context.Context) ListProjectsResourcesItemsOutput {
	return o
}

func (o ListProjectsResourcesItemsOutput) Links() PageLinksPtrOutput {
	return o.ApplyT(func(v ListProjectsResourcesItems) *PageLinks { return v.Links }).(PageLinksPtrOutput)
}

func (o ListProjectsResourcesItemsOutput) Meta() MetaMetaOutput {
	return o.ApplyT(func(v ListProjectsResourcesItems) MetaMeta { return v.Meta }).(MetaMetaOutput)
}

func (o ListProjectsResourcesItemsOutput) Resources() ResourceArrayOutput {
	return o.ApplyT(func(v ListProjectsResourcesItems) []Resource { return v.Resources }).(ResourceArrayOutput)
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

type Project struct {
	// A time value given in ISO8601 combined date and time format that represents when the project was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The description of the project. The maximum length is 255 characters.
	Description *string `pulumi:"description"`
	// The environment of the project's resources.
	Environment *ProjectBaseEnvironment `pulumi:"environment"`
	// The unique universal identifier of this project.
	Id *string `pulumi:"id"`
	// If true, all resources will be added to this project if no project is specified.
	IsDefault *bool `pulumi:"isDefault"`
	// The human-readable name for the project. The maximum length is 175 characters and the name must be unique.
	Name *string `pulumi:"name"`
	// The integer id of the project owner.
	OwnerId *int `pulumi:"ownerId"`
	// The unique universal identifier of the project owner.
	OwnerUuid *string `pulumi:"ownerUuid"`
	// The purpose of the project. The maximum length is 255 characters. It can
	// have one of the following values:
	//
	// - Just trying out DigitalOcean
	// - Class project / Educational purposes
	// - Website or blog
	// - Web Application
	// - Service or API
	// - Mobile Application
	// - Machine learning / AI / Data processing
	// - IoT
	// - Operational / Developer tooling
	//
	// If another value for purpose is specified, for example, "your custom purpose",
	// your purpose will be stored as `Other: your custom purpose`.
	Purpose *string `pulumi:"purpose"`
	// A time value given in ISO8601 combined date and time format that represents when the project was updated.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type ProjectOutput struct{ *pulumi.OutputState }

func (ProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Project)(nil)).Elem()
}

func (o ProjectOutput) ToProjectOutput() ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return o
}

// A time value given in ISO8601 combined date and time format that represents when the project was created.
func (o ProjectOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Project) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The description of the project. The maximum length is 255 characters.
func (o ProjectOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Project) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The environment of the project's resources.
func (o ProjectOutput) Environment() ProjectBaseEnvironmentPtrOutput {
	return o.ApplyT(func(v Project) *ProjectBaseEnvironment { return v.Environment }).(ProjectBaseEnvironmentPtrOutput)
}

// The unique universal identifier of this project.
func (o ProjectOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Project) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// If true, all resources will be added to this project if no project is specified.
func (o ProjectOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Project) *bool { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

// The human-readable name for the project. The maximum length is 175 characters and the name must be unique.
func (o ProjectOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Project) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The integer id of the project owner.
func (o ProjectOutput) OwnerId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Project) *int { return v.OwnerId }).(pulumi.IntPtrOutput)
}

// The unique universal identifier of the project owner.
func (o ProjectOutput) OwnerUuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Project) *string { return v.OwnerUuid }).(pulumi.StringPtrOutput)
}

// The purpose of the project. The maximum length is 255 characters. It can
// have one of the following values:
//
// - Just trying out DigitalOcean
// - Class project / Educational purposes
// - Website or blog
// - Web Application
// - Service or API
// - Mobile Application
// - Machine learning / AI / Data processing
// - IoT
// - Operational / Developer tooling
//
// If another value for purpose is specified, for example, "your custom purpose",
// your purpose will be stored as `Other: your custom purpose`.
func (o ProjectOutput) Purpose() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Project) *string { return v.Purpose }).(pulumi.StringPtrOutput)
}

// A time value given in ISO8601 combined date and time format that represents when the project was updated.
func (o ProjectOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Project) *string { return v.UpdatedAt }).(pulumi.StringPtrOutput)
}

type ProjectPtrOutput struct{ *pulumi.OutputState }

func (ProjectPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (o ProjectPtrOutput) ToProjectPtrOutput() ProjectPtrOutput {
	return o
}

func (o ProjectPtrOutput) ToProjectPtrOutputWithContext(ctx context.Context) ProjectPtrOutput {
	return o
}

func (o ProjectPtrOutput) Elem() ProjectOutput {
	return o.ApplyT(func(v *Project) Project {
		if v != nil {
			return *v
		}
		var ret Project
		return ret
	}).(ProjectOutput)
}

// A time value given in ISO8601 combined date and time format that represents when the project was created.
func (o ProjectPtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

// The description of the project. The maximum length is 255 characters.
func (o ProjectPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

// The environment of the project's resources.
func (o ProjectPtrOutput) Environment() ProjectBaseEnvironmentPtrOutput {
	return o.ApplyT(func(v *Project) *ProjectBaseEnvironment {
		if v == nil {
			return nil
		}
		return v.Environment
	}).(ProjectBaseEnvironmentPtrOutput)
}

// The unique universal identifier of this project.
func (o ProjectPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

// If true, all resources will be added to this project if no project is specified.
func (o ProjectPtrOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Project) *bool {
		if v == nil {
			return nil
		}
		return v.IsDefault
	}).(pulumi.BoolPtrOutput)
}

// The human-readable name for the project. The maximum length is 175 characters and the name must be unique.
func (o ProjectPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

// The integer id of the project owner.
func (o ProjectPtrOutput) OwnerId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Project) *int {
		if v == nil {
			return nil
		}
		return v.OwnerId
	}).(pulumi.IntPtrOutput)
}

// The unique universal identifier of the project owner.
func (o ProjectPtrOutput) OwnerUuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) *string {
		if v == nil {
			return nil
		}
		return v.OwnerUuid
	}).(pulumi.StringPtrOutput)
}

// The purpose of the project. The maximum length is 255 characters. It can
// have one of the following values:
//
// - Just trying out DigitalOcean
// - Class project / Educational purposes
// - Website or blog
// - Web Application
// - Service or API
// - Mobile Application
// - Machine learning / AI / Data processing
// - IoT
// - Operational / Developer tooling
//
// If another value for purpose is specified, for example, "your custom purpose",
// your purpose will be stored as `Other: your custom purpose`.
func (o ProjectPtrOutput) Purpose() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) *string {
		if v == nil {
			return nil
		}
		return v.Purpose
	}).(pulumi.StringPtrOutput)
}

// A time value given in ISO8601 combined date and time format that represents when the project was updated.
func (o ProjectPtrOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) *string {
		if v == nil {
			return nil
		}
		return v.UpdatedAt
	}).(pulumi.StringPtrOutput)
}

type ProjectArrayOutput struct{ *pulumi.OutputState }

func (ProjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Project)(nil)).Elem()
}

func (o ProjectArrayOutput) ToProjectArrayOutput() ProjectArrayOutput {
	return o
}

func (o ProjectArrayOutput) ToProjectArrayOutputWithContext(ctx context.Context) ProjectArrayOutput {
	return o
}

func (o ProjectArrayOutput) Index(i pulumi.IntInput) ProjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Project {
		return vs[0].([]Project)[vs[1].(int)]
	}).(ProjectOutput)
}

type Resource struct {
	// A time value given in ISO8601 combined date and time format that represents when the project was created.
	AssignedAt *string `pulumi:"assignedAt"`
	// The links object contains the `self` object, which contains the resource relationship.
	Links *ResourceLinksProperties `pulumi:"links"`
	// The status of assigning and fetching the resources.
	Status *ResourceStatus `pulumi:"status"`
	// The uniform resource name (URN) for the resource in the format do:resource_type:resource_id.
	Urn *string `pulumi:"urn"`
}

type ResourceOutput struct{ *pulumi.OutputState }

func (ResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Resource)(nil)).Elem()
}

func (o ResourceOutput) ToResourceOutput() ResourceOutput {
	return o
}

func (o ResourceOutput) ToResourceOutputWithContext(ctx context.Context) ResourceOutput {
	return o
}

// A time value given in ISO8601 combined date and time format that represents when the project was created.
func (o ResourceOutput) AssignedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Resource) *string { return v.AssignedAt }).(pulumi.StringPtrOutput)
}

// The links object contains the `self` object, which contains the resource relationship.
func (o ResourceOutput) Links() ResourceLinksPropertiesPtrOutput {
	return o.ApplyT(func(v Resource) *ResourceLinksProperties { return v.Links }).(ResourceLinksPropertiesPtrOutput)
}

// The status of assigning and fetching the resources.
func (o ResourceOutput) Status() ResourceStatusPtrOutput {
	return o.ApplyT(func(v Resource) *ResourceStatus { return v.Status }).(ResourceStatusPtrOutput)
}

// The uniform resource name (URN) for the resource in the format do:resource_type:resource_id.
func (o ResourceOutput) Urn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Resource) *string { return v.Urn }).(pulumi.StringPtrOutput)
}

type ResourceArrayOutput struct{ *pulumi.OutputState }

func (ResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Resource)(nil)).Elem()
}

func (o ResourceArrayOutput) ToResourceArrayOutput() ResourceArrayOutput {
	return o
}

func (o ResourceArrayOutput) ToResourceArrayOutputWithContext(ctx context.Context) ResourceArrayOutput {
	return o
}

func (o ResourceArrayOutput) Index(i pulumi.IntInput) ResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Resource {
		return vs[0].([]Resource)[vs[1].(int)]
	}).(ResourceOutput)
}

// The links object contains the `self` object, which contains the resource relationship.
type ResourceLinksProperties struct {
	// A URI that can be used to retrieve the resource.
	Self *string `pulumi:"self"`
}

// The links object contains the `self` object, which contains the resource relationship.
type ResourceLinksPropertiesOutput struct{ *pulumi.OutputState }

func (ResourceLinksPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceLinksProperties)(nil)).Elem()
}

func (o ResourceLinksPropertiesOutput) ToResourceLinksPropertiesOutput() ResourceLinksPropertiesOutput {
	return o
}

func (o ResourceLinksPropertiesOutput) ToResourceLinksPropertiesOutputWithContext(ctx context.Context) ResourceLinksPropertiesOutput {
	return o
}

// A URI that can be used to retrieve the resource.
func (o ResourceLinksPropertiesOutput) Self() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceLinksProperties) *string { return v.Self }).(pulumi.StringPtrOutput)
}

type ResourceLinksPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ResourceLinksPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceLinksProperties)(nil)).Elem()
}

func (o ResourceLinksPropertiesPtrOutput) ToResourceLinksPropertiesPtrOutput() ResourceLinksPropertiesPtrOutput {
	return o
}

func (o ResourceLinksPropertiesPtrOutput) ToResourceLinksPropertiesPtrOutputWithContext(ctx context.Context) ResourceLinksPropertiesPtrOutput {
	return o
}

func (o ResourceLinksPropertiesPtrOutput) Elem() ResourceLinksPropertiesOutput {
	return o.ApplyT(func(v *ResourceLinksProperties) ResourceLinksProperties {
		if v != nil {
			return *v
		}
		var ret ResourceLinksProperties
		return ret
	}).(ResourceLinksPropertiesOutput)
}

// A URI that can be used to retrieve the resource.
func (o ResourceLinksPropertiesPtrOutput) Self() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceLinksProperties) *string {
		if v == nil {
			return nil
		}
		return v.Self
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProjectPropertiesOutput{})
	pulumi.RegisterOutputType(GetProjectsDefaultPropertiesOutput{})
	pulumi.RegisterOutputType(ListProjectsItemsOutput{})
	pulumi.RegisterOutputType(ListProjectsResourcesDefaultItemsOutput{})
	pulumi.RegisterOutputType(ListProjectsResourcesItemsOutput{})
	pulumi.RegisterOutputType(MetaMetaOutput{})
	pulumi.RegisterOutputType(PageLinksOutput{})
	pulumi.RegisterOutputType(PageLinksPtrOutput{})
	pulumi.RegisterOutputType(PageLinksPagesPropertiesOutput{})
	pulumi.RegisterOutputType(PageLinksPagesPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ProjectOutput{})
	pulumi.RegisterOutputType(ProjectPtrOutput{})
	pulumi.RegisterOutputType(ProjectArrayOutput{})
	pulumi.RegisterOutputType(ResourceOutput{})
	pulumi.RegisterOutputType(ResourceArrayOutput{})
	pulumi.RegisterOutputType(ResourceLinksPropertiesOutput{})
	pulumi.RegisterOutputType(ResourceLinksPropertiesPtrOutput{})
}
