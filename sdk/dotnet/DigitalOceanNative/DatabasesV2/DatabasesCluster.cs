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
    [DigitalOceanNativeResourceType("digitalocean-native:databases/v2:DatabasesCluster")]
    public partial class DatabasesCluster : global::Pulumi.CustomResource
    {
        [Output("backupRestore")]
        public Output<Outputs.DatabaseBackup?> BackupRestore { get; private set; } = null!;

        [Output("connection")]
        public Output<Outputs.DatabaseClusterConnection?> Connection { get; private set; } = null!;

        /// <summary>
        /// A time value given in ISO8601 combined date and time format that represents when the database cluster was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string?> CreatedAt { get; private set; } = null!;

        [Output("database")]
        public Output<Outputs.DatabaseCluster> Database { get; private set; } = null!;

        /// <summary>
        /// An array of strings containing the names of databases created in the database cluster.
        /// </summary>
        [Output("dbNames")]
        public Output<ImmutableArray<string>> DbNames { get; private set; } = null!;

        /// <summary>
        /// A slug representing the database engine used for the cluster. The possible values are: "pg" for PostgreSQL, "mysql" for MySQL, "redis" for Redis, and "mongodb" for MongoDB.
        /// </summary>
        [Output("engine")]
        public Output<CloudySkySoftware.Pulumi.DigitalOceanNative.DatabasesV2.DatabaseClusterEngine?> Engine { get; private set; } = null!;

        [Output("maintenanceWindow")]
        public Output<Outputs.DatabaseClusterMaintenanceWindow?> MaintenanceWindow { get; private set; } = null!;

        /// <summary>
        /// A unique, human-readable name referring to a database cluster.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// The number of nodes in the database cluster.
        /// </summary>
        [Output("numNodes")]
        public Output<int?> NumNodes { get; private set; } = null!;

        [Output("privateConnection")]
        public Output<Outputs.DatabaseClusterPrivateConnection?> PrivateConnection { get; private set; } = null!;

        /// <summary>
        /// A string specifying the UUID of the VPC to which the database cluster will be assigned. If excluded, the cluster when creating a new database cluster, it will be assigned to your account's default VPC for the region.
        /// </summary>
        [Output("privateNetworkUuid")]
        public Output<string?> PrivateNetworkUuid { get; private set; } = null!;

        /// <summary>
        /// The ID of the project that the database cluster is assigned to. If excluded when creating a new database cluster, it will be assigned to your default project.
        /// </summary>
        [Output("projectId")]
        public Output<string?> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The slug identifier for the region where the database cluster is located.
        /// </summary>
        [Output("region")]
        public Output<string?> Region { get; private set; } = null!;

        [Output("rules")]
        public Output<ImmutableArray<Outputs.FirewallRule>> Rules { get; private set; } = null!;

        /// <summary>
        /// A string representing the semantic version of the database engine in use for the cluster.
        /// </summary>
        [Output("semanticVersion")]
        public Output<string?> SemanticVersion { get; private set; } = null!;

        /// <summary>
        /// The slug identifier representing the size of the nodes in the database cluster.
        /// </summary>
        [Output("size")]
        public Output<string?> Size { get; private set; } = null!;

        /// <summary>
        /// A string representing the current status of the database cluster.
        /// </summary>
        [Output("status")]
        public Output<CloudySkySoftware.Pulumi.DigitalOceanNative.DatabasesV2.DatabaseClusterStatus?> Status { get; private set; } = null!;

        /// <summary>
        /// An array of tags that have been applied to the database cluster.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        [Output("users")]
        public Output<ImmutableArray<Outputs.DatabaseUser>> Users { get; private set; } = null!;

        /// <summary>
        /// A string representing the version of the database engine in use for the cluster.
        /// </summary>
        [Output("version")]
        public Output<string?> Version { get; private set; } = null!;

        /// <summary>
        /// A timestamp referring to the date when the particular version will no longer be available for creating new clusters. If null, the version does not have an end of availability timeline.
        /// </summary>
        [Output("versionEndOfAvailability")]
        public Output<string?> VersionEndOfAvailability { get; private set; } = null!;

        /// <summary>
        /// A timestamp referring to the date when the particular version will no longer be supported. If null, the version does not have an end of life timeline.
        /// </summary>
        [Output("versionEndOfLife")]
        public Output<string?> VersionEndOfLife { get; private set; } = null!;


        /// <summary>
        /// Create a DatabasesCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatabasesCluster(string name, DatabasesClusterArgs args, CustomResourceOptions? options = null)
            : base("digitalocean-native:databases/v2:DatabasesCluster", name, args ?? new DatabasesClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatabasesCluster(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("digitalocean-native:databases/v2:DatabasesCluster", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing DatabasesCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatabasesCluster Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DatabasesCluster(name, id, options);
        }
    }

    public sealed class DatabasesClusterArgs : global::Pulumi.ResourceArgs
    {
        [Input("backupRestore")]
        public Input<Inputs.DatabaseBackupArgs>? BackupRestore { get; set; }

        [Input("connection")]
        public Input<Inputs.DatabaseClusterConnectionArgs>? Connection { get; set; }

        /// <summary>
        /// A time value given in ISO8601 combined date and time format that represents when the database cluster was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("dbNames")]
        private InputList<string>? _dbNames;

        /// <summary>
        /// An array of strings containing the names of databases created in the database cluster.
        /// </summary>
        public InputList<string> DbNames
        {
            get => _dbNames ?? (_dbNames = new InputList<string>());
            set => _dbNames = value;
        }

        /// <summary>
        /// A slug representing the database engine used for the cluster. The possible values are: "pg" for PostgreSQL, "mysql" for MySQL, "redis" for Redis, and "mongodb" for MongoDB.
        /// </summary>
        [Input("engine", required: true)]
        public Input<CloudySkySoftware.Pulumi.DigitalOceanNative.DatabasesV2.DatabaseClusterEngine> Engine { get; set; } = null!;

        [Input("maintenanceWindow")]
        public Input<Inputs.DatabaseClusterMaintenanceWindowArgs>? MaintenanceWindow { get; set; }

        /// <summary>
        /// A unique, human-readable name referring to a database cluster.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The number of nodes in the database cluster.
        /// </summary>
        [Input("numNodes", required: true)]
        public Input<int> NumNodes { get; set; } = null!;

        [Input("privateConnection")]
        public Input<Inputs.DatabaseClusterPrivateConnectionArgs>? PrivateConnection { get; set; }

        /// <summary>
        /// A string specifying the UUID of the VPC to which the database cluster will be assigned. If excluded, the cluster when creating a new database cluster, it will be assigned to your account's default VPC for the region.
        /// </summary>
        [Input("privateNetworkUuid")]
        public Input<string>? PrivateNetworkUuid { get; set; }

        /// <summary>
        /// The ID of the project that the database cluster is assigned to. If excluded when creating a new database cluster, it will be assigned to your default project.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The slug identifier for the region where the database cluster is located.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("rules")]
        private InputList<Inputs.FirewallRuleArgs>? _rules;
        public InputList<Inputs.FirewallRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.FirewallRuleArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// A string representing the semantic version of the database engine in use for the cluster.
        /// </summary>
        [Input("semanticVersion")]
        public Input<string>? SemanticVersion { get; set; }

        /// <summary>
        /// The slug identifier representing the size of the nodes in the database cluster.
        /// </summary>
        [Input("size", required: true)]
        public Input<string> Size { get; set; } = null!;

        /// <summary>
        /// A string representing the current status of the database cluster.
        /// </summary>
        [Input("status")]
        public Input<CloudySkySoftware.Pulumi.DigitalOceanNative.DatabasesV2.DatabaseClusterStatus>? Status { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// An array of tags that have been applied to the database cluster.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("users")]
        private InputList<Inputs.DatabaseUserArgs>? _users;
        public InputList<Inputs.DatabaseUserArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.DatabaseUserArgs>());
            set => _users = value;
        }

        /// <summary>
        /// A string representing the version of the database engine in use for the cluster.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// A timestamp referring to the date when the particular version will no longer be available for creating new clusters. If null, the version does not have an end of availability timeline.
        /// </summary>
        [Input("versionEndOfAvailability")]
        public Input<string>? VersionEndOfAvailability { get; set; }

        /// <summary>
        /// A timestamp referring to the date when the particular version will no longer be supported. If null, the version does not have an end of life timeline.
        /// </summary>
        [Input("versionEndOfLife")]
        public Input<string>? VersionEndOfLife { get; set; }

        public DatabasesClusterArgs()
        {
        }
        public static new DatabasesClusterArgs Empty => new DatabasesClusterArgs();
    }
}
