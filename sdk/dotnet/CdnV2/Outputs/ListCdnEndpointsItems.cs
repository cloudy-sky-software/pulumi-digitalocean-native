// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.CdnV2.Outputs
{

    [OutputType]
    public sealed class ListCdnEndpointsItems
    {
        public readonly ImmutableArray<Outputs.CdnEndpoint> Endpoints;
        public readonly Outputs.PageLinks? Links;
        public readonly Outputs.MetaMeta MetaValue;

        [OutputConstructor]
        private ListCdnEndpointsItems(
            ImmutableArray<Outputs.CdnEndpoint> endpoints,

            Outputs.PageLinks? links,

            Outputs.MetaMeta meta)
        {
            Endpoints = endpoints;
            Links = links;
            MetaValue = meta;
        }
    }
}