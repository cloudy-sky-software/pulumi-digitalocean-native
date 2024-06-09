// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.DropletsV2
{
    [DigitalOceanNativeResourceType("digitalocean-native:droplets/v2:Rebuild")]
    public partial class Rebuild : global::Pulumi.CustomResource
    {
        [Output("action")]
        public Output<Outputs.Action?> Action { get; private set; } = null!;

        /// <summary>
        /// The image ID of a public or private image or the slug identifier for a public image. The Droplet will be rebuilt using this image as its base.
        /// </summary>
        [Output("image")]
        public Output<Union<string, int>?> Image { get; private set; } = null!;

        /// <summary>
        /// The type of action to initiate for the Droplet.
        /// </summary>
        [Output("type")]
        public Output<CloudySkySoftware.Pulumi.DigitalOceanNative.DropletsV2.DropletActionType?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Rebuild resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Rebuild(string name, RebuildArgs args, CustomResourceOptions? options = null)
            : base("digitalocean-native:droplets/v2:Rebuild", name, args ?? new RebuildArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Rebuild(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("digitalocean-native:droplets/v2:Rebuild", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Rebuild resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Rebuild Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Rebuild(name, id, options);
        }
    }

    public sealed class RebuildArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique identifier for a Droplet instance.
        /// </summary>
        [Input("dropletId")]
        public Input<string>? DropletId { get; set; }

        /// <summary>
        /// The image ID of a public or private image or the slug identifier for a public image. The Droplet will be rebuilt using this image as its base.
        /// </summary>
        [Input("image")]
        public InputUnion<string, int>? Image { get; set; }

        /// <summary>
        /// The type of action to initiate for the Droplet.
        /// </summary>
        [Input("type", required: true)]
        public Input<CloudySkySoftware.Pulumi.DigitalOceanNative.DropletsV2.DropletActionType> Type { get; set; } = null!;

        public RebuildArgs()
        {
        }
        public static new RebuildArgs Empty => new RebuildArgs();
    }
}