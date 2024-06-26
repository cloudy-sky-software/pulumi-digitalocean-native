// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.FunctionsV2.Outputs
{

    [OutputType]
    public sealed class TriggerInfo
    {
        /// <summary>
        /// UTC time string.
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// Name of function(action) that exists in the given namespace.
        /// </summary>
        public readonly string? Function;
        /// <summary>
        /// Indicates weather the trigger is paused or unpaused.
        /// </summary>
        public readonly bool? IsEnabled;
        /// <summary>
        /// The trigger's unique name within the namespace.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// A unique string format of UUID with a prefix fn-.
        /// </summary>
        public readonly string? Namespace;
        /// <summary>
        /// Trigger details for SCHEDULED type, where body is optional.
        /// </summary>
        public readonly Outputs.ScheduledDetails? ScheduledDetails;
        public readonly Outputs.TriggerInfoScheduledRunsProperties? ScheduledRuns;
        /// <summary>
        /// String which indicates the type of trigger source like SCHEDULED.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// UTC time string.
        /// </summary>
        public readonly string? UpdatedAt;

        [OutputConstructor]
        private TriggerInfo(
            string? createdAt,

            string? function,

            bool? isEnabled,

            string? name,

            string? @namespace,

            Outputs.ScheduledDetails? scheduledDetails,

            Outputs.TriggerInfoScheduledRunsProperties? scheduledRuns,

            string? type,

            string? updatedAt)
        {
            CreatedAt = createdAt;
            Function = function;
            IsEnabled = isEnabled;
            Name = name;
            Namespace = @namespace;
            ScheduledDetails = scheduledDetails;
            ScheduledRuns = scheduledRuns;
            Type = type;
            UpdatedAt = updatedAt;
        }
    }
}
