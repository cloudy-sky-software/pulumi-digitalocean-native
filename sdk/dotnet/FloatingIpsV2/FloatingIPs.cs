// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.FloatingIpsV2
{
    [DigitalOceanNativeResourceType("digitalocean-native:floating_ips/v2:FloatingIPs")]
    public partial class FloatingIPs : global::Pulumi.CustomResource
    {
        [Output("floatingIp")]
        public Output<Outputs.FloatingIp?> FloatingIp { get; private set; } = null!;

        [Output("links")]
        public Output<Outputs.LinksProperties?> Links { get; private set; } = null!;


        /// <summary>
        /// Create a FloatingIPs resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FloatingIPs(string name, FloatingIPsArgs? args = null, CustomResourceOptions? options = null)
            : base("digitalocean-native:floating_ips/v2:FloatingIPs", name, args ?? new FloatingIPsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FloatingIPs(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("digitalocean-native:floating_ips/v2:FloatingIPs", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing FloatingIPs resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FloatingIPs Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new FloatingIPs(name, id, options);
        }
    }

    public sealed class FloatingIPsArgs : global::Pulumi.ResourceArgs
    {
        public FloatingIPsArgs()
        {
        }
        public static new FloatingIPsArgs Empty => new FloatingIPsArgs();
    }
}