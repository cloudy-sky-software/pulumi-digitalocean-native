// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.ImagesV2
{
    public static class ListImageActions
    {
        public static Task<ListImageActionsResult> InvokeAsync(ListImageActionsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListImageActionsResult>("digitalocean-native:images/v2:listImageActions", args ?? new ListImageActionsArgs(), options.WithDefaults());

        public static Output<ListImageActionsResult> Invoke(ListImageActionsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListImageActionsResult>("digitalocean-native:images/v2:listImageActions", args ?? new ListImageActionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListImageActionsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique number that can be used to identify and reference a specific image.
        /// </summary>
        [Input("imageId", required: true)]
        public string ImageId { get; set; } = null!;

        public ListImageActionsArgs()
        {
        }
        public static new ListImageActionsArgs Empty => new ListImageActionsArgs();
    }

    public sealed class ListImageActionsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique number that can be used to identify and reference a specific image.
        /// </summary>
        [Input("imageId", required: true)]
        public Input<string> ImageId { get; set; } = null!;

        public ListImageActionsInvokeArgs()
        {
        }
        public static new ListImageActionsInvokeArgs Empty => new ListImageActionsInvokeArgs();
    }


    [OutputType]
    public sealed class ListImageActionsResult
    {
        public readonly Outputs.ListImageActions Items;

        [OutputConstructor]
        private ListImageActionsResult(Outputs.ListImageActions items)
        {
            Items = items;
        }
    }
}
