// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.FloatingIpsV2
{
    public static class ListFloatingIPs
    {
        public static Task<Outputs.ListFloatingIPsItems> InvokeAsync(ListFloatingIPsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.ListFloatingIPsItems>("digitalocean-native:floating_ips/v2:listFloatingIPs", args ?? new ListFloatingIPsArgs(), options.WithDefaults());

        public static Output<Outputs.ListFloatingIPsItems> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ListFloatingIPsItems>("digitalocean-native:floating_ips/v2:listFloatingIPs", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListFloatingIPsArgs : global::Pulumi.InvokeArgs
    {
        public ListFloatingIPsArgs()
        {
        }
        public static new ListFloatingIPsArgs Empty => new ListFloatingIPsArgs();
    }
}
