// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.ReservedIpsV2.Outputs
{

    /// <summary>
    /// The details of the Droplet's backups feature, if backups are configured for the Droplet. This object contains keys for the start and end times of the window during which the backup will start.
    /// </summary>
    [OutputType]
    public sealed class DropletNextBackupWindowProperties
    {
        /// <summary>
        /// A time value given in ISO8601 combined date and time format specifying the end of the Droplet's backup window.
        /// </summary>
        public readonly string? End;
        /// <summary>
        /// A time value given in ISO8601 combined date and time format specifying the start of the Droplet's backup window.
        /// </summary>
        public readonly string? Start;

        [OutputConstructor]
        private DropletNextBackupWindowProperties(
            string? end,

            string? start)
        {
            End = end;
            Start = start;
        }
    }
}
