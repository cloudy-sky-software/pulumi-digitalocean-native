// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.DatabasesV2
{
    public static class GetDatabasesMigrationStatu
    {
        public static Task<GetDatabasesMigrationStatuResult> InvokeAsync(GetDatabasesMigrationStatuArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatabasesMigrationStatuResult>("digitalocean-native:databases/v2:getDatabasesMigrationStatu", args ?? new GetDatabasesMigrationStatuArgs(), options.WithDefaults());

        public static Output<GetDatabasesMigrationStatuResult> Invoke(GetDatabasesMigrationStatuInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabasesMigrationStatuResult>("digitalocean-native:databases/v2:getDatabasesMigrationStatu", args ?? new GetDatabasesMigrationStatuInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabasesMigrationStatuArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a database cluster.
        /// </summary>
        [Input("databaseClusterUuid", required: true)]
        public string DatabaseClusterUuid { get; set; } = null!;

        public GetDatabasesMigrationStatuArgs()
        {
        }
        public static new GetDatabasesMigrationStatuArgs Empty => new GetDatabasesMigrationStatuArgs();
    }

    public sealed class GetDatabasesMigrationStatuInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a database cluster.
        /// </summary>
        [Input("databaseClusterUuid", required: true)]
        public Input<string> DatabaseClusterUuid { get; set; } = null!;

        public GetDatabasesMigrationStatuInvokeArgs()
        {
        }
        public static new GetDatabasesMigrationStatuInvokeArgs Empty => new GetDatabasesMigrationStatuInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatabasesMigrationStatuResult
    {
        public readonly Outputs.OnlineMigration Items;

        [OutputConstructor]
        private GetDatabasesMigrationStatuResult(Outputs.OnlineMigration items)
        {
            Items = items;
        }
    }
}