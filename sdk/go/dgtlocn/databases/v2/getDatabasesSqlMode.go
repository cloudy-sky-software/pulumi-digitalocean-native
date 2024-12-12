// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabasesSqlMode(ctx *pulumi.Context, args *LookupDatabasesSqlModeArgs, opts ...pulumi.InvokeOption) (*LookupDatabasesSqlModeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDatabasesSqlModeResult
	err := ctx.Invoke("digitalocean-native:databases/v2:getDatabasesSqlMode", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabasesSqlModeArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid string `pulumi:"databaseClusterUuid"`
}

type LookupDatabasesSqlModeResult struct {
	// A string specifying the configured SQL modes for the MySQL cluster.
	SqlMode string `pulumi:"sqlMode"`
}

func LookupDatabasesSqlModeOutput(ctx *pulumi.Context, args LookupDatabasesSqlModeOutputArgs, opts ...pulumi.InvokeOption) LookupDatabasesSqlModeResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDatabasesSqlModeResultOutput, error) {
			args := v.(LookupDatabasesSqlModeArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("digitalocean-native:databases/v2:getDatabasesSqlMode", args, LookupDatabasesSqlModeResultOutput{}, options).(LookupDatabasesSqlModeResultOutput), nil
		}).(LookupDatabasesSqlModeResultOutput)
}

type LookupDatabasesSqlModeOutputArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid pulumi.StringInput `pulumi:"databaseClusterUuid"`
}

func (LookupDatabasesSqlModeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabasesSqlModeArgs)(nil)).Elem()
}

type LookupDatabasesSqlModeResultOutput struct{ *pulumi.OutputState }

func (LookupDatabasesSqlModeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabasesSqlModeResult)(nil)).Elem()
}

func (o LookupDatabasesSqlModeResultOutput) ToLookupDatabasesSqlModeResultOutput() LookupDatabasesSqlModeResultOutput {
	return o
}

func (o LookupDatabasesSqlModeResultOutput) ToLookupDatabasesSqlModeResultOutputWithContext(ctx context.Context) LookupDatabasesSqlModeResultOutput {
	return o
}

// A string specifying the configured SQL modes for the MySQL cluster.
func (o LookupDatabasesSqlModeResultOutput) SqlMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasesSqlModeResult) string { return v.SqlMode }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabasesSqlModeResultOutput{})
}
