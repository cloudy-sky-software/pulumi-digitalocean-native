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
    public static class ListRegistryRepositoryManifests
    {
        public static Task<ListRegistryRepositoryManifestsResult> InvokeAsync(ListRegistryRepositoryManifestsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListRegistryRepositoryManifestsResult>("digitalocean-native:registry/v2:listRegistryRepositoryManifests", args ?? new ListRegistryRepositoryManifestsArgs(), options.WithDefaults());

        public static Output<ListRegistryRepositoryManifestsResult> Invoke(ListRegistryRepositoryManifestsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListRegistryRepositoryManifestsResult>("digitalocean-native:registry/v2:listRegistryRepositoryManifests", args ?? new ListRegistryRepositoryManifestsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListRegistryRepositoryManifestsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of a container registry.
        /// </summary>
        [Input("registryName", required: true)]
        public string RegistryName { get; set; } = null!;

        /// <summary>
        /// The name of a container registry repository. If the name contains `/` characters, they must be URL-encoded, e.g. `%2F`.
        /// </summary>
        [Input("repositoryName", required: true)]
        public string RepositoryName { get; set; } = null!;

        public ListRegistryRepositoryManifestsArgs()
        {
        }
        public static new ListRegistryRepositoryManifestsArgs Empty => new ListRegistryRepositoryManifestsArgs();
    }

    public sealed class ListRegistryRepositoryManifestsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of a container registry.
        /// </summary>
        [Input("registryName", required: true)]
        public Input<string> RegistryName { get; set; } = null!;

        /// <summary>
        /// The name of a container registry repository. If the name contains `/` characters, they must be URL-encoded, e.g. `%2F`.
        /// </summary>
        [Input("repositoryName", required: true)]
        public Input<string> RepositoryName { get; set; } = null!;

        public ListRegistryRepositoryManifestsInvokeArgs()
        {
        }
        public static new ListRegistryRepositoryManifestsInvokeArgs Empty => new ListRegistryRepositoryManifestsInvokeArgs();
    }


    [OutputType]
    public sealed class ListRegistryRepositoryManifestsResult
    {
        public readonly Outputs.ListRegistryRepositoryManifestsItems Items;

        [OutputConstructor]
        private ListRegistryRepositoryManifestsResult(Outputs.ListRegistryRepositoryManifestsItems items)
        {
            Items = items;
        }
    }
}
