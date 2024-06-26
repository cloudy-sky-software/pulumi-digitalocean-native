// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.DatabasesV2.Outputs
{

    [OutputType]
    public sealed class DatabaseVersionAvailability
    {
        /// <summary>
        /// A timestamp referring to the date when the particular version will no longer be available for creating new clusters. If null, the version does not have an end of availability timeline.
        /// </summary>
        public readonly string? EndOfAvailability;
        /// <summary>
        /// A timestamp referring to the date when the particular version will no longer be supported. If null, the version does not have an end of life timeline.
        /// </summary>
        public readonly string? EndOfLife;
        /// <summary>
        /// The engine version.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private DatabaseVersionAvailability(
            string? endOfAvailability,

            string? endOfLife,

            string? version)
        {
            EndOfAvailability = endOfAvailability;
            EndOfLife = endOfLife;
            Version = version;
        }
    }
}
