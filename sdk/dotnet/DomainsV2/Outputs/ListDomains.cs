// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.DomainsV2.Outputs
{

    [OutputType]
    public sealed class ListDomains
    {
        /// <summary>
        /// Array of volumes.
        /// </summary>
        public readonly ImmutableArray<Outputs.Domain> Domains;
        public readonly Outputs.PageLinks? Links;
        public readonly Outputs.MetaMeta MetaValue;

        [OutputConstructor]
        private ListDomains(
            ImmutableArray<Outputs.Domain> domains,

            Outputs.PageLinks? links,

            Outputs.MetaMeta meta)
        {
            Domains = domains;
            Links = links;
            MetaValue = meta;
        }
    }
}