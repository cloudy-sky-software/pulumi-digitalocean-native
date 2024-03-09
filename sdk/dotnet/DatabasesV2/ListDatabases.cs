// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.DatabasesV2
{
    public static class ListDatabases
    {
        public static Task<ListDatabasesResult> InvokeAsync(ListDatabasesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListDatabasesResult>("digitalocean-native:databases/v2:listDatabases", args ?? new ListDatabasesArgs(), options.WithDefaults());

        public static Output<ListDatabasesResult> Invoke(ListDatabasesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListDatabasesResult>("digitalocean-native:databases/v2:listDatabases", args ?? new ListDatabasesInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListDatabasesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a database cluster.
        /// </summary>
        [Input("databaseClusterUuid", required: true)]
        public string DatabaseClusterUuid { get; set; } = null!;

        public ListDatabasesArgs()
        {
        }
        public static new ListDatabasesArgs Empty => new ListDatabasesArgs();
    }

    public sealed class ListDatabasesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a database cluster.
        /// </summary>
        [Input("databaseClusterUuid", required: true)]
        public Input<string> DatabaseClusterUuid { get; set; } = null!;

        public ListDatabasesInvokeArgs()
        {
        }
        public static new ListDatabasesInvokeArgs Empty => new ListDatabasesInvokeArgs();
    }


    [OutputType]
    public sealed class ListDatabasesResult
    {
        public readonly Outputs.ListDatabasesProperties Items;

        [OutputConstructor]
        private ListDatabasesResult(Outputs.ListDatabasesProperties items)
        {
            Items = items;
        }
    }
}
