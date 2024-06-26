// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.TagsV2.Outputs
{

    [OutputType]
    public sealed class ListTagsItems
    {
        public readonly Outputs.PageLinks? Links;
        public readonly Outputs.MetaMeta MetaValue;
        public readonly ImmutableArray<Outputs.Tags> Tags;

        [OutputConstructor]
        private ListTagsItems(
            Outputs.PageLinks? links,

            Outputs.MetaMeta meta,

            ImmutableArray<Outputs.Tags> tags)
        {
            Links = links;
            MetaValue = meta;
            Tags = tags;
        }
    }
}
