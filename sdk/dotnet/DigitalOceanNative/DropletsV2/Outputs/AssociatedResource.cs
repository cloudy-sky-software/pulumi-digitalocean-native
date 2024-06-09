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
    /// An objects containing information about a resource associated with a Droplet.
    /// </summary>
    [OutputType]
    public sealed class AssociatedResource
    {
        /// <summary>
        /// The cost of the resource in USD per month if the resource is retained after the Droplet is destroyed.
        /// </summary>
        public readonly string? Cost;
        /// <summary>
        /// The unique identifier for the resource associated with the Droplet.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The name of the resource associated with the Droplet.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private AssociatedResource(
            string? cost,

            string? id,

            string? name)
        {
            Cost = cost;
            Id = id;
            Name = name;
        }
    }
}