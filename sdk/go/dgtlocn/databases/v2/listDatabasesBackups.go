// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDatabasesBackups(ctx *pulumi.Context, args *ListDatabasesBackupsArgs, opts ...pulumi.InvokeOption) (*ListDatabasesBackupsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListDatabasesBackupsResult
	err := ctx.Invoke("digitalocean-native:databases/v2:listDatabasesBackups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDatabasesBackupsArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid string `pulumi:"databaseClusterUuid"`
}

type ListDatabasesBackupsResult struct {
	Backups []Backup `pulumi:"backups"`
}

func ListDatabasesBackupsOutput(ctx *pulumi.Context, args ListDatabasesBackupsOutputArgs, opts ...pulumi.InvokeOption) ListDatabasesBackupsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListDatabasesBackupsResultOutput, error) {
			args := v.(ListDatabasesBackupsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("digitalocean-native:databases/v2:listDatabasesBackups", args, ListDatabasesBackupsResultOutput{}, options).(ListDatabasesBackupsResultOutput), nil
		}).(ListDatabasesBackupsResultOutput)
}

type ListDatabasesBackupsOutputArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid pulumi.StringInput `pulumi:"databaseClusterUuid"`
}

func (ListDatabasesBackupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabasesBackupsArgs)(nil)).Elem()
}

type ListDatabasesBackupsResultOutput struct{ *pulumi.OutputState }

func (ListDatabasesBackupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabasesBackupsResult)(nil)).Elem()
}

func (o ListDatabasesBackupsResultOutput) ToListDatabasesBackupsResultOutput() ListDatabasesBackupsResultOutput {
	return o
}

func (o ListDatabasesBackupsResultOutput) ToListDatabasesBackupsResultOutputWithContext(ctx context.Context) ListDatabasesBackupsResultOutput {
	return o
}

func (o ListDatabasesBackupsResultOutput) Backups() BackupArrayOutput {
	return o.ApplyT(func(v ListDatabasesBackupsResult) []Backup { return v.Backups }).(BackupArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDatabasesBackupsResultOutput{})
}
