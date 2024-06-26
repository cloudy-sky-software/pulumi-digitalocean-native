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
    public static class ListMonitoringAlertPolicy
    {
        public static Task<ListMonitoringAlertPolicyResult> InvokeAsync(ListMonitoringAlertPolicyArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListMonitoringAlertPolicyResult>("digitalocean-native:monitoring/v2:listMonitoringAlertPolicy", args ?? new ListMonitoringAlertPolicyArgs(), options.WithDefaults());

        public static Output<ListMonitoringAlertPolicyResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListMonitoringAlertPolicyResult>("digitalocean-native:monitoring/v2:listMonitoringAlertPolicy", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListMonitoringAlertPolicyArgs : global::Pulumi.InvokeArgs
    {
        public ListMonitoringAlertPolicyArgs()
        {
        }
        public static new ListMonitoringAlertPolicyArgs Empty => new ListMonitoringAlertPolicyArgs();
    }


    [OutputType]
    public sealed class ListMonitoringAlertPolicyResult
    {
        public readonly Outputs.ListMonitoringAlertPolicyItems Items;

        [OutputConstructor]
        private ListMonitoringAlertPolicyResult(Outputs.ListMonitoringAlertPolicyItems items)
        {
            Items = items;
        }
    }
}
