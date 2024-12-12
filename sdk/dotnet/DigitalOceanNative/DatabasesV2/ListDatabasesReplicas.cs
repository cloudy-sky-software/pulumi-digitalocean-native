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
    public static class ListDatabasesReplicas
    {
        public static Task<Outputs.ListDatabasesReplicasProperties> InvokeAsync(ListDatabasesReplicasArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.ListDatabasesReplicasProperties>("digitalocean-native:databases/v2:listDatabasesReplicas", args ?? new ListDatabasesReplicasArgs(), options.WithDefaults());

        public static Output<Outputs.ListDatabasesReplicasProperties> Invoke(ListDatabasesReplicasInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ListDatabasesReplicasProperties>("digitalocean-native:databases/v2:listDatabasesReplicas", args ?? new ListDatabasesReplicasInvokeArgs(), options.WithDefaults());

        public static Output<Outputs.ListDatabasesReplicasProperties> Invoke(ListDatabasesReplicasInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ListDatabasesReplicasProperties>("digitalocean-native:databases/v2:listDatabasesReplicas", args ?? new ListDatabasesReplicasInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListDatabasesReplicasArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a database cluster.
        /// </summary>
        [Input("databaseClusterUuid", required: true)]
        public string DatabaseClusterUuid { get; set; } = null!;

        public ListDatabasesReplicasArgs()
        {
        }
        public static new ListDatabasesReplicasArgs Empty => new ListDatabasesReplicasArgs();
    }

    public sealed class ListDatabasesReplicasInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a database cluster.
        /// </summary>
        [Input("databaseClusterUuid", required: true)]
        public Input<string> DatabaseClusterUuid { get; set; } = null!;

        public ListDatabasesReplicasInvokeArgs()
        {
        }
        public static new ListDatabasesReplicasInvokeArgs Empty => new ListDatabasesReplicasInvokeArgs();
    }
}
