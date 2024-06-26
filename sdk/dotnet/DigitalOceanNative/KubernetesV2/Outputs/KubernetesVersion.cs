// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.KubernetesV2.Outputs
{

    [OutputType]
    public sealed class KubernetesVersion
    {
        /// <summary>
        /// The upstream version string for the version of Kubernetes provided by a given slug.
        /// </summary>
        public readonly string? KubernetesVersionValue;
        /// <summary>
        /// The slug identifier for an available version of Kubernetes for use when creating or updating a cluster. The string contains both the upstream version of Kubernetes as well as the DigitalOcean revision.
        /// </summary>
        public readonly string? Slug;
        /// <summary>
        /// The features available with the version of Kubernetes provided by a given slug.
        /// </summary>
        public readonly ImmutableArray<string> SupportedFeatures;

        [OutputConstructor]
        private KubernetesVersion(
            string? kubernetesVersion,

            string? slug,

            ImmutableArray<string> supportedFeatures)
        {
            KubernetesVersionValue = kubernetesVersion;
            Slug = slug;
            SupportedFeatures = supportedFeatures;
        }
    }
}
