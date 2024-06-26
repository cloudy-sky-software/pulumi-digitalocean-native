// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Project struct {
	pulumi.CustomResourceState

	// A time value given in ISO8601 combined date and time format that represents when the project was created.
	CreatedAt pulumi.StringPtrOutput `pulumi:"createdAt"`
	// The description of the project. The maximum length is 255 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The environment of the project's resources.
	Environment ProjectBaseEnvironmentPtrOutput `pulumi:"environment"`
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

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		args = &ProjectArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Project
	err := ctx.RegisterResource("digitalocean-native:projects/v2:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProject gets an existing Project resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectState, opts ...pulumi.ResourceOption) (*Project, error) {
	var resource Project
	err := ctx.ReadResource("digitalocean-native:projects/v2:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Project resources.
type projectState struct {
}

type ProjectState struct {
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	// A time value given in ISO8601 combined date and time format that represents when the project was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The description of the project. The maximum length is 255 characters.
	Description *string `pulumi:"description"`
	// The environment of the project's resources.
	Environment *ProjectBaseEnvironment `pulumi:"environment"`
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

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// A time value given in ISO8601 combined date and time format that represents when the project was created.
	CreatedAt pulumi.StringPtrInput
	// The description of the project. The maximum length is 255 characters.
	Description pulumi.StringPtrInput
	// The environment of the project's resources.
	Environment ProjectBaseEnvironmentPtrInput
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

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}

type ProjectInput interface {
	pulumi.Input

	ToProjectOutput() ProjectOutput
	ToProjectOutputWithContext(ctx context.Context) ProjectOutput
}

func (*Project) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (i *Project) ToProjectOutput() ProjectOutput {
	return i.ToProjectOutputWithContext(context.Background())
}

func (i *Project) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectOutput)
}

type ProjectOutput struct{ *pulumi.OutputState }

func (ProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (o ProjectOutput) ToProjectOutput() ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return o
}

// A time value given in ISO8601 combined date and time format that represents when the project was created.
func (o ProjectOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The description of the project. The maximum length is 255 characters.
func (o ProjectOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The environment of the project's resources.
func (o ProjectOutput) Environment() ProjectBaseEnvironmentPtrOutput {
	return o.ApplyT(func(v *Project) ProjectBaseEnvironmentPtrOutput { return v.Environment }).(ProjectBaseEnvironmentPtrOutput)
}

// The human-readable name for the project. The maximum length is 175 characters and the name must be unique.
func (o ProjectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The integer id of the project owner.
func (o ProjectOutput) OwnerId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.IntPtrOutput { return v.OwnerId }).(pulumi.IntPtrOutput)
}

// The unique universal identifier of the project owner.
func (o ProjectOutput) OwnerUuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.OwnerUuid }).(pulumi.StringPtrOutput)
}

func (o ProjectOutput) Project() ProjectTypePtrOutput {
	return o.ApplyT(func(v *Project) ProjectTypePtrOutput { return v.Project }).(ProjectTypePtrOutput)
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
func (o ProjectOutput) Purpose() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Purpose }).(pulumi.StringOutput)
}

// A time value given in ISO8601 combined date and time format that represents when the project was updated.
func (o ProjectOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.UpdatedAt }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectInput)(nil)).Elem(), &Project{})
	pulumi.RegisterOutputType(ProjectOutput{})
}
