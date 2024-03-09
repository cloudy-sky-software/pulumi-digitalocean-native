// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.VolumesV2
{
    public static class ListVolumeSnapshots
    {
        public static Task<ListVolumeSnapshotsResult> InvokeAsync(ListVolumeSnapshotsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListVolumeSnapshotsResult>("digitalocean-native:volumes/v2:listVolumeSnapshots", args ?? new ListVolumeSnapshotsArgs(), options.WithDefaults());

        public static Output<ListVolumeSnapshotsResult> Invoke(ListVolumeSnapshotsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListVolumeSnapshotsResult>("digitalocean-native:volumes/v2:listVolumeSnapshots", args ?? new ListVolumeSnapshotsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListVolumeSnapshotsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the block storage volume.
        /// </summary>
        [Input("volumeId", required: true)]
        public string VolumeId { get; set; } = null!;

        public ListVolumeSnapshotsArgs()
        {
        }
        public static new ListVolumeSnapshotsArgs Empty => new ListVolumeSnapshotsArgs();
    }

    public sealed class ListVolumeSnapshotsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the block storage volume.
        /// </summary>
        [Input("volumeId", required: true)]
        public Input<string> VolumeId { get; set; } = null!;

        public ListVolumeSnapshotsInvokeArgs()
        {
        }
        public static new ListVolumeSnapshotsInvokeArgs Empty => new ListVolumeSnapshotsInvokeArgs();
    }


    [OutputType]
    public sealed class ListVolumeSnapshotsResult
    {
        public readonly Outputs.ListVolumeSnapshots Items;

        [OutputConstructor]
        private ListVolumeSnapshotsResult(Outputs.ListVolumeSnapshots items)
        {
            Items = items;
        }
    }
}