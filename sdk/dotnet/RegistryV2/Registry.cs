// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.RegistryV2
{
    [DigitalOceanNativeResourceType("digitalocean-native:registry/v2:Registry")]
    public partial class Registry : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A globally unique name for the container registry. Must be lowercase and be composed only of numbers, letters and `-`, up to a limit of 63 characters.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Slug of the region where registry data is stored. When not provided, a region will be selected.
        /// </summary>
        [Output("region")]
        public Output<Pulumi.DigitalOceanNative.RegistryV2.Region?> Region { get; private set; } = null!;

        [Output("registry")]
        public Output<Outputs.Registry?> RegistryValue { get; private set; } = null!;

        /// <summary>
        /// The slug of the subscription tier to sign up for. Valid values can be retrieved using the options endpoint.
        /// </summary>
        [Output("subscriptionTierSlug")]
        public Output<Pulumi.DigitalOceanNative.RegistryV2.SubscriptionTierSlug> SubscriptionTierSlug { get; private set; } = null!;


        /// <summary>
        /// Create a Registry resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Registry(string name, RegistryArgs args, CustomResourceOptions? options = null)
            : base("digitalocean-native:registry/v2:Registry", name, args ?? new RegistryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Registry(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("digitalocean-native:registry/v2:Registry", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Registry resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Registry Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Registry(name, id, options);
        }
    }

    public sealed class RegistryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A globally unique name for the container registry. Must be lowercase and be composed only of numbers, letters and `-`, up to a limit of 63 characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Slug of the region where registry data is stored. When not provided, a region will be selected.
        /// </summary>
        [Input("region")]
        public Input<Pulumi.DigitalOceanNative.RegistryV2.Region>? Region { get; set; }

        /// <summary>
        /// The slug of the subscription tier to sign up for. Valid values can be retrieved using the options endpoint.
        /// </summary>
        [Input("subscriptionTierSlug", required: true)]
        public Input<Pulumi.DigitalOceanNative.RegistryV2.SubscriptionTierSlug> SubscriptionTierSlug { get; set; } = null!;

        public RegistryArgs()
        {
        }
        public static new RegistryArgs Empty => new RegistryArgs();
    }
}
