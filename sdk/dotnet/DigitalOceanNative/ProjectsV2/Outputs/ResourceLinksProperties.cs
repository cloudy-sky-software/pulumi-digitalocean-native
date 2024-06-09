// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.ProjectsV2.Outputs
{

    /// <summary>
    /// The links object contains the `self` object, which contains the resource relationship.
    /// </summary>
    [OutputType]
    public sealed class ResourceLinksProperties
    {
        /// <summary>
        /// A URI that can be used to retrieve the resource.
        /// </summary>
        public readonly string? Self;

        [OutputConstructor]
        private ResourceLinksProperties(string? self)
        {
            Self = self;
        }
    }
}