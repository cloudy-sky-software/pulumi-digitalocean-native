// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.FirewallsV2
{
    public static class ListFirewalls
    {
        public static Task<ListFirewallsResult> InvokeAsync(ListFirewallsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListFirewallsResult>("digitalocean-native:firewalls/v2:listFirewalls", args ?? new ListFirewallsArgs(), options.WithDefaults());

        public static Output<ListFirewallsResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListFirewallsResult>("digitalocean-native:firewalls/v2:listFirewalls", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListFirewallsArgs : global::Pulumi.InvokeArgs
    {
        public ListFirewallsArgs()
        {
        }
        public static new ListFirewallsArgs Empty => new ListFirewallsArgs();
    }


    [OutputType]
    public sealed class ListFirewallsResult
    {
        public readonly Outputs.ListFirewalls Items;

        [OutputConstructor]
        private ListFirewallsResult(Outputs.ListFirewalls items)
        {
            Items = items;
        }
    }
}
