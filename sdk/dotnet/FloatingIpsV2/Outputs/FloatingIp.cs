// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.FloatingIpsV2.Outputs
{

    [OutputType]
    public sealed class FloatingIp
    {
        /// <summary>
        /// The Droplet that the floating IP has been assigned to. When you query a floating IP, if it is assigned to a Droplet, the entire Droplet object will be returned. If it is not assigned, the value will be null.
        /// </summary>
        public readonly Outputs.Droplet? Droplet;
        /// <summary>
        /// The public IP address of the floating IP. It also serves as its identifier.
        /// </summary>
        public readonly string? Ip;
        /// <summary>
        /// A boolean value indicating whether or not the floating IP has pending actions preventing new ones from being submitted.
        /// </summary>
        public readonly bool? Locked;
        /// <summary>
        /// The UUID of the project to which the reserved IP currently belongs.
        /// </summary>
        public readonly string? ProjectId;
        public readonly Outputs.FloatingIpRegion? Region;

        [OutputConstructor]
        private FloatingIp(
            Outputs.Droplet? droplet,

            string? ip,

            bool? locked,

            string? projectId,

            Outputs.FloatingIpRegion? region)
        {
            Droplet = droplet;
            Ip = ip;
            Locked = locked;
            ProjectId = projectId;
            Region = region;
        }
    }
}
