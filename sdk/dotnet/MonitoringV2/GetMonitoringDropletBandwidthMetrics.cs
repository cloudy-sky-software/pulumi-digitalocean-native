// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.MonitoringV2
{
    public static class GetMonitoringDropletBandwidthMetrics
    {
        public static Task<GetMonitoringDropletBandwidthMetricsResult> InvokeAsync(GetMonitoringDropletBandwidthMetricsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMonitoringDropletBandwidthMetricsResult>("digitalocean-native:monitoring/v2:getMonitoringDropletBandwidthMetrics", args ?? new GetMonitoringDropletBandwidthMetricsArgs(), options.WithDefaults());

        public static Output<GetMonitoringDropletBandwidthMetricsResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMonitoringDropletBandwidthMetricsResult>("digitalocean-native:monitoring/v2:getMonitoringDropletBandwidthMetrics", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class GetMonitoringDropletBandwidthMetricsArgs : global::Pulumi.InvokeArgs
    {
        public GetMonitoringDropletBandwidthMetricsArgs()
        {
        }
        public static new GetMonitoringDropletBandwidthMetricsArgs Empty => new GetMonitoringDropletBandwidthMetricsArgs();
    }


    [OutputType]
    public sealed class GetMonitoringDropletBandwidthMetricsResult
    {
        public readonly Outputs.Metrics Items;

        [OutputConstructor]
        private GetMonitoringDropletBandwidthMetricsResult(Outputs.Metrics items)
        {
            Items = items;
        }
    }
}
