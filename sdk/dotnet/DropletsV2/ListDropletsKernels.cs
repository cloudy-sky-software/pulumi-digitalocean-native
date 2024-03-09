// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.DropletsV2
{
    public static class ListDropletsKernels
    {
        public static Task<ListDropletsKernelsResult> InvokeAsync(ListDropletsKernelsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListDropletsKernelsResult>("digitalocean-native:droplets/v2:listDropletsKernels", args ?? new ListDropletsKernelsArgs(), options.WithDefaults());

        public static Output<ListDropletsKernelsResult> Invoke(ListDropletsKernelsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListDropletsKernelsResult>("digitalocean-native:droplets/v2:listDropletsKernels", args ?? new ListDropletsKernelsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListDropletsKernelsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a Droplet instance.
        /// </summary>
        [Input("dropletId", required: true)]
        public string DropletId { get; set; } = null!;

        public ListDropletsKernelsArgs()
        {
        }
        public static new ListDropletsKernelsArgs Empty => new ListDropletsKernelsArgs();
    }

    public sealed class ListDropletsKernelsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a Droplet instance.
        /// </summary>
        [Input("dropletId", required: true)]
        public Input<string> DropletId { get; set; } = null!;

        public ListDropletsKernelsInvokeArgs()
        {
        }
        public static new ListDropletsKernelsInvokeArgs Empty => new ListDropletsKernelsInvokeArgs();
    }


    [OutputType]
    public sealed class ListDropletsKernelsResult
    {
        public readonly Outputs.ListDropletsKernels Items;

        [OutputConstructor]
        private ListDropletsKernelsResult(Outputs.ListDropletsKernels items)
        {
            Items = items;
        }
    }
}
