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
    public sealed class AppWorkerSpec
    {
        /// <summary>
        /// An optional build command to run while building this component from source.
        /// </summary>
        public readonly string? BuildCommand;
        /// <summary>
        /// The path to the Dockerfile relative to the root of the repo. If set, it will be used to build this component. Otherwise, App Platform will attempt to build it using buildpacks.
        /// </summary>
        public readonly string? DockerfilePath;
        /// <summary>
        /// An environment slug describing the type of this app. For a full list, please refer to [the product documentation](https://www.digitalocean.com/docs/app-platform/).
        /// </summary>
        public readonly string? EnvironmentSlug;
        /// <summary>
        /// A list of environment variables made available to the component.
        /// </summary>
        public readonly ImmutableArray<Outputs.AppVariableDefinition> Envs;
        public readonly Outputs.AppsGitSourceSpec? Git;
        public readonly Outputs.AppsGithubSourceSpec? Github;
        public readonly Outputs.AppsGitlabSourceSpec? Gitlab;
        public readonly Outputs.AppsImageSourceSpec? Image;
        /// <summary>
        /// The amount of instances that this component should be scaled to. Default: 1
        /// </summary>
        public readonly int? InstanceCount;
        /// <summary>
        /// The instance size to use for this component. Default: `basic-xxs`
        /// </summary>
        public readonly Pulumi.DigitalOceanNative.AppsV2.AppComponentInstanceBaseInstanceSizeSlug? InstanceSizeSlug;
        public readonly Outputs.AppLogDestinationDefinition? LogDestinations;
        /// <summary>
        /// The name. Must be unique across all components within the same app.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// An optional run command to override the component's default.
        /// </summary>
        public readonly string? RunCommand;
        /// <summary>
        /// An optional path to the working directory to use for the build. For Dockerfile builds, this will be used as the build context. Must be relative to the root of the repo.
        /// </summary>
        public readonly string? SourceDir;

        [OutputConstructor]
        private AppWorkerSpec(
            string? buildCommand,

            string? dockerfilePath,

            string? environmentSlug,

            ImmutableArray<Outputs.AppVariableDefinition> envs,

            Outputs.AppsGitSourceSpec? git,

            Outputs.AppsGithubSourceSpec? github,

            Outputs.AppsGitlabSourceSpec? gitlab,

            Outputs.AppsImageSourceSpec? image,

            int? instanceCount,

            Pulumi.DigitalOceanNative.AppsV2.AppComponentInstanceBaseInstanceSizeSlug? instanceSizeSlug,

            Outputs.AppLogDestinationDefinition? logDestinations,

            string? name,

            string? runCommand,

            string? sourceDir)
        {
            BuildCommand = buildCommand;
            DockerfilePath = dockerfilePath;
            EnvironmentSlug = environmentSlug;
            Envs = envs;
            Git = git;
            Github = github;
            Gitlab = gitlab;
            Image = image;
            InstanceCount = instanceCount;
            InstanceSizeSlug = instanceSizeSlug;
            LogDestinations = logDestinations;
            Name = name;
            RunCommand = runCommand;
            SourceDir = sourceDir;
        }
    }
}