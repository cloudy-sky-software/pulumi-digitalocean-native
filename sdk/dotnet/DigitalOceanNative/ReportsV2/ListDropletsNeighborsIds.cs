// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.ReportsV2
{
    public static class ListDropletsNeighborsIds
    {
        public static Task<Outputs.NeighborIds> InvokeAsync(ListDropletsNeighborsIdsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.NeighborIds>("digitalocean-native:reports/v2:listDropletsNeighborsIds", args ?? new ListDropletsNeighborsIdsArgs(), options.WithDefaults());

        public static Output<Outputs.NeighborIds> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.NeighborIds>("digitalocean-native:reports/v2:listDropletsNeighborsIds", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListDropletsNeighborsIdsArgs : global::Pulumi.InvokeArgs
    {
        public ListDropletsNeighborsIdsArgs()
        {
        }
        public static new ListDropletsNeighborsIdsArgs Empty => new ListDropletsNeighborsIdsArgs();
    }
}
