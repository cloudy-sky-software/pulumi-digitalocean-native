// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.KubernetesV2
{
    public static class GetKubernetesCredentials
    {
        public static Task<GetKubernetesCredentialsResult> InvokeAsync(GetKubernetesCredentialsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetKubernetesCredentialsResult>("digitalocean-native:kubernetes/v2:getKubernetesCredentials", args ?? new GetKubernetesCredentialsArgs(), options.WithDefaults());

        public static Output<GetKubernetesCredentialsResult> Invoke(GetKubernetesCredentialsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetKubernetesCredentialsResult>("digitalocean-native:kubernetes/v2:getKubernetesCredentials", args ?? new GetKubernetesCredentialsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetKubernetesCredentialsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique ID that can be used to reference a Kubernetes cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        public GetKubernetesCredentialsArgs()
        {
        }
        public static new GetKubernetesCredentialsArgs Empty => new GetKubernetesCredentialsArgs();
    }

    public sealed class GetKubernetesCredentialsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique ID that can be used to reference a Kubernetes cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        public GetKubernetesCredentialsInvokeArgs()
        {
        }
        public static new GetKubernetesCredentialsInvokeArgs Empty => new GetKubernetesCredentialsInvokeArgs();
    }


    [OutputType]
    public sealed class GetKubernetesCredentialsResult
    {
        public readonly Outputs.Credentials Items;

        [OutputConstructor]
        private GetKubernetesCredentialsResult(Outputs.Credentials items)
        {
            Items = items;
        }
    }
}
