// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"errors"
	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DatabasesSqlMode struct {
	pulumi.CustomResourceState

	// A string specifying the configured SQL modes for the MySQL cluster.
	SqlMode pulumi.StringOutput `pulumi:"sqlMode"`
}

// NewDatabasesSqlMode registers a new resource with the given unique name, arguments, and options.
func NewDatabasesSqlMode(ctx *pulumi.Context,
	name string, args *DatabasesSqlModeArgs, opts ...pulumi.ResourceOption) (*DatabasesSqlMode, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SqlMode == nil {
		return nil, errors.New("invalid value for required argument 'SqlMode'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DatabasesSqlMode
	err := ctx.RegisterResource("digitalocean-native:databases/v2:DatabasesSqlMode", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabasesSqlMode gets an existing DatabasesSqlMode resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabasesSqlMode(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabasesSqlModeState, opts ...pulumi.ResourceOption) (*DatabasesSqlMode, error) {
	var resource DatabasesSqlMode
	err := ctx.ReadResource("digitalocean-native:databases/v2:DatabasesSqlMode", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabasesSqlMode resources.
type databasesSqlModeState struct {
}

type DatabasesSqlModeState struct {
}

func (DatabasesSqlModeState) ElementType() reflect.Type {
	return reflect.TypeOf((*databasesSqlModeState)(nil)).Elem()
}

type databasesSqlModeArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid *string `pulumi:"databaseClusterUuid"`
	// A string specifying the configured SQL modes for the MySQL cluster.
	SqlMode string `pulumi:"sqlMode"`
}

// The set of arguments for constructing a DatabasesSqlMode resource.
type DatabasesSqlModeArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid pulumi.StringPtrInput
	// A string specifying the configured SQL modes for the MySQL cluster.
	SqlMode pulumi.StringInput
}

func (DatabasesSqlModeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databasesSqlModeArgs)(nil)).Elem()
}

type DatabasesSqlModeInput interface {
	pulumi.Input

	ToDatabasesSqlModeOutput() DatabasesSqlModeOutput
	ToDatabasesSqlModeOutputWithContext(ctx context.Context) DatabasesSqlModeOutput
}

func (*DatabasesSqlMode) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabasesSqlMode)(nil)).Elem()
}

func (i *DatabasesSqlMode) ToDatabasesSqlModeOutput() DatabasesSqlModeOutput {
	return i.ToDatabasesSqlModeOutputWithContext(context.Background())
}

func (i *DatabasesSqlMode) ToDatabasesSqlModeOutputWithContext(ctx context.Context) DatabasesSqlModeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasesSqlModeOutput)
}

type DatabasesSqlModeOutput struct{ *pulumi.OutputState }

func (DatabasesSqlModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabasesSqlMode)(nil)).Elem()
}

func (o DatabasesSqlModeOutput) ToDatabasesSqlModeOutput() DatabasesSqlModeOutput {
	return o
}

func (o DatabasesSqlModeOutput) ToDatabasesSqlModeOutputWithContext(ctx context.Context) DatabasesSqlModeOutput {
	return o
}

// A string specifying the configured SQL modes for the MySQL cluster.
func (o DatabasesSqlModeOutput) SqlMode() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabasesSqlMode) pulumi.StringOutput { return v.SqlMode }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabasesSqlModeInput)(nil)).Elem(), &DatabasesSqlMode{})
	pulumi.RegisterOutputType(DatabasesSqlModeOutput{})
}
