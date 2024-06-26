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
    /// An object specifying allow and deny rules to control traffic to the load balancer.
    /// </summary>
    [OutputType]
    public sealed class LbFirewall
    {
        /// <summary>
        /// the rules for allowing traffic to the load balancer (in the form 'ip:1.2.3.4' or 'cidr:1.2.0.0/16')
        /// </summary>
        public readonly ImmutableArray<string> Allow;
        /// <summary>
        /// the rules for denying traffic to the load balancer (in the form 'ip:1.2.3.4' or 'cidr:1.2.0.0/16')
        /// </summary>
        public readonly ImmutableArray<string> Deny;

        [OutputConstructor]
        private LbFirewall(
            ImmutableArray<string> allow,

            ImmutableArray<string> deny)
        {
            Allow = allow;
            Deny = deny;
        }
    }
}
