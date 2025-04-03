// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProjectsDefault struct {
	pulumi.CustomResourceState

	// A time value given in ISO8601 combined date and time format that represents when the project was created.
	CreatedAt pulumi.StringPtrOutput `pulumi:"createdAt"`
	// The description of the project. The maximum length is 255 characters.
	Description pulumi.StringOutput `pulumi:"description"`
	// The environment of the project's resources.
	Environment ProjectBaseEnvironmentOutput `pulumi:"environment"`
	// If true, all resources will be added to this project if no project is specified.
	IsDefault pulumi.BoolOutput `pulumi:"isDefault"`
	// The human-readable name for the project. The maximum length is 175 characters and the name must be unique.
	Name pulumi.StringOutput `pulumi:"name"`
	// The integer id of the project owner.
	OwnerId pulumi.IntPtrOutput `pulumi:"ownerId"`
	// The unique universal identifier of the project owner.
	OwnerUuid pulumi.StringPtrOutput `pulumi:"ownerUuid"`
	Project   ProjectTypePtrOutput   `pulumi:"project"`
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
	Purpose pulumi.StringOutput `pulumi:"purpose"`
	// A time value given in ISO8601 combined date and time format that represents when the project was updated.
	UpdatedAt pulumi.StringPtrOutput `pulumi:"updatedAt"`
}

// NewProjectsDefault registers a new resource with the given unique name, arguments, and options.
func NewProjectsDefault(ctx *pulumi.Context,
	name string, args *ProjectsDefaultArgs, opts ...pulumi.ResourceOption) (*ProjectsDefault, error) {
	if args == nil {
		args = &ProjectsDefaultArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectsDefault
	err := ctx.RegisterResource("digitalocean-native:projects/v2:ProjectsDefault", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectsDefault gets an existing ProjectsDefault resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectsDefault(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectsDefaultState, opts ...pulumi.ResourceOption) (*ProjectsDefault, error) {
	var resource ProjectsDefault
	err := ctx.ReadResource("digitalocean-native:projects/v2:ProjectsDefault", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectsDefault resources.
type projectsDefaultState struct {
}

type ProjectsDefaultState struct {
}

func (ProjectsDefaultState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectsDefaultState)(nil)).Elem()
}

type projectsDefaultArgs struct {
	// A time value given in ISO8601 combined date and time format that represents when the project was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The description of the project. The maximum length is 255 characters.
	Description *string `pulumi:"description"`
	// The environment of the project's resources.
	Environment *ProjectBaseEnvironment `pulumi:"environment"`
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

// The set of arguments for constructing a ProjectsDefault resource.
type ProjectsDefaultArgs struct {
	// A time value given in ISO8601 combined date and time format that represents when the project was created.
	CreatedAt pulumi.StringPtrInput
	// The description of the project. The maximum length is 255 characters.
	Description pulumi.StringPtrInput
	// The environment of the project's resources.
	Environment ProjectBaseEnvironmentPtrInput
	// If true, all resources will be added to this project if no project is specified.
	IsDefault pulumi.BoolPtrInput
	// The human-readable name for the project. The maximum length is 175 characters and the name must be unique.
	Name pulumi.StringPtrInput
	// The integer id of the project owner.
	OwnerId pulumi.IntPtrInput
	// The unique universal identifier of the project owner.
	OwnerUuid pulumi.StringPtrInput
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
	Purpose pulumi.StringPtrInput
	// A time value given in ISO8601 combined date and time format that represents when the project was updated.
	UpdatedAt pulumi.StringPtrInput
}

func (ProjectsDefaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectsDefaultArgs)(nil)).Elem()
}

type ProjectsDefaultInput interface {
	pulumi.Input

	ToProjectsDefaultOutput() ProjectsDefaultOutput
	ToProjectsDefaultOutputWithContext(ctx context.Context) ProjectsDefaultOutput
}

func (*ProjectsDefault) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectsDefault)(nil)).Elem()
}

func (i *ProjectsDefault) ToProjectsDefaultOutput() ProjectsDefaultOutput {
	return i.ToProjectsDefaultOutputWithContext(context.Background())
}

func (i *ProjectsDefault) ToProjectsDefaultOutputWithContext(ctx context.Context) ProjectsDefaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectsDefaultOutput)
}

type ProjectsDefaultOutput struct{ *pulumi.OutputState }

func (ProjectsDefaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectsDefault)(nil)).Elem()
}

func (o ProjectsDefaultOutput) ToProjectsDefaultOutput() ProjectsDefaultOutput {
	return o
}

func (o ProjectsDefaultOutput) ToProjectsDefaultOutputWithContext(ctx context.Context) ProjectsDefaultOutput {
	return o
}

// A time value given in ISO8601 combined date and time format that represents when the project was created.
func (o ProjectsDefaultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectsDefault) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The description of the project. The maximum length is 255 characters.
func (o ProjectsDefaultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectsDefault) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The environment of the project's resources.
func (o ProjectsDefaultOutput) Environment() ProjectBaseEnvironmentOutput {
	return o.ApplyT(func(v *ProjectsDefault) ProjectBaseEnvironmentOutput { return v.Environment }).(ProjectBaseEnvironmentOutput)
}

// If true, all resources will be added to this project if no project is specified.
func (o ProjectsDefaultOutput) IsDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectsDefault) pulumi.BoolOutput { return v.IsDefault }).(pulumi.BoolOutput)
}

// The human-readable name for the project. The maximum length is 175 characters and the name must be unique.
func (o ProjectsDefaultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectsDefault) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The integer id of the project owner.
func (o ProjectsDefaultOutput) OwnerId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ProjectsDefault) pulumi.IntPtrOutput { return v.OwnerId }).(pulumi.IntPtrOutput)
}

// The unique universal identifier of the project owner.
func (o ProjectsDefaultOutput) OwnerUuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectsDefault) pulumi.StringPtrOutput { return v.OwnerUuid }).(pulumi.StringPtrOutput)
}

func (o ProjectsDefaultOutput) Project() ProjectTypePtrOutput {
	return o.ApplyT(func(v *ProjectsDefault) ProjectTypePtrOutput { return v.Project }).(ProjectTypePtrOutput)
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
func (o ProjectsDefaultOutput) Purpose() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectsDefault) pulumi.StringOutput { return v.Purpose }).(pulumi.StringOutput)
}

// A time value given in ISO8601 combined date and time format that represents when the project was updated.
func (o ProjectsDefaultOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectsDefault) pulumi.StringPtrOutput { return v.UpdatedAt }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectsDefaultInput)(nil)).Elem(), &ProjectsDefault{})
	pulumi.RegisterOutputType(ProjectsDefaultOutput{})
}
