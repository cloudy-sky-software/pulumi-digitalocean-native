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

type GetFunctionsNamespaceProperties struct {
	Namespace *NamespaceInfo `pulumi:"namespace"`
}

type GetFunctionsNamespacePropertiesOutput struct{ *pulumi.OutputState }

func (GetFunctionsNamespacePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFunctionsNamespaceProperties)(nil)).Elem()
}

func (o GetFunctionsNamespacePropertiesOutput) ToGetFunctionsNamespacePropertiesOutput() GetFunctionsNamespacePropertiesOutput {
	return o
}

func (o GetFunctionsNamespacePropertiesOutput) ToGetFunctionsNamespacePropertiesOutputWithContext(ctx context.Context) GetFunctionsNamespacePropertiesOutput {
	return o
}

func (o GetFunctionsNamespacePropertiesOutput) Namespace() NamespaceInfoPtrOutput {
	return o.ApplyT(func(v GetFunctionsNamespaceProperties) *NamespaceInfo { return v.Namespace }).(NamespaceInfoPtrOutput)
}

type GetFunctionsTriggerProperties struct {
	Trigger *TriggerInfo `pulumi:"trigger"`
}

type GetFunctionsTriggerPropertiesOutput struct{ *pulumi.OutputState }

func (GetFunctionsTriggerPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFunctionsTriggerProperties)(nil)).Elem()
}

func (o GetFunctionsTriggerPropertiesOutput) ToGetFunctionsTriggerPropertiesOutput() GetFunctionsTriggerPropertiesOutput {
	return o
}

func (o GetFunctionsTriggerPropertiesOutput) ToGetFunctionsTriggerPropertiesOutputWithContext(ctx context.Context) GetFunctionsTriggerPropertiesOutput {
	return o
}

func (o GetFunctionsTriggerPropertiesOutput) Trigger() TriggerInfoPtrOutput {
	return o.ApplyT(func(v GetFunctionsTriggerProperties) *TriggerInfo { return v.Trigger }).(TriggerInfoPtrOutput)
}

type ListFunctionsNamespaces struct {
	Namespaces []NamespaceInfo `pulumi:"namespaces"`
}

type ListFunctionsNamespacesOutput struct{ *pulumi.OutputState }

func (ListFunctionsNamespacesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListFunctionsNamespaces)(nil)).Elem()
}

func (o ListFunctionsNamespacesOutput) ToListFunctionsNamespacesOutput() ListFunctionsNamespacesOutput {
	return o
}

func (o ListFunctionsNamespacesOutput) ToListFunctionsNamespacesOutputWithContext(ctx context.Context) ListFunctionsNamespacesOutput {
	return o
}

func (o ListFunctionsNamespacesOutput) Namespaces() NamespaceInfoArrayOutput {
	return o.ApplyT(func(v ListFunctionsNamespaces) []NamespaceInfo { return v.Namespaces }).(NamespaceInfoArrayOutput)
}

type ListFunctionsTriggers struct {
	Triggers []TriggerInfo `pulumi:"triggers"`
}

type ListFunctionsTriggersOutput struct{ *pulumi.OutputState }

func (ListFunctionsTriggersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListFunctionsTriggers)(nil)).Elem()
}

func (o ListFunctionsTriggersOutput) ToListFunctionsTriggersOutput() ListFunctionsTriggersOutput {
	return o
}

func (o ListFunctionsTriggersOutput) ToListFunctionsTriggersOutputWithContext(ctx context.Context) ListFunctionsTriggersOutput {
	return o
}

func (o ListFunctionsTriggersOutput) Triggers() TriggerInfoArrayOutput {
	return o.ApplyT(func(v ListFunctionsTriggers) []TriggerInfo { return v.Triggers }).(TriggerInfoArrayOutput)
}

type NamespaceInfo struct {
	// The namespace's API hostname. Each function in a namespace is provided an endpoint at the namespace's hostname.
	ApiHost *string `pulumi:"apiHost"`
	// UTC time string.
	CreatedAt *string `pulumi:"createdAt"`
	// A random alpha numeric string. This key is used in conjunction with the namespace's UUID to authenticate
	// a user to use the namespace via `doctl`, DigitalOcean's official CLI.
	Key *string `pulumi:"key"`
	// The namespace's unique name.
	Label *string `pulumi:"label"`
	// A unique string format of UUID with a prefix fn-.
	Namespace *string `pulumi:"namespace"`
	// The namespace's datacenter region.
	Region *string `pulumi:"region"`
	// UTC time string.
	UpdatedAt *string `pulumi:"updatedAt"`
	// The namespace's Universally Unique Identifier.
	Uuid *string `pulumi:"uuid"`
}

type NamespaceInfoOutput struct{ *pulumi.OutputState }

func (NamespaceInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceInfo)(nil)).Elem()
}

func (o NamespaceInfoOutput) ToNamespaceInfoOutput() NamespaceInfoOutput {
	return o
}

func (o NamespaceInfoOutput) ToNamespaceInfoOutputWithContext(ctx context.Context) NamespaceInfoOutput {
	return o
}

// The namespace's API hostname. Each function in a namespace is provided an endpoint at the namespace's hostname.
func (o NamespaceInfoOutput) ApiHost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceInfo) *string { return v.ApiHost }).(pulumi.StringPtrOutput)
}

// UTC time string.
func (o NamespaceInfoOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceInfo) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// A random alpha numeric string. This key is used in conjunction with the namespace's UUID to authenticate
// a user to use the namespace via `doctl`, DigitalOcean's official CLI.
func (o NamespaceInfoOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceInfo) *string { return v.Key }).(pulumi.StringPtrOutput)
}

// The namespace's unique name.
func (o NamespaceInfoOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceInfo) *string { return v.Label }).(pulumi.StringPtrOutput)
}

// A unique string format of UUID with a prefix fn-.
func (o NamespaceInfoOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceInfo) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

// The namespace's datacenter region.
func (o NamespaceInfoOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceInfo) *string { return v.Region }).(pulumi.StringPtrOutput)
}

// UTC time string.
func (o NamespaceInfoOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceInfo) *string { return v.UpdatedAt }).(pulumi.StringPtrOutput)
}

// The namespace's Universally Unique Identifier.
func (o NamespaceInfoOutput) Uuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceInfo) *string { return v.Uuid }).(pulumi.StringPtrOutput)
}

type NamespaceInfoPtrOutput struct{ *pulumi.OutputState }

func (NamespaceInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceInfo)(nil)).Elem()
}

func (o NamespaceInfoPtrOutput) ToNamespaceInfoPtrOutput() NamespaceInfoPtrOutput {
	return o
}

func (o NamespaceInfoPtrOutput) ToNamespaceInfoPtrOutputWithContext(ctx context.Context) NamespaceInfoPtrOutput {
	return o
}

func (o NamespaceInfoPtrOutput) Elem() NamespaceInfoOutput {
	return o.ApplyT(func(v *NamespaceInfo) NamespaceInfo {
		if v != nil {
			return *v
		}
		var ret NamespaceInfo
		return ret
	}).(NamespaceInfoOutput)
}

// The namespace's API hostname. Each function in a namespace is provided an endpoint at the namespace's hostname.
func (o NamespaceInfoPtrOutput) ApiHost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceInfo) *string {
		if v == nil {
			return nil
		}
		return v.ApiHost
	}).(pulumi.StringPtrOutput)
}

// UTC time string.
func (o NamespaceInfoPtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceInfo) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

// A random alpha numeric string. This key is used in conjunction with the namespace's UUID to authenticate
// a user to use the namespace via `doctl`, DigitalOcean's official CLI.
func (o NamespaceInfoPtrOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceInfo) *string {
		if v == nil {
			return nil
		}
		return v.Key
	}).(pulumi.StringPtrOutput)
}

// The namespace's unique name.
func (o NamespaceInfoPtrOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceInfo) *string {
		if v == nil {
			return nil
		}
		return v.Label
	}).(pulumi.StringPtrOutput)
}

// A unique string format of UUID with a prefix fn-.
func (o NamespaceInfoPtrOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceInfo) *string {
		if v == nil {
			return nil
		}
		return v.Namespace
	}).(pulumi.StringPtrOutput)
}

// The namespace's datacenter region.
func (o NamespaceInfoPtrOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceInfo) *string {
		if v == nil {
			return nil
		}
		return v.Region
	}).(pulumi.StringPtrOutput)
}

// UTC time string.
func (o NamespaceInfoPtrOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceInfo) *string {
		if v == nil {
			return nil
		}
		return v.UpdatedAt
	}).(pulumi.StringPtrOutput)
}

// The namespace's Universally Unique Identifier.
func (o NamespaceInfoPtrOutput) Uuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceInfo) *string {
		if v == nil {
			return nil
		}
		return v.Uuid
	}).(pulumi.StringPtrOutput)
}

type NamespaceInfoArrayOutput struct{ *pulumi.OutputState }

func (NamespaceInfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NamespaceInfo)(nil)).Elem()
}

func (o NamespaceInfoArrayOutput) ToNamespaceInfoArrayOutput() NamespaceInfoArrayOutput {
	return o
}

func (o NamespaceInfoArrayOutput) ToNamespaceInfoArrayOutputWithContext(ctx context.Context) NamespaceInfoArrayOutput {
	return o
}

func (o NamespaceInfoArrayOutput) Index(i pulumi.IntInput) NamespaceInfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NamespaceInfo {
		return vs[0].([]NamespaceInfo)[vs[1].(int)]
	}).(NamespaceInfoOutput)
}

// Trigger details for SCHEDULED type, where body is optional.
type ScheduledDetails struct {
	// Optional data to be sent to function while triggering the function.
	Body *ScheduledDetailsBodyProperties `pulumi:"body"`
	// valid cron expression string which is required for SCHEDULED type triggers.
	Cron string `pulumi:"cron"`
}

// ScheduledDetailsInput is an input type that accepts ScheduledDetailsArgs and ScheduledDetailsOutput values.
// You can construct a concrete instance of `ScheduledDetailsInput` via:
//
//	ScheduledDetailsArgs{...}
type ScheduledDetailsInput interface {
	pulumi.Input

	ToScheduledDetailsOutput() ScheduledDetailsOutput
	ToScheduledDetailsOutputWithContext(context.Context) ScheduledDetailsOutput
}

// Trigger details for SCHEDULED type, where body is optional.
type ScheduledDetailsArgs struct {
	// Optional data to be sent to function while triggering the function.
	Body ScheduledDetailsBodyPropertiesPtrInput `pulumi:"body"`
	// valid cron expression string which is required for SCHEDULED type triggers.
	Cron pulumi.StringInput `pulumi:"cron"`
}

func (ScheduledDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledDetails)(nil)).Elem()
}

func (i ScheduledDetailsArgs) ToScheduledDetailsOutput() ScheduledDetailsOutput {
	return i.ToScheduledDetailsOutputWithContext(context.Background())
}

func (i ScheduledDetailsArgs) ToScheduledDetailsOutputWithContext(ctx context.Context) ScheduledDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledDetailsOutput)
}

// Trigger details for SCHEDULED type, where body is optional.
type ScheduledDetailsOutput struct{ *pulumi.OutputState }

func (ScheduledDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledDetails)(nil)).Elem()
}

func (o ScheduledDetailsOutput) ToScheduledDetailsOutput() ScheduledDetailsOutput {
	return o
}

func (o ScheduledDetailsOutput) ToScheduledDetailsOutputWithContext(ctx context.Context) ScheduledDetailsOutput {
	return o
}

// Optional data to be sent to function while triggering the function.
func (o ScheduledDetailsOutput) Body() ScheduledDetailsBodyPropertiesPtrOutput {
	return o.ApplyT(func(v ScheduledDetails) *ScheduledDetailsBodyProperties { return v.Body }).(ScheduledDetailsBodyPropertiesPtrOutput)
}

// valid cron expression string which is required for SCHEDULED type triggers.
func (o ScheduledDetailsOutput) Cron() pulumi.StringOutput {
	return o.ApplyT(func(v ScheduledDetails) string { return v.Cron }).(pulumi.StringOutput)
}

type ScheduledDetailsPtrOutput struct{ *pulumi.OutputState }

func (ScheduledDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledDetails)(nil)).Elem()
}

func (o ScheduledDetailsPtrOutput) ToScheduledDetailsPtrOutput() ScheduledDetailsPtrOutput {
	return o
}

func (o ScheduledDetailsPtrOutput) ToScheduledDetailsPtrOutputWithContext(ctx context.Context) ScheduledDetailsPtrOutput {
	return o
}

func (o ScheduledDetailsPtrOutput) Elem() ScheduledDetailsOutput {
	return o.ApplyT(func(v *ScheduledDetails) ScheduledDetails {
		if v != nil {
			return *v
		}
		var ret ScheduledDetails
		return ret
	}).(ScheduledDetailsOutput)
}

// Optional data to be sent to function while triggering the function.
func (o ScheduledDetailsPtrOutput) Body() ScheduledDetailsBodyPropertiesPtrOutput {
	return o.ApplyT(func(v *ScheduledDetails) *ScheduledDetailsBodyProperties {
		if v == nil {
			return nil
		}
		return v.Body
	}).(ScheduledDetailsBodyPropertiesPtrOutput)
}

// valid cron expression string which is required for SCHEDULED type triggers.
func (o ScheduledDetailsPtrOutput) Cron() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledDetails) *string {
		if v == nil {
			return nil
		}
		return &v.Cron
	}).(pulumi.StringPtrOutput)
}

// Optional data to be sent to function while triggering the function.
type ScheduledDetailsBodyProperties struct {
	Name *string `pulumi:"name"`
}

// ScheduledDetailsBodyPropertiesInput is an input type that accepts ScheduledDetailsBodyPropertiesArgs and ScheduledDetailsBodyPropertiesOutput values.
// You can construct a concrete instance of `ScheduledDetailsBodyPropertiesInput` via:
//
//	ScheduledDetailsBodyPropertiesArgs{...}
type ScheduledDetailsBodyPropertiesInput interface {
	pulumi.Input

	ToScheduledDetailsBodyPropertiesOutput() ScheduledDetailsBodyPropertiesOutput
	ToScheduledDetailsBodyPropertiesOutputWithContext(context.Context) ScheduledDetailsBodyPropertiesOutput
}

// Optional data to be sent to function while triggering the function.
type ScheduledDetailsBodyPropertiesArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (ScheduledDetailsBodyPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledDetailsBodyProperties)(nil)).Elem()
}

func (i ScheduledDetailsBodyPropertiesArgs) ToScheduledDetailsBodyPropertiesOutput() ScheduledDetailsBodyPropertiesOutput {
	return i.ToScheduledDetailsBodyPropertiesOutputWithContext(context.Background())
}

func (i ScheduledDetailsBodyPropertiesArgs) ToScheduledDetailsBodyPropertiesOutputWithContext(ctx context.Context) ScheduledDetailsBodyPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledDetailsBodyPropertiesOutput)
}

func (i ScheduledDetailsBodyPropertiesArgs) ToScheduledDetailsBodyPropertiesPtrOutput() ScheduledDetailsBodyPropertiesPtrOutput {
	return i.ToScheduledDetailsBodyPropertiesPtrOutputWithContext(context.Background())
}

func (i ScheduledDetailsBodyPropertiesArgs) ToScheduledDetailsBodyPropertiesPtrOutputWithContext(ctx context.Context) ScheduledDetailsBodyPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledDetailsBodyPropertiesOutput).ToScheduledDetailsBodyPropertiesPtrOutputWithContext(ctx)
}

// ScheduledDetailsBodyPropertiesPtrInput is an input type that accepts ScheduledDetailsBodyPropertiesArgs, ScheduledDetailsBodyPropertiesPtr and ScheduledDetailsBodyPropertiesPtrOutput values.
// You can construct a concrete instance of `ScheduledDetailsBodyPropertiesPtrInput` via:
//
//	        ScheduledDetailsBodyPropertiesArgs{...}
//
//	or:
//
//	        nil
type ScheduledDetailsBodyPropertiesPtrInput interface {
	pulumi.Input

	ToScheduledDetailsBodyPropertiesPtrOutput() ScheduledDetailsBodyPropertiesPtrOutput
	ToScheduledDetailsBodyPropertiesPtrOutputWithContext(context.Context) ScheduledDetailsBodyPropertiesPtrOutput
}

type scheduledDetailsBodyPropertiesPtrType ScheduledDetailsBodyPropertiesArgs

func ScheduledDetailsBodyPropertiesPtr(v *ScheduledDetailsBodyPropertiesArgs) ScheduledDetailsBodyPropertiesPtrInput {
	return (*scheduledDetailsBodyPropertiesPtrType)(v)
}

func (*scheduledDetailsBodyPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledDetailsBodyProperties)(nil)).Elem()
}

func (i *scheduledDetailsBodyPropertiesPtrType) ToScheduledDetailsBodyPropertiesPtrOutput() ScheduledDetailsBodyPropertiesPtrOutput {
	return i.ToScheduledDetailsBodyPropertiesPtrOutputWithContext(context.Background())
}

func (i *scheduledDetailsBodyPropertiesPtrType) ToScheduledDetailsBodyPropertiesPtrOutputWithContext(ctx context.Context) ScheduledDetailsBodyPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledDetailsBodyPropertiesPtrOutput)
}

// Optional data to be sent to function while triggering the function.
type ScheduledDetailsBodyPropertiesOutput struct{ *pulumi.OutputState }

func (ScheduledDetailsBodyPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledDetailsBodyProperties)(nil)).Elem()
}

func (o ScheduledDetailsBodyPropertiesOutput) ToScheduledDetailsBodyPropertiesOutput() ScheduledDetailsBodyPropertiesOutput {
	return o
}

func (o ScheduledDetailsBodyPropertiesOutput) ToScheduledDetailsBodyPropertiesOutputWithContext(ctx context.Context) ScheduledDetailsBodyPropertiesOutput {
	return o
}

func (o ScheduledDetailsBodyPropertiesOutput) ToScheduledDetailsBodyPropertiesPtrOutput() ScheduledDetailsBodyPropertiesPtrOutput {
	return o.ToScheduledDetailsBodyPropertiesPtrOutputWithContext(context.Background())
}

func (o ScheduledDetailsBodyPropertiesOutput) ToScheduledDetailsBodyPropertiesPtrOutputWithContext(ctx context.Context) ScheduledDetailsBodyPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScheduledDetailsBodyProperties) *ScheduledDetailsBodyProperties {
		return &v
	}).(ScheduledDetailsBodyPropertiesPtrOutput)
}

func (o ScheduledDetailsBodyPropertiesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduledDetailsBodyProperties) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ScheduledDetailsBodyPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ScheduledDetailsBodyPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledDetailsBodyProperties)(nil)).Elem()
}

func (o ScheduledDetailsBodyPropertiesPtrOutput) ToScheduledDetailsBodyPropertiesPtrOutput() ScheduledDetailsBodyPropertiesPtrOutput {
	return o
}

func (o ScheduledDetailsBodyPropertiesPtrOutput) ToScheduledDetailsBodyPropertiesPtrOutputWithContext(ctx context.Context) ScheduledDetailsBodyPropertiesPtrOutput {
	return o
}

func (o ScheduledDetailsBodyPropertiesPtrOutput) Elem() ScheduledDetailsBodyPropertiesOutput {
	return o.ApplyT(func(v *ScheduledDetailsBodyProperties) ScheduledDetailsBodyProperties {
		if v != nil {
			return *v
		}
		var ret ScheduledDetailsBodyProperties
		return ret
	}).(ScheduledDetailsBodyPropertiesOutput)
}

func (o ScheduledDetailsBodyPropertiesPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledDetailsBodyProperties) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type TriggerInfo struct {
	// UTC time string.
	CreatedAt *string `pulumi:"createdAt"`
	// Name of function(action) that exists in the given namespace.
	Function *string `pulumi:"function"`
	// Indicates weather the trigger is paused or unpaused.
	IsEnabled *bool `pulumi:"isEnabled"`
	// The trigger's unique name within the namespace.
	Name *string `pulumi:"name"`
	// A unique string format of UUID with a prefix fn-.
	Namespace *string `pulumi:"namespace"`
	// Trigger details for SCHEDULED type, where body is optional.
	ScheduledDetails *ScheduledDetails                   `pulumi:"scheduledDetails"`
	ScheduledRuns    *TriggerInfoScheduledRunsProperties `pulumi:"scheduledRuns"`
	// String which indicates the type of trigger source like SCHEDULED.
	Type *string `pulumi:"type"`
	// UTC time string.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type TriggerInfoOutput struct{ *pulumi.OutputState }

func (TriggerInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerInfo)(nil)).Elem()
}

func (o TriggerInfoOutput) ToTriggerInfoOutput() TriggerInfoOutput {
	return o
}

func (o TriggerInfoOutput) ToTriggerInfoOutputWithContext(ctx context.Context) TriggerInfoOutput {
	return o
}

// UTC time string.
func (o TriggerInfoOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TriggerInfo) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// Name of function(action) that exists in the given namespace.
func (o TriggerInfoOutput) Function() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TriggerInfo) *string { return v.Function }).(pulumi.StringPtrOutput)
}

// Indicates weather the trigger is paused or unpaused.
func (o TriggerInfoOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v TriggerInfo) *bool { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

// The trigger's unique name within the namespace.
func (o TriggerInfoOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TriggerInfo) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A unique string format of UUID with a prefix fn-.
func (o TriggerInfoOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TriggerInfo) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Trigger details for SCHEDULED type, where body is optional.
func (o TriggerInfoOutput) ScheduledDetails() ScheduledDetailsPtrOutput {
	return o.ApplyT(func(v TriggerInfo) *ScheduledDetails { return v.ScheduledDetails }).(ScheduledDetailsPtrOutput)
}

func (o TriggerInfoOutput) ScheduledRuns() TriggerInfoScheduledRunsPropertiesPtrOutput {
	return o.ApplyT(func(v TriggerInfo) *TriggerInfoScheduledRunsProperties { return v.ScheduledRuns }).(TriggerInfoScheduledRunsPropertiesPtrOutput)
}

// String which indicates the type of trigger source like SCHEDULED.
func (o TriggerInfoOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TriggerInfo) *string { return v.Type }).(pulumi.StringPtrOutput)
}

// UTC time string.
func (o TriggerInfoOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TriggerInfo) *string { return v.UpdatedAt }).(pulumi.StringPtrOutput)
}

type TriggerInfoPtrOutput struct{ *pulumi.OutputState }

func (TriggerInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerInfo)(nil)).Elem()
}

func (o TriggerInfoPtrOutput) ToTriggerInfoPtrOutput() TriggerInfoPtrOutput {
	return o
}

func (o TriggerInfoPtrOutput) ToTriggerInfoPtrOutputWithContext(ctx context.Context) TriggerInfoPtrOutput {
	return o
}

func (o TriggerInfoPtrOutput) Elem() TriggerInfoOutput {
	return o.ApplyT(func(v *TriggerInfo) TriggerInfo {
		if v != nil {
			return *v
		}
		var ret TriggerInfo
		return ret
	}).(TriggerInfoOutput)
}

// UTC time string.
func (o TriggerInfoPtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TriggerInfo) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

// Name of function(action) that exists in the given namespace.
func (o TriggerInfoPtrOutput) Function() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TriggerInfo) *string {
		if v == nil {
			return nil
		}
		return v.Function
	}).(pulumi.StringPtrOutput)
}

// Indicates weather the trigger is paused or unpaused.
func (o TriggerInfoPtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TriggerInfo) *bool {
		if v == nil {
			return nil
		}
		return v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

// The trigger's unique name within the namespace.
func (o TriggerInfoPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TriggerInfo) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

// A unique string format of UUID with a prefix fn-.
func (o TriggerInfoPtrOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TriggerInfo) *string {
		if v == nil {
			return nil
		}
		return v.Namespace
	}).(pulumi.StringPtrOutput)
}

// Trigger details for SCHEDULED type, where body is optional.
func (o TriggerInfoPtrOutput) ScheduledDetails() ScheduledDetailsPtrOutput {
	return o.ApplyT(func(v *TriggerInfo) *ScheduledDetails {
		if v == nil {
			return nil
		}
		return v.ScheduledDetails
	}).(ScheduledDetailsPtrOutput)
}

func (o TriggerInfoPtrOutput) ScheduledRuns() TriggerInfoScheduledRunsPropertiesPtrOutput {
	return o.ApplyT(func(v *TriggerInfo) *TriggerInfoScheduledRunsProperties {
		if v == nil {
			return nil
		}
		return v.ScheduledRuns
	}).(TriggerInfoScheduledRunsPropertiesPtrOutput)
}

// String which indicates the type of trigger source like SCHEDULED.
func (o TriggerInfoPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TriggerInfo) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

// UTC time string.
func (o TriggerInfoPtrOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TriggerInfo) *string {
		if v == nil {
			return nil
		}
		return v.UpdatedAt
	}).(pulumi.StringPtrOutput)
}

type TriggerInfoArrayOutput struct{ *pulumi.OutputState }

func (TriggerInfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TriggerInfo)(nil)).Elem()
}

func (o TriggerInfoArrayOutput) ToTriggerInfoArrayOutput() TriggerInfoArrayOutput {
	return o
}

func (o TriggerInfoArrayOutput) ToTriggerInfoArrayOutputWithContext(ctx context.Context) TriggerInfoArrayOutput {
	return o
}

func (o TriggerInfoArrayOutput) Index(i pulumi.IntInput) TriggerInfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TriggerInfo {
		return vs[0].([]TriggerInfo)[vs[1].(int)]
	}).(TriggerInfoOutput)
}

type TriggerInfoScheduledRunsProperties struct {
	// Indicates last run time. null value indicates trigger not run yet.
	LastRunAt *string `pulumi:"lastRunAt"`
	// Indicates next run time. null value indicates trigger will not run.
	NextRunAt *string `pulumi:"nextRunAt"`
}

type TriggerInfoScheduledRunsPropertiesOutput struct{ *pulumi.OutputState }

func (TriggerInfoScheduledRunsPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerInfoScheduledRunsProperties)(nil)).Elem()
}

func (o TriggerInfoScheduledRunsPropertiesOutput) ToTriggerInfoScheduledRunsPropertiesOutput() TriggerInfoScheduledRunsPropertiesOutput {
	return o
}

func (o TriggerInfoScheduledRunsPropertiesOutput) ToTriggerInfoScheduledRunsPropertiesOutputWithContext(ctx context.Context) TriggerInfoScheduledRunsPropertiesOutput {
	return o
}

// Indicates last run time. null value indicates trigger not run yet.
func (o TriggerInfoScheduledRunsPropertiesOutput) LastRunAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TriggerInfoScheduledRunsProperties) *string { return v.LastRunAt }).(pulumi.StringPtrOutput)
}

// Indicates next run time. null value indicates trigger will not run.
func (o TriggerInfoScheduledRunsPropertiesOutput) NextRunAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TriggerInfoScheduledRunsProperties) *string { return v.NextRunAt }).(pulumi.StringPtrOutput)
}

type TriggerInfoScheduledRunsPropertiesPtrOutput struct{ *pulumi.OutputState }

func (TriggerInfoScheduledRunsPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerInfoScheduledRunsProperties)(nil)).Elem()
}

func (o TriggerInfoScheduledRunsPropertiesPtrOutput) ToTriggerInfoScheduledRunsPropertiesPtrOutput() TriggerInfoScheduledRunsPropertiesPtrOutput {
	return o
}

func (o TriggerInfoScheduledRunsPropertiesPtrOutput) ToTriggerInfoScheduledRunsPropertiesPtrOutputWithContext(ctx context.Context) TriggerInfoScheduledRunsPropertiesPtrOutput {
	return o
}

func (o TriggerInfoScheduledRunsPropertiesPtrOutput) Elem() TriggerInfoScheduledRunsPropertiesOutput {
	return o.ApplyT(func(v *TriggerInfoScheduledRunsProperties) TriggerInfoScheduledRunsProperties {
		if v != nil {
			return *v
		}
		var ret TriggerInfoScheduledRunsProperties
		return ret
	}).(TriggerInfoScheduledRunsPropertiesOutput)
}

// Indicates last run time. null value indicates trigger not run yet.
func (o TriggerInfoScheduledRunsPropertiesPtrOutput) LastRunAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TriggerInfoScheduledRunsProperties) *string {
		if v == nil {
			return nil
		}
		return v.LastRunAt
	}).(pulumi.StringPtrOutput)
}

// Indicates next run time. null value indicates trigger will not run.
func (o TriggerInfoScheduledRunsPropertiesPtrOutput) NextRunAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TriggerInfoScheduledRunsProperties) *string {
		if v == nil {
			return nil
		}
		return v.NextRunAt
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduledDetailsInput)(nil)).Elem(), ScheduledDetailsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduledDetailsBodyPropertiesInput)(nil)).Elem(), ScheduledDetailsBodyPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduledDetailsBodyPropertiesPtrInput)(nil)).Elem(), ScheduledDetailsBodyPropertiesArgs{})
	pulumi.RegisterOutputType(GetFunctionsNamespacePropertiesOutput{})
	pulumi.RegisterOutputType(GetFunctionsTriggerPropertiesOutput{})
	pulumi.RegisterOutputType(ListFunctionsNamespacesOutput{})
	pulumi.RegisterOutputType(ListFunctionsTriggersOutput{})
	pulumi.RegisterOutputType(NamespaceInfoOutput{})
	pulumi.RegisterOutputType(NamespaceInfoPtrOutput{})
	pulumi.RegisterOutputType(NamespaceInfoArrayOutput{})
	pulumi.RegisterOutputType(ScheduledDetailsOutput{})
	pulumi.RegisterOutputType(ScheduledDetailsPtrOutput{})
	pulumi.RegisterOutputType(ScheduledDetailsBodyPropertiesOutput{})
	pulumi.RegisterOutputType(ScheduledDetailsBodyPropertiesPtrOutput{})
	pulumi.RegisterOutputType(TriggerInfoOutput{})
	pulumi.RegisterOutputType(TriggerInfoPtrOutput{})
	pulumi.RegisterOutputType(TriggerInfoArrayOutput{})
	pulumi.RegisterOutputType(TriggerInfoScheduledRunsPropertiesOutput{})
	pulumi.RegisterOutputType(TriggerInfoScheduledRunsPropertiesPtrOutput{})
}
