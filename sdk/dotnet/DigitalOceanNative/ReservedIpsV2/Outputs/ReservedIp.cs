// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.ReservedIpsV2.Outputs
{

    [OutputType]
    public sealed class ReservedIp
    {
        /// <summary>
        /// The Droplet that the reserved IP has been assigned to. When you query a reserved IP, if it is assigned to a Droplet, the entire Droplet object will be returned. If it is not assigned, the value will be null.
        /// </summary>
        public readonly Outputs.Droplet? Droplet;
        /// <summary>
        /// The public IP address of the reserved IP. It also serves as its identifier.
        /// </summary>
        public readonly string? Ip;
        /// <summary>
        /// A boolean value indicating whether or not the reserved IP has pending actions preventing new ones from being submitted.
        /// </summary>
        public readonly bool? Locked;
        /// <summary>
        /// The UUID of the project to which the reserved IP currently belongs.
        /// </summary>
        public readonly string? ProjectId;
        public readonly Outputs.ReservedIpRegion? Region;

        [OutputConstructor]
        private ReservedIp(
            Outputs.Droplet? droplet,

            string? ip,

            bool? locked,

            string? projectId,

            Outputs.ReservedIpRegion? region)
        {
            Droplet = droplet;
            Ip = ip;
            Locked = locked;
            ProjectId = projectId;
            Region = region;
        }
    }
}