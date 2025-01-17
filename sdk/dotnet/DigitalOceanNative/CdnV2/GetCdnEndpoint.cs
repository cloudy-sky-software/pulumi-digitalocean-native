// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.CdnV2
{
    public static class GetCdnEndpoint
    {
        public static Task<Outputs.GetCdnEndpointProperties> InvokeAsync(GetCdnEndpointArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetCdnEndpointProperties>("digitalocean-native:cdn/v2:getCdnEndpoint", args ?? new GetCdnEndpointArgs(), options.WithDefaults());

        public static Output<Outputs.GetCdnEndpointProperties> Invoke(GetCdnEndpointInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetCdnEndpointProperties>("digitalocean-native:cdn/v2:getCdnEndpoint", args ?? new GetCdnEndpointInvokeArgs(), options.WithDefaults());

        public static Output<Outputs.GetCdnEndpointProperties> Invoke(GetCdnEndpointInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetCdnEndpointProperties>("digitalocean-native:cdn/v2:getCdnEndpoint", args ?? new GetCdnEndpointInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCdnEndpointArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a CDN endpoint.
        /// </summary>
        [Input("cdnId", required: true)]
        public string CdnId { get; set; } = null!;

        public GetCdnEndpointArgs()
        {
        }
        public static new GetCdnEndpointArgs Empty => new GetCdnEndpointArgs();
    }

    public sealed class GetCdnEndpointInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a CDN endpoint.
        /// </summary>
        [Input("cdnId", required: true)]
        public Input<string> CdnId { get; set; } = null!;

        public GetCdnEndpointInvokeArgs()
        {
        }
        public static new GetCdnEndpointInvokeArgs Empty => new GetCdnEndpointInvokeArgs();
    }
}
