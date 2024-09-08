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
    public static class GetMonitoringDropletLoad5Metric
    {
        public static Task<Outputs.Metrics> InvokeAsync(GetMonitoringDropletLoad5MetricArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.Metrics>("digitalocean-native:monitoring/v2:getMonitoringDropletLoad5Metric", args ?? new GetMonitoringDropletLoad5MetricArgs(), options.WithDefaults());

        public static Output<Outputs.Metrics> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.Metrics>("digitalocean-native:monitoring/v2:getMonitoringDropletLoad5Metric", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class GetMonitoringDropletLoad5MetricArgs : global::Pulumi.InvokeArgs
    {
        public GetMonitoringDropletLoad5MetricArgs()
        {
        }
        public static new GetMonitoringDropletLoad5MetricArgs Empty => new GetMonitoringDropletLoad5MetricArgs();
    }
}
