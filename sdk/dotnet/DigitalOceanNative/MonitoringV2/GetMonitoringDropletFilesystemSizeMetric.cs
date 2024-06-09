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
    public static class GetMonitoringDropletFilesystemSizeMetric
    {
        public static Task<GetMonitoringDropletFilesystemSizeMetricResult> InvokeAsync(GetMonitoringDropletFilesystemSizeMetricArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMonitoringDropletFilesystemSizeMetricResult>("digitalocean-native:monitoring/v2:getMonitoringDropletFilesystemSizeMetric", args ?? new GetMonitoringDropletFilesystemSizeMetricArgs(), options.WithDefaults());

        public static Output<GetMonitoringDropletFilesystemSizeMetricResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMonitoringDropletFilesystemSizeMetricResult>("digitalocean-native:monitoring/v2:getMonitoringDropletFilesystemSizeMetric", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class GetMonitoringDropletFilesystemSizeMetricArgs : global::Pulumi.InvokeArgs
    {
        public GetMonitoringDropletFilesystemSizeMetricArgs()
        {
        }
        public static new GetMonitoringDropletFilesystemSizeMetricArgs Empty => new GetMonitoringDropletFilesystemSizeMetricArgs();
    }


    [OutputType]
    public sealed class GetMonitoringDropletFilesystemSizeMetricResult
    {
        public readonly Outputs.Metrics Items;

        [OutputConstructor]
        private GetMonitoringDropletFilesystemSizeMetricResult(Outputs.Metrics items)
        {
            Items = items;
        }
    }
}