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
    public static class GetDatabasesUser
    {
        public static Task<Outputs.GetDatabasesUserProperties> InvokeAsync(GetDatabasesUserArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetDatabasesUserProperties>("digitalocean-native:databases/v2:getDatabasesUser", args ?? new GetDatabasesUserArgs(), options.WithDefaults());

        public static Output<Outputs.GetDatabasesUserProperties> Invoke(GetDatabasesUserInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetDatabasesUserProperties>("digitalocean-native:databases/v2:getDatabasesUser", args ?? new GetDatabasesUserInvokeArgs(), options.WithDefaults());

        public static Output<Outputs.GetDatabasesUserProperties> Invoke(GetDatabasesUserInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetDatabasesUserProperties>("digitalocean-native:databases/v2:getDatabasesUser", args ?? new GetDatabasesUserInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabasesUserArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a database cluster.
        /// </summary>
        [Input("databaseClusterUuid", required: true)]
        public string DatabaseClusterUuid { get; set; } = null!;

        /// <summary>
        /// The name of the database user.
        /// </summary>
        [Input("username", required: true)]
        public string Username { get; set; } = null!;

        public GetDatabasesUserArgs()
        {
        }
        public static new GetDatabasesUserArgs Empty => new GetDatabasesUserArgs();
    }

    public sealed class GetDatabasesUserInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a database cluster.
        /// </summary>
        [Input("databaseClusterUuid", required: true)]
        public Input<string> DatabaseClusterUuid { get; set; } = null!;

        /// <summary>
        /// The name of the database user.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public GetDatabasesUserInvokeArgs()
        {
        }
        public static new GetDatabasesUserInvokeArgs Empty => new GetDatabasesUserInvokeArgs();
    }
}
