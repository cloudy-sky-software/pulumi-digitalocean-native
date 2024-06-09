// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.UptimeV2
{
    public static class ListUptimeAlerts
    {
        public static Task<ListUptimeAlertsResult> InvokeAsync(ListUptimeAlertsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListUptimeAlertsResult>("digitalocean-native:uptime/v2:listUptimeAlerts", args ?? new ListUptimeAlertsArgs(), options.WithDefaults());

        public static Output<ListUptimeAlertsResult> Invoke(ListUptimeAlertsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListUptimeAlertsResult>("digitalocean-native:uptime/v2:listUptimeAlerts", args ?? new ListUptimeAlertsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListUptimeAlertsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a check.
        /// </summary>
        [Input("checkId", required: true)]
        public string CheckId { get; set; } = null!;

        public ListUptimeAlertsArgs()
        {
        }
        public static new ListUptimeAlertsArgs Empty => new ListUptimeAlertsArgs();
    }

    public sealed class ListUptimeAlertsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a check.
        /// </summary>
        [Input("checkId", required: true)]
        public Input<string> CheckId { get; set; } = null!;

        public ListUptimeAlertsInvokeArgs()
        {
        }
        public static new ListUptimeAlertsInvokeArgs Empty => new ListUptimeAlertsInvokeArgs();
    }


    [OutputType]
    public sealed class ListUptimeAlertsResult
    {
        public readonly Outputs.ListUptimeAlertsItems Items;

        [OutputConstructor]
        private ListUptimeAlertsResult(Outputs.ListUptimeAlertsItems items)
        {
            Items = items;
        }
    }
}