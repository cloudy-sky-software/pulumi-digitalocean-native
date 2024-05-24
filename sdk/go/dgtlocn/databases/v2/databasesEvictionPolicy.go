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

type DatabasesEvictionPolicy struct {
	pulumi.CustomResourceState

	// A string specifying the desired eviction policy for the Redis cluster.
	//
	// - `noeviction`: Don't evict any data, returns error when memory limit is reached.
	// - `allkeys_lru:` Evict any key, least recently used (LRU) first.
	// - `allkeys_random`: Evict keys in a random order.
	// - `volatile_lru`: Evict keys with expiration only, least recently used (LRU) first.
	// - `volatile_random`: Evict keys with expiration only in a random order.
	// - `volatile_ttl`: Evict keys with expiration only, shortest time-to-live (TTL) first.
	EvictionPolicy DatabasesEvictionPolicyEvictionPolicyOutput `pulumi:"evictionPolicy"`
}

// NewDatabasesEvictionPolicy registers a new resource with the given unique name, arguments, and options.
func NewDatabasesEvictionPolicy(ctx *pulumi.Context,
	name string, args *DatabasesEvictionPolicyArgs, opts ...pulumi.ResourceOption) (*DatabasesEvictionPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EvictionPolicy == nil {
		return nil, errors.New("invalid value for required argument 'EvictionPolicy'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DatabasesEvictionPolicy
	err := ctx.RegisterResource("digitalocean-native:databases/v2:DatabasesEvictionPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabasesEvictionPolicy gets an existing DatabasesEvictionPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabasesEvictionPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabasesEvictionPolicyState, opts ...pulumi.ResourceOption) (*DatabasesEvictionPolicy, error) {
	var resource DatabasesEvictionPolicy
	err := ctx.ReadResource("digitalocean-native:databases/v2:DatabasesEvictionPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabasesEvictionPolicy resources.
type databasesEvictionPolicyState struct {
}

type DatabasesEvictionPolicyState struct {
}

func (DatabasesEvictionPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*databasesEvictionPolicyState)(nil)).Elem()
}

type databasesEvictionPolicyArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid *string `pulumi:"databaseClusterUuid"`
	// A string specifying the desired eviction policy for the Redis cluster.
	//
	// - `noeviction`: Don't evict any data, returns error when memory limit is reached.
	// - `allkeys_lru:` Evict any key, least recently used (LRU) first.
	// - `allkeys_random`: Evict keys in a random order.
	// - `volatile_lru`: Evict keys with expiration only, least recently used (LRU) first.
	// - `volatile_random`: Evict keys with expiration only in a random order.
	// - `volatile_ttl`: Evict keys with expiration only, shortest time-to-live (TTL) first.
	EvictionPolicy DatabasesEvictionPolicyEvictionPolicy `pulumi:"evictionPolicy"`
}

// The set of arguments for constructing a DatabasesEvictionPolicy resource.
type DatabasesEvictionPolicyArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid pulumi.StringPtrInput
	// A string specifying the desired eviction policy for the Redis cluster.
	//
	// - `noeviction`: Don't evict any data, returns error when memory limit is reached.
	// - `allkeys_lru:` Evict any key, least recently used (LRU) first.
	// - `allkeys_random`: Evict keys in a random order.
	// - `volatile_lru`: Evict keys with expiration only, least recently used (LRU) first.
	// - `volatile_random`: Evict keys with expiration only in a random order.
	// - `volatile_ttl`: Evict keys with expiration only, shortest time-to-live (TTL) first.
	EvictionPolicy DatabasesEvictionPolicyEvictionPolicyInput
}

func (DatabasesEvictionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databasesEvictionPolicyArgs)(nil)).Elem()
}

type DatabasesEvictionPolicyInput interface {
	pulumi.Input

	ToDatabasesEvictionPolicyOutput() DatabasesEvictionPolicyOutput
	ToDatabasesEvictionPolicyOutputWithContext(ctx context.Context) DatabasesEvictionPolicyOutput
}

func (*DatabasesEvictionPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabasesEvictionPolicy)(nil)).Elem()
}

func (i *DatabasesEvictionPolicy) ToDatabasesEvictionPolicyOutput() DatabasesEvictionPolicyOutput {
	return i.ToDatabasesEvictionPolicyOutputWithContext(context.Background())
}

func (i *DatabasesEvictionPolicy) ToDatabasesEvictionPolicyOutputWithContext(ctx context.Context) DatabasesEvictionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasesEvictionPolicyOutput)
}

type DatabasesEvictionPolicyOutput struct{ *pulumi.OutputState }

func (DatabasesEvictionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabasesEvictionPolicy)(nil)).Elem()
}

func (o DatabasesEvictionPolicyOutput) ToDatabasesEvictionPolicyOutput() DatabasesEvictionPolicyOutput {
	return o
}

func (o DatabasesEvictionPolicyOutput) ToDatabasesEvictionPolicyOutputWithContext(ctx context.Context) DatabasesEvictionPolicyOutput {
	return o
}

// A string specifying the desired eviction policy for the Redis cluster.
//
// - `noeviction`: Don't evict any data, returns error when memory limit is reached.
// - `allkeys_lru:` Evict any key, least recently used (LRU) first.
// - `allkeys_random`: Evict keys in a random order.
// - `volatile_lru`: Evict keys with expiration only, least recently used (LRU) first.
// - `volatile_random`: Evict keys with expiration only in a random order.
// - `volatile_ttl`: Evict keys with expiration only, shortest time-to-live (TTL) first.
func (o DatabasesEvictionPolicyOutput) EvictionPolicy() DatabasesEvictionPolicyEvictionPolicyOutput {
	return o.ApplyT(func(v *DatabasesEvictionPolicy) DatabasesEvictionPolicyEvictionPolicyOutput { return v.EvictionPolicy }).(DatabasesEvictionPolicyEvictionPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabasesEvictionPolicyInput)(nil)).Elem(), &DatabasesEvictionPolicy{})
	pulumi.RegisterOutputType(DatabasesEvictionPolicyOutput{})
}
