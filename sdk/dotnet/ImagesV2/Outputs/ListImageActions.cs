// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.ImagesV2.Outputs
{

    [OutputType]
    public sealed class ListImageActions
    {
        public readonly ImmutableArray<Outputs.Action> Actions;
        public readonly Outputs.PageLinks? Links;
        public readonly Outputs.MetaMeta MetaValue;

        [OutputConstructor]
        private ListImageActions(
            ImmutableArray<Outputs.Action> actions,

            Outputs.PageLinks? links,

            Outputs.MetaMeta meta)
        {
            Actions = actions;
            Links = links;
            MetaValue = meta;
        }
    }
}
