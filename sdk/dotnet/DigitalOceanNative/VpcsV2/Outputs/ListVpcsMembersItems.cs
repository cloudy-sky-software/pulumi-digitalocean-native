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
    public sealed class ListVpcsMembersItems
    {
        public readonly Outputs.PageLinks? Links;
        public readonly ImmutableArray<Outputs.VpcMember> Members;
        public readonly Outputs.MetaMeta MetaValue;

        [OutputConstructor]
        private ListVpcsMembersItems(
            Outputs.PageLinks? links,

            ImmutableArray<Outputs.VpcMember> members,

            Outputs.MetaMeta meta)
        {
            Links = links;
            Members = members;
            MetaValue = meta;
        }
    }
}
