// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetKubernetesClusterLintResult(ctx *pulumi.Context, args *GetKubernetesClusterLintResultArgs, opts ...pulumi.InvokeOption) (*GetKubernetesClusterLintResultResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetKubernetesClusterLintResultResult
	err := ctx.Invoke("digitalocean-native:kubernetes/v2:getKubernetesClusterLintResult", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetKubernetesClusterLintResultArgs struct {
	// A unique ID that can be used to reference a Kubernetes cluster.
	ClusterId string `pulumi:"clusterId"`
}

type GetKubernetesClusterLintResultResult struct {
	// A time value given in ISO8601 combined date and time format that represents when the schedule clusterlint run request was completed.
	CompletedAt *string `pulumi:"completedAt"`
	// An array of diagnostics reporting potential problems for the given cluster.
	Diagnostics []ClusterlintResultsDiagnosticsItemProperties `pulumi:"diagnostics"`
	// A time value given in ISO8601 combined date and time format that represents when the schedule clusterlint run request was made.
	RequestedAt *string `pulumi:"requestedAt"`
	// Id of the clusterlint run that can be used later to fetch the diagnostics.
	RunId *string `pulumi:"runId"`
}

func GetKubernetesClusterLintResultOutput(ctx *pulumi.Context, args GetKubernetesClusterLintResultOutputArgs, opts ...pulumi.InvokeOption) GetKubernetesClusterLintResultResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetKubernetesClusterLintResultResult, error) {
			args := v.(GetKubernetesClusterLintResultArgs)
			r, err := GetKubernetesClusterLintResult(ctx, &args, opts...)
			var s GetKubernetesClusterLintResultResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetKubernetesClusterLintResultResultOutput)
}

type GetKubernetesClusterLintResultOutputArgs struct {
	// A unique ID that can be used to reference a Kubernetes cluster.
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
}

func (GetKubernetesClusterLintResultOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKubernetesClusterLintResultArgs)(nil)).Elem()
}

type GetKubernetesClusterLintResultResultOutput struct{ *pulumi.OutputState }

func (GetKubernetesClusterLintResultResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKubernetesClusterLintResultResult)(nil)).Elem()
}

func (o GetKubernetesClusterLintResultResultOutput) ToGetKubernetesClusterLintResultResultOutput() GetKubernetesClusterLintResultResultOutput {
	return o
}

func (o GetKubernetesClusterLintResultResultOutput) ToGetKubernetesClusterLintResultResultOutputWithContext(ctx context.Context) GetKubernetesClusterLintResultResultOutput {
	return o
}

// A time value given in ISO8601 combined date and time format that represents when the schedule clusterlint run request was completed.
func (o GetKubernetesClusterLintResultResultOutput) CompletedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKubernetesClusterLintResultResult) *string { return v.CompletedAt }).(pulumi.StringPtrOutput)
}

// An array of diagnostics reporting potential problems for the given cluster.
func (o GetKubernetesClusterLintResultResultOutput) Diagnostics() ClusterlintResultsDiagnosticsItemPropertiesArrayOutput {
	return o.ApplyT(func(v GetKubernetesClusterLintResultResult) []ClusterlintResultsDiagnosticsItemProperties {
		return v.Diagnostics
	}).(ClusterlintResultsDiagnosticsItemPropertiesArrayOutput)
}

// A time value given in ISO8601 combined date and time format that represents when the schedule clusterlint run request was made.
func (o GetKubernetesClusterLintResultResultOutput) RequestedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKubernetesClusterLintResultResult) *string { return v.RequestedAt }).(pulumi.StringPtrOutput)
}

// Id of the clusterlint run that can be used later to fetch the diagnostics.
func (o GetKubernetesClusterLintResultResultOutput) RunId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKubernetesClusterLintResultResult) *string { return v.RunId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetKubernetesClusterLintResultResultOutput{})
}
