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
    public static class GetDatabase
    {
        public static Task<Outputs.GetDatabaseProperties> InvokeAsync(GetDatabaseArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetDatabaseProperties>("digitalocean-native:databases/v2:getDatabase", args ?? new GetDatabaseArgs(), options.WithDefaults());

        public static Output<Outputs.GetDatabaseProperties> Invoke(GetDatabaseInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetDatabaseProperties>("digitalocean-native:databases/v2:getDatabase", args ?? new GetDatabaseInvokeArgs(), options.WithDefaults());

        public static Output<Outputs.GetDatabaseProperties> Invoke(GetDatabaseInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetDatabaseProperties>("digitalocean-native:databases/v2:getDatabase", args ?? new GetDatabaseInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabaseArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a database cluster.
        /// </summary>
        [Input("databaseClusterUuid", required: true)]
        public string DatabaseClusterUuid { get; set; } = null!;

        /// <summary>
        /// The name of the database.
        /// </summary>
        [Input("databaseName", required: true)]
        public string DatabaseName { get; set; } = null!;

        public GetDatabaseArgs()
        {
        }
        public static new GetDatabaseArgs Empty => new GetDatabaseArgs();
    }

    public sealed class GetDatabaseInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a database cluster.
        /// </summary>
        [Input("databaseClusterUuid", required: true)]
        public Input<string> DatabaseClusterUuid { get; set; } = null!;

        /// <summary>
        /// The name of the database.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        public GetDatabaseInvokeArgs()
        {
        }
        public static new GetDatabaseInvokeArgs Empty => new GetDatabaseInvokeArgs();
    }
}
