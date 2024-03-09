// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The current status of this garbage collection.
type GarbageCollectionStatus string

const (
	GarbageCollectionStatusRequested                   = GarbageCollectionStatus("requested")
	GarbageCollectionStatusWaitingForWriteJWTsToExpire = GarbageCollectionStatus("waiting for write JWTs to expire")
	GarbageCollectionStatusScanningManifests           = GarbageCollectionStatus("scanning manifests")
	GarbageCollectionStatusDeletingUnreferencedBlobs   = GarbageCollectionStatus("deleting unreferenced blobs")
	GarbageCollectionStatusCancelling                  = GarbageCollectionStatus("cancelling")
	GarbageCollectionStatusFailed                      = GarbageCollectionStatus("failed")
	GarbageCollectionStatusSucceeded                   = GarbageCollectionStatus("succeeded")
	GarbageCollectionStatusCancelled                   = GarbageCollectionStatus("cancelled")
)

type GarbageCollectionStatusOutput struct{ *pulumi.OutputState }

func (GarbageCollectionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GarbageCollectionStatus)(nil)).Elem()
}

func (o GarbageCollectionStatusOutput) ToGarbageCollectionStatusOutput() GarbageCollectionStatusOutput {
	return o
}

func (o GarbageCollectionStatusOutput) ToGarbageCollectionStatusOutputWithContext(ctx context.Context) GarbageCollectionStatusOutput {
	return o
}

func (o GarbageCollectionStatusOutput) ToGarbageCollectionStatusPtrOutput() GarbageCollectionStatusPtrOutput {
	return o.ToGarbageCollectionStatusPtrOutputWithContext(context.Background())
}

func (o GarbageCollectionStatusOutput) ToGarbageCollectionStatusPtrOutputWithContext(ctx context.Context) GarbageCollectionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GarbageCollectionStatus) *GarbageCollectionStatus {
		return &v
	}).(GarbageCollectionStatusPtrOutput)
}

func (o GarbageCollectionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o GarbageCollectionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GarbageCollectionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o GarbageCollectionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GarbageCollectionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GarbageCollectionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type GarbageCollectionStatusPtrOutput struct{ *pulumi.OutputState }

func (GarbageCollectionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GarbageCollectionStatus)(nil)).Elem()
}

func (o GarbageCollectionStatusPtrOutput) ToGarbageCollectionStatusPtrOutput() GarbageCollectionStatusPtrOutput {
	return o
}

func (o GarbageCollectionStatusPtrOutput) ToGarbageCollectionStatusPtrOutputWithContext(ctx context.Context) GarbageCollectionStatusPtrOutput {
	return o
}

func (o GarbageCollectionStatusPtrOutput) Elem() GarbageCollectionStatusOutput {
	return o.ApplyT(func(v *GarbageCollectionStatus) GarbageCollectionStatus {
		if v != nil {
			return *v
		}
		var ret GarbageCollectionStatus
		return ret
	}).(GarbageCollectionStatusOutput)
}

func (o GarbageCollectionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GarbageCollectionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *GarbageCollectionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// Slug of the region where registry data is stored. When not provided, a region will be selected.
type Region string

const (
	RegionNyc3 = Region("nyc3")
	RegionSfo3 = Region("sfo3")
	RegionAms3 = Region("ams3")
	RegionSgp1 = Region("sgp1")
	RegionFra1 = Region("fra1")
)

func (Region) ElementType() reflect.Type {
	return reflect.TypeOf((*Region)(nil)).Elem()
}

func (e Region) ToRegionOutput() RegionOutput {
	return pulumi.ToOutput(e).(RegionOutput)
}

func (e Region) ToRegionOutputWithContext(ctx context.Context) RegionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RegionOutput)
}

func (e Region) ToRegionPtrOutput() RegionPtrOutput {
	return e.ToRegionPtrOutputWithContext(context.Background())
}

func (e Region) ToRegionPtrOutputWithContext(ctx context.Context) RegionPtrOutput {
	return Region(e).ToRegionOutputWithContext(ctx).ToRegionPtrOutputWithContext(ctx)
}

func (e Region) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Region) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Region) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Region) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RegionOutput struct{ *pulumi.OutputState }

func (RegionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Region)(nil)).Elem()
}

func (o RegionOutput) ToRegionOutput() RegionOutput {
	return o
}

func (o RegionOutput) ToRegionOutputWithContext(ctx context.Context) RegionOutput {
	return o
}

func (o RegionOutput) ToRegionPtrOutput() RegionPtrOutput {
	return o.ToRegionPtrOutputWithContext(context.Background())
}

func (o RegionOutput) ToRegionPtrOutputWithContext(ctx context.Context) RegionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Region) *Region {
		return &v
	}).(RegionPtrOutput)
}

func (o RegionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RegionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Region) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RegionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RegionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Region) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RegionPtrOutput struct{ *pulumi.OutputState }

func (RegionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Region)(nil)).Elem()
}

func (o RegionPtrOutput) ToRegionPtrOutput() RegionPtrOutput {
	return o
}

func (o RegionPtrOutput) ToRegionPtrOutputWithContext(ctx context.Context) RegionPtrOutput {
	return o
}

func (o RegionPtrOutput) Elem() RegionOutput {
	return o.ApplyT(func(v *Region) Region {
		if v != nil {
			return *v
		}
		var ret Region
		return ret
	}).(RegionOutput)
}

func (o RegionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RegionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Region) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// RegionInput is an input type that accepts values of the Region enum
// A concrete instance of `RegionInput` can be one of the following:
//
//	RegionNyc3
//	RegionSfo3
//	RegionAms3
//	RegionSgp1
//	RegionFra1
type RegionInput interface {
	pulumi.Input

	ToRegionOutput() RegionOutput
	ToRegionOutputWithContext(context.Context) RegionOutput
}

var regionPtrType = reflect.TypeOf((**Region)(nil)).Elem()

type RegionPtrInput interface {
	pulumi.Input

	ToRegionPtrOutput() RegionPtrOutput
	ToRegionPtrOutputWithContext(context.Context) RegionPtrOutput
}

type regionPtr string

func RegionPtr(v string) RegionPtrInput {
	return (*regionPtr)(&v)
}

func (*regionPtr) ElementType() reflect.Type {
	return regionPtrType
}

func (in *regionPtr) ToRegionPtrOutput() RegionPtrOutput {
	return pulumi.ToOutput(in).(RegionPtrOutput)
}

func (in *regionPtr) ToRegionPtrOutputWithContext(ctx context.Context) RegionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RegionPtrOutput)
}

type SubscriptionTierExtendedEligibilityReasonsItem string

const (
	SubscriptionTierExtendedEligibilityReasonsItemOverRepositoryLimit = SubscriptionTierExtendedEligibilityReasonsItem("OverRepositoryLimit")
	SubscriptionTierExtendedEligibilityReasonsItemOverStorageLimit    = SubscriptionTierExtendedEligibilityReasonsItem("OverStorageLimit")
)

type SubscriptionTierExtendedEligibilityReasonsItemOutput struct{ *pulumi.OutputState }

func (SubscriptionTierExtendedEligibilityReasonsItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionTierExtendedEligibilityReasonsItem)(nil)).Elem()
}

func (o SubscriptionTierExtendedEligibilityReasonsItemOutput) ToSubscriptionTierExtendedEligibilityReasonsItemOutput() SubscriptionTierExtendedEligibilityReasonsItemOutput {
	return o
}

func (o SubscriptionTierExtendedEligibilityReasonsItemOutput) ToSubscriptionTierExtendedEligibilityReasonsItemOutputWithContext(ctx context.Context) SubscriptionTierExtendedEligibilityReasonsItemOutput {
	return o
}

func (o SubscriptionTierExtendedEligibilityReasonsItemOutput) ToSubscriptionTierExtendedEligibilityReasonsItemPtrOutput() SubscriptionTierExtendedEligibilityReasonsItemPtrOutput {
	return o.ToSubscriptionTierExtendedEligibilityReasonsItemPtrOutputWithContext(context.Background())
}

func (o SubscriptionTierExtendedEligibilityReasonsItemOutput) ToSubscriptionTierExtendedEligibilityReasonsItemPtrOutputWithContext(ctx context.Context) SubscriptionTierExtendedEligibilityReasonsItemPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubscriptionTierExtendedEligibilityReasonsItem) *SubscriptionTierExtendedEligibilityReasonsItem {
		return &v
	}).(SubscriptionTierExtendedEligibilityReasonsItemPtrOutput)
}

func (o SubscriptionTierExtendedEligibilityReasonsItemOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SubscriptionTierExtendedEligibilityReasonsItemOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SubscriptionTierExtendedEligibilityReasonsItem) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SubscriptionTierExtendedEligibilityReasonsItemOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SubscriptionTierExtendedEligibilityReasonsItemOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SubscriptionTierExtendedEligibilityReasonsItem) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SubscriptionTierExtendedEligibilityReasonsItemPtrOutput struct{ *pulumi.OutputState }

func (SubscriptionTierExtendedEligibilityReasonsItemPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionTierExtendedEligibilityReasonsItem)(nil)).Elem()
}

func (o SubscriptionTierExtendedEligibilityReasonsItemPtrOutput) ToSubscriptionTierExtendedEligibilityReasonsItemPtrOutput() SubscriptionTierExtendedEligibilityReasonsItemPtrOutput {
	return o
}

func (o SubscriptionTierExtendedEligibilityReasonsItemPtrOutput) ToSubscriptionTierExtendedEligibilityReasonsItemPtrOutputWithContext(ctx context.Context) SubscriptionTierExtendedEligibilityReasonsItemPtrOutput {
	return o
}

func (o SubscriptionTierExtendedEligibilityReasonsItemPtrOutput) Elem() SubscriptionTierExtendedEligibilityReasonsItemOutput {
	return o.ApplyT(func(v *SubscriptionTierExtendedEligibilityReasonsItem) SubscriptionTierExtendedEligibilityReasonsItem {
		if v != nil {
			return *v
		}
		var ret SubscriptionTierExtendedEligibilityReasonsItem
		return ret
	}).(SubscriptionTierExtendedEligibilityReasonsItemOutput)
}

func (o SubscriptionTierExtendedEligibilityReasonsItemPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SubscriptionTierExtendedEligibilityReasonsItemPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SubscriptionTierExtendedEligibilityReasonsItem) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SubscriptionTierExtendedEligibilityReasonsItemArrayOutput struct{ *pulumi.OutputState }

func (SubscriptionTierExtendedEligibilityReasonsItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubscriptionTierExtendedEligibilityReasonsItem)(nil)).Elem()
}

func (o SubscriptionTierExtendedEligibilityReasonsItemArrayOutput) ToSubscriptionTierExtendedEligibilityReasonsItemArrayOutput() SubscriptionTierExtendedEligibilityReasonsItemArrayOutput {
	return o
}

func (o SubscriptionTierExtendedEligibilityReasonsItemArrayOutput) ToSubscriptionTierExtendedEligibilityReasonsItemArrayOutputWithContext(ctx context.Context) SubscriptionTierExtendedEligibilityReasonsItemArrayOutput {
	return o
}

func (o SubscriptionTierExtendedEligibilityReasonsItemArrayOutput) Index(i pulumi.IntInput) SubscriptionTierExtendedEligibilityReasonsItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubscriptionTierExtendedEligibilityReasonsItem {
		return vs[0].([]SubscriptionTierExtendedEligibilityReasonsItem)[vs[1].(int)]
	}).(SubscriptionTierExtendedEligibilityReasonsItemOutput)
}

// The slug of the subscription tier to sign up for. Valid values can be retrieved using the options endpoint.
type SubscriptionTierSlug string

const (
	SubscriptionTierSlugStarter      = SubscriptionTierSlug("starter")
	SubscriptionTierSlugBasic        = SubscriptionTierSlug("basic")
	SubscriptionTierSlugProfessional = SubscriptionTierSlug("professional")
)

func (SubscriptionTierSlug) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionTierSlug)(nil)).Elem()
}

func (e SubscriptionTierSlug) ToSubscriptionTierSlugOutput() SubscriptionTierSlugOutput {
	return pulumi.ToOutput(e).(SubscriptionTierSlugOutput)
}

func (e SubscriptionTierSlug) ToSubscriptionTierSlugOutputWithContext(ctx context.Context) SubscriptionTierSlugOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SubscriptionTierSlugOutput)
}

func (e SubscriptionTierSlug) ToSubscriptionTierSlugPtrOutput() SubscriptionTierSlugPtrOutput {
	return e.ToSubscriptionTierSlugPtrOutputWithContext(context.Background())
}

func (e SubscriptionTierSlug) ToSubscriptionTierSlugPtrOutputWithContext(ctx context.Context) SubscriptionTierSlugPtrOutput {
	return SubscriptionTierSlug(e).ToSubscriptionTierSlugOutputWithContext(ctx).ToSubscriptionTierSlugPtrOutputWithContext(ctx)
}

func (e SubscriptionTierSlug) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SubscriptionTierSlug) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SubscriptionTierSlug) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SubscriptionTierSlug) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SubscriptionTierSlugOutput struct{ *pulumi.OutputState }

func (SubscriptionTierSlugOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionTierSlug)(nil)).Elem()
}

func (o SubscriptionTierSlugOutput) ToSubscriptionTierSlugOutput() SubscriptionTierSlugOutput {
	return o
}

func (o SubscriptionTierSlugOutput) ToSubscriptionTierSlugOutputWithContext(ctx context.Context) SubscriptionTierSlugOutput {
	return o
}

func (o SubscriptionTierSlugOutput) ToSubscriptionTierSlugPtrOutput() SubscriptionTierSlugPtrOutput {
	return o.ToSubscriptionTierSlugPtrOutputWithContext(context.Background())
}

func (o SubscriptionTierSlugOutput) ToSubscriptionTierSlugPtrOutputWithContext(ctx context.Context) SubscriptionTierSlugPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubscriptionTierSlug) *SubscriptionTierSlug {
		return &v
	}).(SubscriptionTierSlugPtrOutput)
}

func (o SubscriptionTierSlugOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SubscriptionTierSlugOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SubscriptionTierSlug) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SubscriptionTierSlugOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SubscriptionTierSlugOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SubscriptionTierSlug) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SubscriptionTierSlugPtrOutput struct{ *pulumi.OutputState }

func (SubscriptionTierSlugPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionTierSlug)(nil)).Elem()
}

func (o SubscriptionTierSlugPtrOutput) ToSubscriptionTierSlugPtrOutput() SubscriptionTierSlugPtrOutput {
	return o
}

func (o SubscriptionTierSlugPtrOutput) ToSubscriptionTierSlugPtrOutputWithContext(ctx context.Context) SubscriptionTierSlugPtrOutput {
	return o
}

func (o SubscriptionTierSlugPtrOutput) Elem() SubscriptionTierSlugOutput {
	return o.ApplyT(func(v *SubscriptionTierSlug) SubscriptionTierSlug {
		if v != nil {
			return *v
		}
		var ret SubscriptionTierSlug
		return ret
	}).(SubscriptionTierSlugOutput)
}

func (o SubscriptionTierSlugPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SubscriptionTierSlugPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SubscriptionTierSlug) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SubscriptionTierSlugInput is an input type that accepts values of the SubscriptionTierSlug enum
// A concrete instance of `SubscriptionTierSlugInput` can be one of the following:
//
//	SubscriptionTierSlugStarter
//	SubscriptionTierSlugBasic
//	SubscriptionTierSlugProfessional
type SubscriptionTierSlugInput interface {
	pulumi.Input

	ToSubscriptionTierSlugOutput() SubscriptionTierSlugOutput
	ToSubscriptionTierSlugOutputWithContext(context.Context) SubscriptionTierSlugOutput
}

var subscriptionTierSlugPtrType = reflect.TypeOf((**SubscriptionTierSlug)(nil)).Elem()

type SubscriptionTierSlugPtrInput interface {
	pulumi.Input

	ToSubscriptionTierSlugPtrOutput() SubscriptionTierSlugPtrOutput
	ToSubscriptionTierSlugPtrOutputWithContext(context.Context) SubscriptionTierSlugPtrOutput
}

type subscriptionTierSlugPtr string

func SubscriptionTierSlugPtr(v string) SubscriptionTierSlugPtrInput {
	return (*subscriptionTierSlugPtr)(&v)
}

func (*subscriptionTierSlugPtr) ElementType() reflect.Type {
	return subscriptionTierSlugPtrType
}

func (in *subscriptionTierSlugPtr) ToSubscriptionTierSlugPtrOutput() SubscriptionTierSlugPtrOutput {
	return pulumi.ToOutput(in).(SubscriptionTierSlugPtrOutput)
}

func (in *subscriptionTierSlugPtr) ToSubscriptionTierSlugPtrOutputWithContext(ctx context.Context) SubscriptionTierSlugPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SubscriptionTierSlugPtrOutput)
}

// The slug of the subscription tier to sign up for.
type TierSlug string

const (
	TierSlugStarter      = TierSlug("starter")
	TierSlugBasic        = TierSlug("basic")
	TierSlugProfessional = TierSlug("professional")
)

func (TierSlug) ElementType() reflect.Type {
	return reflect.TypeOf((*TierSlug)(nil)).Elem()
}

func (e TierSlug) ToTierSlugOutput() TierSlugOutput {
	return pulumi.ToOutput(e).(TierSlugOutput)
}

func (e TierSlug) ToTierSlugOutputWithContext(ctx context.Context) TierSlugOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TierSlugOutput)
}

func (e TierSlug) ToTierSlugPtrOutput() TierSlugPtrOutput {
	return e.ToTierSlugPtrOutputWithContext(context.Background())
}

func (e TierSlug) ToTierSlugPtrOutputWithContext(ctx context.Context) TierSlugPtrOutput {
	return TierSlug(e).ToTierSlugOutputWithContext(ctx).ToTierSlugPtrOutputWithContext(ctx)
}

func (e TierSlug) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TierSlug) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TierSlug) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TierSlug) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TierSlugOutput struct{ *pulumi.OutputState }

func (TierSlugOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TierSlug)(nil)).Elem()
}

func (o TierSlugOutput) ToTierSlugOutput() TierSlugOutput {
	return o
}

func (o TierSlugOutput) ToTierSlugOutputWithContext(ctx context.Context) TierSlugOutput {
	return o
}

func (o TierSlugOutput) ToTierSlugPtrOutput() TierSlugPtrOutput {
	return o.ToTierSlugPtrOutputWithContext(context.Background())
}

func (o TierSlugOutput) ToTierSlugPtrOutputWithContext(ctx context.Context) TierSlugPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TierSlug) *TierSlug {
		return &v
	}).(TierSlugPtrOutput)
}

func (o TierSlugOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TierSlugOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TierSlug) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TierSlugOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TierSlugOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TierSlug) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TierSlugPtrOutput struct{ *pulumi.OutputState }

func (TierSlugPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TierSlug)(nil)).Elem()
}

func (o TierSlugPtrOutput) ToTierSlugPtrOutput() TierSlugPtrOutput {
	return o
}

func (o TierSlugPtrOutput) ToTierSlugPtrOutputWithContext(ctx context.Context) TierSlugPtrOutput {
	return o
}

func (o TierSlugPtrOutput) Elem() TierSlugOutput {
	return o.ApplyT(func(v *TierSlug) TierSlug {
		if v != nil {
			return *v
		}
		var ret TierSlug
		return ret
	}).(TierSlugOutput)
}

func (o TierSlugPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TierSlugPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TierSlug) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// TierSlugInput is an input type that accepts values of the TierSlug enum
// A concrete instance of `TierSlugInput` can be one of the following:
//
//	TierSlugStarter
//	TierSlugBasic
//	TierSlugProfessional
type TierSlugInput interface {
	pulumi.Input

	ToTierSlugOutput() TierSlugOutput
	ToTierSlugOutputWithContext(context.Context) TierSlugOutput
}

var tierSlugPtrType = reflect.TypeOf((**TierSlug)(nil)).Elem()

type TierSlugPtrInput interface {
	pulumi.Input

	ToTierSlugPtrOutput() TierSlugPtrOutput
	ToTierSlugPtrOutputWithContext(context.Context) TierSlugPtrOutput
}

type tierSlugPtr string

func TierSlugPtr(v string) TierSlugPtrInput {
	return (*tierSlugPtr)(&v)
}

func (*tierSlugPtr) ElementType() reflect.Type {
	return tierSlugPtrType
}

func (in *tierSlugPtr) ToTierSlugPtrOutput() TierSlugPtrOutput {
	return pulumi.ToOutput(in).(TierSlugPtrOutput)
}

func (in *tierSlugPtr) ToTierSlugPtrOutputWithContext(ctx context.Context) TierSlugPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TierSlugPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegionInput)(nil)).Elem(), Region("nyc3"))
	pulumi.RegisterInputType(reflect.TypeOf((*RegionPtrInput)(nil)).Elem(), Region("nyc3"))
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriptionTierSlugInput)(nil)).Elem(), SubscriptionTierSlug("starter"))
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriptionTierSlugPtrInput)(nil)).Elem(), SubscriptionTierSlug("starter"))
	pulumi.RegisterInputType(reflect.TypeOf((*TierSlugInput)(nil)).Elem(), TierSlug("starter"))
	pulumi.RegisterInputType(reflect.TypeOf((*TierSlugPtrInput)(nil)).Elem(), TierSlug("starter"))
	pulumi.RegisterOutputType(GarbageCollectionStatusOutput{})
	pulumi.RegisterOutputType(GarbageCollectionStatusPtrOutput{})
	pulumi.RegisterOutputType(RegionOutput{})
	pulumi.RegisterOutputType(RegionPtrOutput{})
	pulumi.RegisterOutputType(SubscriptionTierExtendedEligibilityReasonsItemOutput{})
	pulumi.RegisterOutputType(SubscriptionTierExtendedEligibilityReasonsItemPtrOutput{})
	pulumi.RegisterOutputType(SubscriptionTierExtendedEligibilityReasonsItemArrayOutput{})
	pulumi.RegisterOutputType(SubscriptionTierSlugOutput{})
	pulumi.RegisterOutputType(SubscriptionTierSlugPtrOutput{})
	pulumi.RegisterOutputType(TierSlugOutput{})
	pulumi.RegisterOutputType(TierSlugPtrOutput{})
}
