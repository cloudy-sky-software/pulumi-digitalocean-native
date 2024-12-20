// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDatabasesReplicas(ctx *pulumi.Context, args *ListDatabasesReplicasArgs, opts ...pulumi.InvokeOption) (*ListDatabasesReplicasResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListDatabasesReplicasResult
	err := ctx.Invoke("digitalocean-native:databases/v2:listDatabasesReplicas", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDatabasesReplicasArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid string `pulumi:"databaseClusterUuid"`
}

type ListDatabasesReplicasResult struct {
	Replicas []DatabaseReplica `pulumi:"replicas"`
}

func ListDatabasesReplicasOutput(ctx *pulumi.Context, args ListDatabasesReplicasOutputArgs, opts ...pulumi.InvokeOption) ListDatabasesReplicasResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListDatabasesReplicasResultOutput, error) {
			args := v.(ListDatabasesReplicasArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("digitalocean-native:databases/v2:listDatabasesReplicas", args, ListDatabasesReplicasResultOutput{}, options).(ListDatabasesReplicasResultOutput), nil
		}).(ListDatabasesReplicasResultOutput)
}

type ListDatabasesReplicasOutputArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid pulumi.StringInput `pulumi:"databaseClusterUuid"`
}

func (ListDatabasesReplicasOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabasesReplicasArgs)(nil)).Elem()
}

type ListDatabasesReplicasResultOutput struct{ *pulumi.OutputState }

func (ListDatabasesReplicasResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabasesReplicasResult)(nil)).Elem()
}

func (o ListDatabasesReplicasResultOutput) ToListDatabasesReplicasResultOutput() ListDatabasesReplicasResultOutput {
	return o
}

func (o ListDatabasesReplicasResultOutput) ToListDatabasesReplicasResultOutputWithContext(ctx context.Context) ListDatabasesReplicasResultOutput {
	return o
}

func (o ListDatabasesReplicasResultOutput) Replicas() DatabaseReplicaArrayOutput {
	return o.ApplyT(func(v ListDatabasesReplicasResult) []DatabaseReplica { return v.Replicas }).(DatabaseReplicaArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDatabasesReplicasResultOutput{})
}
