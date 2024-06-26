// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.VpcsV2.Outputs
{

    [OutputType]
    public sealed class Vpc
    {
        /// <summary>
        /// A time value given in ISO8601 combined date and time format.
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// A boolean value indicating whether or not the VPC is the default network for the region. All applicable resources are placed into the default VPC network unless otherwise specified during their creation. The `default` field cannot be unset from `true`. If you want to set a new default VPC network, update the `default` field of another VPC network in the same region. The previous network's `default` field will be set to `false` when a new default VPC has been defined.
        /// </summary>
        public readonly bool? Default;
        /// <summary>
        /// A free-form text field for describing the VPC's purpose. It may be a maximum of 255 characters.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// A unique ID that can be used to identify and reference the VPC.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The range of IP addresses in the VPC in CIDR notation. Network ranges cannot overlap with other networks in the same account and must be in range of private addresses as defined in RFC1918. It may not be smaller than `/28` nor larger than `/16`. If no IP range is specified, a `/20` network range is generated that won't conflict with other VPC networks in your account.
        /// </summary>
        public readonly string? IpRange;
        /// <summary>
        /// The name of the VPC. Must be unique and may only contain alphanumeric characters, dashes, and periods.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The slug identifier for the region where the VPC will be created.
        /// </summary>
        public readonly string? Region;
        /// <summary>
        /// The uniform resource name (URN) for the resource in the format do:resource_type:resource_id.
        /// </summary>
        public readonly string? Urn;

        [OutputConstructor]
        private Vpc(
            string? createdAt,

            bool? @default,

            string? description,

            string? id,

            string? ipRange,

            string? name,

            string? region,

            string? urn)
        {
            CreatedAt = createdAt;
            Default = @default;
            Description = description;
            Id = id;
            IpRange = ipRange;
            Name = name;
            Region = region;
            Urn = urn;
        }
    }
}
