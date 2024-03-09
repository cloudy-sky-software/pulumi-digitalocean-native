// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AppsAssignAlertDestinations struct {
	pulumi.CustomResourceState

	Alert         AppAlertPtrOutput               `pulumi:"alert"`
	Emails        pulumi.StringArrayOutput        `pulumi:"emails"`
	SlackWebhooks AppAlertSlackWebhookArrayOutput `pulumi:"slackWebhooks"`
}

// NewAppsAssignAlertDestinations registers a new resource with the given unique name, arguments, and options.
func NewAppsAssignAlertDestinations(ctx *pulumi.Context,
	name string, args *AppsAssignAlertDestinationsArgs, opts ...pulumi.ResourceOption) (*AppsAssignAlertDestinations, error) {
	if args == nil {
		args = &AppsAssignAlertDestinationsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AppsAssignAlertDestinations
	err := ctx.RegisterResource("digitalocean-native:apps/v2:AppsAssignAlertDestinations", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppsAssignAlertDestinations gets an existing AppsAssignAlertDestinations resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppsAssignAlertDestinations(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppsAssignAlertDestinationsState, opts ...pulumi.ResourceOption) (*AppsAssignAlertDestinations, error) {
	var resource AppsAssignAlertDestinations
	err := ctx.ReadResource("digitalocean-native:apps/v2:AppsAssignAlertDestinations", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppsAssignAlertDestinations resources.
type appsAssignAlertDestinationsState struct {
}

type AppsAssignAlertDestinationsState struct {
}

func (AppsAssignAlertDestinationsState) ElementType() reflect.Type {
	return reflect.TypeOf((*appsAssignAlertDestinationsState)(nil)).Elem()
}

type appsAssignAlertDestinationsArgs struct {
	// The alert ID
	AlertId *string `pulumi:"alertId"`
	// The app ID
	AppId         *string                `pulumi:"appId"`
	Emails        []string               `pulumi:"emails"`
	SlackWebhooks []AppAlertSlackWebhook `pulumi:"slackWebhooks"`
}

// The set of arguments for constructing a AppsAssignAlertDestinations resource.
type AppsAssignAlertDestinationsArgs struct {
	// The alert ID
	AlertId pulumi.StringPtrInput
	// The app ID
	AppId         pulumi.StringPtrInput
	Emails        pulumi.StringArrayInput
	SlackWebhooks AppAlertSlackWebhookArrayInput
}

func (AppsAssignAlertDestinationsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appsAssignAlertDestinationsArgs)(nil)).Elem()
}

type AppsAssignAlertDestinationsInput interface {
	pulumi.Input

	ToAppsAssignAlertDestinationsOutput() AppsAssignAlertDestinationsOutput
	ToAppsAssignAlertDestinationsOutputWithContext(ctx context.Context) AppsAssignAlertDestinationsOutput
}

func (*AppsAssignAlertDestinations) ElementType() reflect.Type {
	return reflect.TypeOf((**AppsAssignAlertDestinations)(nil)).Elem()
}

func (i *AppsAssignAlertDestinations) ToAppsAssignAlertDestinationsOutput() AppsAssignAlertDestinationsOutput {
	return i.ToAppsAssignAlertDestinationsOutputWithContext(context.Background())
}

func (i *AppsAssignAlertDestinations) ToAppsAssignAlertDestinationsOutputWithContext(ctx context.Context) AppsAssignAlertDestinationsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppsAssignAlertDestinationsOutput)
}

type AppsAssignAlertDestinationsOutput struct{ *pulumi.OutputState }

func (AppsAssignAlertDestinationsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppsAssignAlertDestinations)(nil)).Elem()
}

func (o AppsAssignAlertDestinationsOutput) ToAppsAssignAlertDestinationsOutput() AppsAssignAlertDestinationsOutput {
	return o
}

func (o AppsAssignAlertDestinationsOutput) ToAppsAssignAlertDestinationsOutputWithContext(ctx context.Context) AppsAssignAlertDestinationsOutput {
	return o
}

func (o AppsAssignAlertDestinationsOutput) Alert() AppAlertPtrOutput {
	return o.ApplyT(func(v *AppsAssignAlertDestinations) AppAlertPtrOutput { return v.Alert }).(AppAlertPtrOutput)
}

func (o AppsAssignAlertDestinationsOutput) Emails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AppsAssignAlertDestinations) pulumi.StringArrayOutput { return v.Emails }).(pulumi.StringArrayOutput)
}

func (o AppsAssignAlertDestinationsOutput) SlackWebhooks() AppAlertSlackWebhookArrayOutput {
	return o.ApplyT(func(v *AppsAssignAlertDestinations) AppAlertSlackWebhookArrayOutput { return v.SlackWebhooks }).(AppAlertSlackWebhookArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppsAssignAlertDestinationsInput)(nil)).Elem(), &AppsAssignAlertDestinations{})
	pulumi.RegisterOutputType(AppsAssignAlertDestinationsOutput{})
}