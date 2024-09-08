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
    public static class GetVolumeAction
    {
        public static Task<Outputs.GetVolumeActionProperties> InvokeAsync(GetVolumeActionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetVolumeActionProperties>("digitalocean-native:volumes/v2:getVolumeAction", args ?? new GetVolumeActionArgs(), options.WithDefaults());

        public static Output<Outputs.GetVolumeActionProperties> Invoke(GetVolumeActionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetVolumeActionProperties>("digitalocean-native:volumes/v2:getVolumeAction", args ?? new GetVolumeActionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVolumeActionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique numeric ID that can be used to identify and reference an action.
        /// </summary>
        [Input("actionId", required: true)]
        public string ActionId { get; set; } = null!;

        /// <summary>
        /// The ID of the block storage volume.
        /// </summary>
        [Input("volumeId", required: true)]
        public string VolumeId { get; set; } = null!;

        public GetVolumeActionArgs()
        {
        }
        public static new GetVolumeActionArgs Empty => new GetVolumeActionArgs();
    }

    public sealed class GetVolumeActionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique numeric ID that can be used to identify and reference an action.
        /// </summary>
        [Input("actionId", required: true)]
        public Input<string> ActionId { get; set; } = null!;

        /// <summary>
        /// The ID of the block storage volume.
        /// </summary>
        [Input("volumeId", required: true)]
        public Input<string> VolumeId { get; set; } = null!;

        public GetVolumeActionInvokeArgs()
        {
        }
        public static new GetVolumeActionInvokeArgs Empty => new GetVolumeActionInvokeArgs();
    }
}
