// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.DropletsV2.Outputs
{

    [OutputType]
    public sealed class FirewallRulesOutboundRulesItem
    {
        /// <summary>
        /// The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. "8000-9000"), or "0" when all ports are open for a protocol. For ICMP rules this parameter will always return "0".
        /// </summary>
        public readonly string Ports;
        /// <summary>
        /// The type of traffic to be allowed. This may be one of `tcp`, `udp`, or `icmp`.
        /// </summary>
        public readonly CloudySkySoftware.Pulumi.DigitalOceanNative.DropletsV2.FirewallRuleBaseProtocol Protocol;

        [OutputConstructor]
        private FirewallRulesOutboundRulesItem(
            string ports,

            CloudySkySoftware.Pulumi.DigitalOceanNative.DropletsV2.FirewallRuleBaseProtocol protocol)
        {
            Ports = ports;
            Protocol = protocol;
        }
    }
}