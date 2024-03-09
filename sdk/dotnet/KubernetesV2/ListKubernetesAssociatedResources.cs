// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.KubernetesV2
{
    public static class ListKubernetesAssociatedResources
    {
        public static Task<ListKubernetesAssociatedResourcesResult> InvokeAsync(ListKubernetesAssociatedResourcesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListKubernetesAssociatedResourcesResult>("digitalocean-native:kubernetes/v2:listKubernetesAssociatedResources", args ?? new ListKubernetesAssociatedResourcesArgs(), options.WithDefaults());

        public static Output<ListKubernetesAssociatedResourcesResult> Invoke(ListKubernetesAssociatedResourcesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListKubernetesAssociatedResourcesResult>("digitalocean-native:kubernetes/v2:listKubernetesAssociatedResources", args ?? new ListKubernetesAssociatedResourcesInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListKubernetesAssociatedResourcesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique ID that can be used to reference a Kubernetes cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        public ListKubernetesAssociatedResourcesArgs()
        {
        }
        public static new ListKubernetesAssociatedResourcesArgs Empty => new ListKubernetesAssociatedResourcesArgs();
    }

    public sealed class ListKubernetesAssociatedResourcesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique ID that can be used to reference a Kubernetes cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        public ListKubernetesAssociatedResourcesInvokeArgs()
        {
        }
        public static new ListKubernetesAssociatedResourcesInvokeArgs Empty => new ListKubernetesAssociatedResourcesInvokeArgs();
    }


    [OutputType]
    public sealed class ListKubernetesAssociatedResourcesResult
    {
        public readonly Outputs.AssociatedKubernetesResources Items;

        [OutputConstructor]
        private ListKubernetesAssociatedResourcesResult(Outputs.AssociatedKubernetesResources items)
        {
            Items = items;
        }
    }
}
