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
    public static class GetAppsLog
    {
        public static Task<Outputs.AppsGetLogsResponse> InvokeAsync(GetAppsLogArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.AppsGetLogsResponse>("digitalocean-native:apps/v2:getAppsLog", args ?? new GetAppsLogArgs(), options.WithDefaults());

        public static Output<Outputs.AppsGetLogsResponse> Invoke(GetAppsLogInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.AppsGetLogsResponse>("digitalocean-native:apps/v2:getAppsLog", args ?? new GetAppsLogInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppsLogArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The app ID
        /// </summary>
        [Input("appId", required: true)]
        public string AppId { get; set; } = null!;

        /// <summary>
        /// An optional component name. If set, logs will be limited to this component only.
        /// </summary>
        [Input("componentName", required: true)]
        public string ComponentName { get; set; } = null!;

        /// <summary>
        /// The deployment ID
        /// </summary>
        [Input("deploymentId", required: true)]
        public string DeploymentId { get; set; } = null!;

        public GetAppsLogArgs()
        {
        }
        public static new GetAppsLogArgs Empty => new GetAppsLogArgs();
    }

    public sealed class GetAppsLogInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The app ID
        /// </summary>
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        /// <summary>
        /// An optional component name. If set, logs will be limited to this component only.
        /// </summary>
        [Input("componentName", required: true)]
        public Input<string> ComponentName { get; set; } = null!;

        /// <summary>
        /// The deployment ID
        /// </summary>
        [Input("deploymentId", required: true)]
        public Input<string> DeploymentId { get; set; } = null!;

        public GetAppsLogInvokeArgs()
        {
        }
        public static new GetAppsLogInvokeArgs Empty => new GetAppsLogInvokeArgs();
    }
}
