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

type Firewall struct {
	// A time value given in ISO8601 combined date and time format that represents when the firewall was created.
	CreatedAt *string `pulumi:"createdAt"`
	// An array containing the IDs of the Droplets assigned to the firewall.
	DropletIds []int `pulumi:"dropletIds"`
	// A unique ID that can be used to identify and reference a firewall.
	Id           *string                         `pulumi:"id"`
	InboundRules []FirewallRulesInboundRulesItem `pulumi:"inboundRules"`
	// A human-readable name for a firewall. The name must begin with an alphanumeric character. Subsequent characters must either be alphanumeric characters, a period (.), or a dash (-).
	Name          *string                          `pulumi:"name"`
	OutboundRules []FirewallRulesOutboundRulesItem `pulumi:"outboundRules"`
	// An array of objects each containing the fields "droplet_id", "removing", and "status". It is provided to detail exactly which Droplets are having their security policies updated. When empty, all changes have been successfully applied.
	PendingChanges []FirewallPropertiesPendingChangesItemProperties `pulumi:"pendingChanges"`
	// A status string indicating the current state of the firewall. This can be "waiting", "succeeded", or "failed".
	Status *FirewallPropertiesStatus `pulumi:"status"`
	Tags   *FirewallPropertiesTags   `pulumi:"tags"`
}

type FirewallOutput struct{ *pulumi.OutputState }

func (FirewallOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Firewall)(nil)).Elem()
}

func (o FirewallOutput) ToFirewallOutput() FirewallOutput {
	return o
}

func (o FirewallOutput) ToFirewallOutputWithContext(ctx context.Context) FirewallOutput {
	return o
}

// A time value given in ISO8601 combined date and time format that represents when the firewall was created.
func (o FirewallOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Firewall) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// An array containing the IDs of the Droplets assigned to the firewall.
func (o FirewallOutput) DropletIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v Firewall) []int { return v.DropletIds }).(pulumi.IntArrayOutput)
}

// A unique ID that can be used to identify and reference a firewall.
func (o FirewallOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Firewall) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o FirewallOutput) InboundRules() FirewallRulesInboundRulesItemArrayOutput {
	return o.ApplyT(func(v Firewall) []FirewallRulesInboundRulesItem { return v.InboundRules }).(FirewallRulesInboundRulesItemArrayOutput)
}

// A human-readable name for a firewall. The name must begin with an alphanumeric character. Subsequent characters must either be alphanumeric characters, a period (.), or a dash (-).
func (o FirewallOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Firewall) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FirewallOutput) OutboundRules() FirewallRulesOutboundRulesItemArrayOutput {
	return o.ApplyT(func(v Firewall) []FirewallRulesOutboundRulesItem { return v.OutboundRules }).(FirewallRulesOutboundRulesItemArrayOutput)
}

// An array of objects each containing the fields "droplet_id", "removing", and "status". It is provided to detail exactly which Droplets are having their security policies updated. When empty, all changes have been successfully applied.
func (o FirewallOutput) PendingChanges() FirewallPropertiesPendingChangesItemPropertiesArrayOutput {
	return o.ApplyT(func(v Firewall) []FirewallPropertiesPendingChangesItemProperties { return v.PendingChanges }).(FirewallPropertiesPendingChangesItemPropertiesArrayOutput)
}

// A status string indicating the current state of the firewall. This can be "waiting", "succeeded", or "failed".
func (o FirewallOutput) Status() FirewallPropertiesStatusPtrOutput {
	return o.ApplyT(func(v Firewall) *FirewallPropertiesStatus { return v.Status }).(FirewallPropertiesStatusPtrOutput)
}

func (o FirewallOutput) Tags() FirewallPropertiesTagsPtrOutput {
	return o.ApplyT(func(v Firewall) *FirewallPropertiesTags { return v.Tags }).(FirewallPropertiesTagsPtrOutput)
}

type FirewallPtrOutput struct{ *pulumi.OutputState }

func (FirewallPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Firewall)(nil)).Elem()
}

func (o FirewallPtrOutput) ToFirewallPtrOutput() FirewallPtrOutput {
	return o
}

func (o FirewallPtrOutput) ToFirewallPtrOutputWithContext(ctx context.Context) FirewallPtrOutput {
	return o
}

func (o FirewallPtrOutput) Elem() FirewallOutput {
	return o.ApplyT(func(v *Firewall) Firewall {
		if v != nil {
			return *v
		}
		var ret Firewall
		return ret
	}).(FirewallOutput)
}

// A time value given in ISO8601 combined date and time format that represents when the firewall was created.
func (o FirewallPtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Firewall) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

// An array containing the IDs of the Droplets assigned to the firewall.
func (o FirewallPtrOutput) DropletIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *Firewall) []int {
		if v == nil {
			return nil
		}
		return v.DropletIds
	}).(pulumi.IntArrayOutput)
}

// A unique ID that can be used to identify and reference a firewall.
func (o FirewallPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Firewall) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o FirewallPtrOutput) InboundRules() FirewallRulesInboundRulesItemArrayOutput {
	return o.ApplyT(func(v *Firewall) []FirewallRulesInboundRulesItem {
		if v == nil {
			return nil
		}
		return v.InboundRules
	}).(FirewallRulesInboundRulesItemArrayOutput)
}

// A human-readable name for a firewall. The name must begin with an alphanumeric character. Subsequent characters must either be alphanumeric characters, a period (.), or a dash (-).
func (o FirewallPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Firewall) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o FirewallPtrOutput) OutboundRules() FirewallRulesOutboundRulesItemArrayOutput {
	return o.ApplyT(func(v *Firewall) []FirewallRulesOutboundRulesItem {
		if v == nil {
			return nil
		}
		return v.OutboundRules
	}).(FirewallRulesOutboundRulesItemArrayOutput)
}

// An array of objects each containing the fields "droplet_id", "removing", and "status". It is provided to detail exactly which Droplets are having their security policies updated. When empty, all changes have been successfully applied.
func (o FirewallPtrOutput) PendingChanges() FirewallPropertiesPendingChangesItemPropertiesArrayOutput {
	return o.ApplyT(func(v *Firewall) []FirewallPropertiesPendingChangesItemProperties {
		if v == nil {
			return nil
		}
		return v.PendingChanges
	}).(FirewallPropertiesPendingChangesItemPropertiesArrayOutput)
}

// A status string indicating the current state of the firewall. This can be "waiting", "succeeded", or "failed".
func (o FirewallPtrOutput) Status() FirewallPropertiesStatusPtrOutput {
	return o.ApplyT(func(v *Firewall) *FirewallPropertiesStatus {
		if v == nil {
			return nil
		}
		return v.Status
	}).(FirewallPropertiesStatusPtrOutput)
}

func (o FirewallPtrOutput) Tags() FirewallPropertiesTagsPtrOutput {
	return o.ApplyT(func(v *Firewall) *FirewallPropertiesTags {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(FirewallPropertiesTagsPtrOutput)
}

type FirewallArrayOutput struct{ *pulumi.OutputState }

func (FirewallArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Firewall)(nil)).Elem()
}

func (o FirewallArrayOutput) ToFirewallArrayOutput() FirewallArrayOutput {
	return o
}

func (o FirewallArrayOutput) ToFirewallArrayOutputWithContext(ctx context.Context) FirewallArrayOutput {
	return o
}

func (o FirewallArrayOutput) Index(i pulumi.IntInput) FirewallOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Firewall {
		return vs[0].([]Firewall)[vs[1].(int)]
	}).(FirewallOutput)
}

type FirewallPropertiesPendingChangesItemProperties struct {
	DropletId *int    `pulumi:"dropletId"`
	Removing  *bool   `pulumi:"removing"`
	Status    *string `pulumi:"status"`
}

// FirewallPropertiesPendingChangesItemPropertiesInput is an input type that accepts FirewallPropertiesPendingChangesItemPropertiesArgs and FirewallPropertiesPendingChangesItemPropertiesOutput values.
// You can construct a concrete instance of `FirewallPropertiesPendingChangesItemPropertiesInput` via:
//
//	FirewallPropertiesPendingChangesItemPropertiesArgs{...}
type FirewallPropertiesPendingChangesItemPropertiesInput interface {
	pulumi.Input

	ToFirewallPropertiesPendingChangesItemPropertiesOutput() FirewallPropertiesPendingChangesItemPropertiesOutput
	ToFirewallPropertiesPendingChangesItemPropertiesOutputWithContext(context.Context) FirewallPropertiesPendingChangesItemPropertiesOutput
}

type FirewallPropertiesPendingChangesItemPropertiesArgs struct {
	DropletId pulumi.IntPtrInput    `pulumi:"dropletId"`
	Removing  pulumi.BoolPtrInput   `pulumi:"removing"`
	Status    pulumi.StringPtrInput `pulumi:"status"`
}

func (FirewallPropertiesPendingChangesItemPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPropertiesPendingChangesItemProperties)(nil)).Elem()
}

func (i FirewallPropertiesPendingChangesItemPropertiesArgs) ToFirewallPropertiesPendingChangesItemPropertiesOutput() FirewallPropertiesPendingChangesItemPropertiesOutput {
	return i.ToFirewallPropertiesPendingChangesItemPropertiesOutputWithContext(context.Background())
}

func (i FirewallPropertiesPendingChangesItemPropertiesArgs) ToFirewallPropertiesPendingChangesItemPropertiesOutputWithContext(ctx context.Context) FirewallPropertiesPendingChangesItemPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPropertiesPendingChangesItemPropertiesOutput)
}

// FirewallPropertiesPendingChangesItemPropertiesArrayInput is an input type that accepts FirewallPropertiesPendingChangesItemPropertiesArray and FirewallPropertiesPendingChangesItemPropertiesArrayOutput values.
// You can construct a concrete instance of `FirewallPropertiesPendingChangesItemPropertiesArrayInput` via:
//
//	FirewallPropertiesPendingChangesItemPropertiesArray{ FirewallPropertiesPendingChangesItemPropertiesArgs{...} }
type FirewallPropertiesPendingChangesItemPropertiesArrayInput interface {
	pulumi.Input

	ToFirewallPropertiesPendingChangesItemPropertiesArrayOutput() FirewallPropertiesPendingChangesItemPropertiesArrayOutput
	ToFirewallPropertiesPendingChangesItemPropertiesArrayOutputWithContext(context.Context) FirewallPropertiesPendingChangesItemPropertiesArrayOutput
}

type FirewallPropertiesPendingChangesItemPropertiesArray []FirewallPropertiesPendingChangesItemPropertiesInput

func (FirewallPropertiesPendingChangesItemPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallPropertiesPendingChangesItemProperties)(nil)).Elem()
}

func (i FirewallPropertiesPendingChangesItemPropertiesArray) ToFirewallPropertiesPendingChangesItemPropertiesArrayOutput() FirewallPropertiesPendingChangesItemPropertiesArrayOutput {
	return i.ToFirewallPropertiesPendingChangesItemPropertiesArrayOutputWithContext(context.Background())
}

func (i FirewallPropertiesPendingChangesItemPropertiesArray) ToFirewallPropertiesPendingChangesItemPropertiesArrayOutputWithContext(ctx context.Context) FirewallPropertiesPendingChangesItemPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPropertiesPendingChangesItemPropertiesArrayOutput)
}

type FirewallPropertiesPendingChangesItemPropertiesOutput struct{ *pulumi.OutputState }

func (FirewallPropertiesPendingChangesItemPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPropertiesPendingChangesItemProperties)(nil)).Elem()
}

func (o FirewallPropertiesPendingChangesItemPropertiesOutput) ToFirewallPropertiesPendingChangesItemPropertiesOutput() FirewallPropertiesPendingChangesItemPropertiesOutput {
	return o
}

func (o FirewallPropertiesPendingChangesItemPropertiesOutput) ToFirewallPropertiesPendingChangesItemPropertiesOutputWithContext(ctx context.Context) FirewallPropertiesPendingChangesItemPropertiesOutput {
	return o
}

func (o FirewallPropertiesPendingChangesItemPropertiesOutput) DropletId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FirewallPropertiesPendingChangesItemProperties) *int { return v.DropletId }).(pulumi.IntPtrOutput)
}

func (o FirewallPropertiesPendingChangesItemPropertiesOutput) Removing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FirewallPropertiesPendingChangesItemProperties) *bool { return v.Removing }).(pulumi.BoolPtrOutput)
}

func (o FirewallPropertiesPendingChangesItemPropertiesOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FirewallPropertiesPendingChangesItemProperties) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type FirewallPropertiesPendingChangesItemPropertiesArrayOutput struct{ *pulumi.OutputState }

func (FirewallPropertiesPendingChangesItemPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallPropertiesPendingChangesItemProperties)(nil)).Elem()
}

func (o FirewallPropertiesPendingChangesItemPropertiesArrayOutput) ToFirewallPropertiesPendingChangesItemPropertiesArrayOutput() FirewallPropertiesPendingChangesItemPropertiesArrayOutput {
	return o
}

func (o FirewallPropertiesPendingChangesItemPropertiesArrayOutput) ToFirewallPropertiesPendingChangesItemPropertiesArrayOutputWithContext(ctx context.Context) FirewallPropertiesPendingChangesItemPropertiesArrayOutput {
	return o
}

func (o FirewallPropertiesPendingChangesItemPropertiesArrayOutput) Index(i pulumi.IntInput) FirewallPropertiesPendingChangesItemPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FirewallPropertiesPendingChangesItemProperties {
		return vs[0].([]FirewallPropertiesPendingChangesItemProperties)[vs[1].(int)]
	}).(FirewallPropertiesPendingChangesItemPropertiesOutput)
}

type FirewallPropertiesTags struct {
}

// FirewallPropertiesTagsInput is an input type that accepts FirewallPropertiesTagsArgs and FirewallPropertiesTagsOutput values.
// You can construct a concrete instance of `FirewallPropertiesTagsInput` via:
//
//	FirewallPropertiesTagsArgs{...}
type FirewallPropertiesTagsInput interface {
	pulumi.Input

	ToFirewallPropertiesTagsOutput() FirewallPropertiesTagsOutput
	ToFirewallPropertiesTagsOutputWithContext(context.Context) FirewallPropertiesTagsOutput
}

type FirewallPropertiesTagsArgs struct {
}

func (FirewallPropertiesTagsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPropertiesTags)(nil)).Elem()
}

func (i FirewallPropertiesTagsArgs) ToFirewallPropertiesTagsOutput() FirewallPropertiesTagsOutput {
	return i.ToFirewallPropertiesTagsOutputWithContext(context.Background())
}

func (i FirewallPropertiesTagsArgs) ToFirewallPropertiesTagsOutputWithContext(ctx context.Context) FirewallPropertiesTagsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPropertiesTagsOutput)
}

func (i FirewallPropertiesTagsArgs) ToFirewallPropertiesTagsPtrOutput() FirewallPropertiesTagsPtrOutput {
	return i.ToFirewallPropertiesTagsPtrOutputWithContext(context.Background())
}

func (i FirewallPropertiesTagsArgs) ToFirewallPropertiesTagsPtrOutputWithContext(ctx context.Context) FirewallPropertiesTagsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPropertiesTagsOutput).ToFirewallPropertiesTagsPtrOutputWithContext(ctx)
}

// FirewallPropertiesTagsPtrInput is an input type that accepts FirewallPropertiesTagsArgs, FirewallPropertiesTagsPtr and FirewallPropertiesTagsPtrOutput values.
// You can construct a concrete instance of `FirewallPropertiesTagsPtrInput` via:
//
//	        FirewallPropertiesTagsArgs{...}
//
//	or:
//
//	        nil
type FirewallPropertiesTagsPtrInput interface {
	pulumi.Input

	ToFirewallPropertiesTagsPtrOutput() FirewallPropertiesTagsPtrOutput
	ToFirewallPropertiesTagsPtrOutputWithContext(context.Context) FirewallPropertiesTagsPtrOutput
}

type firewallPropertiesTagsPtrType FirewallPropertiesTagsArgs

func FirewallPropertiesTagsPtr(v *FirewallPropertiesTagsArgs) FirewallPropertiesTagsPtrInput {
	return (*firewallPropertiesTagsPtrType)(v)
}

func (*firewallPropertiesTagsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPropertiesTags)(nil)).Elem()
}

func (i *firewallPropertiesTagsPtrType) ToFirewallPropertiesTagsPtrOutput() FirewallPropertiesTagsPtrOutput {
	return i.ToFirewallPropertiesTagsPtrOutputWithContext(context.Background())
}

func (i *firewallPropertiesTagsPtrType) ToFirewallPropertiesTagsPtrOutputWithContext(ctx context.Context) FirewallPropertiesTagsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPropertiesTagsPtrOutput)
}

type FirewallPropertiesTagsOutput struct{ *pulumi.OutputState }

func (FirewallPropertiesTagsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPropertiesTags)(nil)).Elem()
}

func (o FirewallPropertiesTagsOutput) ToFirewallPropertiesTagsOutput() FirewallPropertiesTagsOutput {
	return o
}

func (o FirewallPropertiesTagsOutput) ToFirewallPropertiesTagsOutputWithContext(ctx context.Context) FirewallPropertiesTagsOutput {
	return o
}

func (o FirewallPropertiesTagsOutput) ToFirewallPropertiesTagsPtrOutput() FirewallPropertiesTagsPtrOutput {
	return o.ToFirewallPropertiesTagsPtrOutputWithContext(context.Background())
}

func (o FirewallPropertiesTagsOutput) ToFirewallPropertiesTagsPtrOutputWithContext(ctx context.Context) FirewallPropertiesTagsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FirewallPropertiesTags) *FirewallPropertiesTags {
		return &v
	}).(FirewallPropertiesTagsPtrOutput)
}

type FirewallPropertiesTagsPtrOutput struct{ *pulumi.OutputState }

func (FirewallPropertiesTagsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPropertiesTags)(nil)).Elem()
}

func (o FirewallPropertiesTagsPtrOutput) ToFirewallPropertiesTagsPtrOutput() FirewallPropertiesTagsPtrOutput {
	return o
}

func (o FirewallPropertiesTagsPtrOutput) ToFirewallPropertiesTagsPtrOutputWithContext(ctx context.Context) FirewallPropertiesTagsPtrOutput {
	return o
}

func (o FirewallPropertiesTagsPtrOutput) Elem() FirewallPropertiesTagsOutput {
	return o.ApplyT(func(v *FirewallPropertiesTags) FirewallPropertiesTags {
		if v != nil {
			return *v
		}
		var ret FirewallPropertiesTags
		return ret
	}).(FirewallPropertiesTagsOutput)
}

type FirewallRulesInboundRulesItem struct {
	// The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. "8000-9000"), or "0" when all ports are open for a protocol. For ICMP rules this parameter will always return "0".
	Ports string `pulumi:"ports"`
	// The type of traffic to be allowed. This may be one of `tcp`, `udp`, or `icmp`.
	Protocol FirewallRuleBaseProtocol `pulumi:"protocol"`
}

// FirewallRulesInboundRulesItemInput is an input type that accepts FirewallRulesInboundRulesItemArgs and FirewallRulesInboundRulesItemOutput values.
// You can construct a concrete instance of `FirewallRulesInboundRulesItemInput` via:
//
//	FirewallRulesInboundRulesItemArgs{...}
type FirewallRulesInboundRulesItemInput interface {
	pulumi.Input

	ToFirewallRulesInboundRulesItemOutput() FirewallRulesInboundRulesItemOutput
	ToFirewallRulesInboundRulesItemOutputWithContext(context.Context) FirewallRulesInboundRulesItemOutput
}

type FirewallRulesInboundRulesItemArgs struct {
	// The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. "8000-9000"), or "0" when all ports are open for a protocol. For ICMP rules this parameter will always return "0".
	Ports pulumi.StringInput `pulumi:"ports"`
	// The type of traffic to be allowed. This may be one of `tcp`, `udp`, or `icmp`.
	Protocol FirewallRuleBaseProtocolInput `pulumi:"protocol"`
}

func (FirewallRulesInboundRulesItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallRulesInboundRulesItem)(nil)).Elem()
}

func (i FirewallRulesInboundRulesItemArgs) ToFirewallRulesInboundRulesItemOutput() FirewallRulesInboundRulesItemOutput {
	return i.ToFirewallRulesInboundRulesItemOutputWithContext(context.Background())
}

func (i FirewallRulesInboundRulesItemArgs) ToFirewallRulesInboundRulesItemOutputWithContext(ctx context.Context) FirewallRulesInboundRulesItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRulesInboundRulesItemOutput)
}

// FirewallRulesInboundRulesItemArrayInput is an input type that accepts FirewallRulesInboundRulesItemArray and FirewallRulesInboundRulesItemArrayOutput values.
// You can construct a concrete instance of `FirewallRulesInboundRulesItemArrayInput` via:
//
//	FirewallRulesInboundRulesItemArray{ FirewallRulesInboundRulesItemArgs{...} }
type FirewallRulesInboundRulesItemArrayInput interface {
	pulumi.Input

	ToFirewallRulesInboundRulesItemArrayOutput() FirewallRulesInboundRulesItemArrayOutput
	ToFirewallRulesInboundRulesItemArrayOutputWithContext(context.Context) FirewallRulesInboundRulesItemArrayOutput
}

type FirewallRulesInboundRulesItemArray []FirewallRulesInboundRulesItemInput

func (FirewallRulesInboundRulesItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallRulesInboundRulesItem)(nil)).Elem()
}

func (i FirewallRulesInboundRulesItemArray) ToFirewallRulesInboundRulesItemArrayOutput() FirewallRulesInboundRulesItemArrayOutput {
	return i.ToFirewallRulesInboundRulesItemArrayOutputWithContext(context.Background())
}

func (i FirewallRulesInboundRulesItemArray) ToFirewallRulesInboundRulesItemArrayOutputWithContext(ctx context.Context) FirewallRulesInboundRulesItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRulesInboundRulesItemArrayOutput)
}

type FirewallRulesInboundRulesItemOutput struct{ *pulumi.OutputState }

func (FirewallRulesInboundRulesItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallRulesInboundRulesItem)(nil)).Elem()
}

func (o FirewallRulesInboundRulesItemOutput) ToFirewallRulesInboundRulesItemOutput() FirewallRulesInboundRulesItemOutput {
	return o
}

func (o FirewallRulesInboundRulesItemOutput) ToFirewallRulesInboundRulesItemOutputWithContext(ctx context.Context) FirewallRulesInboundRulesItemOutput {
	return o
}

// The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. "8000-9000"), or "0" when all ports are open for a protocol. For ICMP rules this parameter will always return "0".
func (o FirewallRulesInboundRulesItemOutput) Ports() pulumi.StringOutput {
	return o.ApplyT(func(v FirewallRulesInboundRulesItem) string { return v.Ports }).(pulumi.StringOutput)
}

// The type of traffic to be allowed. This may be one of `tcp`, `udp`, or `icmp`.
func (o FirewallRulesInboundRulesItemOutput) Protocol() FirewallRuleBaseProtocolOutput {
	return o.ApplyT(func(v FirewallRulesInboundRulesItem) FirewallRuleBaseProtocol { return v.Protocol }).(FirewallRuleBaseProtocolOutput)
}

type FirewallRulesInboundRulesItemArrayOutput struct{ *pulumi.OutputState }

func (FirewallRulesInboundRulesItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallRulesInboundRulesItem)(nil)).Elem()
}

func (o FirewallRulesInboundRulesItemArrayOutput) ToFirewallRulesInboundRulesItemArrayOutput() FirewallRulesInboundRulesItemArrayOutput {
	return o
}

func (o FirewallRulesInboundRulesItemArrayOutput) ToFirewallRulesInboundRulesItemArrayOutputWithContext(ctx context.Context) FirewallRulesInboundRulesItemArrayOutput {
	return o
}

func (o FirewallRulesInboundRulesItemArrayOutput) Index(i pulumi.IntInput) FirewallRulesInboundRulesItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FirewallRulesInboundRulesItem {
		return vs[0].([]FirewallRulesInboundRulesItem)[vs[1].(int)]
	}).(FirewallRulesInboundRulesItemOutput)
}

type FirewallRulesOutboundRulesItem struct {
	// The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. "8000-9000"), or "0" when all ports are open for a protocol. For ICMP rules this parameter will always return "0".
	Ports string `pulumi:"ports"`
	// The type of traffic to be allowed. This may be one of `tcp`, `udp`, or `icmp`.
	Protocol FirewallRuleBaseProtocol `pulumi:"protocol"`
}

// FirewallRulesOutboundRulesItemInput is an input type that accepts FirewallRulesOutboundRulesItemArgs and FirewallRulesOutboundRulesItemOutput values.
// You can construct a concrete instance of `FirewallRulesOutboundRulesItemInput` via:
//
//	FirewallRulesOutboundRulesItemArgs{...}
type FirewallRulesOutboundRulesItemInput interface {
	pulumi.Input

	ToFirewallRulesOutboundRulesItemOutput() FirewallRulesOutboundRulesItemOutput
	ToFirewallRulesOutboundRulesItemOutputWithContext(context.Context) FirewallRulesOutboundRulesItemOutput
}

type FirewallRulesOutboundRulesItemArgs struct {
	// The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. "8000-9000"), or "0" when all ports are open for a protocol. For ICMP rules this parameter will always return "0".
	Ports pulumi.StringInput `pulumi:"ports"`
	// The type of traffic to be allowed. This may be one of `tcp`, `udp`, or `icmp`.
	Protocol FirewallRuleBaseProtocolInput `pulumi:"protocol"`
}

func (FirewallRulesOutboundRulesItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallRulesOutboundRulesItem)(nil)).Elem()
}

func (i FirewallRulesOutboundRulesItemArgs) ToFirewallRulesOutboundRulesItemOutput() FirewallRulesOutboundRulesItemOutput {
	return i.ToFirewallRulesOutboundRulesItemOutputWithContext(context.Background())
}

func (i FirewallRulesOutboundRulesItemArgs) ToFirewallRulesOutboundRulesItemOutputWithContext(ctx context.Context) FirewallRulesOutboundRulesItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRulesOutboundRulesItemOutput)
}

// FirewallRulesOutboundRulesItemArrayInput is an input type that accepts FirewallRulesOutboundRulesItemArray and FirewallRulesOutboundRulesItemArrayOutput values.
// You can construct a concrete instance of `FirewallRulesOutboundRulesItemArrayInput` via:
//
//	FirewallRulesOutboundRulesItemArray{ FirewallRulesOutboundRulesItemArgs{...} }
type FirewallRulesOutboundRulesItemArrayInput interface {
	pulumi.Input

	ToFirewallRulesOutboundRulesItemArrayOutput() FirewallRulesOutboundRulesItemArrayOutput
	ToFirewallRulesOutboundRulesItemArrayOutputWithContext(context.Context) FirewallRulesOutboundRulesItemArrayOutput
}

type FirewallRulesOutboundRulesItemArray []FirewallRulesOutboundRulesItemInput

func (FirewallRulesOutboundRulesItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallRulesOutboundRulesItem)(nil)).Elem()
}

func (i FirewallRulesOutboundRulesItemArray) ToFirewallRulesOutboundRulesItemArrayOutput() FirewallRulesOutboundRulesItemArrayOutput {
	return i.ToFirewallRulesOutboundRulesItemArrayOutputWithContext(context.Background())
}

func (i FirewallRulesOutboundRulesItemArray) ToFirewallRulesOutboundRulesItemArrayOutputWithContext(ctx context.Context) FirewallRulesOutboundRulesItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRulesOutboundRulesItemArrayOutput)
}

type FirewallRulesOutboundRulesItemOutput struct{ *pulumi.OutputState }

func (FirewallRulesOutboundRulesItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallRulesOutboundRulesItem)(nil)).Elem()
}

func (o FirewallRulesOutboundRulesItemOutput) ToFirewallRulesOutboundRulesItemOutput() FirewallRulesOutboundRulesItemOutput {
	return o
}

func (o FirewallRulesOutboundRulesItemOutput) ToFirewallRulesOutboundRulesItemOutputWithContext(ctx context.Context) FirewallRulesOutboundRulesItemOutput {
	return o
}

// The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. "8000-9000"), or "0" when all ports are open for a protocol. For ICMP rules this parameter will always return "0".
func (o FirewallRulesOutboundRulesItemOutput) Ports() pulumi.StringOutput {
	return o.ApplyT(func(v FirewallRulesOutboundRulesItem) string { return v.Ports }).(pulumi.StringOutput)
}

// The type of traffic to be allowed. This may be one of `tcp`, `udp`, or `icmp`.
func (o FirewallRulesOutboundRulesItemOutput) Protocol() FirewallRuleBaseProtocolOutput {
	return o.ApplyT(func(v FirewallRulesOutboundRulesItem) FirewallRuleBaseProtocol { return v.Protocol }).(FirewallRuleBaseProtocolOutput)
}

type FirewallRulesOutboundRulesItemArrayOutput struct{ *pulumi.OutputState }

func (FirewallRulesOutboundRulesItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallRulesOutboundRulesItem)(nil)).Elem()
}

func (o FirewallRulesOutboundRulesItemArrayOutput) ToFirewallRulesOutboundRulesItemArrayOutput() FirewallRulesOutboundRulesItemArrayOutput {
	return o
}

func (o FirewallRulesOutboundRulesItemArrayOutput) ToFirewallRulesOutboundRulesItemArrayOutputWithContext(ctx context.Context) FirewallRulesOutboundRulesItemArrayOutput {
	return o
}

func (o FirewallRulesOutboundRulesItemArrayOutput) Index(i pulumi.IntInput) FirewallRulesOutboundRulesItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FirewallRulesOutboundRulesItem {
		return vs[0].([]FirewallRulesOutboundRulesItem)[vs[1].(int)]
	}).(FirewallRulesOutboundRulesItemOutput)
}

type GetFirewallsProperties struct {
	Firewall *Firewall `pulumi:"firewall"`
}

type GetFirewallsPropertiesOutput struct{ *pulumi.OutputState }

func (GetFirewallsPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallsProperties)(nil)).Elem()
}

func (o GetFirewallsPropertiesOutput) ToGetFirewallsPropertiesOutput() GetFirewallsPropertiesOutput {
	return o
}

func (o GetFirewallsPropertiesOutput) ToGetFirewallsPropertiesOutputWithContext(ctx context.Context) GetFirewallsPropertiesOutput {
	return o
}

func (o GetFirewallsPropertiesOutput) Firewall() FirewallPtrOutput {
	return o.ApplyT(func(v GetFirewallsProperties) *Firewall { return v.Firewall }).(FirewallPtrOutput)
}

type ListFirewallsItems struct {
	Firewalls []Firewall `pulumi:"firewalls"`
	Links     *PageLinks `pulumi:"links"`
	Meta      MetaMeta   `pulumi:"meta"`
}

type ListFirewallsItemsOutput struct{ *pulumi.OutputState }

func (ListFirewallsItemsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListFirewallsItems)(nil)).Elem()
}

func (o ListFirewallsItemsOutput) ToListFirewallsItemsOutput() ListFirewallsItemsOutput {
	return o
}

func (o ListFirewallsItemsOutput) ToListFirewallsItemsOutputWithContext(ctx context.Context) ListFirewallsItemsOutput {
	return o
}

func (o ListFirewallsItemsOutput) Firewalls() FirewallArrayOutput {
	return o.ApplyT(func(v ListFirewallsItems) []Firewall { return v.Firewalls }).(FirewallArrayOutput)
}

func (o ListFirewallsItemsOutput) Links() PageLinksPtrOutput {
	return o.ApplyT(func(v ListFirewallsItems) *PageLinks { return v.Links }).(PageLinksPtrOutput)
}

func (o ListFirewallsItemsOutput) Meta() MetaMetaOutput {
	return o.ApplyT(func(v ListFirewallsItems) MetaMeta { return v.Meta }).(MetaMetaOutput)
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

type Tags struct {
}

// TagsInput is an input type that accepts TagsArgs and TagsOutput values.
// You can construct a concrete instance of `TagsInput` via:
//
//	TagsArgs{...}
type TagsInput interface {
	pulumi.Input

	ToTagsOutput() TagsOutput
	ToTagsOutputWithContext(context.Context) TagsOutput
}

type TagsArgs struct {
}

func (TagsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Tags)(nil)).Elem()
}

func (i TagsArgs) ToTagsOutput() TagsOutput {
	return i.ToTagsOutputWithContext(context.Background())
}

func (i TagsArgs) ToTagsOutputWithContext(ctx context.Context) TagsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagsOutput)
}

type TagsOutput struct{ *pulumi.OutputState }

func (TagsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Tags)(nil)).Elem()
}

func (o TagsOutput) ToTagsOutput() TagsOutput {
	return o
}

func (o TagsOutput) ToTagsOutputWithContext(ctx context.Context) TagsOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallPropertiesPendingChangesItemPropertiesInput)(nil)).Elem(), FirewallPropertiesPendingChangesItemPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallPropertiesPendingChangesItemPropertiesArrayInput)(nil)).Elem(), FirewallPropertiesPendingChangesItemPropertiesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallPropertiesTagsInput)(nil)).Elem(), FirewallPropertiesTagsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallPropertiesTagsPtrInput)(nil)).Elem(), FirewallPropertiesTagsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallRulesInboundRulesItemInput)(nil)).Elem(), FirewallRulesInboundRulesItemArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallRulesInboundRulesItemArrayInput)(nil)).Elem(), FirewallRulesInboundRulesItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallRulesOutboundRulesItemInput)(nil)).Elem(), FirewallRulesOutboundRulesItemArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallRulesOutboundRulesItemArrayInput)(nil)).Elem(), FirewallRulesOutboundRulesItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TagsInput)(nil)).Elem(), TagsArgs{})
	pulumi.RegisterOutputType(FirewallOutput{})
	pulumi.RegisterOutputType(FirewallPtrOutput{})
	pulumi.RegisterOutputType(FirewallArrayOutput{})
	pulumi.RegisterOutputType(FirewallPropertiesPendingChangesItemPropertiesOutput{})
	pulumi.RegisterOutputType(FirewallPropertiesPendingChangesItemPropertiesArrayOutput{})
	pulumi.RegisterOutputType(FirewallPropertiesTagsOutput{})
	pulumi.RegisterOutputType(FirewallPropertiesTagsPtrOutput{})
	pulumi.RegisterOutputType(FirewallRulesInboundRulesItemOutput{})
	pulumi.RegisterOutputType(FirewallRulesInboundRulesItemArrayOutput{})
	pulumi.RegisterOutputType(FirewallRulesOutboundRulesItemOutput{})
	pulumi.RegisterOutputType(FirewallRulesOutboundRulesItemArrayOutput{})
	pulumi.RegisterOutputType(GetFirewallsPropertiesOutput{})
	pulumi.RegisterOutputType(ListFirewallsItemsOutput{})
	pulumi.RegisterOutputType(MetaMetaOutput{})
	pulumi.RegisterOutputType(PageLinksOutput{})
	pulumi.RegisterOutputType(PageLinksPtrOutput{})
	pulumi.RegisterOutputType(PageLinksPagesPropertiesOutput{})
	pulumi.RegisterOutputType(PageLinksPagesPropertiesPtrOutput{})
	pulumi.RegisterOutputType(TagsOutput{})
}
