// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.DatabasesV2.Outputs
{

    [OutputType]
    public sealed class FirewallRule
    {
        /// <summary>
        /// A unique ID for the database cluster to which the rule is applied.
        /// </summary>
        public readonly string? ClusterUuid;
        /// <summary>
        /// A time value given in ISO8601 combined date and time format that represents when the firewall rule was created.
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// The type of resource that the firewall rule allows to access the database cluster.
        /// </summary>
        public readonly CloudySkySoftware.Pulumi.DigitalOceanNative.DatabasesV2.FirewallRuleType Type;
        /// <summary>
        /// A unique ID for the firewall rule itself.
        /// </summary>
        public readonly string? Uuid;
        /// <summary>
        /// The ID of the specific resource, the name of a tag applied to a group of resources, or the IP address that the firewall rule allows to access the database cluster.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private FirewallRule(
            string? clusterUuid,

            string? createdAt,

            CloudySkySoftware.Pulumi.DigitalOceanNative.DatabasesV2.FirewallRuleType type,

            string? uuid,

            string value)
        {
            ClusterUuid = clusterUuid;
            CreatedAt = createdAt;
            Type = type;
            Uuid = uuid;
            Value = value;
        }
    }
}
