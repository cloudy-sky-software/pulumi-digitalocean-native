// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.AppsV2
{
    public static class ListApps
    {
        public static Task<ListAppsResult> InvokeAsync(ListAppsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListAppsResult>("digitalocean-native:apps/v2:listApps", args ?? new ListAppsArgs(), options.WithDefaults());

        public static Output<ListAppsResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListAppsResult>("digitalocean-native:apps/v2:listApps", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListAppsArgs : global::Pulumi.InvokeArgs
    {
        public ListAppsArgs()
        {
        }
        public static new ListAppsArgs Empty => new ListAppsArgs();
    }


    [OutputType]
    public sealed class ListAppsResult
    {
        public readonly Outputs.AppsResponse Items;

        [OutputConstructor]
        private ListAppsResult(Outputs.AppsResponse items)
        {
            Items = items;
        }
    }
}
