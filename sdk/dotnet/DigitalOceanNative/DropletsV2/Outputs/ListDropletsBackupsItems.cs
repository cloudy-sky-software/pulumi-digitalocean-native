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
    public sealed class ListDropletsBackupsItems
    {
        public readonly ImmutableArray<Outputs.DropletSnapshot> Backups;
        public readonly Outputs.PageLinks? Links;
        public readonly Outputs.MetaMeta MetaValue;

        [OutputConstructor]
        private ListDropletsBackupsItems(
            ImmutableArray<Outputs.DropletSnapshot> backups,

            Outputs.PageLinks? links,

            Outputs.MetaMeta meta)
        {
            Backups = backups;
            Links = links;
            MetaValue = meta;
        }
    }
}
