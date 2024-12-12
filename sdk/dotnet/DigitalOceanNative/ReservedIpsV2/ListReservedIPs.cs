// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.ReservedIpsV2
{
    public static class ListReservedIPs
    {
        public static Task<Outputs.ListReservedIPsItems> InvokeAsync(ListReservedIPsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.ListReservedIPsItems>("digitalocean-native:reserved_ips/v2:listReservedIPs", args ?? new ListReservedIPsArgs(), options.WithDefaults());

        public static Output<Outputs.ListReservedIPsItems> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ListReservedIPsItems>("digitalocean-native:reserved_ips/v2:listReservedIPs", InvokeArgs.Empty, options.WithDefaults());

        public static Output<Outputs.ListReservedIPsItems> Invoke(InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ListReservedIPsItems>("digitalocean-native:reserved_ips/v2:listReservedIPs", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListReservedIPsArgs : global::Pulumi.InvokeArgs
    {
        public ListReservedIPsArgs()
        {
        }
        public static new ListReservedIPsArgs Empty => new ListReservedIPsArgs();
    }
}
