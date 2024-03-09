// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.DatabasesV2
{
    public static class GetDatabasesUser
    {
        public static Task<GetDatabasesUserResult> InvokeAsync(GetDatabasesUserArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatabasesUserResult>("digitalocean-native:databases/v2:getDatabasesUser", args ?? new GetDatabasesUserArgs(), options.WithDefaults());

        public static Output<GetDatabasesUserResult> Invoke(GetDatabasesUserInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabasesUserResult>("digitalocean-native:databases/v2:getDatabasesUser", args ?? new GetDatabasesUserInvokeArgs(), options.WithDefaults());
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


    [OutputType]
    public sealed class GetDatabasesUserResult
    {
        public readonly Outputs.GetDatabasesUserProperties Items;

        [OutputConstructor]
        private GetDatabasesUserResult(Outputs.GetDatabasesUserProperties items)
        {
            Items = items;
        }
    }
}
