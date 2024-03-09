// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.LoadBalancersV2.Outputs
{

    [OutputType]
    public sealed class ListLoadBalancers
    {
        public readonly Outputs.PageLinks? Links;
        public readonly ImmutableArray<Outputs.LoadBalancer> LoadBalancers;
        public readonly Outputs.MetaMeta MetaValue;

        [OutputConstructor]
        private ListLoadBalancers(
            Outputs.PageLinks? links,

            ImmutableArray<Outputs.LoadBalancer> loadBalancers,

            Outputs.MetaMeta meta)
        {
            Links = links;
            LoadBalancers = loadBalancers;
            MetaValue = meta;
        }
    }
}