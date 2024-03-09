// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.ReservedIpsV2.Outputs
{

    /// <summary>
    /// The linked actions can be used to check the status of a Droplet's create event.
    /// </summary>
    [OutputType]
    public sealed class ActionLink
    {
        /// <summary>
        /// A URL that can be used to access the action.
        /// </summary>
        public readonly string? Href;
        /// <summary>
        /// A unique numeric ID that can be used to identify and reference an action.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// A string specifying the type of the related action.
        /// </summary>
        public readonly string? Rel;

        [OutputConstructor]
        private ActionLink(
            string? href,

            int? id,

            string? rel)
        {
            Href = href;
            Id = id;
            Rel = rel;
        }
    }
}
