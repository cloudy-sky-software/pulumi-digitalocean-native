// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.KubernetesV2
{
    public static class GetKubernetesNodePool
    {
        public static Task<GetKubernetesNodePoolResult> InvokeAsync(GetKubernetesNodePoolArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetKubernetesNodePoolResult>("digitalocean-native:kubernetes/v2:getKubernetesNodePool", args ?? new GetKubernetesNodePoolArgs(), options.WithDefaults());

        public static Output<GetKubernetesNodePoolResult> Invoke(GetKubernetesNodePoolInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetKubernetesNodePoolResult>("digitalocean-native:kubernetes/v2:getKubernetesNodePool", args ?? new GetKubernetesNodePoolInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetKubernetesNodePoolArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique ID that can be used to reference a Kubernetes cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// A unique ID that can be used to reference a Kubernetes node pool.
        /// </summary>
        [Input("nodePoolId", required: true)]
        public string NodePoolId { get; set; } = null!;

        public GetKubernetesNodePoolArgs()
        {
        }
        public static new GetKubernetesNodePoolArgs Empty => new GetKubernetesNodePoolArgs();
    }

    public sealed class GetKubernetesNodePoolInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique ID that can be used to reference a Kubernetes cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// A unique ID that can be used to reference a Kubernetes node pool.
        /// </summary>
        [Input("nodePoolId", required: true)]
        public Input<string> NodePoolId { get; set; } = null!;

        public GetKubernetesNodePoolInvokeArgs()
        {
        }
        public static new GetKubernetesNodePoolInvokeArgs Empty => new GetKubernetesNodePoolInvokeArgs();
    }


    [OutputType]
    public sealed class GetKubernetesNodePoolResult
    {
        public readonly Outputs.GetKubernetesNodePoolProperties Items;

        [OutputConstructor]
        private GetKubernetesNodePoolResult(Outputs.GetKubernetesNodePoolProperties items)
        {
            Items = items;
        }
    }
}
