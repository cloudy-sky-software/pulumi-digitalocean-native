// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.SnapshotsV2
{
    public static class GetSnapshot
    {
        public static Task<Outputs.GetSnapshotProperties> InvokeAsync(GetSnapshotArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetSnapshotProperties>("digitalocean-native:snapshots/v2:getSnapshot", args ?? new GetSnapshotArgs(), options.WithDefaults());

        public static Output<Outputs.GetSnapshotProperties> Invoke(GetSnapshotInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetSnapshotProperties>("digitalocean-native:snapshots/v2:getSnapshot", args ?? new GetSnapshotInvokeArgs(), options.WithDefaults());

        public static Output<Outputs.GetSnapshotProperties> Invoke(GetSnapshotInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetSnapshotProperties>("digitalocean-native:snapshots/v2:getSnapshot", args ?? new GetSnapshotInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSnapshotArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Either the ID of an existing snapshot. This will be an integer for a Droplet snapshot or a string for a volume snapshot.
        /// </summary>
        [Input("snapshotId", required: true)]
        public string SnapshotId { get; set; } = null!;

        public GetSnapshotArgs()
        {
        }
        public static new GetSnapshotArgs Empty => new GetSnapshotArgs();
    }

    public sealed class GetSnapshotInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Either the ID of an existing snapshot. This will be an integer for a Droplet snapshot or a string for a volume snapshot.
        /// </summary>
        [Input("snapshotId", required: true)]
        public Input<string> SnapshotId { get; set; } = null!;

        public GetSnapshotInvokeArgs()
        {
        }
        public static new GetSnapshotInvokeArgs Empty => new GetSnapshotInvokeArgs();
    }
}
