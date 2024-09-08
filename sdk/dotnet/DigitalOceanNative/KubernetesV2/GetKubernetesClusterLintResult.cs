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
    public static class GetKubernetesClusterLintResult
    {
        public static Task<Outputs.ClusterlintResults> InvokeAsync(GetKubernetesClusterLintResultArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.ClusterlintResults>("digitalocean-native:kubernetes/v2:getKubernetesClusterLintResult", args ?? new GetKubernetesClusterLintResultArgs(), options.WithDefaults());

        public static Output<Outputs.ClusterlintResults> Invoke(GetKubernetesClusterLintResultInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ClusterlintResults>("digitalocean-native:kubernetes/v2:getKubernetesClusterLintResult", args ?? new GetKubernetesClusterLintResultInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetKubernetesClusterLintResultArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique ID that can be used to reference a Kubernetes cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        public GetKubernetesClusterLintResultArgs()
        {
        }
        public static new GetKubernetesClusterLintResultArgs Empty => new GetKubernetesClusterLintResultArgs();
    }

    public sealed class GetKubernetesClusterLintResultInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique ID that can be used to reference a Kubernetes cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        public GetKubernetesClusterLintResultInvokeArgs()
        {
        }
        public static new GetKubernetesClusterLintResultInvokeArgs Empty => new GetKubernetesClusterLintResultInvokeArgs();
    }
}
