// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.RegionsV2
{
    public static class ListRegions
    {
        public static Task<Outputs.ListRegionsItems> InvokeAsync(ListRegionsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.ListRegionsItems>("digitalocean-native:regions/v2:listRegions", args ?? new ListRegionsArgs(), options.WithDefaults());

        public static Output<Outputs.ListRegionsItems> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ListRegionsItems>("digitalocean-native:regions/v2:listRegions", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListRegionsArgs : global::Pulumi.InvokeArgs
    {
        public ListRegionsArgs()
        {
        }
        public static new ListRegionsArgs Empty => new ListRegionsArgs();
    }
}
