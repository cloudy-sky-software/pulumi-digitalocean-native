// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.LoadBalancersV2.Outputs
{

    /// <summary>
    /// An object specifying a forwarding rule for a load balancer.
    /// </summary>
    [OutputType]
    public sealed class ForwardingRule
    {
        /// <summary>
        /// The ID of the TLS certificate used for SSL termination if enabled.
        /// </summary>
        public readonly string? CertificateId;
        /// <summary>
        /// An integer representing the port on which the load balancer instance will listen.
        /// </summary>
        public readonly int EntryPort;
        /// <summary>
        /// The protocol used for traffic to the load balancer. The possible values are: `http`, `https`, `http2`, `http3`, `tcp`, or `udp`. If you set the  `entry_protocol` to `udp`, the `target_protocol` must be set to `udp`.  When using UDP, the load balancer requires that you set up a health  check with a port that uses TCP, HTTP, or HTTPS to work properly.
        /// </summary>
        public readonly CloudySkySoftware.Pulumi.DigitalOceanNative.LoadBalancersV2.ForwardingRuleEntryProtocol EntryProtocol;
        /// <summary>
        /// An integer representing the port on the backend Droplets to which the load balancer will send traffic.
        /// </summary>
        public readonly int TargetPort;
        /// <summary>
        /// The protocol used for traffic from the load balancer to the backend Droplets. The possible values are: `http`, `https`, `http2`, `tcp`, or `udp`. If you set the `target_protocol` to `udp`, the `entry_protocol` must be set to  `udp`. When using UDP, the load balancer requires that you set up a health  check with a port that uses TCP, HTTP, or HTTPS to work properly.
        /// </summary>
        public readonly CloudySkySoftware.Pulumi.DigitalOceanNative.LoadBalancersV2.ForwardingRuleTargetProtocol TargetProtocol;
        /// <summary>
        /// A boolean value indicating whether SSL encrypted traffic will be passed through to the backend Droplets.
        /// </summary>
        public readonly bool? TlsPassthrough;

        [OutputConstructor]
        private ForwardingRule(
            string? certificateId,

            int entryPort,

            CloudySkySoftware.Pulumi.DigitalOceanNative.LoadBalancersV2.ForwardingRuleEntryProtocol entryProtocol,

            int targetPort,

            CloudySkySoftware.Pulumi.DigitalOceanNative.LoadBalancersV2.ForwardingRuleTargetProtocol targetProtocol,

            bool? tlsPassthrough)
        {
            CertificateId = certificateId;
            EntryPort = entryPort;
            EntryProtocol = entryProtocol;
            TargetPort = targetPort;
            TargetProtocol = targetProtocol;
            TlsPassthrough = tlsPassthrough;
        }
    }
}
