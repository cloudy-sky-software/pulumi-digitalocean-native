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
    public static class GetAppsTier
    {
        public static Task<Outputs.AppsGetTierResponse> InvokeAsync(GetAppsTierArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.AppsGetTierResponse>("digitalocean-native:apps/v2:getAppsTier", args ?? new GetAppsTierArgs(), options.WithDefaults());

        public static Output<Outputs.AppsGetTierResponse> Invoke(GetAppsTierInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.AppsGetTierResponse>("digitalocean-native:apps/v2:getAppsTier", args ?? new GetAppsTierInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppsTierArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The slug of the tier
        /// </summary>
        [Input("slug", required: true)]
        public string Slug { get; set; } = null!;

        public GetAppsTierArgs()
        {
        }
        public static new GetAppsTierArgs Empty => new GetAppsTierArgs();
    }

    public sealed class GetAppsTierInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The slug of the tier
        /// </summary>
        [Input("slug", required: true)]
        public Input<string> Slug { get; set; } = null!;

        public GetAppsTierInvokeArgs()
        {
        }
        public static new GetAppsTierInvokeArgs Empty => new GetAppsTierInvokeArgs();
    }
}
