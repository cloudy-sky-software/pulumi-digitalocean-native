// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.OneClicksV2
{
    public static class ListOneClicks
    {
        public static Task<Outputs.ListOneClicksProperties> InvokeAsync(ListOneClicksArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.ListOneClicksProperties>("digitalocean-native:1-clicks/v2:listOneClicks", args ?? new ListOneClicksArgs(), options.WithDefaults());

        public static Output<Outputs.ListOneClicksProperties> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ListOneClicksProperties>("digitalocean-native:1-clicks/v2:listOneClicks", InvokeArgs.Empty, options.WithDefaults());

        public static Output<Outputs.ListOneClicksProperties> Invoke(InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ListOneClicksProperties>("digitalocean-native:1-clicks/v2:listOneClicks", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListOneClicksArgs : global::Pulumi.InvokeArgs
    {
        public ListOneClicksArgs()
        {
        }
        public static new ListOneClicksArgs Empty => new ListOneClicksArgs();
    }
}
