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
    public static class GetUptimeAlert
    {
        public static Task<Outputs.GetUptimeAlertProperties> InvokeAsync(GetUptimeAlertArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetUptimeAlertProperties>("digitalocean-native:uptime/v2:getUptimeAlert", args ?? new GetUptimeAlertArgs(), options.WithDefaults());

        public static Output<Outputs.GetUptimeAlertProperties> Invoke(GetUptimeAlertInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetUptimeAlertProperties>("digitalocean-native:uptime/v2:getUptimeAlert", args ?? new GetUptimeAlertInvokeArgs(), options.WithDefaults());

        public static Output<Outputs.GetUptimeAlertProperties> Invoke(GetUptimeAlertInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetUptimeAlertProperties>("digitalocean-native:uptime/v2:getUptimeAlert", args ?? new GetUptimeAlertInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUptimeAlertArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for an alert.
        /// </summary>
        [Input("alertId", required: true)]
        public string AlertId { get; set; } = null!;

        /// <summary>
        /// A unique identifier for a check.
        /// </summary>
        [Input("checkId", required: true)]
        public string CheckId { get; set; } = null!;

        public GetUptimeAlertArgs()
        {
        }
        public static new GetUptimeAlertArgs Empty => new GetUptimeAlertArgs();
    }

    public sealed class GetUptimeAlertInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for an alert.
        /// </summary>
        [Input("alertId", required: true)]
        public Input<string> AlertId { get; set; } = null!;

        /// <summary>
        /// A unique identifier for a check.
        /// </summary>
        [Input("checkId", required: true)]
        public Input<string> CheckId { get; set; } = null!;

        public GetUptimeAlertInvokeArgs()
        {
        }
        public static new GetUptimeAlertInvokeArgs Empty => new GetUptimeAlertInvokeArgs();
    }
}
