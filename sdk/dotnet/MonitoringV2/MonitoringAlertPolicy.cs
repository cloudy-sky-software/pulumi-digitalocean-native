// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.MonitoringV2
{
    [DigitalOceanNativeResourceType("digitalocean-native:monitoring/v2:MonitoringAlertPolicy")]
    public partial class MonitoringAlertPolicy : global::Pulumi.CustomResource
    {
        [Output("alerts")]
        public Output<Outputs.Alerts> Alerts { get; private set; } = null!;

        [Output("compare")]
        public Output<Pulumi.DigitalOceanNative.MonitoringV2.Compare> Compare { get; private set; } = null!;

        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        [Output("entities")]
        public Output<ImmutableArray<string>> Entities { get; private set; } = null!;

        [Output("policy")]
        public Output<Outputs.AlertPolicy?> Policy { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        [Output("type")]
        public Output<Pulumi.DigitalOceanNative.MonitoringV2.Type> Type { get; private set; } = null!;

        [Output("value")]
        public Output<double> Value { get; private set; } = null!;

        [Output("window")]
        public Output<Pulumi.DigitalOceanNative.MonitoringV2.Window> Window { get; private set; } = null!;


        /// <summary>
        /// Create a MonitoringAlertPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MonitoringAlertPolicy(string name, MonitoringAlertPolicyArgs args, CustomResourceOptions? options = null)
            : base("digitalocean-native:monitoring/v2:MonitoringAlertPolicy", name, args ?? new MonitoringAlertPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MonitoringAlertPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("digitalocean-native:monitoring/v2:MonitoringAlertPolicy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing MonitoringAlertPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MonitoringAlertPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new MonitoringAlertPolicy(name, id, options);
        }
    }

    public sealed class MonitoringAlertPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("alerts", required: true)]
        public Input<Inputs.AlertsArgs> Alerts { get; set; } = null!;

        [Input("compare", required: true)]
        public Input<Pulumi.DigitalOceanNative.MonitoringV2.Compare> Compare { get; set; } = null!;

        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("entities", required: true)]
        private InputList<string>? _entities;
        public InputList<string> Entities
        {
            get => _entities ?? (_entities = new InputList<string>());
            set => _entities = value;
        }

        [Input("tags", required: true)]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("type", required: true)]
        public Input<Pulumi.DigitalOceanNative.MonitoringV2.Type> Type { get; set; } = null!;

        [Input("value", required: true)]
        public Input<double> Value { get; set; } = null!;

        [Input("window", required: true)]
        public Input<Pulumi.DigitalOceanNative.MonitoringV2.Window> Window { get; set; } = null!;

        public MonitoringAlertPolicyArgs()
        {
        }
        public static new MonitoringAlertPolicyArgs Empty => new MonitoringAlertPolicyArgs();
    }
}