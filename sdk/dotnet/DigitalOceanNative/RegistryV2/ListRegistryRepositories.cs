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
    public static class ListRegistryRepositories
    {
        public static Task<ListRegistryRepositoriesResult> InvokeAsync(ListRegistryRepositoriesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListRegistryRepositoriesResult>("digitalocean-native:registry/v2:listRegistryRepositories", args ?? new ListRegistryRepositoriesArgs(), options.WithDefaults());

        public static Output<ListRegistryRepositoriesResult> Invoke(ListRegistryRepositoriesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListRegistryRepositoriesResult>("digitalocean-native:registry/v2:listRegistryRepositories", args ?? new ListRegistryRepositoriesInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListRegistryRepositoriesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of a container registry.
        /// </summary>
        [Input("registryName", required: true)]
        public string RegistryName { get; set; } = null!;

        public ListRegistryRepositoriesArgs()
        {
        }
        public static new ListRegistryRepositoriesArgs Empty => new ListRegistryRepositoriesArgs();
    }

    public sealed class ListRegistryRepositoriesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of a container registry.
        /// </summary>
        [Input("registryName", required: true)]
        public Input<string> RegistryName { get; set; } = null!;

        public ListRegistryRepositoriesInvokeArgs()
        {
        }
        public static new ListRegistryRepositoriesInvokeArgs Empty => new ListRegistryRepositoriesInvokeArgs();
    }


    [OutputType]
    public sealed class ListRegistryRepositoriesResult
    {
        public readonly Outputs.ListRegistryRepositoriesItems Items;

        [OutputConstructor]
        private ListRegistryRepositoriesResult(Outputs.ListRegistryRepositoriesItems items)
        {
            Items = items;
        }
    }
}
