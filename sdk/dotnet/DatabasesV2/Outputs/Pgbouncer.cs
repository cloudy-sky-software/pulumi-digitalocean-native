// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.DatabasesV2.Outputs
{

    /// <summary>
    /// PGBouncer connection pooling settings
    /// </summary>
    [OutputType]
    public sealed class Pgbouncer
    {
        /// <summary>
        /// If the automatically-created database pools have been unused this many seconds, they are freed. If 0, timeout is disabled.
        /// </summary>
        public readonly int? AutodbIdleTimeout;
        /// <summary>
        /// Only allows a maximum this many server connections per database (regardless of user). If 0, allows unlimited connections.
        /// </summary>
        public readonly int? AutodbMaxDbConnections;
        /// <summary>
        /// PGBouncer pool mode
        /// </summary>
        public readonly Pulumi.DigitalOceanNative.DatabasesV2.PgbouncerAutodbPoolMode? AutodbPoolMode;
        /// <summary>
        /// If non-zero, automatically creates a pool of that size per user when a pool doesn't exist.
        /// </summary>
        public readonly int? AutodbPoolSize;
        /// <summary>
        /// List of parameters to ignore when given in startup packet.
        /// </summary>
        public readonly ImmutableArray<Pulumi.DigitalOceanNative.DatabasesV2.PgbouncerIgnoreStartupParametersItem> IgnoreStartupParameters;
        /// <summary>
        /// If current server connections are below this number, adds more. Improves behavior when usual load comes suddenly back after period of total inactivity. The value is effectively capped at the pool size.
        /// </summary>
        public readonly int? MinPoolSize;
        /// <summary>
        /// Drops server connections if they have been idle more than this many seconds.  If 0, timeout is disabled. 
        /// </summary>
        public readonly int? ServerIdleTimeout;
        /// <summary>
        /// The pooler closes any unused server connection that has been connected longer than this amount of seconds.
        /// </summary>
        public readonly int? ServerLifetime;
        /// <summary>
        /// Run server_reset_query (DISCARD ALL) in all pooling modes.
        /// </summary>
        public readonly bool? ServerResetQueryAlways;

        [OutputConstructor]
        private Pgbouncer(
            int? autodbIdleTimeout,

            int? autodbMaxDbConnections,

            Pulumi.DigitalOceanNative.DatabasesV2.PgbouncerAutodbPoolMode? autodbPoolMode,

            int? autodbPoolSize,

            ImmutableArray<Pulumi.DigitalOceanNative.DatabasesV2.PgbouncerIgnoreStartupParametersItem> ignoreStartupParameters,

            int? minPoolSize,

            int? serverIdleTimeout,

            int? serverLifetime,

            bool? serverResetQueryAlways)
        {
            AutodbIdleTimeout = autodbIdleTimeout;
            AutodbMaxDbConnections = autodbMaxDbConnections;
            AutodbPoolMode = autodbPoolMode;
            AutodbPoolSize = autodbPoolSize;
            IgnoreStartupParameters = ignoreStartupParameters;
            MinPoolSize = minPoolSize;
            ServerIdleTimeout = serverIdleTimeout;
            ServerLifetime = serverLifetime;
            ServerResetQueryAlways = serverResetQueryAlways;
        }
    }
}