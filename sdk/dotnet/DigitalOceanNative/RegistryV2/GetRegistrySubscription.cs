// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.RegistryV2
{
    public static class GetRegistrySubscription
    {
        public static Task<Outputs.GetRegistrySubscriptionProperties> InvokeAsync(GetRegistrySubscriptionArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetRegistrySubscriptionProperties>("digitalocean-native:registry/v2:getRegistrySubscription", args ?? new GetRegistrySubscriptionArgs(), options.WithDefaults());

        public static Output<Outputs.GetRegistrySubscriptionProperties> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetRegistrySubscriptionProperties>("digitalocean-native:registry/v2:getRegistrySubscription", InvokeArgs.Empty, options.WithDefaults());

        public static Output<Outputs.GetRegistrySubscriptionProperties> Invoke(InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetRegistrySubscriptionProperties>("digitalocean-native:registry/v2:getRegistrySubscription", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class GetRegistrySubscriptionArgs : global::Pulumi.InvokeArgs
    {
        public GetRegistrySubscriptionArgs()
        {
        }
        public static new GetRegistrySubscriptionArgs Empty => new GetRegistrySubscriptionArgs();
    }
}
