// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.AppsV2.Inputs
{

    public sealed class AppJobSpecArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional build command to run while building this component from source.
        /// </summary>
        [Input("buildCommand")]
        public Input<string>? BuildCommand { get; set; }

        /// <summary>
        /// The path to the Dockerfile relative to the root of the repo. If set, it will be used to build this component. Otherwise, App Platform will attempt to build it using buildpacks.
        /// </summary>
        [Input("dockerfilePath")]
        public Input<string>? DockerfilePath { get; set; }

        /// <summary>
        /// An environment slug describing the type of this app. For a full list, please refer to [the product documentation](https://www.digitalocean.com/docs/app-platform/).
        /// </summary>
        [Input("environmentSlug")]
        public Input<string>? EnvironmentSlug { get; set; }

        [Input("envs")]
        private InputList<Inputs.AppVariableDefinitionArgs>? _envs;

        /// <summary>
        /// A list of environment variables made available to the component.
        /// </summary>
        public InputList<Inputs.AppVariableDefinitionArgs> Envs
        {
            get => _envs ?? (_envs = new InputList<Inputs.AppVariableDefinitionArgs>());
            set => _envs = value;
        }

        [Input("git")]
        public Input<Inputs.AppsGitSourceSpecArgs>? Git { get; set; }

        [Input("github")]
        public Input<Inputs.AppsGithubSourceSpecArgs>? Github { get; set; }

        [Input("gitlab")]
        public Input<Inputs.AppsGitlabSourceSpecArgs>? Gitlab { get; set; }

        [Input("image")]
        public Input<Inputs.AppsImageSourceSpecArgs>? Image { get; set; }

        /// <summary>
        /// The amount of instances that this component should be scaled to. Default: 1
        /// </summary>
        [Input("instanceCount")]
        public Input<int>? InstanceCount { get; set; }

        /// <summary>
        /// The instance size to use for this component. Default: `basic-xxs`
        /// </summary>
        [Input("instanceSizeSlug")]
        public Input<Pulumi.DigitalOceanNative.AppsV2.AppComponentInstanceBaseInstanceSizeSlug>? InstanceSizeSlug { get; set; }

        /// <summary>
        /// - UNSPECIFIED: Default job type, will auto-complete to POST_DEPLOY kind.
        /// - PRE_DEPLOY: Indicates a job that runs before an app deployment.
        /// - POST_DEPLOY: Indicates a job that runs after an app deployment.
        /// - FAILED_DEPLOY: Indicates a job that runs after a component fails to deploy.
        /// </summary>
        [Input("kind")]
        public Input<Pulumi.DigitalOceanNative.AppsV2.AppJobSpecPropertiesKind>? Kind { get; set; }

        [Input("logDestinations")]
        public Input<Inputs.AppLogDestinationDefinitionArgs>? LogDestinations { get; set; }

        /// <summary>
        /// The name. Must be unique across all components within the same app.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// An optional run command to override the component's default.
        /// </summary>
        [Input("runCommand")]
        public Input<string>? RunCommand { get; set; }

        /// <summary>
        /// An optional path to the working directory to use for the build. For Dockerfile builds, this will be used as the build context. Must be relative to the root of the repo.
        /// </summary>
        [Input("sourceDir")]
        public Input<string>? SourceDir { get; set; }

        public AppJobSpecArgs()
        {
            InstanceCount = 1;
            InstanceSizeSlug = Pulumi.DigitalOceanNative.AppsV2.AppComponentInstanceBaseInstanceSizeSlug.BasicXxs;
            Kind = Pulumi.DigitalOceanNative.AppsV2.AppJobSpecPropertiesKind.Unspecified;
        }
        public static new AppJobSpecArgs Empty => new AppJobSpecArgs();
    }
}
