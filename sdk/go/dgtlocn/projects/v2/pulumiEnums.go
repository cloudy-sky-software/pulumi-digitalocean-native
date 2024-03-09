// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The environment of the project's resources.
type ProjectBaseEnvironment string

const (
	ProjectBaseEnvironmentDevelopment = ProjectBaseEnvironment("Development")
	ProjectBaseEnvironmentStaging     = ProjectBaseEnvironment("Staging")
	ProjectBaseEnvironmentProduction  = ProjectBaseEnvironment("Production")
)

func (ProjectBaseEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectBaseEnvironment)(nil)).Elem()
}

func (e ProjectBaseEnvironment) ToProjectBaseEnvironmentOutput() ProjectBaseEnvironmentOutput {
	return pulumi.ToOutput(e).(ProjectBaseEnvironmentOutput)
}

func (e ProjectBaseEnvironment) ToProjectBaseEnvironmentOutputWithContext(ctx context.Context) ProjectBaseEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProjectBaseEnvironmentOutput)
}

func (e ProjectBaseEnvironment) ToProjectBaseEnvironmentPtrOutput() ProjectBaseEnvironmentPtrOutput {
	return e.ToProjectBaseEnvironmentPtrOutputWithContext(context.Background())
}

func (e ProjectBaseEnvironment) ToProjectBaseEnvironmentPtrOutputWithContext(ctx context.Context) ProjectBaseEnvironmentPtrOutput {
	return ProjectBaseEnvironment(e).ToProjectBaseEnvironmentOutputWithContext(ctx).ToProjectBaseEnvironmentPtrOutputWithContext(ctx)
}

func (e ProjectBaseEnvironment) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProjectBaseEnvironment) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProjectBaseEnvironment) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProjectBaseEnvironment) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProjectBaseEnvironmentOutput struct{ *pulumi.OutputState }

func (ProjectBaseEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectBaseEnvironment)(nil)).Elem()
}

func (o ProjectBaseEnvironmentOutput) ToProjectBaseEnvironmentOutput() ProjectBaseEnvironmentOutput {
	return o
}

func (o ProjectBaseEnvironmentOutput) ToProjectBaseEnvironmentOutputWithContext(ctx context.Context) ProjectBaseEnvironmentOutput {
	return o
}

func (o ProjectBaseEnvironmentOutput) ToProjectBaseEnvironmentPtrOutput() ProjectBaseEnvironmentPtrOutput {
	return o.ToProjectBaseEnvironmentPtrOutputWithContext(context.Background())
}

func (o ProjectBaseEnvironmentOutput) ToProjectBaseEnvironmentPtrOutputWithContext(ctx context.Context) ProjectBaseEnvironmentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProjectBaseEnvironment) *ProjectBaseEnvironment {
		return &v
	}).(ProjectBaseEnvironmentPtrOutput)
}

func (o ProjectBaseEnvironmentOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProjectBaseEnvironmentOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProjectBaseEnvironment) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProjectBaseEnvironmentOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProjectBaseEnvironmentOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProjectBaseEnvironment) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProjectBaseEnvironmentPtrOutput struct{ *pulumi.OutputState }

func (ProjectBaseEnvironmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectBaseEnvironment)(nil)).Elem()
}

func (o ProjectBaseEnvironmentPtrOutput) ToProjectBaseEnvironmentPtrOutput() ProjectBaseEnvironmentPtrOutput {
	return o
}

func (o ProjectBaseEnvironmentPtrOutput) ToProjectBaseEnvironmentPtrOutputWithContext(ctx context.Context) ProjectBaseEnvironmentPtrOutput {
	return o
}

func (o ProjectBaseEnvironmentPtrOutput) Elem() ProjectBaseEnvironmentOutput {
	return o.ApplyT(func(v *ProjectBaseEnvironment) ProjectBaseEnvironment {
		if v != nil {
			return *v
		}
		var ret ProjectBaseEnvironment
		return ret
	}).(ProjectBaseEnvironmentOutput)
}

func (o ProjectBaseEnvironmentPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProjectBaseEnvironmentPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProjectBaseEnvironment) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ProjectBaseEnvironmentInput is an input type that accepts values of the ProjectBaseEnvironment enum
// A concrete instance of `ProjectBaseEnvironmentInput` can be one of the following:
//
//	ProjectBaseEnvironmentDevelopment
//	ProjectBaseEnvironmentStaging
//	ProjectBaseEnvironmentProduction
type ProjectBaseEnvironmentInput interface {
	pulumi.Input

	ToProjectBaseEnvironmentOutput() ProjectBaseEnvironmentOutput
	ToProjectBaseEnvironmentOutputWithContext(context.Context) ProjectBaseEnvironmentOutput
}

var projectBaseEnvironmentPtrType = reflect.TypeOf((**ProjectBaseEnvironment)(nil)).Elem()

type ProjectBaseEnvironmentPtrInput interface {
	pulumi.Input

	ToProjectBaseEnvironmentPtrOutput() ProjectBaseEnvironmentPtrOutput
	ToProjectBaseEnvironmentPtrOutputWithContext(context.Context) ProjectBaseEnvironmentPtrOutput
}

type projectBaseEnvironmentPtr string

func ProjectBaseEnvironmentPtr(v string) ProjectBaseEnvironmentPtrInput {
	return (*projectBaseEnvironmentPtr)(&v)
}

func (*projectBaseEnvironmentPtr) ElementType() reflect.Type {
	return projectBaseEnvironmentPtrType
}

func (in *projectBaseEnvironmentPtr) ToProjectBaseEnvironmentPtrOutput() ProjectBaseEnvironmentPtrOutput {
	return pulumi.ToOutput(in).(ProjectBaseEnvironmentPtrOutput)
}

func (in *projectBaseEnvironmentPtr) ToProjectBaseEnvironmentPtrOutputWithContext(ctx context.Context) ProjectBaseEnvironmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProjectBaseEnvironmentPtrOutput)
}

// The status of assigning and fetching the resources.
type ResourceStatus string

const (
	ResourceStatusOk              = ResourceStatus("ok")
	ResourceStatusNotFound        = ResourceStatus("not_found")
	ResourceStatusAssigned        = ResourceStatus("assigned")
	ResourceStatusAlreadyAssigned = ResourceStatus("already_assigned")
	ResourceStatusServiceDown     = ResourceStatus("service_down")
)

type ResourceStatusOutput struct{ *pulumi.OutputState }

func (ResourceStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceStatus)(nil)).Elem()
}

func (o ResourceStatusOutput) ToResourceStatusOutput() ResourceStatusOutput {
	return o
}

func (o ResourceStatusOutput) ToResourceStatusOutputWithContext(ctx context.Context) ResourceStatusOutput {
	return o
}

func (o ResourceStatusOutput) ToResourceStatusPtrOutput() ResourceStatusPtrOutput {
	return o.ToResourceStatusPtrOutputWithContext(context.Background())
}

func (o ResourceStatusOutput) ToResourceStatusPtrOutputWithContext(ctx context.Context) ResourceStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceStatus) *ResourceStatus {
		return &v
	}).(ResourceStatusPtrOutput)
}

func (o ResourceStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceStatusPtrOutput struct{ *pulumi.OutputState }

func (ResourceStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceStatus)(nil)).Elem()
}

func (o ResourceStatusPtrOutput) ToResourceStatusPtrOutput() ResourceStatusPtrOutput {
	return o
}

func (o ResourceStatusPtrOutput) ToResourceStatusPtrOutputWithContext(ctx context.Context) ResourceStatusPtrOutput {
	return o
}

func (o ResourceStatusPtrOutput) Elem() ResourceStatusOutput {
	return o.ApplyT(func(v *ResourceStatus) ResourceStatus {
		if v != nil {
			return *v
		}
		var ret ResourceStatus
		return ret
	}).(ResourceStatusOutput)
}

func (o ResourceStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectBaseEnvironmentInput)(nil)).Elem(), ProjectBaseEnvironment("Development"))
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectBaseEnvironmentPtrInput)(nil)).Elem(), ProjectBaseEnvironment("Development"))
	pulumi.RegisterOutputType(ProjectBaseEnvironmentOutput{})
	pulumi.RegisterOutputType(ProjectBaseEnvironmentPtrOutput{})
	pulumi.RegisterOutputType(ResourceStatusOutput{})
	pulumi.RegisterOutputType(ResourceStatusPtrOutput{})
}
