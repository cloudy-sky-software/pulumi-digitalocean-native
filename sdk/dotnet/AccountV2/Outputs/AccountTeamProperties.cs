// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.AccountV2.Outputs
{

    /// <summary>
    /// When authorized in a team context, includes information about the current team.
    /// </summary>
    [OutputType]
    public sealed class AccountTeamProperties
    {
        /// <summary>
        /// The name for the current team.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The unique universal identifier for the current team.
        /// </summary>
        public readonly string? Uuid;

        [OutputConstructor]
        private AccountTeamProperties(
            string? name,

            string? uuid)
        {
            Name = name;
            Uuid = uuid;
        }
    }
}
