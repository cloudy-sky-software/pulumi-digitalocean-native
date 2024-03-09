// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.AppsV2.Inputs
{

    public sealed class AppsImageSourceSpecArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The registry name. Must be left empty for the `DOCR` registry type.
        /// </summary>
        [Input("registry")]
        public Input<string>? Registry { get; set; }

        /// <summary>
        /// - DOCKER_HUB: The DockerHub container registry type.
        /// - DOCR: The DigitalOcean container registry type.
        /// </summary>
        [Input("registryType")]
        public Input<Pulumi.DigitalOceanNative.AppsV2.AppsImageSourceSpecRegistryType>? RegistryType { get; set; }

        /// <summary>
        /// The repository name.
        /// </summary>
        [Input("repository")]
        public Input<string>? Repository { get; set; }

        /// <summary>
        /// The repository tag. Defaults to `latest` if not provided.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        public AppsImageSourceSpecArgs()
        {
            Tag = "latest";
        }
        public static new AppsImageSourceSpecArgs Empty => new AppsImageSourceSpecArgs();
    }
}