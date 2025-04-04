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
    public static class GetApp
    {
        public static Task<Outputs.AppResponse> InvokeAsync(GetAppArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.AppResponse>("digitalocean-native:apps/v2:getApp", args ?? new GetAppArgs(), options.WithDefaults());

        public static Output<Outputs.AppResponse> Invoke(GetAppInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.AppResponse>("digitalocean-native:apps/v2:getApp", args ?? new GetAppInvokeArgs(), options.WithDefaults());

        public static Output<Outputs.AppResponse> Invoke(GetAppInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.AppResponse>("digitalocean-native:apps/v2:getApp", args ?? new GetAppInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the app
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetAppArgs()
        {
        }
        public static new GetAppArgs Empty => new GetAppArgs();
    }

    public sealed class GetAppInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the app
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetAppInvokeArgs()
        {
        }
        public static new GetAppInvokeArgs Empty => new GetAppInvokeArgs();
    }
}
