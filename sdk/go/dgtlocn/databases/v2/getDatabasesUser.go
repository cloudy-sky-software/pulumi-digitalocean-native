// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabasesUser(ctx *pulumi.Context, args *LookupDatabasesUserArgs, opts ...pulumi.InvokeOption) (*LookupDatabasesUserResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDatabasesUserResult
	err := ctx.Invoke("digitalocean-native:databases/v2:getDatabasesUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabasesUserArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid string `pulumi:"databaseClusterUuid"`
	// The name of the database user.
	Username string `pulumi:"username"`
}

type LookupDatabasesUserResult struct {
	User DatabaseUser `pulumi:"user"`
}

func LookupDatabasesUserOutput(ctx *pulumi.Context, args LookupDatabasesUserOutputArgs, opts ...pulumi.InvokeOption) LookupDatabasesUserResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDatabasesUserResultOutput, error) {
			args := v.(LookupDatabasesUserArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("digitalocean-native:databases/v2:getDatabasesUser", args, LookupDatabasesUserResultOutput{}, options).(LookupDatabasesUserResultOutput), nil
		}).(LookupDatabasesUserResultOutput)
}

type LookupDatabasesUserOutputArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid pulumi.StringInput `pulumi:"databaseClusterUuid"`
	// The name of the database user.
	Username pulumi.StringInput `pulumi:"username"`
}

func (LookupDatabasesUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabasesUserArgs)(nil)).Elem()
}

type LookupDatabasesUserResultOutput struct{ *pulumi.OutputState }

func (LookupDatabasesUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabasesUserResult)(nil)).Elem()
}

func (o LookupDatabasesUserResultOutput) ToLookupDatabasesUserResultOutput() LookupDatabasesUserResultOutput {
	return o
}

func (o LookupDatabasesUserResultOutput) ToLookupDatabasesUserResultOutputWithContext(ctx context.Context) LookupDatabasesUserResultOutput {
	return o
}

func (o LookupDatabasesUserResultOutput) User() DatabaseUserOutput {
	return o.ApplyT(func(v LookupDatabasesUserResult) DatabaseUser { return v.User }).(DatabaseUserOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabasesUserResultOutput{})
}
