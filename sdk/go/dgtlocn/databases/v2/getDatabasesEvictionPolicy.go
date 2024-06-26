// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabasesEvictionPolicy(ctx *pulumi.Context, args *LookupDatabasesEvictionPolicyArgs, opts ...pulumi.InvokeOption) (*LookupDatabasesEvictionPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDatabasesEvictionPolicyResult
	err := ctx.Invoke("digitalocean-native:databases/v2:getDatabasesEvictionPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabasesEvictionPolicyArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid string `pulumi:"databaseClusterUuid"`
}

type LookupDatabasesEvictionPolicyResult struct {
	Items GetDatabasesEvictionPolicyProperties `pulumi:"items"`
}

func LookupDatabasesEvictionPolicyOutput(ctx *pulumi.Context, args LookupDatabasesEvictionPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupDatabasesEvictionPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabasesEvictionPolicyResult, error) {
			args := v.(LookupDatabasesEvictionPolicyArgs)
			r, err := LookupDatabasesEvictionPolicy(ctx, &args, opts...)
			var s LookupDatabasesEvictionPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabasesEvictionPolicyResultOutput)
}

type LookupDatabasesEvictionPolicyOutputArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid pulumi.StringInput `pulumi:"databaseClusterUuid"`
}

func (LookupDatabasesEvictionPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabasesEvictionPolicyArgs)(nil)).Elem()
}

type LookupDatabasesEvictionPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupDatabasesEvictionPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabasesEvictionPolicyResult)(nil)).Elem()
}

func (o LookupDatabasesEvictionPolicyResultOutput) ToLookupDatabasesEvictionPolicyResultOutput() LookupDatabasesEvictionPolicyResultOutput {
	return o
}

func (o LookupDatabasesEvictionPolicyResultOutput) ToLookupDatabasesEvictionPolicyResultOutputWithContext(ctx context.Context) LookupDatabasesEvictionPolicyResultOutput {
	return o
}

func (o LookupDatabasesEvictionPolicyResultOutput) Items() GetDatabasesEvictionPolicyPropertiesOutput {
	return o.ApplyT(func(v LookupDatabasesEvictionPolicyResult) GetDatabasesEvictionPolicyProperties { return v.Items }).(GetDatabasesEvictionPolicyPropertiesOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabasesEvictionPolicyResultOutput{})
}
