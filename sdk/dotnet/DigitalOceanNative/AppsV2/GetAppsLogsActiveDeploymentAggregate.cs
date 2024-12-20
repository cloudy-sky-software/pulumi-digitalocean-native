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
    public static class GetAppsLogsActiveDeploymentAggregate
    {
        public static Task<Outputs.AppsGetLogsResponse> InvokeAsync(GetAppsLogsActiveDeploymentAggregateArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.AppsGetLogsResponse>("digitalocean-native:apps/v2:getAppsLogsActiveDeploymentAggregate", args ?? new GetAppsLogsActiveDeploymentAggregateArgs(), options.WithDefaults());

        public static Output<Outputs.AppsGetLogsResponse> Invoke(GetAppsLogsActiveDeploymentAggregateInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.AppsGetLogsResponse>("digitalocean-native:apps/v2:getAppsLogsActiveDeploymentAggregate", args ?? new GetAppsLogsActiveDeploymentAggregateInvokeArgs(), options.WithDefaults());

        public static Output<Outputs.AppsGetLogsResponse> Invoke(GetAppsLogsActiveDeploymentAggregateInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.AppsGetLogsResponse>("digitalocean-native:apps/v2:getAppsLogsActiveDeploymentAggregate", args ?? new GetAppsLogsActiveDeploymentAggregateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppsLogsActiveDeploymentAggregateArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The app ID
        /// </summary>
        [Input("appId", required: true)]
        public string AppId { get; set; } = null!;

        public GetAppsLogsActiveDeploymentAggregateArgs()
        {
        }
        public static new GetAppsLogsActiveDeploymentAggregateArgs Empty => new GetAppsLogsActiveDeploymentAggregateArgs();
    }

    public sealed class GetAppsLogsActiveDeploymentAggregateInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The app ID
        /// </summary>
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        public GetAppsLogsActiveDeploymentAggregateInvokeArgs()
        {
        }
        public static new GetAppsLogsActiveDeploymentAggregateInvokeArgs Empty => new GetAppsLogsActiveDeploymentAggregateInvokeArgs();
    }
}
