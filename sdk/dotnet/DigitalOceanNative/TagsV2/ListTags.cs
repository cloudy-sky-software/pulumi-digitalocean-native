// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.TagsV2
{
    public static class ListTags
    {
        public static Task<ListTagsResult> InvokeAsync(ListTagsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListTagsResult>("digitalocean-native:tags/v2:listTags", args ?? new ListTagsArgs(), options.WithDefaults());

        public static Output<ListTagsResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListTagsResult>("digitalocean-native:tags/v2:listTags", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListTagsArgs : global::Pulumi.InvokeArgs
    {
        public ListTagsArgs()
        {
        }
        public static new ListTagsArgs Empty => new ListTagsArgs();
    }


    [OutputType]
    public sealed class ListTagsResult
    {
        public readonly Outputs.ListTagsItems Items;

        [OutputConstructor]
        private ListTagsResult(Outputs.ListTagsItems items)
        {
            Items = items;
        }
    }
}
