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
    public static class GetMonitoringAlertPolicy
    {
        public static Task<Outputs.GetMonitoringAlertPolicyProperties> InvokeAsync(GetMonitoringAlertPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetMonitoringAlertPolicyProperties>("digitalocean-native:monitoring/v2:getMonitoringAlertPolicy", args ?? new GetMonitoringAlertPolicyArgs(), options.WithDefaults());

        public static Output<Outputs.GetMonitoringAlertPolicyProperties> Invoke(GetMonitoringAlertPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetMonitoringAlertPolicyProperties>("digitalocean-native:monitoring/v2:getMonitoringAlertPolicy", args ?? new GetMonitoringAlertPolicyInvokeArgs(), options.WithDefaults());

        public static Output<Outputs.GetMonitoringAlertPolicyProperties> Invoke(GetMonitoringAlertPolicyInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetMonitoringAlertPolicyProperties>("digitalocean-native:monitoring/v2:getMonitoringAlertPolicy", args ?? new GetMonitoringAlertPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMonitoringAlertPolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for an alert policy.
        /// </summary>
        [Input("alertUuid", required: true)]
        public string AlertUuid { get; set; } = null!;

        public GetMonitoringAlertPolicyArgs()
        {
        }
        public static new GetMonitoringAlertPolicyArgs Empty => new GetMonitoringAlertPolicyArgs();
    }

    public sealed class GetMonitoringAlertPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for an alert policy.
        /// </summary>
        [Input("alertUuid", required: true)]
        public Input<string> AlertUuid { get; set; } = null!;

        public GetMonitoringAlertPolicyInvokeArgs()
        {
        }
        public static new GetMonitoringAlertPolicyInvokeArgs Empty => new GetMonitoringAlertPolicyInvokeArgs();
    }
}
