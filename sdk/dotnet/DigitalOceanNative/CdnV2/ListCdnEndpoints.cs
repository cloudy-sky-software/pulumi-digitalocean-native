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
    public static class ListCdnEndpoints
    {
        public static Task<ListCdnEndpointsResult> InvokeAsync(ListCdnEndpointsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListCdnEndpointsResult>("digitalocean-native:cdn/v2:listCdnEndpoints", args ?? new ListCdnEndpointsArgs(), options.WithDefaults());

        public static Output<ListCdnEndpointsResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListCdnEndpointsResult>("digitalocean-native:cdn/v2:listCdnEndpoints", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListCdnEndpointsArgs : global::Pulumi.InvokeArgs
    {
        public ListCdnEndpointsArgs()
        {
        }
        public static new ListCdnEndpointsArgs Empty => new ListCdnEndpointsArgs();
    }


    [OutputType]
    public sealed class ListCdnEndpointsResult
    {
        public readonly Outputs.ListCdnEndpointsItems Items;

        [OutputConstructor]
        private ListCdnEndpointsResult(Outputs.ListCdnEndpointsItems items)
        {
            Items = items;
        }
    }
}