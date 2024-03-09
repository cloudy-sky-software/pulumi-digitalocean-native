// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProjectsAssignResources struct {
	pulumi.CustomResourceState

	Resources ResourceArrayOutput `pulumi:"resources"`
}

// NewProjectsAssignResources registers a new resource with the given unique name, arguments, and options.
func NewProjectsAssignResources(ctx *pulumi.Context,
	name string, args *ProjectsAssignResourcesArgs, opts ...pulumi.ResourceOption) (*ProjectsAssignResources, error) {
	if args == nil {
		args = &ProjectsAssignResourcesArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectsAssignResources
	err := ctx.RegisterResource("digitalocean-native:projects/v2:ProjectsAssignResources", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectsAssignResources gets an existing ProjectsAssignResources resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectsAssignResources(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectsAssignResourcesState, opts ...pulumi.ResourceOption) (*ProjectsAssignResources, error) {
	var resource ProjectsAssignResources
	err := ctx.ReadResource("digitalocean-native:projects/v2:ProjectsAssignResources", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectsAssignResources resources.
type projectsAssignResourcesState struct {
}

type ProjectsAssignResourcesState struct {
}

func (ProjectsAssignResourcesState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectsAssignResourcesState)(nil)).Elem()
}

type projectsAssignResourcesArgs struct {
	// A unique identifier for a project.
	ProjectId *string `pulumi:"projectId"`
	// A list of uniform resource names (URNs) to be added to a project.
	Resources []string `pulumi:"resources"`
}

// The set of arguments for constructing a ProjectsAssignResources resource.
type ProjectsAssignResourcesArgs struct {
	// A unique identifier for a project.
	ProjectId pulumi.StringPtrInput
	// A list of uniform resource names (URNs) to be added to a project.
	Resources pulumi.StringArrayInput
}

func (ProjectsAssignResourcesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectsAssignResourcesArgs)(nil)).Elem()
}

type ProjectsAssignResourcesInput interface {
	pulumi.Input

	ToProjectsAssignResourcesOutput() ProjectsAssignResourcesOutput
	ToProjectsAssignResourcesOutputWithContext(ctx context.Context) ProjectsAssignResourcesOutput
}

func (*ProjectsAssignResources) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectsAssignResources)(nil)).Elem()
}

func (i *ProjectsAssignResources) ToProjectsAssignResourcesOutput() ProjectsAssignResourcesOutput {
	return i.ToProjectsAssignResourcesOutputWithContext(context.Background())
}

func (i *ProjectsAssignResources) ToProjectsAssignResourcesOutputWithContext(ctx context.Context) ProjectsAssignResourcesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectsAssignResourcesOutput)
}

type ProjectsAssignResourcesOutput struct{ *pulumi.OutputState }

func (ProjectsAssignResourcesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectsAssignResources)(nil)).Elem()
}

func (o ProjectsAssignResourcesOutput) ToProjectsAssignResourcesOutput() ProjectsAssignResourcesOutput {
	return o
}

func (o ProjectsAssignResourcesOutput) ToProjectsAssignResourcesOutputWithContext(ctx context.Context) ProjectsAssignResourcesOutput {
	return o
}

func (o ProjectsAssignResourcesOutput) Resources() ResourceArrayOutput {
	return o.ApplyT(func(v *ProjectsAssignResources) ResourceArrayOutput { return v.Resources }).(ResourceArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectsAssignResourcesInput)(nil)).Elem(), &ProjectsAssignResources{})
	pulumi.RegisterOutputType(ProjectsAssignResourcesOutput{})
}