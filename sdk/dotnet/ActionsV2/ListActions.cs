// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.ActionsV2
{
    public static class ListActions
    {
        public static Task<ListActionsResult> InvokeAsync(ListActionsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListActionsResult>("digitalocean-native:actions/v2:listActions", args ?? new ListActionsArgs(), options.WithDefaults());

        public static Output<ListActionsResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListActionsResult>("digitalocean-native:actions/v2:listActions", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListActionsArgs : global::Pulumi.InvokeArgs
    {
        public ListActionsArgs()
        {
        }
        public static new ListActionsArgs Empty => new ListActionsArgs();
    }


    [OutputType]
    public sealed class ListActionsResult
    {
        public readonly Outputs.ListActionsItems Items;

        [OutputConstructor]
        private ListActionsResult(Outputs.ListActionsItems items)
        {
            Items = items;
        }
    }
}
