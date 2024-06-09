// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.RegistryV2.Outputs
{

    [OutputType]
    public sealed class Subscription
    {
        /// <summary>
        /// The time at which the subscription was created.
        /// </summary>
        public readonly string? CreatedAt;
        public readonly Outputs.SubscriptionTierBase? Tier;
        /// <summary>
        /// The time at which the subscription was last updated.
        /// </summary>
        public readonly string? UpdatedAt;

        [OutputConstructor]
        private Subscription(
            string? createdAt,

            Outputs.SubscriptionTierBase? tier,

            string? updatedAt)
        {
            CreatedAt = createdAt;
            Tier = tier;
            UpdatedAt = updatedAt;
        }
    }
}