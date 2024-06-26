// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.ReservedIpsV2
{
    [DigitalOceanNativeResourceType("digitalocean-native:reserved_ips/v2:ReservedIP")]
    public partial class ReservedIP : global::Pulumi.CustomResource
    {
        [Output("links")]
        public Output<Outputs.LinksProperties?> Links { get; private set; } = null!;

        /// <summary>
        /// The UUID of the project to which the reserved IP will be assigned.
        /// </summary>
        [Output("projectId")]
        public Output<string?> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The slug identifier for the region the reserved IP will be reserved to.
        /// </summary>
        [Output("region")]
        public Output<string?> Region { get; private set; } = null!;

        [Output("reservedIp")]
        public Output<Outputs.ReservedIp?> ReservedIp { get; private set; } = null!;


        /// <summary>
        /// Create a ReservedIP resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReservedIP(string name, ReservedIPArgs? args = null, CustomResourceOptions? options = null)
            : base("digitalocean-native:reserved_ips/v2:ReservedIP", name, args ?? new ReservedIPArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReservedIP(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("digitalocean-native:reserved_ips/v2:ReservedIP", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ReservedIP resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReservedIP Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ReservedIP(name, id, options);
        }
    }

    public sealed class ReservedIPArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The UUID of the project to which the reserved IP will be assigned.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The slug identifier for the region the reserved IP will be reserved to.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public ReservedIPArgs()
        {
        }
        public static new ReservedIPArgs Empty => new ReservedIPArgs();
    }
}
