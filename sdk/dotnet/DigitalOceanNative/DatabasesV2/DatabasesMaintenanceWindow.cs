// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.DatabasesV2
{
    [DigitalOceanNativeResourceType("digitalocean-native:databases/v2:DatabasesMaintenanceWindow")]
    public partial class DatabasesMaintenanceWindow : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The day of the week on which to apply maintenance updates.
        /// </summary>
        [Output("day")]
        public Output<string> Day { get; private set; } = null!;

        /// <summary>
        /// A list of strings, each containing information about a pending maintenance update.
        /// </summary>
        [Output("description")]
        public Output<ImmutableArray<string>> Description { get; private set; } = null!;

        /// <summary>
        /// The hour in UTC at which maintenance updates will be applied in 24 hour format.
        /// </summary>
        [Output("hour")]
        public Output<string> Hour { get; private set; } = null!;

        /// <summary>
        /// A boolean value indicating whether any maintenance is scheduled to be performed in the next window.
        /// </summary>
        [Output("pending")]
        public Output<bool?> Pending { get; private set; } = null!;


        /// <summary>
        /// Create a DatabasesMaintenanceWindow resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatabasesMaintenanceWindow(string name, DatabasesMaintenanceWindowArgs args, CustomResourceOptions? options = null)
            : base("digitalocean-native:databases/v2:DatabasesMaintenanceWindow", name, args ?? new DatabasesMaintenanceWindowArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatabasesMaintenanceWindow(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("digitalocean-native:databases/v2:DatabasesMaintenanceWindow", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing DatabasesMaintenanceWindow resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatabasesMaintenanceWindow Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DatabasesMaintenanceWindow(name, id, options);
        }
    }

    public sealed class DatabasesMaintenanceWindowArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique identifier for a database cluster.
        /// </summary>
        [Input("databaseClusterUuid")]
        public Input<string>? DatabaseClusterUuid { get; set; }

        /// <summary>
        /// The day of the week on which to apply maintenance updates.
        /// </summary>
        [Input("day", required: true)]
        public Input<string> Day { get; set; } = null!;

        /// <summary>
        /// The hour in UTC at which maintenance updates will be applied in 24 hour format.
        /// </summary>
        [Input("hour", required: true)]
        public Input<string> Hour { get; set; } = null!;

        public DatabasesMaintenanceWindowArgs()
        {
        }
        public static new DatabasesMaintenanceWindowArgs Empty => new DatabasesMaintenanceWindowArgs();
    }
}
