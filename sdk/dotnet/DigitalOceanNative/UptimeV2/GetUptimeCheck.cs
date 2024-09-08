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
    public static class GetUptimeCheck
    {
        public static Task<Outputs.GetUptimeCheckProperties> InvokeAsync(GetUptimeCheckArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetUptimeCheckProperties>("digitalocean-native:uptime/v2:getUptimeCheck", args ?? new GetUptimeCheckArgs(), options.WithDefaults());

        public static Output<Outputs.GetUptimeCheckProperties> Invoke(GetUptimeCheckInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetUptimeCheckProperties>("digitalocean-native:uptime/v2:getUptimeCheck", args ?? new GetUptimeCheckInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUptimeCheckArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a check.
        /// </summary>
        [Input("checkId", required: true)]
        public string CheckId { get; set; } = null!;

        public GetUptimeCheckArgs()
        {
        }
        public static new GetUptimeCheckArgs Empty => new GetUptimeCheckArgs();
    }

    public sealed class GetUptimeCheckInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a check.
        /// </summary>
        [Input("checkId", required: true)]
        public Input<string> CheckId { get; set; } = null!;

        public GetUptimeCheckInvokeArgs()
        {
        }
        public static new GetUptimeCheckInvokeArgs Empty => new GetUptimeCheckInvokeArgs();
    }
}
