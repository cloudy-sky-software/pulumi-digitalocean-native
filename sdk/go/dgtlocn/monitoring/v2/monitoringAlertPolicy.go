// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"errors"
	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MonitoringAlertPolicy struct {
	pulumi.CustomResourceState

	Alerts      AlertsOutput             `pulumi:"alerts"`
	Compare     CompareOutput            `pulumi:"compare"`
	Description pulumi.StringOutput      `pulumi:"description"`
	Enabled     pulumi.BoolOutput        `pulumi:"enabled"`
	Entities    pulumi.StringArrayOutput `pulumi:"entities"`
	Policy      AlertPolicyPtrOutput     `pulumi:"policy"`
	Tags        pulumi.StringArrayOutput `pulumi:"tags"`
	Type        TypeOutput               `pulumi:"type"`
	Value       pulumi.Float64Output     `pulumi:"value"`
	Window      WindowOutput             `pulumi:"window"`
}

// NewMonitoringAlertPolicy registers a new resource with the given unique name, arguments, and options.
func NewMonitoringAlertPolicy(ctx *pulumi.Context,
	name string, args *MonitoringAlertPolicyArgs, opts ...pulumi.ResourceOption) (*MonitoringAlertPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Alerts == nil {
		return nil, errors.New("invalid value for required argument 'Alerts'")
	}
	if args.Compare == nil {
		return nil, errors.New("invalid value for required argument 'Compare'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.Entities == nil {
		return nil, errors.New("invalid value for required argument 'Entities'")
	}
	if args.Tags == nil {
		return nil, errors.New("invalid value for required argument 'Tags'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	if args.Window == nil {
		return nil, errors.New("invalid value for required argument 'Window'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MonitoringAlertPolicy
	err := ctx.RegisterResource("digitalocean-native:monitoring/v2:MonitoringAlertPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMonitoringAlertPolicy gets an existing MonitoringAlertPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMonitoringAlertPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MonitoringAlertPolicyState, opts ...pulumi.ResourceOption) (*MonitoringAlertPolicy, error) {
	var resource MonitoringAlertPolicy
	err := ctx.ReadResource("digitalocean-native:monitoring/v2:MonitoringAlertPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MonitoringAlertPolicy resources.
type monitoringAlertPolicyState struct {
}

type MonitoringAlertPolicyState struct {
}

func (MonitoringAlertPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*monitoringAlertPolicyState)(nil)).Elem()
}

type monitoringAlertPolicyArgs struct {
	Alerts      Alerts   `pulumi:"alerts"`
	Compare     Compare  `pulumi:"compare"`
	Description string   `pulumi:"description"`
	Enabled     bool     `pulumi:"enabled"`
	Entities    []string `pulumi:"entities"`
	Tags        []string `pulumi:"tags"`
	Type        Type     `pulumi:"type"`
	Value       float64  `pulumi:"value"`
	Window      Window   `pulumi:"window"`
}

// The set of arguments for constructing a MonitoringAlertPolicy resource.
type MonitoringAlertPolicyArgs struct {
	Alerts      AlertsInput
	Compare     CompareInput
	Description pulumi.StringInput
	Enabled     pulumi.BoolInput
	Entities    pulumi.StringArrayInput
	Tags        pulumi.StringArrayInput
	Type        TypeInput
	Value       pulumi.Float64Input
	Window      WindowInput
}

func (MonitoringAlertPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*monitoringAlertPolicyArgs)(nil)).Elem()
}

type MonitoringAlertPolicyInput interface {
	pulumi.Input

	ToMonitoringAlertPolicyOutput() MonitoringAlertPolicyOutput
	ToMonitoringAlertPolicyOutputWithContext(ctx context.Context) MonitoringAlertPolicyOutput
}

func (*MonitoringAlertPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringAlertPolicy)(nil)).Elem()
}

func (i *MonitoringAlertPolicy) ToMonitoringAlertPolicyOutput() MonitoringAlertPolicyOutput {
	return i.ToMonitoringAlertPolicyOutputWithContext(context.Background())
}

func (i *MonitoringAlertPolicy) ToMonitoringAlertPolicyOutputWithContext(ctx context.Context) MonitoringAlertPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringAlertPolicyOutput)
}

type MonitoringAlertPolicyOutput struct{ *pulumi.OutputState }

func (MonitoringAlertPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringAlertPolicy)(nil)).Elem()
}

func (o MonitoringAlertPolicyOutput) ToMonitoringAlertPolicyOutput() MonitoringAlertPolicyOutput {
	return o
}

func (o MonitoringAlertPolicyOutput) ToMonitoringAlertPolicyOutputWithContext(ctx context.Context) MonitoringAlertPolicyOutput {
	return o
}

func (o MonitoringAlertPolicyOutput) Alerts() AlertsOutput {
	return o.ApplyT(func(v *MonitoringAlertPolicy) AlertsOutput { return v.Alerts }).(AlertsOutput)
}

func (o MonitoringAlertPolicyOutput) Compare() CompareOutput {
	return o.ApplyT(func(v *MonitoringAlertPolicy) CompareOutput { return v.Compare }).(CompareOutput)
}

func (o MonitoringAlertPolicyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringAlertPolicy) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o MonitoringAlertPolicyOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *MonitoringAlertPolicy) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

func (o MonitoringAlertPolicyOutput) Entities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MonitoringAlertPolicy) pulumi.StringArrayOutput { return v.Entities }).(pulumi.StringArrayOutput)
}

func (o MonitoringAlertPolicyOutput) Policy() AlertPolicyPtrOutput {
	return o.ApplyT(func(v *MonitoringAlertPolicy) AlertPolicyPtrOutput { return v.Policy }).(AlertPolicyPtrOutput)
}

func (o MonitoringAlertPolicyOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MonitoringAlertPolicy) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o MonitoringAlertPolicyOutput) Type() TypeOutput {
	return o.ApplyT(func(v *MonitoringAlertPolicy) TypeOutput { return v.Type }).(TypeOutput)
}

func (o MonitoringAlertPolicyOutput) Value() pulumi.Float64Output {
	return o.ApplyT(func(v *MonitoringAlertPolicy) pulumi.Float64Output { return v.Value }).(pulumi.Float64Output)
}

func (o MonitoringAlertPolicyOutput) Window() WindowOutput {
	return o.ApplyT(func(v *MonitoringAlertPolicy) WindowOutput { return v.Window }).(WindowOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MonitoringAlertPolicyInput)(nil)).Elem(), &MonitoringAlertPolicy{})
	pulumi.RegisterOutputType(MonitoringAlertPolicyOutput{})
}