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
    public static class GetDatabasesMigrationStatu
    {
        public static Task<Outputs.OnlineMigration> InvokeAsync(GetDatabasesMigrationStatuArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.OnlineMigration>("digitalocean-native:databases/v2:getDatabasesMigrationStatu", args ?? new GetDatabasesMigrationStatuArgs(), options.WithDefaults());

        public static Output<Outputs.OnlineMigration> Invoke(GetDatabasesMigrationStatuInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.OnlineMigration>("digitalocean-native:databases/v2:getDatabasesMigrationStatu", args ?? new GetDatabasesMigrationStatuInvokeArgs(), options.WithDefaults());

        public static Output<Outputs.OnlineMigration> Invoke(GetDatabasesMigrationStatuInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.OnlineMigration>("digitalocean-native:databases/v2:getDatabasesMigrationStatu", args ?? new GetDatabasesMigrationStatuInvokeArgs(), options.WithDefaults());
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
}
