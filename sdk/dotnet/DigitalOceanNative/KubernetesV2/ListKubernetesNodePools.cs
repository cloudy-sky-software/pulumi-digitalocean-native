// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.KubernetesV2
{
    public static class ListKubernetesNodePools
    {
        public static Task<ListKubernetesNodePoolsResult> InvokeAsync(ListKubernetesNodePoolsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListKubernetesNodePoolsResult>("digitalocean-native:kubernetes/v2:listKubernetesNodePools", args ?? new ListKubernetesNodePoolsArgs(), options.WithDefaults());

        public static Output<ListKubernetesNodePoolsResult> Invoke(ListKubernetesNodePoolsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListKubernetesNodePoolsResult>("digitalocean-native:kubernetes/v2:listKubernetesNodePools", args ?? new ListKubernetesNodePoolsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListKubernetesNodePoolsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique ID that can be used to reference a Kubernetes cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        public ListKubernetesNodePoolsArgs()
        {
        }
        public static new ListKubernetesNodePoolsArgs Empty => new ListKubernetesNodePoolsArgs();
    }

    public sealed class ListKubernetesNodePoolsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique ID that can be used to reference a Kubernetes cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        public ListKubernetesNodePoolsInvokeArgs()
        {
        }
        public static new ListKubernetesNodePoolsInvokeArgs Empty => new ListKubernetesNodePoolsInvokeArgs();
    }


    [OutputType]
    public sealed class ListKubernetesNodePoolsResult
    {
        public readonly Outputs.ListKubernetesNodePoolsProperties Items;

        [OutputConstructor]
        private ListKubernetesNodePoolsResult(Outputs.ListKubernetesNodePoolsProperties items)
        {
            Items = items;
        }
    }
}