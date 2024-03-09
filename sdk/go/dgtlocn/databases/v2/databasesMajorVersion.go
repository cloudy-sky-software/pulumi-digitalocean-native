// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DatabasesMajorVersion struct {
	pulumi.CustomResourceState

	// A string representing the version of the database engine in use for the cluster.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewDatabasesMajorVersion registers a new resource with the given unique name, arguments, and options.
func NewDatabasesMajorVersion(ctx *pulumi.Context,
	name string, args *DatabasesMajorVersionArgs, opts ...pulumi.ResourceOption) (*DatabasesMajorVersion, error) {
	if args == nil {
		args = &DatabasesMajorVersionArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DatabasesMajorVersion
	err := ctx.RegisterResource("digitalocean-native:databases/v2:DatabasesMajorVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabasesMajorVersion gets an existing DatabasesMajorVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabasesMajorVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabasesMajorVersionState, opts ...pulumi.ResourceOption) (*DatabasesMajorVersion, error) {
	var resource DatabasesMajorVersion
	err := ctx.ReadResource("digitalocean-native:databases/v2:DatabasesMajorVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabasesMajorVersion resources.
type databasesMajorVersionState struct {
}

type DatabasesMajorVersionState struct {
}

func (DatabasesMajorVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*databasesMajorVersionState)(nil)).Elem()
}

type databasesMajorVersionArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid *string `pulumi:"databaseClusterUuid"`
	// A string representing the version of the database engine in use for the cluster.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a DatabasesMajorVersion resource.
type DatabasesMajorVersionArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid pulumi.StringPtrInput
	// A string representing the version of the database engine in use for the cluster.
	Version pulumi.StringPtrInput
}

func (DatabasesMajorVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databasesMajorVersionArgs)(nil)).Elem()
}

type DatabasesMajorVersionInput interface {
	pulumi.Input

	ToDatabasesMajorVersionOutput() DatabasesMajorVersionOutput
	ToDatabasesMajorVersionOutputWithContext(ctx context.Context) DatabasesMajorVersionOutput
}

func (*DatabasesMajorVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabasesMajorVersion)(nil)).Elem()
}

func (i *DatabasesMajorVersion) ToDatabasesMajorVersionOutput() DatabasesMajorVersionOutput {
	return i.ToDatabasesMajorVersionOutputWithContext(context.Background())
}

func (i *DatabasesMajorVersion) ToDatabasesMajorVersionOutputWithContext(ctx context.Context) DatabasesMajorVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasesMajorVersionOutput)
}

type DatabasesMajorVersionOutput struct{ *pulumi.OutputState }

func (DatabasesMajorVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabasesMajorVersion)(nil)).Elem()
}

func (o DatabasesMajorVersionOutput) ToDatabasesMajorVersionOutput() DatabasesMajorVersionOutput {
	return o
}

func (o DatabasesMajorVersionOutput) ToDatabasesMajorVersionOutputWithContext(ctx context.Context) DatabasesMajorVersionOutput {
	return o
}

// A string representing the version of the database engine in use for the cluster.
func (o DatabasesMajorVersionOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabasesMajorVersion) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabasesMajorVersionInput)(nil)).Elem(), &DatabasesMajorVersion{})
	pulumi.RegisterOutputType(DatabasesMajorVersionOutput{})
}
