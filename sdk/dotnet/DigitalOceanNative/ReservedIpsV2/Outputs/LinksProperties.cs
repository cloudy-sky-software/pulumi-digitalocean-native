// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.ReservedIpsV2.Outputs
{

    [OutputType]
    public sealed class LinksProperties
    {
        public readonly ImmutableArray<Outputs.ActionLink> Actions;
        public readonly ImmutableArray<Outputs.ActionLink> Droplets;

        [OutputConstructor]
        private LinksProperties(
            ImmutableArray<Outputs.ActionLink> actions,

            ImmutableArray<Outputs.ActionLink> droplets)
        {
            Actions = actions;
            Droplets = droplets;
        }
    }
}