// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.ImagesV2
{
    public static class GetImageAction
    {
        public static Task<Outputs.Action> InvokeAsync(GetImageActionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.Action>("digitalocean-native:images/v2:getImageAction", args ?? new GetImageActionArgs(), options.WithDefaults());

        public static Output<Outputs.Action> Invoke(GetImageActionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.Action>("digitalocean-native:images/v2:getImageAction", args ?? new GetImageActionInvokeArgs(), options.WithDefaults());

        public static Output<Outputs.Action> Invoke(GetImageActionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.Action>("digitalocean-native:images/v2:getImageAction", args ?? new GetImageActionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetImageActionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique numeric ID that can be used to identify and reference an action.
        /// </summary>
        [Input("actionId", required: true)]
        public string ActionId { get; set; } = null!;

        /// <summary>
        /// A unique number that can be used to identify and reference a specific image.
        /// </summary>
        [Input("imageId", required: true)]
        public string ImageId { get; set; } = null!;

        public GetImageActionArgs()
        {
        }
        public static new GetImageActionArgs Empty => new GetImageActionArgs();
    }

    public sealed class GetImageActionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique numeric ID that can be used to identify and reference an action.
        /// </summary>
        [Input("actionId", required: true)]
        public Input<string> ActionId { get; set; } = null!;

        /// <summary>
        /// A unique number that can be used to identify and reference a specific image.
        /// </summary>
        [Input("imageId", required: true)]
        public Input<string> ImageId { get; set; } = null!;

        public GetImageActionInvokeArgs()
        {
        }
        public static new GetImageActionInvokeArgs Empty => new GetImageActionInvokeArgs();
    }
}
