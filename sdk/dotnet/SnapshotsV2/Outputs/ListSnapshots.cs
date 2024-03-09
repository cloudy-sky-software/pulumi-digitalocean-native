// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.SnapshotsV2.Outputs
{

    [OutputType]
    public sealed class ListSnapshots
    {
        public readonly Outputs.PageLinks? Links;
        public readonly Outputs.MetaMeta MetaValue;
        public readonly ImmutableArray<Outputs.Snapshots> Snapshots;

        [OutputConstructor]
        private ListSnapshots(
            Outputs.PageLinks? links,

            Outputs.MetaMeta meta,

            ImmutableArray<Outputs.Snapshots> snapshots)
        {
            Links = links;
            MetaValue = meta;
            Snapshots = snapshots;
        }
    }
}
