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
    public static class GetKubernetesCredential
    {
        public static Task<GetKubernetesCredentialResult> InvokeAsync(GetKubernetesCredentialArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetKubernetesCredentialResult>("digitalocean-native:kubernetes/v2:getKubernetesCredential", args ?? new GetKubernetesCredentialArgs(), options.WithDefaults());

        public static Output<GetKubernetesCredentialResult> Invoke(GetKubernetesCredentialInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetKubernetesCredentialResult>("digitalocean-native:kubernetes/v2:getKubernetesCredential", args ?? new GetKubernetesCredentialInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetKubernetesCredentialArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique ID that can be used to reference a Kubernetes cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        public GetKubernetesCredentialArgs()
        {
        }
        public static new GetKubernetesCredentialArgs Empty => new GetKubernetesCredentialArgs();
    }

    public sealed class GetKubernetesCredentialInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique ID that can be used to reference a Kubernetes cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        public GetKubernetesCredentialInvokeArgs()
        {
        }
        public static new GetKubernetesCredentialInvokeArgs Empty => new GetKubernetesCredentialInvokeArgs();
    }


    [OutputType]
    public sealed class GetKubernetesCredentialResult
    {
        public readonly Outputs.Credentials Items;

        [OutputConstructor]
        private GetKubernetesCredentialResult(Outputs.Credentials items)
        {
            Items = items;
        }
    }
}