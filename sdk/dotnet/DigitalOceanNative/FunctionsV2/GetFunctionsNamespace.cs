// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.FunctionsV2
{
    public static class GetFunctionsNamespace
    {
        public static Task<Outputs.GetFunctionsNamespaceProperties> InvokeAsync(GetFunctionsNamespaceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetFunctionsNamespaceProperties>("digitalocean-native:functions/v2:getFunctionsNamespace", args ?? new GetFunctionsNamespaceArgs(), options.WithDefaults());

        public static Output<Outputs.GetFunctionsNamespaceProperties> Invoke(GetFunctionsNamespaceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetFunctionsNamespaceProperties>("digitalocean-native:functions/v2:getFunctionsNamespace", args ?? new GetFunctionsNamespaceInvokeArgs(), options.WithDefaults());

        public static Output<Outputs.GetFunctionsNamespaceProperties> Invoke(GetFunctionsNamespaceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetFunctionsNamespaceProperties>("digitalocean-native:functions/v2:getFunctionsNamespace", args ?? new GetFunctionsNamespaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFunctionsNamespaceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the namespace to be managed.
        /// </summary>
        [Input("namespaceId", required: true)]
        public string NamespaceId { get; set; } = null!;

        public GetFunctionsNamespaceArgs()
        {
        }
        public static new GetFunctionsNamespaceArgs Empty => new GetFunctionsNamespaceArgs();
    }

    public sealed class GetFunctionsNamespaceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the namespace to be managed.
        /// </summary>
        [Input("namespaceId", required: true)]
        public Input<string> NamespaceId { get; set; } = null!;

        public GetFunctionsNamespaceInvokeArgs()
        {
        }
        public static new GetFunctionsNamespaceInvokeArgs Empty => new GetFunctionsNamespaceInvokeArgs();
    }
}
