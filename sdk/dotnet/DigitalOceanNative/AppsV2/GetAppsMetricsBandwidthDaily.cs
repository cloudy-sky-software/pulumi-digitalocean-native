// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.AppsV2
{
    public static class GetAppsMetricsBandwidthDaily
    {
        public static Task<GetAppsMetricsBandwidthDailyResult> InvokeAsync(GetAppsMetricsBandwidthDailyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAppsMetricsBandwidthDailyResult>("digitalocean-native:apps/v2:getAppsMetricsBandwidthDaily", args ?? new GetAppsMetricsBandwidthDailyArgs(), options.WithDefaults());

        public static Output<GetAppsMetricsBandwidthDailyResult> Invoke(GetAppsMetricsBandwidthDailyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAppsMetricsBandwidthDailyResult>("digitalocean-native:apps/v2:getAppsMetricsBandwidthDaily", args ?? new GetAppsMetricsBandwidthDailyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppsMetricsBandwidthDailyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The app ID
        /// </summary>
        [Input("appId", required: true)]
        public string AppId { get; set; } = null!;

        public GetAppsMetricsBandwidthDailyArgs()
        {
        }
        public static new GetAppsMetricsBandwidthDailyArgs Empty => new GetAppsMetricsBandwidthDailyArgs();
    }

    public sealed class GetAppsMetricsBandwidthDailyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The app ID
        /// </summary>
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        public GetAppsMetricsBandwidthDailyInvokeArgs()
        {
        }
        public static new GetAppsMetricsBandwidthDailyInvokeArgs Empty => new GetAppsMetricsBandwidthDailyInvokeArgs();
    }


    [OutputType]
    public sealed class GetAppsMetricsBandwidthDailyResult
    {
        public readonly Outputs.AppMetricsBandwidthUsage Items;

        [OutputConstructor]
        private GetAppsMetricsBandwidthDailyResult(Outputs.AppMetricsBandwidthUsage items)
        {
            Items = items;
        }
    }
}
