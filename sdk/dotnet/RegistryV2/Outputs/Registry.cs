// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.RegistryV2.Outputs
{

    [OutputType]
    public sealed class Registry
    {
        /// <summary>
        /// A time value given in ISO8601 combined date and time format that represents when the registry was created.
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// A globally unique name for the container registry. Must be lowercase and be composed only of numbers, letters and `-`, up to a limit of 63 characters.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Slug of the region where registry data is stored
        /// </summary>
        public readonly string? Region;
        /// <summary>
        /// The amount of storage used in the registry in bytes.
        /// </summary>
        public readonly int? StorageUsageBytes;
        /// <summary>
        /// The time at which the storage usage was updated. Storage usage is calculated asynchronously, and may not immediately reflect pushes to the registry.
        /// </summary>
        public readonly string? StorageUsageBytesUpdatedAt;
        public readonly Outputs.RegistrySubscription? Subscription;

        [OutputConstructor]
        private Registry(
            string? createdAt,

            string? name,

            string? region,

            int? storageUsageBytes,

            string? storageUsageBytesUpdatedAt,

            Outputs.RegistrySubscription? subscription)
        {
            CreatedAt = createdAt;
            Name = name;
            Region = region;
            StorageUsageBytes = storageUsageBytes;
            StorageUsageBytesUpdatedAt = storageUsageBytesUpdatedAt;
            Subscription = subscription;
        }
    }
}
