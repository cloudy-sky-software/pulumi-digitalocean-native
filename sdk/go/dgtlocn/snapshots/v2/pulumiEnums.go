// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The type of resource that the snapshot originated from.
type SnapshotsPropertiesResourceType string

const (
	SnapshotsPropertiesResourceTypeDroplet = SnapshotsPropertiesResourceType("droplet")
	SnapshotsPropertiesResourceTypeVolume  = SnapshotsPropertiesResourceType("volume")
)

type SnapshotsPropertiesResourceTypeOutput struct{ *pulumi.OutputState }

func (SnapshotsPropertiesResourceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SnapshotsPropertiesResourceType)(nil)).Elem()
}

func (o SnapshotsPropertiesResourceTypeOutput) ToSnapshotsPropertiesResourceTypeOutput() SnapshotsPropertiesResourceTypeOutput {
	return o
}

func (o SnapshotsPropertiesResourceTypeOutput) ToSnapshotsPropertiesResourceTypeOutputWithContext(ctx context.Context) SnapshotsPropertiesResourceTypeOutput {
	return o
}

func (o SnapshotsPropertiesResourceTypeOutput) ToSnapshotsPropertiesResourceTypePtrOutput() SnapshotsPropertiesResourceTypePtrOutput {
	return o.ToSnapshotsPropertiesResourceTypePtrOutputWithContext(context.Background())
}

func (o SnapshotsPropertiesResourceTypeOutput) ToSnapshotsPropertiesResourceTypePtrOutputWithContext(ctx context.Context) SnapshotsPropertiesResourceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SnapshotsPropertiesResourceType) *SnapshotsPropertiesResourceType {
		return &v
	}).(SnapshotsPropertiesResourceTypePtrOutput)
}

func (o SnapshotsPropertiesResourceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SnapshotsPropertiesResourceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SnapshotsPropertiesResourceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SnapshotsPropertiesResourceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SnapshotsPropertiesResourceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SnapshotsPropertiesResourceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SnapshotsPropertiesResourceTypePtrOutput struct{ *pulumi.OutputState }

func (SnapshotsPropertiesResourceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotsPropertiesResourceType)(nil)).Elem()
}

func (o SnapshotsPropertiesResourceTypePtrOutput) ToSnapshotsPropertiesResourceTypePtrOutput() SnapshotsPropertiesResourceTypePtrOutput {
	return o
}

func (o SnapshotsPropertiesResourceTypePtrOutput) ToSnapshotsPropertiesResourceTypePtrOutputWithContext(ctx context.Context) SnapshotsPropertiesResourceTypePtrOutput {
	return o
}

func (o SnapshotsPropertiesResourceTypePtrOutput) Elem() SnapshotsPropertiesResourceTypeOutput {
	return o.ApplyT(func(v *SnapshotsPropertiesResourceType) SnapshotsPropertiesResourceType {
		if v != nil {
			return *v
		}
		var ret SnapshotsPropertiesResourceType
		return ret
	}).(SnapshotsPropertiesResourceTypeOutput)
}

func (o SnapshotsPropertiesResourceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SnapshotsPropertiesResourceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SnapshotsPropertiesResourceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SnapshotsPropertiesResourceTypeOutput{})
	pulumi.RegisterOutputType(SnapshotsPropertiesResourceTypePtrOutput{})
}