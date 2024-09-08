// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListFloatingIPsAction(ctx *pulumi.Context, args *ListFloatingIPsActionArgs, opts ...pulumi.InvokeOption) (*ListFloatingIPsActionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListFloatingIPsActionResult
	err := ctx.Invoke("digitalocean-native:floating_ips/v2:listFloatingIPsAction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListFloatingIPsActionArgs struct {
	// A floating IP address.
	FloatingIp string `pulumi:"floatingIp"`
}

type ListFloatingIPsActionResult struct {
	Actions []Action   `pulumi:"actions"`
	Links   *PageLinks `pulumi:"links"`
	Meta    MetaMeta   `pulumi:"meta"`
}

func ListFloatingIPsActionOutput(ctx *pulumi.Context, args ListFloatingIPsActionOutputArgs, opts ...pulumi.InvokeOption) ListFloatingIPsActionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListFloatingIPsActionResult, error) {
			args := v.(ListFloatingIPsActionArgs)
			r, err := ListFloatingIPsAction(ctx, &args, opts...)
			var s ListFloatingIPsActionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListFloatingIPsActionResultOutput)
}

type ListFloatingIPsActionOutputArgs struct {
	// A floating IP address.
	FloatingIp pulumi.StringInput `pulumi:"floatingIp"`
}

func (ListFloatingIPsActionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListFloatingIPsActionArgs)(nil)).Elem()
}

type ListFloatingIPsActionResultOutput struct{ *pulumi.OutputState }

func (ListFloatingIPsActionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListFloatingIPsActionResult)(nil)).Elem()
}

func (o ListFloatingIPsActionResultOutput) ToListFloatingIPsActionResultOutput() ListFloatingIPsActionResultOutput {
	return o
}

func (o ListFloatingIPsActionResultOutput) ToListFloatingIPsActionResultOutputWithContext(ctx context.Context) ListFloatingIPsActionResultOutput {
	return o
}

func (o ListFloatingIPsActionResultOutput) Actions() ActionArrayOutput {
	return o.ApplyT(func(v ListFloatingIPsActionResult) []Action { return v.Actions }).(ActionArrayOutput)
}

func (o ListFloatingIPsActionResultOutput) Links() PageLinksPtrOutput {
	return o.ApplyT(func(v ListFloatingIPsActionResult) *PageLinks { return v.Links }).(PageLinksPtrOutput)
}

func (o ListFloatingIPsActionResultOutput) Meta() MetaMetaOutput {
	return o.ApplyT(func(v ListFloatingIPsActionResult) MetaMeta { return v.Meta }).(MetaMetaOutput)
}

func init() {
	pulumi.RegisterOutputType(ListFloatingIPsActionResultOutput{})
}
