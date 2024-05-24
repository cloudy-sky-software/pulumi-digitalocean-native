// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.AppsV2.Outputs
{

    [OutputType]
    public sealed class AppsImageSourceSpec
    {
        /// <summary>
        /// The registry name. Must be left empty for the `DOCR` registry type.
        /// </summary>
        public readonly string? Registry;
        /// <summary>
        /// - DOCKER_HUB: The DockerHub container registry type.
        /// - DOCR: The DigitalOcean container registry type.
        /// </summary>
        public readonly Pulumi.DigitalOceanNative.AppsV2.AppsValidateAppSpecAppsImageSourceSpecRegistryType? RegistryType;
        /// <summary>
        /// The repository name.
        /// </summary>
        public readonly string? Repository;
        /// <summary>
        /// The repository tag. Defaults to `latest` if not provided.
        /// </summary>
        public readonly string? Tag;

        [OutputConstructor]
        private AppsImageSourceSpec(
            string? registry,

            Pulumi.DigitalOceanNative.AppsV2.AppsValidateAppSpecAppsImageSourceSpecRegistryType? registryType,

            string? repository,

            string? tag)
        {
            Registry = registry;
            RegistryType = registryType;
            Repository = repository;
            Tag = tag;
        }
    }
}
