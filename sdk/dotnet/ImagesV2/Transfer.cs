// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.ImagesV2
{
    [DigitalOceanNativeResourceType("digitalocean-native:images/v2:Transfer")]
    public partial class Transfer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A time value given in ISO8601 combined date and time format that represents when the action was completed.
        /// </summary>
        [Output("completedAt")]
        public Output<string?> CompletedAt { get; private set; } = null!;

        [Output("region")]
        public Output<Outputs.Region?> Region { get; private set; } = null!;

        [Output("regionSlug")]
        public Output<Outputs.RegionSlug?> RegionSlug { get; private set; } = null!;

        /// <summary>
        /// A unique identifier for the resource that the action is associated with.
        /// </summary>
        [Output("resourceId")]
        public Output<int?> ResourceId { get; private set; } = null!;

        /// <summary>
        /// The type of resource that the action is associated with.
        /// </summary>
        [Output("resourceType")]
        public Output<string?> ResourceType { get; private set; } = null!;

        /// <summary>
        /// A time value given in ISO8601 combined date and time format that represents when the action was initiated.
        /// </summary>
        [Output("startedAt")]
        public Output<string?> StartedAt { get; private set; } = null!;

        /// <summary>
        /// The current status of the action. This can be "in-progress", "completed", or "errored".
        /// </summary>
        [Output("status")]
        public Output<Pulumi.DigitalOceanNative.ImagesV2.Status?> Status { get; private set; } = null!;

        /// <summary>
        /// This is the type of action that the object represents. For example, this could be "transfer" to represent the state of an image transfer action.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Transfer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Transfer(string name, TransferArgs args, CustomResourceOptions? options = null)
            : base("digitalocean-native:images/v2:Transfer", name, args ?? new TransferArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Transfer(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("digitalocean-native:images/v2:Transfer", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Transfer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Transfer Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Transfer(name, id, options);
        }
    }

    public sealed class TransferArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique number that can be used to identify and reference a specific image.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// The slug identifier for the region where the resource will initially be  available.
        /// </summary>
        [Input("region", required: true)]
        public Input<Pulumi.DigitalOceanNative.ImagesV2.TransferPropertiesRegion> Region { get; set; } = null!;

        /// <summary>
        /// The action to be taken on the image. Can be either `convert` or `transfer`.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.DigitalOceanNative.ImagesV2.ImageActionBaseType> Type { get; set; } = null!;

        public TransferArgs()
        {
        }
        public static new TransferArgs Empty => new TransferArgs();
    }
}
