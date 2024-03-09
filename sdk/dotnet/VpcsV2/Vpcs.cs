// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.VpcsV2
{
    [DigitalOceanNativeResourceType("digitalocean-native:vpcs/v2:Vpcs")]
    public partial class Vpcs : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A free-form text field for describing the VPC's purpose. It may be a maximum of 255 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The range of IP addresses in the VPC in CIDR notation. Network ranges cannot overlap with other networks in the same account and must be in range of private addresses as defined in RFC1918. It may not be smaller than `/28` nor larger than `/16`. If no IP range is specified, a `/20` network range is generated that won't conflict with other VPC networks in your account.
        /// </summary>
        [Output("ipRange")]
        public Output<string?> IpRange { get; private set; } = null!;

        /// <summary>
        /// The name of the VPC. Must be unique and may only contain alphanumeric characters, dashes, and periods.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The slug identifier for the region where the VPC will be created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        [Output("vpc")]
        public Output<Outputs.Vpc?> Vpc { get; private set; } = null!;


        /// <summary>
        /// Create a Vpcs resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Vpcs(string name, VpcsArgs? args = null, CustomResourceOptions? options = null)
            : base("digitalocean-native:vpcs/v2:Vpcs", name, args ?? new VpcsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Vpcs(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("digitalocean-native:vpcs/v2:Vpcs", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Vpcs resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Vpcs Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Vpcs(name, id, options);
        }
    }

    public sealed class VpcsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A free-form text field for describing the VPC's purpose. It may be a maximum of 255 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The range of IP addresses in the VPC in CIDR notation. Network ranges cannot overlap with other networks in the same account and must be in range of private addresses as defined in RFC1918. It may not be smaller than `/28` nor larger than `/16`. If no IP range is specified, a `/20` network range is generated that won't conflict with other VPC networks in your account.
        /// </summary>
        [Input("ipRange")]
        public Input<string>? IpRange { get; set; }

        /// <summary>
        /// The name of the VPC. Must be unique and may only contain alphanumeric characters, dashes, and periods.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The slug identifier for the region where the VPC will be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public VpcsArgs()
        {
        }
        public static new VpcsArgs Empty => new VpcsArgs();
    }
}