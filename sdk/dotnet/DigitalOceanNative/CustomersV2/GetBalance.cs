// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.CustomersV2
{
    public static class GetBalance
    {
        public static Task<Outputs.Balance> InvokeAsync(GetBalanceArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.Balance>("digitalocean-native:customers/v2:getBalance", args ?? new GetBalanceArgs(), options.WithDefaults());

        public static Output<Outputs.Balance> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.Balance>("digitalocean-native:customers/v2:getBalance", InvokeArgs.Empty, options.WithDefaults());

        public static Output<Outputs.Balance> Invoke(InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.Balance>("digitalocean-native:customers/v2:getBalance", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class GetBalanceArgs : global::Pulumi.InvokeArgs
    {
        public GetBalanceArgs()
        {
        }
        public static new GetBalanceArgs Empty => new GetBalanceArgs();
    }
}
