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

type AlertPolicy struct {
	Alerts      Alerts             `pulumi:"alerts"`
	Compare     AlertPolicyCompare `pulumi:"compare"`
	Description string             `pulumi:"description"`
	Enabled     bool               `pulumi:"enabled"`
	Entities    []string           `pulumi:"entities"`
	Tags        []string           `pulumi:"tags"`
	Type        AlertPolicyType    `pulumi:"type"`
	Uuid        string             `pulumi:"uuid"`
	Value       float64            `pulumi:"value"`
	Window      AlertPolicyWindow  `pulumi:"window"`
}

type AlertPolicyOutput struct{ *pulumi.OutputState }

func (AlertPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertPolicy)(nil)).Elem()
}

func (o AlertPolicyOutput) ToAlertPolicyOutput() AlertPolicyOutput {
	return o
}

func (o AlertPolicyOutput) ToAlertPolicyOutputWithContext(ctx context.Context) AlertPolicyOutput {
	return o
}

func (o AlertPolicyOutput) Alerts() AlertsOutput {
	return o.ApplyT(func(v AlertPolicy) Alerts { return v.Alerts }).(AlertsOutput)
}

func (o AlertPolicyOutput) Compare() AlertPolicyCompareOutput {
	return o.ApplyT(func(v AlertPolicy) AlertPolicyCompare { return v.Compare }).(AlertPolicyCompareOutput)
}

func (o AlertPolicyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v AlertPolicy) string { return v.Description }).(pulumi.StringOutput)
}

func (o AlertPolicyOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v AlertPolicy) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o AlertPolicyOutput) Entities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AlertPolicy) []string { return v.Entities }).(pulumi.StringArrayOutput)
}

func (o AlertPolicyOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AlertPolicy) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o AlertPolicyOutput) Type() AlertPolicyTypeOutput {
	return o.ApplyT(func(v AlertPolicy) AlertPolicyType { return v.Type }).(AlertPolicyTypeOutput)
}

func (o AlertPolicyOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v AlertPolicy) string { return v.Uuid }).(pulumi.StringOutput)
}

func (o AlertPolicyOutput) Value() pulumi.Float64Output {
	return o.ApplyT(func(v AlertPolicy) float64 { return v.Value }).(pulumi.Float64Output)
}

func (o AlertPolicyOutput) Window() AlertPolicyWindowOutput {
	return o.ApplyT(func(v AlertPolicy) AlertPolicyWindow { return v.Window }).(AlertPolicyWindowOutput)
}

type AlertPolicyPtrOutput struct{ *pulumi.OutputState }

func (AlertPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertPolicy)(nil)).Elem()
}

func (o AlertPolicyPtrOutput) ToAlertPolicyPtrOutput() AlertPolicyPtrOutput {
	return o
}

func (o AlertPolicyPtrOutput) ToAlertPolicyPtrOutputWithContext(ctx context.Context) AlertPolicyPtrOutput {
	return o
}

func (o AlertPolicyPtrOutput) Elem() AlertPolicyOutput {
	return o.ApplyT(func(v *AlertPolicy) AlertPolicy {
		if v != nil {
			return *v
		}
		var ret AlertPolicy
		return ret
	}).(AlertPolicyOutput)
}

func (o AlertPolicyPtrOutput) Alerts() AlertsPtrOutput {
	return o.ApplyT(func(v *AlertPolicy) *Alerts {
		if v == nil {
			return nil
		}
		return &v.Alerts
	}).(AlertsPtrOutput)
}

func (o AlertPolicyPtrOutput) Compare() AlertPolicyComparePtrOutput {
	return o.ApplyT(func(v *AlertPolicy) *AlertPolicyCompare {
		if v == nil {
			return nil
		}
		return &v.Compare
	}).(AlertPolicyComparePtrOutput)
}

func (o AlertPolicyPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlertPolicy) *string {
		if v == nil {
			return nil
		}
		return &v.Description
	}).(pulumi.StringPtrOutput)
}

func (o AlertPolicyPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AlertPolicy) *bool {
		if v == nil {
			return nil
		}
		return &v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o AlertPolicyPtrOutput) Entities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AlertPolicy) []string {
		if v == nil {
			return nil
		}
		return v.Entities
	}).(pulumi.StringArrayOutput)
}

func (o AlertPolicyPtrOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AlertPolicy) []string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringArrayOutput)
}

func (o AlertPolicyPtrOutput) Type() AlertPolicyTypePtrOutput {
	return o.ApplyT(func(v *AlertPolicy) *AlertPolicyType {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(AlertPolicyTypePtrOutput)
}

func (o AlertPolicyPtrOutput) Uuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlertPolicy) *string {
		if v == nil {
			return nil
		}
		return &v.Uuid
	}).(pulumi.StringPtrOutput)
}

func (o AlertPolicyPtrOutput) Value() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *AlertPolicy) *float64 {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.Float64PtrOutput)
}

func (o AlertPolicyPtrOutput) Window() AlertPolicyWindowPtrOutput {
	return o.ApplyT(func(v *AlertPolicy) *AlertPolicyWindow {
		if v == nil {
			return nil
		}
		return &v.Window
	}).(AlertPolicyWindowPtrOutput)
}

type AlertPolicyArrayOutput struct{ *pulumi.OutputState }

func (AlertPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AlertPolicy)(nil)).Elem()
}

func (o AlertPolicyArrayOutput) ToAlertPolicyArrayOutput() AlertPolicyArrayOutput {
	return o
}

func (o AlertPolicyArrayOutput) ToAlertPolicyArrayOutputWithContext(ctx context.Context) AlertPolicyArrayOutput {
	return o
}

func (o AlertPolicyArrayOutput) Index(i pulumi.IntInput) AlertPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AlertPolicy {
		return vs[0].([]AlertPolicy)[vs[1].(int)]
	}).(AlertPolicyOutput)
}

type Alerts struct {
	// An email to notify on an alert trigger.
	Email []string `pulumi:"email"`
	// Slack integration details.
	Slack []SlackDetails `pulumi:"slack"`
}

// AlertsInput is an input type that accepts AlertsArgs and AlertsOutput values.
// You can construct a concrete instance of `AlertsInput` via:
//
//	AlertsArgs{...}
type AlertsInput interface {
	pulumi.Input

	ToAlertsOutput() AlertsOutput
	ToAlertsOutputWithContext(context.Context) AlertsOutput
}

type AlertsArgs struct {
	// An email to notify on an alert trigger.
	Email pulumi.StringArrayInput `pulumi:"email"`
	// Slack integration details.
	Slack SlackDetailsArrayInput `pulumi:"slack"`
}

func (AlertsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Alerts)(nil)).Elem()
}

func (i AlertsArgs) ToAlertsOutput() AlertsOutput {
	return i.ToAlertsOutputWithContext(context.Background())
}

func (i AlertsArgs) ToAlertsOutputWithContext(ctx context.Context) AlertsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertsOutput)
}

type AlertsOutput struct{ *pulumi.OutputState }

func (AlertsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Alerts)(nil)).Elem()
}

func (o AlertsOutput) ToAlertsOutput() AlertsOutput {
	return o
}

func (o AlertsOutput) ToAlertsOutputWithContext(ctx context.Context) AlertsOutput {
	return o
}

// An email to notify on an alert trigger.
func (o AlertsOutput) Email() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Alerts) []string { return v.Email }).(pulumi.StringArrayOutput)
}

// Slack integration details.
func (o AlertsOutput) Slack() SlackDetailsArrayOutput {
	return o.ApplyT(func(v Alerts) []SlackDetails { return v.Slack }).(SlackDetailsArrayOutput)
}

type AlertsPtrOutput struct{ *pulumi.OutputState }

func (AlertsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Alerts)(nil)).Elem()
}

func (o AlertsPtrOutput) ToAlertsPtrOutput() AlertsPtrOutput {
	return o
}

func (o AlertsPtrOutput) ToAlertsPtrOutputWithContext(ctx context.Context) AlertsPtrOutput {
	return o
}

func (o AlertsPtrOutput) Elem() AlertsOutput {
	return o.ApplyT(func(v *Alerts) Alerts {
		if v != nil {
			return *v
		}
		var ret Alerts
		return ret
	}).(AlertsOutput)
}

// An email to notify on an alert trigger.
func (o AlertsPtrOutput) Email() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Alerts) []string {
		if v == nil {
			return nil
		}
		return v.Email
	}).(pulumi.StringArrayOutput)
}

// Slack integration details.
func (o AlertsPtrOutput) Slack() SlackDetailsArrayOutput {
	return o.ApplyT(func(v *Alerts) []SlackDetails {
		if v == nil {
			return nil
		}
		return v.Slack
	}).(SlackDetailsArrayOutput)
}

type GetMonitoringAlertPolicyProperties struct {
	Policy *AlertPolicy `pulumi:"policy"`
}

type GetMonitoringAlertPolicyPropertiesOutput struct{ *pulumi.OutputState }

func (GetMonitoringAlertPolicyPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMonitoringAlertPolicyProperties)(nil)).Elem()
}

func (o GetMonitoringAlertPolicyPropertiesOutput) ToGetMonitoringAlertPolicyPropertiesOutput() GetMonitoringAlertPolicyPropertiesOutput {
	return o
}

func (o GetMonitoringAlertPolicyPropertiesOutput) ToGetMonitoringAlertPolicyPropertiesOutputWithContext(ctx context.Context) GetMonitoringAlertPolicyPropertiesOutput {
	return o
}

func (o GetMonitoringAlertPolicyPropertiesOutput) Policy() AlertPolicyPtrOutput {
	return o.ApplyT(func(v GetMonitoringAlertPolicyProperties) *AlertPolicy { return v.Policy }).(AlertPolicyPtrOutput)
}

type ListMonitoringAlertPolicyItems struct {
	Links    *PageLinks    `pulumi:"links"`
	Meta     MetaMeta      `pulumi:"meta"`
	Policies []AlertPolicy `pulumi:"policies"`
}

type ListMonitoringAlertPolicyItemsOutput struct{ *pulumi.OutputState }

func (ListMonitoringAlertPolicyItemsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMonitoringAlertPolicyItems)(nil)).Elem()
}

func (o ListMonitoringAlertPolicyItemsOutput) ToListMonitoringAlertPolicyItemsOutput() ListMonitoringAlertPolicyItemsOutput {
	return o
}

func (o ListMonitoringAlertPolicyItemsOutput) ToListMonitoringAlertPolicyItemsOutputWithContext(ctx context.Context) ListMonitoringAlertPolicyItemsOutput {
	return o
}

func (o ListMonitoringAlertPolicyItemsOutput) Links() PageLinksPtrOutput {
	return o.ApplyT(func(v ListMonitoringAlertPolicyItems) *PageLinks { return v.Links }).(PageLinksPtrOutput)
}

func (o ListMonitoringAlertPolicyItemsOutput) Meta() MetaMetaOutput {
	return o.ApplyT(func(v ListMonitoringAlertPolicyItems) MetaMeta { return v.Meta }).(MetaMetaOutput)
}

func (o ListMonitoringAlertPolicyItemsOutput) Policies() AlertPolicyArrayOutput {
	return o.ApplyT(func(v ListMonitoringAlertPolicyItems) []AlertPolicy { return v.Policies }).(AlertPolicyArrayOutput)
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

type Metrics struct {
	Data   MetricsData   `pulumi:"data"`
	Status MetricsStatus `pulumi:"status"`
}

type MetricsOutput struct{ *pulumi.OutputState }

func (MetricsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Metrics)(nil)).Elem()
}

func (o MetricsOutput) ToMetricsOutput() MetricsOutput {
	return o
}

func (o MetricsOutput) ToMetricsOutputWithContext(ctx context.Context) MetricsOutput {
	return o
}

func (o MetricsOutput) Data() MetricsDataOutput {
	return o.ApplyT(func(v Metrics) MetricsData { return v.Data }).(MetricsDataOutput)
}

func (o MetricsOutput) Status() MetricsStatusOutput {
	return o.ApplyT(func(v Metrics) MetricsStatus { return v.Status }).(MetricsStatusOutput)
}

type MetricsData struct {
	// Result of query.
	Result     []MetricsResult       `pulumi:"result"`
	ResultType MetricsDataResultType `pulumi:"resultType"`
}

type MetricsDataOutput struct{ *pulumi.OutputState }

func (MetricsDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricsData)(nil)).Elem()
}

func (o MetricsDataOutput) ToMetricsDataOutput() MetricsDataOutput {
	return o
}

func (o MetricsDataOutput) ToMetricsDataOutputWithContext(ctx context.Context) MetricsDataOutput {
	return o
}

// Result of query.
func (o MetricsDataOutput) Result() MetricsResultArrayOutput {
	return o.ApplyT(func(v MetricsData) []MetricsResult { return v.Result }).(MetricsResultArrayOutput)
}

func (o MetricsDataOutput) ResultType() MetricsDataResultTypeOutput {
	return o.ApplyT(func(v MetricsData) MetricsDataResultType { return v.ResultType }).(MetricsDataResultTypeOutput)
}

type MetricsResult struct {
	// An object containing the metric labels.
	Metric interface{}     `pulumi:"metric"`
	Values [][]interface{} `pulumi:"values"`
}

type MetricsResultOutput struct{ *pulumi.OutputState }

func (MetricsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricsResult)(nil)).Elem()
}

func (o MetricsResultOutput) ToMetricsResultOutput() MetricsResultOutput {
	return o
}

func (o MetricsResultOutput) ToMetricsResultOutputWithContext(ctx context.Context) MetricsResultOutput {
	return o
}

// An object containing the metric labels.
func (o MetricsResultOutput) Metric() pulumi.AnyOutput {
	return o.ApplyT(func(v MetricsResult) interface{} { return v.Metric }).(pulumi.AnyOutput)
}

func (o MetricsResultOutput) Values() pulumi.ArrayArrayOutput {
	return o.ApplyT(func(v MetricsResult) [][]interface{} { return v.Values }).(pulumi.ArrayArrayOutput)
}

type MetricsResultArrayOutput struct{ *pulumi.OutputState }

func (MetricsResultArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricsResult)(nil)).Elem()
}

func (o MetricsResultArrayOutput) ToMetricsResultArrayOutput() MetricsResultArrayOutput {
	return o
}

func (o MetricsResultArrayOutput) ToMetricsResultArrayOutputWithContext(ctx context.Context) MetricsResultArrayOutput {
	return o
}

func (o MetricsResultArrayOutput) Index(i pulumi.IntInput) MetricsResultOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MetricsResult {
		return vs[0].([]MetricsResult)[vs[1].(int)]
	}).(MetricsResultOutput)
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

type SlackDetails struct {
	// Slack channel to notify of an alert trigger.
	Channel string `pulumi:"channel"`
	// Slack Webhook URL.
	Url string `pulumi:"url"`
}

// SlackDetailsInput is an input type that accepts SlackDetailsArgs and SlackDetailsOutput values.
// You can construct a concrete instance of `SlackDetailsInput` via:
//
//	SlackDetailsArgs{...}
type SlackDetailsInput interface {
	pulumi.Input

	ToSlackDetailsOutput() SlackDetailsOutput
	ToSlackDetailsOutputWithContext(context.Context) SlackDetailsOutput
}

type SlackDetailsArgs struct {
	// Slack channel to notify of an alert trigger.
	Channel pulumi.StringInput `pulumi:"channel"`
	// Slack Webhook URL.
	Url pulumi.StringInput `pulumi:"url"`
}

func (SlackDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SlackDetails)(nil)).Elem()
}

func (i SlackDetailsArgs) ToSlackDetailsOutput() SlackDetailsOutput {
	return i.ToSlackDetailsOutputWithContext(context.Background())
}

func (i SlackDetailsArgs) ToSlackDetailsOutputWithContext(ctx context.Context) SlackDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlackDetailsOutput)
}

// SlackDetailsArrayInput is an input type that accepts SlackDetailsArray and SlackDetailsArrayOutput values.
// You can construct a concrete instance of `SlackDetailsArrayInput` via:
//
//	SlackDetailsArray{ SlackDetailsArgs{...} }
type SlackDetailsArrayInput interface {
	pulumi.Input

	ToSlackDetailsArrayOutput() SlackDetailsArrayOutput
	ToSlackDetailsArrayOutputWithContext(context.Context) SlackDetailsArrayOutput
}

type SlackDetailsArray []SlackDetailsInput

func (SlackDetailsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SlackDetails)(nil)).Elem()
}

func (i SlackDetailsArray) ToSlackDetailsArrayOutput() SlackDetailsArrayOutput {
	return i.ToSlackDetailsArrayOutputWithContext(context.Background())
}

func (i SlackDetailsArray) ToSlackDetailsArrayOutputWithContext(ctx context.Context) SlackDetailsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlackDetailsArrayOutput)
}

type SlackDetailsOutput struct{ *pulumi.OutputState }

func (SlackDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SlackDetails)(nil)).Elem()
}

func (o SlackDetailsOutput) ToSlackDetailsOutput() SlackDetailsOutput {
	return o
}

func (o SlackDetailsOutput) ToSlackDetailsOutputWithContext(ctx context.Context) SlackDetailsOutput {
	return o
}

// Slack channel to notify of an alert trigger.
func (o SlackDetailsOutput) Channel() pulumi.StringOutput {
	return o.ApplyT(func(v SlackDetails) string { return v.Channel }).(pulumi.StringOutput)
}

// Slack Webhook URL.
func (o SlackDetailsOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v SlackDetails) string { return v.Url }).(pulumi.StringOutput)
}

type SlackDetailsArrayOutput struct{ *pulumi.OutputState }

func (SlackDetailsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SlackDetails)(nil)).Elem()
}

func (o SlackDetailsArrayOutput) ToSlackDetailsArrayOutput() SlackDetailsArrayOutput {
	return o
}

func (o SlackDetailsArrayOutput) ToSlackDetailsArrayOutputWithContext(ctx context.Context) SlackDetailsArrayOutput {
	return o
}

func (o SlackDetailsArrayOutput) Index(i pulumi.IntInput) SlackDetailsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SlackDetails {
		return vs[0].([]SlackDetails)[vs[1].(int)]
	}).(SlackDetailsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AlertsInput)(nil)).Elem(), AlertsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SlackDetailsInput)(nil)).Elem(), SlackDetailsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SlackDetailsArrayInput)(nil)).Elem(), SlackDetailsArray{})
	pulumi.RegisterOutputType(AlertPolicyOutput{})
	pulumi.RegisterOutputType(AlertPolicyPtrOutput{})
	pulumi.RegisterOutputType(AlertPolicyArrayOutput{})
	pulumi.RegisterOutputType(AlertsOutput{})
	pulumi.RegisterOutputType(AlertsPtrOutput{})
	pulumi.RegisterOutputType(GetMonitoringAlertPolicyPropertiesOutput{})
	pulumi.RegisterOutputType(ListMonitoringAlertPolicyItemsOutput{})
	pulumi.RegisterOutputType(MetaMetaOutput{})
	pulumi.RegisterOutputType(MetricsOutput{})
	pulumi.RegisterOutputType(MetricsDataOutput{})
	pulumi.RegisterOutputType(MetricsResultOutput{})
	pulumi.RegisterOutputType(MetricsResultArrayOutput{})
	pulumi.RegisterOutputType(PageLinksOutput{})
	pulumi.RegisterOutputType(PageLinksPtrOutput{})
	pulumi.RegisterOutputType(PageLinksPagesPropertiesOutput{})
	pulumi.RegisterOutputType(PageLinksPagesPropertiesPtrOutput{})
	pulumi.RegisterOutputType(SlackDetailsOutput{})
	pulumi.RegisterOutputType(SlackDetailsArrayOutput{})
}
