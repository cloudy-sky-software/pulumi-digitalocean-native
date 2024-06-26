// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.OneClicksV2.Outputs
{

    [OutputType]
    public sealed class OneClicks
    {
        /// <summary>
        /// The slug identifier for the 1-Click application.
        /// </summary>
        public readonly string Slug;
        /// <summary>
        /// The type of the 1-Click application.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private OneClicks(
            string slug,

            string type)
        {
            Slug = slug;
            Type = type;
        }
    }
}
