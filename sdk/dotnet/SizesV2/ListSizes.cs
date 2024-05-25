// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.SizesV2
{
    public static class ListSizes
    {
        public static Task<ListSizesResult> InvokeAsync(ListSizesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListSizesResult>("digitalocean-native:sizes/v2:listSizes", args ?? new ListSizesArgs(), options.WithDefaults());

        public static Output<ListSizesResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListSizesResult>("digitalocean-native:sizes/v2:listSizes", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListSizesArgs : global::Pulumi.InvokeArgs
    {
        public ListSizesArgs()
        {
        }
        public static new ListSizesArgs Empty => new ListSizesArgs();
    }


    [OutputType]
    public sealed class ListSizesResult
    {
        public readonly Outputs.ListSizesItems Items;

        [OutputConstructor]
        private ListSizesResult(Outputs.ListSizesItems items)
        {
            Items = items;
        }
    }
}
