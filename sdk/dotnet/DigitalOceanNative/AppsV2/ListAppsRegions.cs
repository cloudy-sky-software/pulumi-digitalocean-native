// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.AppsV2
{
    public static class ListAppsRegions
    {
        public static Task<Outputs.AppsListRegionsResponse> InvokeAsync(ListAppsRegionsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.AppsListRegionsResponse>("digitalocean-native:apps/v2:listAppsRegions", args ?? new ListAppsRegionsArgs(), options.WithDefaults());

        public static Output<Outputs.AppsListRegionsResponse> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.AppsListRegionsResponse>("digitalocean-native:apps/v2:listAppsRegions", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListAppsRegionsArgs : global::Pulumi.InvokeArgs
    {
        public ListAppsRegionsArgs()
        {
        }
        public static new ListAppsRegionsArgs Empty => new ListAppsRegionsArgs();
    }
}
