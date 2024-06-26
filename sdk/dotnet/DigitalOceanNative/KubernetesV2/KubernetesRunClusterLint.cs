// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.KubernetesV2
{
    [DigitalOceanNativeResourceType("digitalocean-native:kubernetes/v2:KubernetesRunClusterLint")]
    public partial class KubernetesRunClusterLint : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An array of checks that will be run when clusterlint executes checks.
        /// </summary>
        [Output("excludeChecks")]
        public Output<ImmutableArray<string>> ExcludeChecks { get; private set; } = null!;

        /// <summary>
        /// An array of check groups that will be omitted when clusterlint executes checks.
        /// </summary>
        [Output("excludeGroups")]
        public Output<ImmutableArray<string>> ExcludeGroups { get; private set; } = null!;

        /// <summary>
        /// An array of checks that will be run when clusterlint executes checks.
        /// </summary>
        [Output("includeChecks")]
        public Output<ImmutableArray<string>> IncludeChecks { get; private set; } = null!;

        /// <summary>
        /// An array of check groups that will be run when clusterlint executes checks.
        /// </summary>
        [Output("includeGroups")]
        public Output<ImmutableArray<string>> IncludeGroups { get; private set; } = null!;

        /// <summary>
        /// ID of the clusterlint run that can be used later to fetch the diagnostics.
        /// </summary>
        [Output("runId")]
        public Output<string?> RunId { get; private set; } = null!;


        /// <summary>
        /// Create a KubernetesRunClusterLint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KubernetesRunClusterLint(string name, KubernetesRunClusterLintArgs? args = null, CustomResourceOptions? options = null)
            : base("digitalocean-native:kubernetes/v2:KubernetesRunClusterLint", name, args ?? new KubernetesRunClusterLintArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KubernetesRunClusterLint(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("digitalocean-native:kubernetes/v2:KubernetesRunClusterLint", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing KubernetesRunClusterLint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KubernetesRunClusterLint Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new KubernetesRunClusterLint(name, id, options);
        }
    }

    public sealed class KubernetesRunClusterLintArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique ID that can be used to reference a Kubernetes cluster.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        [Input("excludeChecks")]
        private InputList<string>? _excludeChecks;

        /// <summary>
        /// An array of checks that will be run when clusterlint executes checks.
        /// </summary>
        public InputList<string> ExcludeChecks
        {
            get => _excludeChecks ?? (_excludeChecks = new InputList<string>());
            set => _excludeChecks = value;
        }

        [Input("excludeGroups")]
        private InputList<string>? _excludeGroups;

        /// <summary>
        /// An array of check groups that will be omitted when clusterlint executes checks.
        /// </summary>
        public InputList<string> ExcludeGroups
        {
            get => _excludeGroups ?? (_excludeGroups = new InputList<string>());
            set => _excludeGroups = value;
        }

        [Input("includeChecks")]
        private InputList<string>? _includeChecks;

        /// <summary>
        /// An array of checks that will be run when clusterlint executes checks.
        /// </summary>
        public InputList<string> IncludeChecks
        {
            get => _includeChecks ?? (_includeChecks = new InputList<string>());
            set => _includeChecks = value;
        }

        [Input("includeGroups")]
        private InputList<string>? _includeGroups;

        /// <summary>
        /// An array of check groups that will be run when clusterlint executes checks.
        /// </summary>
        public InputList<string> IncludeGroups
        {
            get => _includeGroups ?? (_includeGroups = new InputList<string>());
            set => _includeGroups = value;
        }

        public KubernetesRunClusterLintArgs()
        {
        }
        public static new KubernetesRunClusterLintArgs Empty => new KubernetesRunClusterLintArgs();
    }
}
