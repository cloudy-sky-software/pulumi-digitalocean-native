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
    public static class GetDatabasesReplica
    {
        public static Task<GetDatabasesReplicaResult> InvokeAsync(GetDatabasesReplicaArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatabasesReplicaResult>("digitalocean-native:databases/v2:getDatabasesReplica", args ?? new GetDatabasesReplicaArgs(), options.WithDefaults());

        public static Output<GetDatabasesReplicaResult> Invoke(GetDatabasesReplicaInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabasesReplicaResult>("digitalocean-native:databases/v2:getDatabasesReplica", args ?? new GetDatabasesReplicaInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabasesReplicaArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a database cluster.
        /// </summary>
        [Input("databaseClusterUuid", required: true)]
        public string DatabaseClusterUuid { get; set; } = null!;

        /// <summary>
        /// The name of the database replica.
        /// </summary>
        [Input("replicaName", required: true)]
        public string ReplicaName { get; set; } = null!;

        public GetDatabasesReplicaArgs()
        {
        }
        public static new GetDatabasesReplicaArgs Empty => new GetDatabasesReplicaArgs();
    }

    public sealed class GetDatabasesReplicaInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a database cluster.
        /// </summary>
        [Input("databaseClusterUuid", required: true)]
        public Input<string> DatabaseClusterUuid { get; set; } = null!;

        /// <summary>
        /// The name of the database replica.
        /// </summary>
        [Input("replicaName", required: true)]
        public Input<string> ReplicaName { get; set; } = null!;

        public GetDatabasesReplicaInvokeArgs()
        {
        }
        public static new GetDatabasesReplicaInvokeArgs Empty => new GetDatabasesReplicaInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatabasesReplicaResult
    {
        public readonly Outputs.GetDatabasesReplicaProperties Items;

        [OutputConstructor]
        private GetDatabasesReplicaResult(Outputs.GetDatabasesReplicaProperties items)
        {
            Items = items;
        }
    }
}
