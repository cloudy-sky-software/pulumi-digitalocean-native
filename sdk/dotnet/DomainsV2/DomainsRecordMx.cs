// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.DomainsV2
{
    [DigitalOceanNativeResourceType("digitalocean-native:domains/v2:DomainsRecordMx")]
    public partial class DomainsRecordMx : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Variable data depending on record type. For example, the "data" value for an A record would be the IPv4 address to which the domain will be mapped. For a CAA record, it would contain the domain name of the CA being granted permission to issue certificates.
        /// </summary>
        [Output("data")]
        public Output<string?> Data { get; private set; } = null!;

        [Output("domainRecord")]
        public Output<Outputs.DomainRecord?> DomainRecord { get; private set; } = null!;

        /// <summary>
        /// An unsigned integer between 0-255 used for CAA records.
        /// </summary>
        [Output("flags")]
        public Output<int?> Flags { get; private set; } = null!;

        /// <summary>
        /// The host name, alias, or service being defined by the record.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// The port for SRV records.
        /// </summary>
        [Output("port")]
        public Output<int?> Port { get; private set; } = null!;

        /// <summary>
        /// The priority for SRV and MX records.
        /// </summary>
        [Output("priority")]
        public Output<int?> Priority { get; private set; } = null!;

        /// <summary>
        /// The parameter tag for CAA records. Valid values are "issue", "issuewild", or "iodef"
        /// </summary>
        [Output("tag")]
        public Output<string?> Tag { get; private set; } = null!;

        /// <summary>
        /// This value is the time to live for the record, in seconds. This defines the time frame that clients can cache queried information before a refresh should be requested.
        /// </summary>
        [Output("ttl")]
        public Output<int?> Ttl { get; private set; } = null!;

        /// <summary>
        /// The type of the DNS record. For example: A, CNAME, TXT, ...
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;

        /// <summary>
        /// The weight for SRV records.
        /// </summary>
        [Output("weight")]
        public Output<int?> Weight { get; private set; } = null!;


        /// <summary>
        /// Create a DomainsRecordMx resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DomainsRecordMx(string name, DomainsRecordMxArgs args, CustomResourceOptions? options = null)
            : base("digitalocean-native:domains/v2:DomainsRecordMx", name, args ?? new DomainsRecordMxArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DomainsRecordMx(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("digitalocean-native:domains/v2:DomainsRecordMx", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/cloudy-sky-software/pulumi-digitalocean-native",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DomainsRecordMx resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DomainsRecordMx Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DomainsRecordMx(name, id, options);
        }
    }

    public sealed class DomainsRecordMxArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Variable data depending on record type. For example, the "data" value for an A record would be the IPv4 address to which the domain will be mapped. For a CAA record, it would contain the domain name of the CA being granted permission to issue certificates.
        /// </summary>
        [Input("data")]
        public Input<string>? Data { get; set; }

        /// <summary>
        /// The name of the domain itself.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// An unsigned integer between 0-255 used for CAA records.
        /// </summary>
        [Input("flags")]
        public Input<int>? Flags { get; set; }

        /// <summary>
        /// The host name, alias, or service being defined by the record.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The port for SRV records.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The priority for SRV and MX records.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The parameter tag for CAA records. Valid values are "issue", "issuewild", or "iodef"
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        /// <summary>
        /// This value is the time to live for the record, in seconds. This defines the time frame that clients can cache queried information before a refresh should be requested.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// The type of the DNS record. For example: A, CNAME, TXT, ...
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The weight for SRV records.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public DomainsRecordMxArgs()
        {
        }
        public static new DomainsRecordMxArgs Empty => new DomainsRecordMxArgs();
    }
}
