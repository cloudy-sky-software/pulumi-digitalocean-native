// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.DropletsV2
{
    public static class ListDropletActions
    {
        public static Task<ListDropletActionsResult> InvokeAsync(ListDropletActionsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListDropletActionsResult>("digitalocean-native:droplets/v2:listDropletActions", args ?? new ListDropletActionsArgs(), options.WithDefaults());

        public static Output<ListDropletActionsResult> Invoke(ListDropletActionsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListDropletActionsResult>("digitalocean-native:droplets/v2:listDropletActions", args ?? new ListDropletActionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListDropletActionsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a Droplet instance.
        /// </summary>
        [Input("dropletId", required: true)]
        public string DropletId { get; set; } = null!;

        public ListDropletActionsArgs()
        {
        }
        public static new ListDropletActionsArgs Empty => new ListDropletActionsArgs();
    }

    public sealed class ListDropletActionsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a Droplet instance.
        /// </summary>
        [Input("dropletId", required: true)]
        public Input<string> DropletId { get; set; } = null!;

        public ListDropletActionsInvokeArgs()
        {
        }
        public static new ListDropletActionsInvokeArgs Empty => new ListDropletActionsInvokeArgs();
    }


    [OutputType]
    public sealed class ListDropletActionsResult
    {
        public readonly Outputs.ListDropletActions Items;

        [OutputConstructor]
        private ListDropletActionsResult(Outputs.ListDropletActions items)
        {
            Items = items;
        }
    }
}
