// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.DomainsV2
{
    public static class GetDomains
    {
        public static Task<GetDomainsResult> InvokeAsync(GetDomainsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDomainsResult>("digitalocean-native:domains/v2:getDomains", args ?? new GetDomainsArgs(), options.WithDefaults());

        public static Output<GetDomainsResult> Invoke(GetDomainsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDomainsResult>("digitalocean-native:domains/v2:getDomains", args ?? new GetDomainsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDomainsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the domain itself.
        /// </summary>
        [Input("domainName", required: true)]
        public string DomainName { get; set; } = null!;

        public GetDomainsArgs()
        {
        }
        public static new GetDomainsArgs Empty => new GetDomainsArgs();
    }

    public sealed class GetDomainsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the domain itself.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        public GetDomainsInvokeArgs()
        {
        }
        public static new GetDomainsInvokeArgs Empty => new GetDomainsInvokeArgs();
    }


    [OutputType]
    public sealed class GetDomainsResult
    {
        public readonly Outputs.GetDomainsProperties Items;

        [OutputConstructor]
        private GetDomainsResult(Outputs.GetDomainsProperties items)
        {
            Items = items;
        }
    }
}
