// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.TagsV2.Inputs
{

    public sealed class ResourcesItemPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identifier of a resource.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Input("resourceType")]
        public Input<Pulumi.DigitalOceanNative.TagsV2.ResourcesItemPropertiesResourceType>? ResourceType { get; set; }

        public ResourcesItemPropertiesArgs()
        {
        }
        public static new ResourcesItemPropertiesArgs Empty => new ResourcesItemPropertiesArgs();
    }
}
