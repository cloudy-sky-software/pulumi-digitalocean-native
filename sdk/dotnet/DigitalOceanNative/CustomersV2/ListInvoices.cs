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
    public static class ListInvoices
    {
        public static Task<Outputs.ListInvoicesItems> InvokeAsync(ListInvoicesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.ListInvoicesItems>("digitalocean-native:customers/v2:listInvoices", args ?? new ListInvoicesArgs(), options.WithDefaults());

        public static Output<Outputs.ListInvoicesItems> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ListInvoicesItems>("digitalocean-native:customers/v2:listInvoices", InvokeArgs.Empty, options.WithDefaults());

        public static Output<Outputs.ListInvoicesItems> Invoke(InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ListInvoicesItems>("digitalocean-native:customers/v2:listInvoices", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListInvoicesArgs : global::Pulumi.InvokeArgs
    {
        public ListInvoicesArgs()
        {
        }
        public static new ListInvoicesArgs Empty => new ListInvoicesArgs();
    }
}
