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
    public static class GetRegistryDockerCredential
    {
        public static Task<GetRegistryDockerCredentialResult> InvokeAsync(GetRegistryDockerCredentialArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRegistryDockerCredentialResult>("digitalocean-native:registry/v2:getRegistryDockerCredential", args ?? new GetRegistryDockerCredentialArgs(), options.WithDefaults());

        public static Output<GetRegistryDockerCredentialResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRegistryDockerCredentialResult>("digitalocean-native:registry/v2:getRegistryDockerCredential", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class GetRegistryDockerCredentialArgs : global::Pulumi.InvokeArgs
    {
        public GetRegistryDockerCredentialArgs()
        {
        }
        public static new GetRegistryDockerCredentialArgs Empty => new GetRegistryDockerCredentialArgs();
    }


    [OutputType]
    public sealed class GetRegistryDockerCredentialResult
    {
        public readonly Outputs.DockerCredentials Items;

        [OutputConstructor]
        private GetRegistryDockerCredentialResult(Outputs.DockerCredentials items)
        {
            Items = items;
        }
    }
}
