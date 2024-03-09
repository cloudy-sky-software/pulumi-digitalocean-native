// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Databases struct {
	pulumi.CustomResourceState

	Db DatabaseOutput `pulumi:"db"`
	// The name of the database.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewDatabases registers a new resource with the given unique name, arguments, and options.
func NewDatabases(ctx *pulumi.Context,
	name string, args *DatabasesArgs, opts ...pulumi.ResourceOption) (*Databases, error) {
	if args == nil {
		args = &DatabasesArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Databases
	err := ctx.RegisterResource("digitalocean-native:databases/v2:Databases", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabases gets an existing Databases resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabases(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabasesState, opts ...pulumi.ResourceOption) (*Databases, error) {
	var resource Databases
	err := ctx.ReadResource("digitalocean-native:databases/v2:Databases", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Databases resources.
type databasesState struct {
}

type DatabasesState struct {
}

func (DatabasesState) ElementType() reflect.Type {
	return reflect.TypeOf((*databasesState)(nil)).Elem()
}

type databasesArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid *string `pulumi:"databaseClusterUuid"`
	// The name of the database.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Databases resource.
type DatabasesArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid pulumi.StringPtrInput
	// The name of the database.
	Name pulumi.StringPtrInput
}

func (DatabasesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databasesArgs)(nil)).Elem()
}

type DatabasesInput interface {
	pulumi.Input

	ToDatabasesOutput() DatabasesOutput
	ToDatabasesOutputWithContext(ctx context.Context) DatabasesOutput
}

func (*Databases) ElementType() reflect.Type {
	return reflect.TypeOf((**Databases)(nil)).Elem()
}

func (i *Databases) ToDatabasesOutput() DatabasesOutput {
	return i.ToDatabasesOutputWithContext(context.Background())
}

func (i *Databases) ToDatabasesOutputWithContext(ctx context.Context) DatabasesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasesOutput)
}

type DatabasesOutput struct{ *pulumi.OutputState }

func (DatabasesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Databases)(nil)).Elem()
}

func (o DatabasesOutput) ToDatabasesOutput() DatabasesOutput {
	return o
}

func (o DatabasesOutput) ToDatabasesOutputWithContext(ctx context.Context) DatabasesOutput {
	return o
}

func (o DatabasesOutput) Db() DatabaseOutput {
	return o.ApplyT(func(v *Databases) DatabaseOutput { return v.Db }).(DatabaseOutput)
}

// The name of the database.
func (o DatabasesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Databases) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabasesInput)(nil)).Elem(), &Databases{})
	pulumi.RegisterOutputType(DatabasesOutput{})
}
