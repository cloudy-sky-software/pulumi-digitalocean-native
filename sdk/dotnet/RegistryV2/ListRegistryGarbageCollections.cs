// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.RegistryV2
{
    public static class ListRegistryGarbageCollections
    {
        public static Task<ListRegistryGarbageCollectionsResult> InvokeAsync(ListRegistryGarbageCollectionsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListRegistryGarbageCollectionsResult>("digitalocean-native:registry/v2:listRegistryGarbageCollections", args ?? new ListRegistryGarbageCollectionsArgs(), options.WithDefaults());

        public static Output<ListRegistryGarbageCollectionsResult> Invoke(ListRegistryGarbageCollectionsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListRegistryGarbageCollectionsResult>("digitalocean-native:registry/v2:listRegistryGarbageCollections", args ?? new ListRegistryGarbageCollectionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListRegistryGarbageCollectionsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of a container registry.
        /// </summary>
        [Input("registryName", required: true)]
        public string RegistryName { get; set; } = null!;

        public ListRegistryGarbageCollectionsArgs()
        {
        }
        public static new ListRegistryGarbageCollectionsArgs Empty => new ListRegistryGarbageCollectionsArgs();
    }

    public sealed class ListRegistryGarbageCollectionsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of a container registry.
        /// </summary>
        [Input("registryName", required: true)]
        public Input<string> RegistryName { get; set; } = null!;

        public ListRegistryGarbageCollectionsInvokeArgs()
        {
        }
        public static new ListRegistryGarbageCollectionsInvokeArgs Empty => new ListRegistryGarbageCollectionsInvokeArgs();
    }


    [OutputType]
    public sealed class ListRegistryGarbageCollectionsResult
    {
        public readonly Outputs.ListRegistryGarbageCollectionsProperties Items;

        [OutputConstructor]
        private ListRegistryGarbageCollectionsResult(Outputs.ListRegistryGarbageCollectionsProperties items)
        {
            Items = items;
        }
    }
}
