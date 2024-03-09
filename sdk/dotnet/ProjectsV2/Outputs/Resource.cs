// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.ProjectsV2.Outputs
{

    [OutputType]
    public sealed class Resource
    {
        /// <summary>
        /// A time value given in ISO8601 combined date and time format that represents when the project was created.
        /// </summary>
        public readonly string? AssignedAt;
        /// <summary>
        /// The links object contains the `self` object, which contains the resource relationship.
        /// </summary>
        public readonly Outputs.ResourceLinksProperties? Links;
        /// <summary>
        /// The status of assigning and fetching the resources.
        /// </summary>
        public readonly Pulumi.DigitalOceanNative.ProjectsV2.ResourceStatus? Status;
        /// <summary>
        /// The uniform resource name (URN) for the resource in the format do:resource_type:resource_id.
        /// </summary>
        public readonly string? Urn;

        [OutputConstructor]
        private Resource(
            string? assignedAt,

            Outputs.ResourceLinksProperties? links,

            Pulumi.DigitalOceanNative.ProjectsV2.ResourceStatus? status,

            string? urn)
        {
            AssignedAt = assignedAt;
            Links = links;
            Status = status;
            Urn = urn;
        }
    }
}