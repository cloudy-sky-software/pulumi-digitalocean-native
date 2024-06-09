// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.FirewallsV2
{
    [DigitalOceanNativeResourceType("digitalocean-native:firewalls/v2:FirewallsTag")]
    public partial class FirewallsTag : global::Pulumi.CustomResource
    {
        [Output("tags")]
        public Output<Outputs.Tags> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallsTag resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallsTag(string name, FirewallsTagArgs args, CustomResourceOptions? options = null)
            : base("digitalocean-native:firewalls/v2:FirewallsTag", name, args ?? new FirewallsTagArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallsTag(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("digitalocean-native:firewalls/v2:FirewallsTag", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing FirewallsTag resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallsTag Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new FirewallsTag(name, id, options);
        }
    }

    public sealed class FirewallsTagArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique ID that can be used to identify and reference a firewall.
        /// </summary>
        [Input("firewallId")]
        public Input<string>? FirewallId { get; set; }

        [Input("tags", required: true)]
        public Input<Inputs.TagsArgs> Tags { get; set; } = null!;

        public FirewallsTagArgs()
        {
        }
        public static new FirewallsTagArgs Empty => new FirewallsTagArgs();
    }
}