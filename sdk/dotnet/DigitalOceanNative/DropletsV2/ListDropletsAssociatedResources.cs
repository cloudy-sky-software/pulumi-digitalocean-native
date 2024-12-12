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
    public static class ListDropletsAssociatedResources
    {
        public static Task<Outputs.ListDropletsAssociatedResourcesItems> InvokeAsync(ListDropletsAssociatedResourcesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.ListDropletsAssociatedResourcesItems>("digitalocean-native:droplets/v2:listDropletsAssociatedResources", args ?? new ListDropletsAssociatedResourcesArgs(), options.WithDefaults());

        public static Output<Outputs.ListDropletsAssociatedResourcesItems> Invoke(ListDropletsAssociatedResourcesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ListDropletsAssociatedResourcesItems>("digitalocean-native:droplets/v2:listDropletsAssociatedResources", args ?? new ListDropletsAssociatedResourcesInvokeArgs(), options.WithDefaults());

        public static Output<Outputs.ListDropletsAssociatedResourcesItems> Invoke(ListDropletsAssociatedResourcesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ListDropletsAssociatedResourcesItems>("digitalocean-native:droplets/v2:listDropletsAssociatedResources", args ?? new ListDropletsAssociatedResourcesInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListDropletsAssociatedResourcesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a Droplet instance.
        /// </summary>
        [Input("dropletId", required: true)]
        public string DropletId { get; set; } = null!;

        public ListDropletsAssociatedResourcesArgs()
        {
        }
        public static new ListDropletsAssociatedResourcesArgs Empty => new ListDropletsAssociatedResourcesArgs();
    }

    public sealed class ListDropletsAssociatedResourcesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a Droplet instance.
        /// </summary>
        [Input("dropletId", required: true)]
        public Input<string> DropletId { get; set; } = null!;

        public ListDropletsAssociatedResourcesInvokeArgs()
        {
        }
        public static new ListDropletsAssociatedResourcesInvokeArgs Empty => new ListDropletsAssociatedResourcesInvokeArgs();
    }
}
