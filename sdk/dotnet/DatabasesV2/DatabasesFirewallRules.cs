// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.DatabasesV2
{
    [DigitalOceanNativeResourceType("digitalocean-native:databases/v2:DatabasesFirewallRules")]
    public partial class DatabasesFirewallRules : global::Pulumi.CustomResource
    {
        [Output("rules")]
        public Output<ImmutableArray<Outputs.FirewallRule>> Rules { get; private set; } = null!;


        /// <summary>
        /// Create a DatabasesFirewallRules resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatabasesFirewallRules(string name, DatabasesFirewallRulesArgs? args = null, CustomResourceOptions? options = null)
            : base("digitalocean-native:databases/v2:DatabasesFirewallRules", name, args ?? new DatabasesFirewallRulesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatabasesFirewallRules(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("digitalocean-native:databases/v2:DatabasesFirewallRules", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing DatabasesFirewallRules resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatabasesFirewallRules Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DatabasesFirewallRules(name, id, options);
        }
    }

    public sealed class DatabasesFirewallRulesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique identifier for a database cluster.
        /// </summary>
        [Input("databaseClusterUuid")]
        public Input<string>? DatabaseClusterUuid { get; set; }

        [Input("rules")]
        private InputList<Inputs.FirewallRuleArgs>? _rules;
        public InputList<Inputs.FirewallRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.FirewallRuleArgs>());
            set => _rules = value;
        }

        public DatabasesFirewallRulesArgs()
        {
        }
        public static new DatabasesFirewallRulesArgs Empty => new DatabasesFirewallRulesArgs();
    }
}
