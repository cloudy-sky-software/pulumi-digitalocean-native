// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.RegistryV2.Outputs
{

    [OutputType]
    public sealed class DockerCredentials
    {
        public readonly Outputs.DockerCredentialsAuthsProperties? Auths;

        [OutputConstructor]
        private DockerCredentials(Outputs.DockerCredentialsAuthsProperties? auths)
        {
            Auths = auths;
        }
    }
}
