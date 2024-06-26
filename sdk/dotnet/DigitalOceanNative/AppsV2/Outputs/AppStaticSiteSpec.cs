// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.AppsV2.Outputs
{

    [OutputType]
    public sealed class AppStaticSiteSpec
    {
        /// <summary>
        /// An optional build command to run while building this component from source.
        /// </summary>
        public readonly string? BuildCommand;
        /// <summary>
        /// The name of the document to use as the fallback for any requests to documents that are not found when serving this static site. Only 1 of `catchall_document` or `error_document` can be set.
        /// </summary>
        public readonly string? CatchallDocument;
        public readonly Outputs.AppsCorsPolicy? Cors;
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
        /// <summary>
        /// The name of the error document to use when serving this static site. Default: 404.html. If no such file exists within the built assets, App Platform will supply one.
        /// </summary>
        public readonly string? ErrorDocument;
        public readonly Outputs.AppsGitSourceSpec? Git;
        public readonly Outputs.AppsGithubSourceSpec? Github;
        public readonly Outputs.AppsGitlabSourceSpec? Gitlab;
        public readonly Outputs.AppsImageSourceSpec? Image;
        /// <summary>
        /// The name of the index document to use when serving this static site. Default: index.html
        /// </summary>
        public readonly string? IndexDocument;
        public readonly Outputs.AppLogDestinationDefinition? LogDestinations;
        /// <summary>
        /// The name. Must be unique across all components within the same app.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// An optional path to where the built assets will be located, relative to the build context. If not set, App Platform will automatically scan for these directory names: `_static`, `dist`, `public`, `build`.
        /// </summary>
        public readonly string? OutputDir;
        /// <summary>
        /// A list of HTTP routes that should be routed to this component.
        /// </summary>
        public readonly ImmutableArray<Outputs.AppRouteSpec> Routes;
        /// <summary>
        /// An optional run command to override the component's default.
        /// </summary>
        public readonly string? RunCommand;
        /// <summary>
        /// An optional path to the working directory to use for the build. For Dockerfile builds, this will be used as the build context. Must be relative to the root of the repo.
        /// </summary>
        public readonly string? SourceDir;

        [OutputConstructor]
        private AppStaticSiteSpec(
            string? buildCommand,

            string? catchallDocument,

            Outputs.AppsCorsPolicy? cors,

            string? dockerfilePath,

            string? environmentSlug,

            ImmutableArray<Outputs.AppVariableDefinition> envs,

            string? errorDocument,

            Outputs.AppsGitSourceSpec? git,

            Outputs.AppsGithubSourceSpec? github,

            Outputs.AppsGitlabSourceSpec? gitlab,

            Outputs.AppsImageSourceSpec? image,

            string? indexDocument,

            Outputs.AppLogDestinationDefinition? logDestinations,

            string? name,

            string? outputDir,

            ImmutableArray<Outputs.AppRouteSpec> routes,

            string? runCommand,

            string? sourceDir)
        {
            BuildCommand = buildCommand;
            CatchallDocument = catchallDocument;
            Cors = cors;
            DockerfilePath = dockerfilePath;
            EnvironmentSlug = environmentSlug;
            Envs = envs;
            ErrorDocument = errorDocument;
            Git = git;
            Github = github;
            Gitlab = gitlab;
            Image = image;
            IndexDocument = indexDocument;
            LogDestinations = logDestinations;
            Name = name;
            OutputDir = outputDir;
            Routes = routes;
            RunCommand = runCommand;
            SourceDir = sourceDir;
        }
    }
}
