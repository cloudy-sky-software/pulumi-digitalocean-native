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
    public static class ListDropletsBackups
    {
        public static Task<ListDropletsBackupsResult> InvokeAsync(ListDropletsBackupsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListDropletsBackupsResult>("digitalocean-native:droplets/v2:listDropletsBackups", args ?? new ListDropletsBackupsArgs(), options.WithDefaults());

        public static Output<ListDropletsBackupsResult> Invoke(ListDropletsBackupsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListDropletsBackupsResult>("digitalocean-native:droplets/v2:listDropletsBackups", args ?? new ListDropletsBackupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListDropletsBackupsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a Droplet instance.
        /// </summary>
        [Input("dropletId", required: true)]
        public string DropletId { get; set; } = null!;

        public ListDropletsBackupsArgs()
        {
        }
        public static new ListDropletsBackupsArgs Empty => new ListDropletsBackupsArgs();
    }

    public sealed class ListDropletsBackupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a Droplet instance.
        /// </summary>
        [Input("dropletId", required: true)]
        public Input<string> DropletId { get; set; } = null!;

        public ListDropletsBackupsInvokeArgs()
        {
        }
        public static new ListDropletsBackupsInvokeArgs Empty => new ListDropletsBackupsInvokeArgs();
    }


    [OutputType]
    public sealed class ListDropletsBackupsResult
    {
        public readonly Outputs.ListDropletsBackupsItems Items;

        [OutputConstructor]
        private ListDropletsBackupsResult(Outputs.ListDropletsBackupsItems items)
        {
            Items = items;
        }
    }
}
