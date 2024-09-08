// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabasesEvictionPolicy(ctx *pulumi.Context, args *LookupDatabasesEvictionPolicyArgs, opts ...pulumi.InvokeOption) (*LookupDatabasesEvictionPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDatabasesEvictionPolicyResult
	err := ctx.Invoke("digitalocean-native:databases/v2:getDatabasesEvictionPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabasesEvictionPolicyArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid string `pulumi:"databaseClusterUuid"`
}

type LookupDatabasesEvictionPolicyResult struct {
	// A string specifying the desired eviction policy for the Redis cluster.
	//
	// - `noeviction`: Don't evict any data, returns error when memory limit is reached.
	// - `allkeys_lru:` Evict any key, least recently used (LRU) first.
	// - `allkeys_random`: Evict keys in a random order.
	// - `volatile_lru`: Evict keys with expiration only, least recently used (LRU) first.
	// - `volatile_random`: Evict keys with expiration only in a random order.
	// - `volatile_ttl`: Evict keys with expiration only, shortest time-to-live (TTL) first.
	EvictionPolicy GetDatabasesEvictionPolicyPropertiesEvictionPolicy `pulumi:"evictionPolicy"`
}

func LookupDatabasesEvictionPolicyOutput(ctx *pulumi.Context, args LookupDatabasesEvictionPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupDatabasesEvictionPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabasesEvictionPolicyResult, error) {
			args := v.(LookupDatabasesEvictionPolicyArgs)
			r, err := LookupDatabasesEvictionPolicy(ctx, &args, opts...)
			var s LookupDatabasesEvictionPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabasesEvictionPolicyResultOutput)
}

type LookupDatabasesEvictionPolicyOutputArgs struct {
	// A unique identifier for a database cluster.
	DatabaseClusterUuid pulumi.StringInput `pulumi:"databaseClusterUuid"`
}

func (LookupDatabasesEvictionPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabasesEvictionPolicyArgs)(nil)).Elem()
}

type LookupDatabasesEvictionPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupDatabasesEvictionPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabasesEvictionPolicyResult)(nil)).Elem()
}

func (o LookupDatabasesEvictionPolicyResultOutput) ToLookupDatabasesEvictionPolicyResultOutput() LookupDatabasesEvictionPolicyResultOutput {
	return o
}

func (o LookupDatabasesEvictionPolicyResultOutput) ToLookupDatabasesEvictionPolicyResultOutputWithContext(ctx context.Context) LookupDatabasesEvictionPolicyResultOutput {
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
func (o LookupDatabasesEvictionPolicyResultOutput) EvictionPolicy() GetDatabasesEvictionPolicyPropertiesEvictionPolicyOutput {
	return o.ApplyT(func(v LookupDatabasesEvictionPolicyResult) GetDatabasesEvictionPolicyPropertiesEvictionPolicy {
		return v.EvictionPolicy
	}).(GetDatabasesEvictionPolicyPropertiesEvictionPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabasesEvictionPolicyResultOutput{})
}
