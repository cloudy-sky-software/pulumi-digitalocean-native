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
    [DigitalOceanNativeResourceType("digitalocean-native:databases/v2:DatabasesOnlineMigration")]
    public partial class DatabasesOnlineMigration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enables SSL encryption when connecting to the source database.
        /// </summary>
        [Output("disableSsl")]
        public Output<bool?> DisableSsl { get; private set; } = null!;

        [Output("source")]
        public Output<Outputs.SourceProperties?> Source { get; private set; } = null!;


        /// <summary>
        /// Create a DatabasesOnlineMigration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatabasesOnlineMigration(string name, DatabasesOnlineMigrationArgs? args = null, CustomResourceOptions? options = null)
            : base("digitalocean-native:databases/v2:DatabasesOnlineMigration", name, args ?? new DatabasesOnlineMigrationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatabasesOnlineMigration(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("digitalocean-native:databases/v2:DatabasesOnlineMigration", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing DatabasesOnlineMigration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatabasesOnlineMigration Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DatabasesOnlineMigration(name, id, options);
        }
    }

    public sealed class DatabasesOnlineMigrationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique identifier for a database cluster.
        /// </summary>
        [Input("databaseClusterUuid")]
        public Input<string>? DatabaseClusterUuid { get; set; }

        /// <summary>
        /// Enables SSL encryption when connecting to the source database.
        /// </summary>
        [Input("disableSsl")]
        public Input<bool>? DisableSsl { get; set; }

        [Input("source")]
        public Input<Inputs.SourcePropertiesArgs>? Source { get; set; }

        public DatabasesOnlineMigrationArgs()
        {
        }
        public static new DatabasesOnlineMigrationArgs Empty => new DatabasesOnlineMigrationArgs();
    }
}