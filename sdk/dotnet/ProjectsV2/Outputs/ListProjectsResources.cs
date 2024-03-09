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
    public sealed class ListProjectsResources
    {
        public readonly Outputs.PageLinks? Links;
        public readonly Outputs.MetaMeta MetaValue;
        public readonly ImmutableArray<Outputs.Resource> Resources;

        [OutputConstructor]
        private ListProjectsResources(
            Outputs.PageLinks? links,

            Outputs.MetaMeta meta,

            ImmutableArray<Outputs.Resource> resources)
        {
            Links = links;
            MetaValue = meta;
            Resources = resources;
        }
    }
}
