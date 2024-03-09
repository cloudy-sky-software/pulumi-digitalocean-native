// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.DomainsV2
{
    public static class GetDomainsRecord
    {
        public static Task<GetDomainsRecordResult> InvokeAsync(GetDomainsRecordArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDomainsRecordResult>("digitalocean-native:domains/v2:getDomainsRecord", args ?? new GetDomainsRecordArgs(), options.WithDefaults());

        public static Output<GetDomainsRecordResult> Invoke(GetDomainsRecordInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDomainsRecordResult>("digitalocean-native:domains/v2:getDomainsRecord", args ?? new GetDomainsRecordInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDomainsRecordArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the domain itself.
        /// </summary>
        [Input("domainName", required: true)]
        public string DomainName { get; set; } = null!;

        /// <summary>
        /// The unique identifier of the domain record.
        /// </summary>
        [Input("domainRecordId", required: true)]
        public string DomainRecordId { get; set; } = null!;

        public GetDomainsRecordArgs()
        {
        }
        public static new GetDomainsRecordArgs Empty => new GetDomainsRecordArgs();
    }

    public sealed class GetDomainsRecordInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the domain itself.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        /// <summary>
        /// The unique identifier of the domain record.
        /// </summary>
        [Input("domainRecordId", required: true)]
        public Input<string> DomainRecordId { get; set; } = null!;

        public GetDomainsRecordInvokeArgs()
        {
        }
        public static new GetDomainsRecordInvokeArgs Empty => new GetDomainsRecordInvokeArgs();
    }


    [OutputType]
    public sealed class GetDomainsRecordResult
    {
        public readonly Outputs.GetDomainsRecordProperties Items;

        [OutputConstructor]
        private GetDomainsRecordResult(Outputs.GetDomainsRecordProperties items)
        {
            Items = items;
        }
    }
}