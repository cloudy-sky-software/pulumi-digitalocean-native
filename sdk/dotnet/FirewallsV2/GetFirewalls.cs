// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.FirewallsV2
{
    public static class GetFirewalls
    {
        public static Task<GetFirewallsResult> InvokeAsync(GetFirewallsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFirewallsResult>("digitalocean-native:firewalls/v2:getFirewalls", args ?? new GetFirewallsArgs(), options.WithDefaults());

        public static Output<GetFirewallsResult> Invoke(GetFirewallsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallsResult>("digitalocean-native:firewalls/v2:getFirewalls", args ?? new GetFirewallsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirewallsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique ID that can be used to identify and reference a firewall.
        /// </summary>
        [Input("firewallId", required: true)]
        public string FirewallId { get; set; } = null!;

        public GetFirewallsArgs()
        {
        }
        public static new GetFirewallsArgs Empty => new GetFirewallsArgs();
    }

    public sealed class GetFirewallsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique ID that can be used to identify and reference a firewall.
        /// </summary>
        [Input("firewallId", required: true)]
        public Input<string> FirewallId { get; set; } = null!;

        public GetFirewallsInvokeArgs()
        {
        }
        public static new GetFirewallsInvokeArgs Empty => new GetFirewallsInvokeArgs();
    }


    [OutputType]
    public sealed class GetFirewallsResult
    {
        public readonly Outputs.GetFirewallsProperties Items;

        [OutputConstructor]
        private GetFirewallsResult(Outputs.GetFirewallsProperties items)
        {
            Items = items;
        }
    }
}