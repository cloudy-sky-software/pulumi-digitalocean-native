// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.FloatingIpsV2
{
    public static class ListFloatingIPsAction
    {
        public static Task<ListFloatingIPsActionResult> InvokeAsync(ListFloatingIPsActionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListFloatingIPsActionResult>("digitalocean-native:floating_ips/v2:listFloatingIPsAction", args ?? new ListFloatingIPsActionArgs(), options.WithDefaults());

        public static Output<ListFloatingIPsActionResult> Invoke(ListFloatingIPsActionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListFloatingIPsActionResult>("digitalocean-native:floating_ips/v2:listFloatingIPsAction", args ?? new ListFloatingIPsActionInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListFloatingIPsActionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A floating IP address.
        /// </summary>
        [Input("floatingIp", required: true)]
        public string FloatingIp { get; set; } = null!;

        public ListFloatingIPsActionArgs()
        {
        }
        public static new ListFloatingIPsActionArgs Empty => new ListFloatingIPsActionArgs();
    }

    public sealed class ListFloatingIPsActionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A floating IP address.
        /// </summary>
        [Input("floatingIp", required: true)]
        public Input<string> FloatingIp { get; set; } = null!;

        public ListFloatingIPsActionInvokeArgs()
        {
        }
        public static new ListFloatingIPsActionInvokeArgs Empty => new ListFloatingIPsActionInvokeArgs();
    }


    [OutputType]
    public sealed class ListFloatingIPsActionResult
    {
        public readonly Outputs.ListFloatingIPsAction Items;

        [OutputConstructor]
        private ListFloatingIPsActionResult(Outputs.ListFloatingIPsAction items)
        {
            Items = items;
        }
    }
}
