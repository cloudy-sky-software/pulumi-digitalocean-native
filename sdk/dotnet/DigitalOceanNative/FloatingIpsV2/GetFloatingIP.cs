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
    public static class GetFloatingIP
    {
        public static Task<Outputs.GetFloatingIPProperties> InvokeAsync(GetFloatingIPArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetFloatingIPProperties>("digitalocean-native:floating_ips/v2:getFloatingIP", args ?? new GetFloatingIPArgs(), options.WithDefaults());

        public static Output<Outputs.GetFloatingIPProperties> Invoke(GetFloatingIPInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetFloatingIPProperties>("digitalocean-native:floating_ips/v2:getFloatingIP", args ?? new GetFloatingIPInvokeArgs(), options.WithDefaults());

        public static Output<Outputs.GetFloatingIPProperties> Invoke(GetFloatingIPInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetFloatingIPProperties>("digitalocean-native:floating_ips/v2:getFloatingIP", args ?? new GetFloatingIPInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFloatingIPArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A floating IP address.
        /// </summary>
        [Input("floatingIp", required: true)]
        public string FloatingIp { get; set; } = null!;

        public GetFloatingIPArgs()
        {
        }
        public static new GetFloatingIPArgs Empty => new GetFloatingIPArgs();
    }

    public sealed class GetFloatingIPInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A floating IP address.
        /// </summary>
        [Input("floatingIp", required: true)]
        public Input<string> FloatingIp { get; set; } = null!;

        public GetFloatingIPInvokeArgs()
        {
        }
        public static new GetFloatingIPInvokeArgs Empty => new GetFloatingIPInvokeArgs();
    }
}
