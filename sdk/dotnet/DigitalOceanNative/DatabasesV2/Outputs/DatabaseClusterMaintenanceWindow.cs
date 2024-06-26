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
    public sealed class DatabaseClusterMaintenanceWindow
    {
        /// <summary>
        /// The day of the week on which to apply maintenance updates.
        /// </summary>
        public readonly string Day;
        /// <summary>
        /// A list of strings, each containing information about a pending maintenance update.
        /// </summary>
        public readonly ImmutableArray<string> Description;
        /// <summary>
        /// The hour in UTC at which maintenance updates will be applied in 24 hour format.
        /// </summary>
        public readonly string Hour;
        /// <summary>
        /// A boolean value indicating whether any maintenance is scheduled to be performed in the next window.
        /// </summary>
        public readonly bool? Pending;

        [OutputConstructor]
        private DatabaseClusterMaintenanceWindow(
            string day,

            ImmutableArray<string> description,

            string hour,

            bool? pending)
        {
            Day = day;
            Description = description;
            Hour = hour;
            Pending = pending;
        }
    }
}
