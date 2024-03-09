// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.SnapshotsV2
{
    public static class ListSnapshots
    {
        public static Task<ListSnapshotsResult> InvokeAsync(ListSnapshotsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListSnapshotsResult>("digitalocean-native:snapshots/v2:listSnapshots", args ?? new ListSnapshotsArgs(), options.WithDefaults());

        public static Output<ListSnapshotsResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListSnapshotsResult>("digitalocean-native:snapshots/v2:listSnapshots", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListSnapshotsArgs : global::Pulumi.InvokeArgs
    {
        public ListSnapshotsArgs()
        {
        }
        public static new ListSnapshotsArgs Empty => new ListSnapshotsArgs();
    }


    [OutputType]
    public sealed class ListSnapshotsResult
    {
        public readonly Outputs.ListSnapshots Items;

        [OutputConstructor]
        private ListSnapshotsResult(Outputs.ListSnapshots items)
        {
            Items = items;
        }
    }
}
