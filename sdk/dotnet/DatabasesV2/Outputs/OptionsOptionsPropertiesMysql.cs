// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.DatabasesV2.Outputs
{

    [OutputType]
    public sealed class OptionsOptionsPropertiesMysql
    {
        /// <summary>
        /// An array of objects, each indicating the node sizes (otherwise referred to as slugs) that are available with various numbers of nodes in the database cluster. Each slugs denotes the node's identifier, CPU, and RAM (in that order).
        /// </summary>
        public readonly ImmutableArray<Outputs.DatabaseLayoutOption> Layouts;
        /// <summary>
        /// An array of strings containing the names of available regions
        /// </summary>
        public readonly ImmutableArray<string> Regions;
        /// <summary>
        /// An array of strings containing the names of available regions
        /// </summary>
        public readonly ImmutableArray<string> Versions;

        [OutputConstructor]
        private OptionsOptionsPropertiesMysql(
            ImmutableArray<Outputs.DatabaseLayoutOption> layouts,

            ImmutableArray<string> regions,

            ImmutableArray<string> versions)
        {
            Layouts = layouts;
            Regions = regions;
            Versions = versions;
        }
    }
}
