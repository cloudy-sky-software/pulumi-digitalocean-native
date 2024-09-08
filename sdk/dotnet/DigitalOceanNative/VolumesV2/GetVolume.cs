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
    public static class GetVolume
    {
        public static Task<Outputs.GetVolumeProperties> InvokeAsync(GetVolumeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetVolumeProperties>("digitalocean-native:volumes/v2:getVolume", args ?? new GetVolumeArgs(), options.WithDefaults());

        public static Output<Outputs.GetVolumeProperties> Invoke(GetVolumeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetVolumeProperties>("digitalocean-native:volumes/v2:getVolume", args ?? new GetVolumeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVolumeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the block storage volume.
        /// </summary>
        [Input("volumeId", required: true)]
        public string VolumeId { get; set; } = null!;

        public GetVolumeArgs()
        {
        }
        public static new GetVolumeArgs Empty => new GetVolumeArgs();
    }

    public sealed class GetVolumeInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the block storage volume.
        /// </summary>
        [Input("volumeId", required: true)]
        public Input<string> VolumeId { get; set; } = null!;

        public GetVolumeInvokeArgs()
        {
        }
        public static new GetVolumeInvokeArgs Empty => new GetVolumeInvokeArgs();
    }
}
