// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDatabasesUsers(ctx *pulumi.Context, args *ListDatabasesUsersArgs, opts ...pulumi.InvokeOption) (*ListDatabasesUsersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListDatabasesUsersResult
	err := ctx.Invoke("digitalocean-native:databases/v2:listDatabasesUsers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDatabasesUsersArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid string `pulumi:"databaseClusterUuid"`
}

type ListDatabasesUsersResult struct {
	Users []DatabaseUser `pulumi:"users"`
}

func ListDatabasesUsersOutput(ctx *pulumi.Context, args ListDatabasesUsersOutputArgs, opts ...pulumi.InvokeOption) ListDatabasesUsersResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListDatabasesUsersResultOutput, error) {
			args := v.(ListDatabasesUsersArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("digitalocean-native:databases/v2:listDatabasesUsers", args, ListDatabasesUsersResultOutput{}, options).(ListDatabasesUsersResultOutput), nil
		}).(ListDatabasesUsersResultOutput)
}

type ListDatabasesUsersOutputArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid pulumi.StringInput `pulumi:"databaseClusterUuid"`
}

func (ListDatabasesUsersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabasesUsersArgs)(nil)).Elem()
}

type ListDatabasesUsersResultOutput struct{ *pulumi.OutputState }

func (ListDatabasesUsersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabasesUsersResult)(nil)).Elem()
}

func (o ListDatabasesUsersResultOutput) ToListDatabasesUsersResultOutput() ListDatabasesUsersResultOutput {
	return o
}

func (o ListDatabasesUsersResultOutput) ToListDatabasesUsersResultOutputWithContext(ctx context.Context) ListDatabasesUsersResultOutput {
	return o
}

func (o ListDatabasesUsersResultOutput) Users() DatabaseUserArrayOutput {
	return o.ApplyT(func(v ListDatabasesUsersResult) []DatabaseUser { return v.Users }).(DatabaseUserArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDatabasesUsersResultOutput{})
}
