// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.KubernetesV2
{
    public static class GetKubernetesClusterLintResults
    {
        public static Task<GetKubernetesClusterLintResultsResult> InvokeAsync(GetKubernetesClusterLintResultsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetKubernetesClusterLintResultsResult>("digitalocean-native:kubernetes/v2:getKubernetesClusterLintResults", args ?? new GetKubernetesClusterLintResultsArgs(), options.WithDefaults());

        public static Output<GetKubernetesClusterLintResultsResult> Invoke(GetKubernetesClusterLintResultsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetKubernetesClusterLintResultsResult>("digitalocean-native:kubernetes/v2:getKubernetesClusterLintResults", args ?? new GetKubernetesClusterLintResultsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetKubernetesClusterLintResultsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique ID that can be used to reference a Kubernetes cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        public GetKubernetesClusterLintResultsArgs()
        {
        }
        public static new GetKubernetesClusterLintResultsArgs Empty => new GetKubernetesClusterLintResultsArgs();
    }

    public sealed class GetKubernetesClusterLintResultsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique ID that can be used to reference a Kubernetes cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        public GetKubernetesClusterLintResultsInvokeArgs()
        {
        }
        public static new GetKubernetesClusterLintResultsInvokeArgs Empty => new GetKubernetesClusterLintResultsInvokeArgs();
    }


    [OutputType]
    public sealed class GetKubernetesClusterLintResultsResult
    {
        public readonly Outputs.ClusterlintResults Items;

        [OutputConstructor]
        private GetKubernetesClusterLintResultsResult(Outputs.ClusterlintResults items)
        {
            Items = items;
        }
    }
}
