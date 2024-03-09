// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.ReservedIpsV2.Outputs
{

    [OutputType]
    public sealed class GetReservedIPsActionsPropertiesAction
    {
        /// <summary>
        /// A time value given in ISO8601 combined date and time format that represents when the action was completed.
        /// </summary>
        public readonly string? CompletedAt;
        /// <summary>
        /// A unique numeric ID that can be used to identify and reference an action.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// The UUID of the project to which the reserved IP currently belongs.
        /// </summary>
        public readonly string? ProjectId;
        public readonly Outputs.Region? Region;
        public readonly Outputs.ActionRegionSlug? RegionSlug;
        /// <summary>
        /// A unique identifier for the resource that the action is associated with.
        /// </summary>
        public readonly int? ResourceId;
        /// <summary>
        /// The type of resource that the action is associated with.
        /// </summary>
        public readonly string? ResourceType;
        /// <summary>
        /// A time value given in ISO8601 combined date and time format that represents when the action was initiated.
        /// </summary>
        public readonly string? StartedAt;
        /// <summary>
        /// The current status of the action. This can be "in-progress", "completed", or "errored".
        /// </summary>
        public readonly Pulumi.DigitalOceanNative.ReservedIpsV2.ActionStatus? Status;
        /// <summary>
        /// This is the type of action that the object represents. For example, this could be "transfer" to represent the state of an image transfer action.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private GetReservedIPsActionsPropertiesAction(
            string? completedAt,

            int? id,

            string? projectId,

            Outputs.Region? region,

            Outputs.ActionRegionSlug? regionSlug,

            int? resourceId,

            string? resourceType,

            string? startedAt,

            Pulumi.DigitalOceanNative.ReservedIpsV2.ActionStatus? status,

            string? type)
        {
            CompletedAt = completedAt;
            Id = id;
            ProjectId = projectId;
            Region = region;
            RegionSlug = regionSlug;
            ResourceId = resourceId;
            ResourceType = resourceType;
            StartedAt = startedAt;
            Status = status;
            Type = type;
        }
    }
}
