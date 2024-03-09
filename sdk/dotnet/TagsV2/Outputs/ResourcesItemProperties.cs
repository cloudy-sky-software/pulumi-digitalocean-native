// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.TagsV2.Outputs
{

    [OutputType]
    public sealed class ResourcesItemProperties
    {
        /// <summary>
        /// The identifier of a resource.
        /// </summary>
        public readonly string? ResourceId;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly Pulumi.DigitalOceanNative.TagsV2.ResourcesItemPropertiesResourceType? ResourceType;

        [OutputConstructor]
        private ResourcesItemProperties(
            string? resourceId,

            Pulumi.DigitalOceanNative.TagsV2.ResourcesItemPropertiesResourceType? resourceType)
        {
            ResourceId = resourceId;
            ResourceType = resourceType;
        }
    }
}
