// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.RegionsV2
{
    public static class ListRegions
    {
        public static Task<ListRegionsResult> InvokeAsync(ListRegionsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListRegionsResult>("digitalocean-native:regions/v2:listRegions", args ?? new ListRegionsArgs(), options.WithDefaults());

        public static Output<ListRegionsResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListRegionsResult>("digitalocean-native:regions/v2:listRegions", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListRegionsArgs : global::Pulumi.InvokeArgs
    {
        public ListRegionsArgs()
        {
        }
        public static new ListRegionsArgs Empty => new ListRegionsArgs();
    }


    [OutputType]
    public sealed class ListRegionsResult
    {
        public readonly Outputs.ListRegions Items;

        [OutputConstructor]
        private ListRegionsResult(Outputs.ListRegions items)
        {
            Items = items;
        }
    }
}
