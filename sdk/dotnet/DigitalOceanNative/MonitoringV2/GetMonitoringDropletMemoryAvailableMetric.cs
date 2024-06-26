// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.MonitoringV2
{
    public static class GetMonitoringDropletMemoryAvailableMetric
    {
        public static Task<GetMonitoringDropletMemoryAvailableMetricResult> InvokeAsync(GetMonitoringDropletMemoryAvailableMetricArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMonitoringDropletMemoryAvailableMetricResult>("digitalocean-native:monitoring/v2:getMonitoringDropletMemoryAvailableMetric", args ?? new GetMonitoringDropletMemoryAvailableMetricArgs(), options.WithDefaults());

        public static Output<GetMonitoringDropletMemoryAvailableMetricResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMonitoringDropletMemoryAvailableMetricResult>("digitalocean-native:monitoring/v2:getMonitoringDropletMemoryAvailableMetric", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class GetMonitoringDropletMemoryAvailableMetricArgs : global::Pulumi.InvokeArgs
    {
        public GetMonitoringDropletMemoryAvailableMetricArgs()
        {
        }
        public static new GetMonitoringDropletMemoryAvailableMetricArgs Empty => new GetMonitoringDropletMemoryAvailableMetricArgs();
    }


    [OutputType]
    public sealed class GetMonitoringDropletMemoryAvailableMetricResult
    {
        public readonly Outputs.Metrics Items;

        [OutputConstructor]
        private GetMonitoringDropletMemoryAvailableMetricResult(Outputs.Metrics items)
        {
            Items = items;
        }
    }
}
