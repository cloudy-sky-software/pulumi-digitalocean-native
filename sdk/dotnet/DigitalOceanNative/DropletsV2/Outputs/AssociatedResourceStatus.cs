// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.DropletsV2.Outputs
{

    /// <summary>
    /// An objects containing information about a resources scheduled for deletion.
    /// </summary>
    [OutputType]
    public sealed class AssociatedResourceStatus
    {
        /// <summary>
        /// A time value given in ISO8601 combined date and time format indicating when the requested action was completed.
        /// </summary>
        public readonly string? CompletedAt;
        /// <summary>
        /// An object containing information about a resource scheduled for deletion.
        /// </summary>
        public readonly Outputs.DestroyedAssociatedResource? Droplet;
        /// <summary>
        /// A count of the associated resources that failed to be destroyed, if any.
        /// </summary>
        public readonly int? Failures;
        /// <summary>
        /// An object containing additional information about resource related to a Droplet requested to be destroyed.
        /// </summary>
        public readonly Outputs.AssociatedResourceStatusResourcesProperties? Resources;

        [OutputConstructor]
        private AssociatedResourceStatus(
            string? completedAt,

            Outputs.DestroyedAssociatedResource? droplet,

            int? failures,

            Outputs.AssociatedResourceStatusResourcesProperties? resources)
        {
            CompletedAt = completedAt;
            Droplet = droplet;
            Failures = failures;
            Resources = resources;
        }
    }
}
