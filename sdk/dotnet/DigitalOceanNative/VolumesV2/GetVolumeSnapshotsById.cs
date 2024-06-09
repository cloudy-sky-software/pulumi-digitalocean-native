// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.VolumesV2
{
    public static class GetVolumeSnapshotsById
    {
        public static Task<GetVolumeSnapshotsByIdResult> InvokeAsync(GetVolumeSnapshotsByIdArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVolumeSnapshotsByIdResult>("digitalocean-native:volumes/v2:getVolumeSnapshotsById", args ?? new GetVolumeSnapshotsByIdArgs(), options.WithDefaults());

        public static Output<GetVolumeSnapshotsByIdResult> Invoke(GetVolumeSnapshotsByIdInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVolumeSnapshotsByIdResult>("digitalocean-native:volumes/v2:getVolumeSnapshotsById", args ?? new GetVolumeSnapshotsByIdInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVolumeSnapshotsByIdArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Either the ID of an existing snapshot. This will be an integer for a Droplet snapshot or a string for a volume snapshot.
        /// </summary>
        [Input("snapshotId", required: true)]
        public string SnapshotId { get; set; } = null!;

        public GetVolumeSnapshotsByIdArgs()
        {
        }
        public static new GetVolumeSnapshotsByIdArgs Empty => new GetVolumeSnapshotsByIdArgs();
    }

    public sealed class GetVolumeSnapshotsByIdInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Either the ID of an existing snapshot. This will be an integer for a Droplet snapshot or a string for a volume snapshot.
        /// </summary>
        [Input("snapshotId", required: true)]
        public Input<string> SnapshotId { get; set; } = null!;

        public GetVolumeSnapshotsByIdInvokeArgs()
        {
        }
        public static new GetVolumeSnapshotsByIdInvokeArgs Empty => new GetVolumeSnapshotsByIdInvokeArgs();
    }


    [OutputType]
    public sealed class GetVolumeSnapshotsByIdResult
    {
        public readonly Outputs.GetVolumeSnapshotsByIdProperties Items;

        [OutputConstructor]
        private GetVolumeSnapshotsByIdResult(Outputs.GetVolumeSnapshotsByIdProperties items)
        {
            Items = items;
        }
    }
}