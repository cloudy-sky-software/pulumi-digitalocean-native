// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.ImagesV2
{
    public static class GetImages
    {
        public static Task<GetImagesResult> InvokeAsync(GetImagesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetImagesResult>("digitalocean-native:images/v2:getImages", args ?? new GetImagesArgs(), options.WithDefaults());

        public static Output<GetImagesResult> Invoke(GetImagesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetImagesResult>("digitalocean-native:images/v2:getImages", args ?? new GetImagesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetImagesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique number (id) or string (slug) used to identify and reference a
        /// specific image.
        /// 
        /// **Public** images can be identified by image `id` or `slug`.
        /// 
        /// **Private** images *must* be identified by image `id`.
        /// </summary>
        [Input("imageId", required: true)]
        public string ImageId { get; set; } = null!;

        public GetImagesArgs()
        {
        }
        public static new GetImagesArgs Empty => new GetImagesArgs();
    }

    public sealed class GetImagesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique number (id) or string (slug) used to identify and reference a
        /// specific image.
        /// 
        /// **Public** images can be identified by image `id` or `slug`.
        /// 
        /// **Private** images *must* be identified by image `id`.
        /// </summary>
        [Input("imageId", required: true)]
        public Input<string> ImageId { get; set; } = null!;

        public GetImagesInvokeArgs()
        {
        }
        public static new GetImagesInvokeArgs Empty => new GetImagesInvokeArgs();
    }


    [OutputType]
    public sealed class GetImagesResult
    {
        public readonly Outputs.GetImagesProperties Items;

        [OutputConstructor]
        private GetImagesResult(Outputs.GetImagesProperties items)
        {
            Items = items;
        }
    }
}