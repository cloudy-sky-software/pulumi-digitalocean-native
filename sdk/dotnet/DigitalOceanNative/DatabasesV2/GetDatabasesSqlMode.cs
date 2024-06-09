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
    public static class GetDatabasesSqlMode
    {
        public static Task<GetDatabasesSqlModeResult> InvokeAsync(GetDatabasesSqlModeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatabasesSqlModeResult>("digitalocean-native:databases/v2:getDatabasesSqlMode", args ?? new GetDatabasesSqlModeArgs(), options.WithDefaults());

        public static Output<GetDatabasesSqlModeResult> Invoke(GetDatabasesSqlModeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabasesSqlModeResult>("digitalocean-native:databases/v2:getDatabasesSqlMode", args ?? new GetDatabasesSqlModeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabasesSqlModeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a database cluster.
        /// </summary>
        [Input("databaseClusterUuid", required: true)]
        public string DatabaseClusterUuid { get; set; } = null!;

        public GetDatabasesSqlModeArgs()
        {
        }
        public static new GetDatabasesSqlModeArgs Empty => new GetDatabasesSqlModeArgs();
    }

    public sealed class GetDatabasesSqlModeInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a database cluster.
        /// </summary>
        [Input("databaseClusterUuid", required: true)]
        public Input<string> DatabaseClusterUuid { get; set; } = null!;

        public GetDatabasesSqlModeInvokeArgs()
        {
        }
        public static new GetDatabasesSqlModeInvokeArgs Empty => new GetDatabasesSqlModeInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatabasesSqlModeResult
    {
        public readonly Outputs.SqlMode Items;

        [OutputConstructor]
        private GetDatabasesSqlModeResult(Outputs.SqlMode items)
        {
            Items = items;
        }
    }
}