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
    public static class ListAppsDeployments
    {
        public static Task<Outputs.AppsDeploymentsResponse> InvokeAsync(ListAppsDeploymentsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.AppsDeploymentsResponse>("digitalocean-native:apps/v2:listAppsDeployments", args ?? new ListAppsDeploymentsArgs(), options.WithDefaults());

        public static Output<Outputs.AppsDeploymentsResponse> Invoke(ListAppsDeploymentsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.AppsDeploymentsResponse>("digitalocean-native:apps/v2:listAppsDeployments", args ?? new ListAppsDeploymentsInvokeArgs(), options.WithDefaults());

        public static Output<Outputs.AppsDeploymentsResponse> Invoke(ListAppsDeploymentsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.AppsDeploymentsResponse>("digitalocean-native:apps/v2:listAppsDeployments", args ?? new ListAppsDeploymentsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListAppsDeploymentsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The app ID
        /// </summary>
        [Input("appId", required: true)]
        public string AppId { get; set; } = null!;

        public ListAppsDeploymentsArgs()
        {
        }
        public static new ListAppsDeploymentsArgs Empty => new ListAppsDeploymentsArgs();
    }

    public sealed class ListAppsDeploymentsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The app ID
        /// </summary>
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        public ListAppsDeploymentsInvokeArgs()
        {
        }
        public static new ListAppsDeploymentsInvokeArgs Empty => new ListAppsDeploymentsInvokeArgs();
    }
}
