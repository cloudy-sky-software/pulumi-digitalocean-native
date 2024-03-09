// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.MonitoringV2
{
    public static class GetMonitoringDropletFilesystemFreeMetrics
    {
        public static Task<GetMonitoringDropletFilesystemFreeMetricsResult> InvokeAsync(GetMonitoringDropletFilesystemFreeMetricsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMonitoringDropletFilesystemFreeMetricsResult>("digitalocean-native:monitoring/v2:getMonitoringDropletFilesystemFreeMetrics", args ?? new GetMonitoringDropletFilesystemFreeMetricsArgs(), options.WithDefaults());

        public static Output<GetMonitoringDropletFilesystemFreeMetricsResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMonitoringDropletFilesystemFreeMetricsResult>("digitalocean-native:monitoring/v2:getMonitoringDropletFilesystemFreeMetrics", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class GetMonitoringDropletFilesystemFreeMetricsArgs : global::Pulumi.InvokeArgs
    {
        public GetMonitoringDropletFilesystemFreeMetricsArgs()
        {
        }
        public static new GetMonitoringDropletFilesystemFreeMetricsArgs Empty => new GetMonitoringDropletFilesystemFreeMetricsArgs();
    }


    [OutputType]
    public sealed class GetMonitoringDropletFilesystemFreeMetricsResult
    {
        public readonly Outputs.Metrics Items;

        [OutputConstructor]
        private GetMonitoringDropletFilesystemFreeMetricsResult(Outputs.Metrics items)
        {
            Items = items;
        }
    }
}