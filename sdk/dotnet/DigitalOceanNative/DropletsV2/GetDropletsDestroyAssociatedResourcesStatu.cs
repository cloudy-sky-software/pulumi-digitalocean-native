// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.DropletsV2
{
    public static class GetDropletsDestroyAssociatedResourcesStatu
    {
        public static Task<GetDropletsDestroyAssociatedResourcesStatuResult> InvokeAsync(GetDropletsDestroyAssociatedResourcesStatuArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDropletsDestroyAssociatedResourcesStatuResult>("digitalocean-native:droplets/v2:getDropletsDestroyAssociatedResourcesStatu", args ?? new GetDropletsDestroyAssociatedResourcesStatuArgs(), options.WithDefaults());

        public static Output<GetDropletsDestroyAssociatedResourcesStatuResult> Invoke(GetDropletsDestroyAssociatedResourcesStatuInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDropletsDestroyAssociatedResourcesStatuResult>("digitalocean-native:droplets/v2:getDropletsDestroyAssociatedResourcesStatu", args ?? new GetDropletsDestroyAssociatedResourcesStatuInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDropletsDestroyAssociatedResourcesStatuArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a Droplet instance.
        /// </summary>
        [Input("dropletId", required: true)]
        public string DropletId { get; set; } = null!;

        public GetDropletsDestroyAssociatedResourcesStatuArgs()
        {
        }
        public static new GetDropletsDestroyAssociatedResourcesStatuArgs Empty => new GetDropletsDestroyAssociatedResourcesStatuArgs();
    }

    public sealed class GetDropletsDestroyAssociatedResourcesStatuInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a Droplet instance.
        /// </summary>
        [Input("dropletId", required: true)]
        public Input<string> DropletId { get; set; } = null!;

        public GetDropletsDestroyAssociatedResourcesStatuInvokeArgs()
        {
        }
        public static new GetDropletsDestroyAssociatedResourcesStatuInvokeArgs Empty => new GetDropletsDestroyAssociatedResourcesStatuInvokeArgs();
    }


    [OutputType]
    public sealed class GetDropletsDestroyAssociatedResourcesStatuResult
    {
        public readonly Outputs.AssociatedResourceStatus Items;

        [OutputConstructor]
        private GetDropletsDestroyAssociatedResourcesStatuResult(Outputs.AssociatedResourceStatus items)
        {
            Items = items;
        }
    }
}
