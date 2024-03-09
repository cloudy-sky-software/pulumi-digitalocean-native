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
    public sealed class SubscriptionTierBase
    {
        /// <summary>
        /// A boolean indicating whether the subscription tier supports additional storage above what is included in the base plan at an additional cost per GiB used.
        /// </summary>
        public readonly bool? AllowStorageOverage;
        /// <summary>
        /// The amount of outbound data transfer included in the subscription tier in bytes.
        /// </summary>
        public readonly int? IncludedBandwidthBytes;
        /// <summary>
        /// The number of repositories included in the subscription tier. `0` indicates that the subscription tier includes unlimited repositories.
        /// </summary>
        public readonly int? IncludedRepositories;
        /// <summary>
        /// The amount of storage included in the subscription tier in bytes.
        /// </summary>
        public readonly int? IncludedStorageBytes;
        /// <summary>
        /// The monthly cost of the subscription tier in cents.
        /// </summary>
        public readonly int? MonthlyPriceInCents;
        /// <summary>
        /// The name of the subscription tier.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The slug identifier of the subscription tier.
        /// </summary>
        public readonly string? Slug;
        /// <summary>
        /// The price paid in cents per GiB for additional storage beyond what is included in the subscription plan.
        /// </summary>
        public readonly int? StorageOveragePriceInCents;

        [OutputConstructor]
        private SubscriptionTierBase(
            bool? allowStorageOverage,

            int? includedBandwidthBytes,

            int? includedRepositories,

            int? includedStorageBytes,

            int? monthlyPriceInCents,

            string? name,

            string? slug,

            int? storageOveragePriceInCents)
        {
            AllowStorageOverage = allowStorageOverage;
            IncludedBandwidthBytes = includedBandwidthBytes;
            IncludedRepositories = includedRepositories;
            IncludedStorageBytes = includedStorageBytes;
            MonthlyPriceInCents = monthlyPriceInCents;
            Name = name;
            Slug = slug;
            StorageOveragePriceInCents = storageOveragePriceInCents;
        }
    }
}