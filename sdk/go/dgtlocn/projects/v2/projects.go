// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Projects struct {
	pulumi.CustomResourceState

	// A time value given in ISO8601 combined date and time format that represents when the project was created.
	CreatedAt pulumi.StringPtrOutput `pulumi:"createdAt"`
	// The description of the project. The maximum length is 255 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The environment of the project's resources.
	Environment ProjectsProjectBaseEnvironmentPtrOutput `pulumi:"environment"`
	// The human-readable name for the project. The maximum length is 175 characters and the name must be unique.
	Name pulumi.StringOutput `pulumi:"name"`
	// The integer id of the project owner.
	OwnerId pulumi.IntPtrOutput `pulumi:"ownerId"`
	// The unique universal identifier of the project owner.
	OwnerUuid pulumi.StringPtrOutput `pulumi:"ownerUuid"`
	Project   ProjectPtrOutput       `pulumi:"project"`
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

// NewProjects registers a new resource with the given unique name, arguments, and options.
func NewProjects(ctx *pulumi.Context,
	name string, args *ProjectsArgs, opts ...pulumi.ResourceOption) (*Projects, error) {
	if args == nil {
		args = &ProjectsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Projects
	err := ctx.RegisterResource("digitalocean-native:projects/v2:Projects", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjects gets an existing Projects resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjects(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectsState, opts ...pulumi.ResourceOption) (*Projects, error) {
	var resource Projects
	err := ctx.ReadResource("digitalocean-native:projects/v2:Projects", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Projects resources.
type projectsState struct {
}

type ProjectsState struct {
}

func (ProjectsState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectsState)(nil)).Elem()
}

type projectsArgs struct {
	// A time value given in ISO8601 combined date and time format that represents when the project was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The description of the project. The maximum length is 255 characters.
	Description *string `pulumi:"description"`
	// The environment of the project's resources.
	Environment *ProjectsProjectBaseEnvironment `pulumi:"environment"`
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

// The set of arguments for constructing a Projects resource.
type ProjectsArgs struct {
	// A time value given in ISO8601 combined date and time format that represents when the project was created.
	CreatedAt pulumi.StringPtrInput
	// The description of the project. The maximum length is 255 characters.
	Description pulumi.StringPtrInput
	// The environment of the project's resources.
	Environment ProjectsProjectBaseEnvironmentPtrInput
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

func (ProjectsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectsArgs)(nil)).Elem()
}

type ProjectsInput interface {
	pulumi.Input

	ToProjectsOutput() ProjectsOutput
	ToProjectsOutputWithContext(ctx context.Context) ProjectsOutput
}

func (*Projects) ElementType() reflect.Type {
	return reflect.TypeOf((**Projects)(nil)).Elem()
}

func (i *Projects) ToProjectsOutput() ProjectsOutput {
	return i.ToProjectsOutputWithContext(context.Background())
}

func (i *Projects) ToProjectsOutputWithContext(ctx context.Context) ProjectsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectsOutput)
}

type ProjectsOutput struct{ *pulumi.OutputState }

func (ProjectsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Projects)(nil)).Elem()
}

func (o ProjectsOutput) ToProjectsOutput() ProjectsOutput {
	return o
}

func (o ProjectsOutput) ToProjectsOutputWithContext(ctx context.Context) ProjectsOutput {
	return o
}

// A time value given in ISO8601 combined date and time format that represents when the project was created.
func (o ProjectsOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Projects) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The description of the project. The maximum length is 255 characters.
func (o ProjectsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Projects) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The environment of the project's resources.
func (o ProjectsOutput) Environment() ProjectsProjectBaseEnvironmentPtrOutput {
	return o.ApplyT(func(v *Projects) ProjectsProjectBaseEnvironmentPtrOutput { return v.Environment }).(ProjectsProjectBaseEnvironmentPtrOutput)
}

// The human-readable name for the project. The maximum length is 175 characters and the name must be unique.
func (o ProjectsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Projects) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The integer id of the project owner.
func (o ProjectsOutput) OwnerId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Projects) pulumi.IntPtrOutput { return v.OwnerId }).(pulumi.IntPtrOutput)
}

// The unique universal identifier of the project owner.
func (o ProjectsOutput) OwnerUuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Projects) pulumi.StringPtrOutput { return v.OwnerUuid }).(pulumi.StringPtrOutput)
}

func (o ProjectsOutput) Project() ProjectPtrOutput {
	return o.ApplyT(func(v *Projects) ProjectPtrOutput { return v.Project }).(ProjectPtrOutput)
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
func (o ProjectsOutput) Purpose() pulumi.StringOutput {
	return o.ApplyT(func(v *Projects) pulumi.StringOutput { return v.Purpose }).(pulumi.StringOutput)
}

// A time value given in ISO8601 combined date and time format that represents when the project was updated.
func (o ProjectsOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Projects) pulumi.StringPtrOutput { return v.UpdatedAt }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectsInput)(nil)).Elem(), &Projects{})
	pulumi.RegisterOutputType(ProjectsOutput{})
}
