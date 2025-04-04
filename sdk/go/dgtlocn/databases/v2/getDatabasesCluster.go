// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabasesCluster(ctx *pulumi.Context, args *LookupDatabasesClusterArgs, opts ...pulumi.InvokeOption) (*LookupDatabasesClusterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDatabasesClusterResult
	err := ctx.Invoke("digitalocean-native:databases/v2:getDatabasesCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabasesClusterArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid string `pulumi:"databaseClusterUuid"`
}

type LookupDatabasesClusterResult struct {
	Database DatabaseCluster `pulumi:"database"`
}

func LookupDatabasesClusterOutput(ctx *pulumi.Context, args LookupDatabasesClusterOutputArgs, opts ...pulumi.InvokeOption) LookupDatabasesClusterResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDatabasesClusterResultOutput, error) {
			args := v.(LookupDatabasesClusterArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("digitalocean-native:databases/v2:getDatabasesCluster", args, LookupDatabasesClusterResultOutput{}, options).(LookupDatabasesClusterResultOutput), nil
		}).(LookupDatabasesClusterResultOutput)
}

type LookupDatabasesClusterOutputArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid pulumi.StringInput `pulumi:"databaseClusterUuid"`
}

func (LookupDatabasesClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabasesClusterArgs)(nil)).Elem()
}

type LookupDatabasesClusterResultOutput struct{ *pulumi.OutputState }

func (LookupDatabasesClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabasesClusterResult)(nil)).Elem()
}

func (o LookupDatabasesClusterResultOutput) ToLookupDatabasesClusterResultOutput() LookupDatabasesClusterResultOutput {
	return o
}

func (o LookupDatabasesClusterResultOutput) ToLookupDatabasesClusterResultOutputWithContext(ctx context.Context) LookupDatabasesClusterResultOutput {
	return o
}

func (o LookupDatabasesClusterResultOutput) Database() DatabaseClusterOutput {
	return o.ApplyT(func(v LookupDatabasesClusterResult) DatabaseCluster { return v.Database }).(DatabaseClusterOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabasesClusterResultOutput{})
}
