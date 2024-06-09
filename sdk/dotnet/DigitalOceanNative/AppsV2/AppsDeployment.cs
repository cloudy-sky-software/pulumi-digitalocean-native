// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.AppsV2
{
    [DigitalOceanNativeResourceType("digitalocean-native:apps/v2:AppsDeployment")]
    public partial class AppsDeployment : global::Pulumi.CustomResource
    {
        [Output("deployment")]
        public Output<Outputs.AppsDeployment?> Deployment { get; private set; } = null!;

        [Output("forceBuild")]
        public Output<bool?> ForceBuild { get; private set; } = null!;


        /// <summary>
        /// Create a AppsDeployment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppsDeployment(string name, AppsDeploymentArgs? args = null, CustomResourceOptions? options = null)
            : base("digitalocean-native:apps/v2:AppsDeployment", name, args ?? new AppsDeploymentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppsDeployment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("digitalocean-native:apps/v2:AppsDeployment", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/cloudy-sky-software/pulumi-digitalocean-native",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AppsDeployment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppsDeployment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AppsDeployment(name, id, options);
        }
    }

    public sealed class AppsDeploymentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The app ID
        /// </summary>
        [Input("appId")]
        public Input<string>? AppId { get; set; }

        [Input("forceBuild")]
        public Input<bool>? ForceBuild { get; set; }

        public AppsDeploymentArgs()
        {
        }
        public static new AppsDeploymentArgs Empty => new AppsDeploymentArgs();
    }
}