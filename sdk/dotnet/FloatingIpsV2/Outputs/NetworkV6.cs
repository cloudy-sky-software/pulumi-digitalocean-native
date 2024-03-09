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
    public sealed class NetworkV6
    {
        /// <summary>
        /// The gateway of the specified IPv6 network interface.
        /// </summary>
        public readonly string? Gateway;
        /// <summary>
        /// The IP address of the IPv6 network interface.
        /// </summary>
        public readonly string? IpAddress;
        /// <summary>
        /// The netmask of the IPv6 network interface.
        /// </summary>
        public readonly int? Netmask;
        /// <summary>
        /// The type of the IPv6 network interface.
        /// 
        /// **Note**: IPv6 private  networking is not currently supported.
        /// </summary>
        public readonly Pulumi.DigitalOceanNative.FloatingIpsV2.NetworkV6Type? Type;

        [OutputConstructor]
        private NetworkV6(
            string? gateway,

            string? ipAddress,

            int? netmask,

            Pulumi.DigitalOceanNative.FloatingIpsV2.NetworkV6Type? type)
        {
            Gateway = gateway;
            IpAddress = ipAddress;
            Netmask = netmask;
            Type = type;
        }
    }
}
