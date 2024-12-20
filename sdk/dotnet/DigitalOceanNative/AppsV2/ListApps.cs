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
    public static class ListApps
    {
        public static Task<Outputs.AppsResponse> InvokeAsync(ListAppsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.AppsResponse>("digitalocean-native:apps/v2:listApps", args ?? new ListAppsArgs(), options.WithDefaults());

        public static Output<Outputs.AppsResponse> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.AppsResponse>("digitalocean-native:apps/v2:listApps", InvokeArgs.Empty, options.WithDefaults());

        public static Output<Outputs.AppsResponse> Invoke(InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.AppsResponse>("digitalocean-native:apps/v2:listApps", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListAppsArgs : global::Pulumi.InvokeArgs
    {
        public ListAppsArgs()
        {
        }
        public static new ListAppsArgs Empty => new ListAppsArgs();
    }
}
