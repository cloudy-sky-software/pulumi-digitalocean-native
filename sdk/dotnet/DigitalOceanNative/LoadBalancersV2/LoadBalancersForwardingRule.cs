// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.LoadBalancersV2
{
    [DigitalOceanNativeResourceType("digitalocean-native:load_balancers/v2:LoadBalancersForwardingRule")]
    public partial class LoadBalancersForwardingRule : global::Pulumi.CustomResource
    {
        [Output("forwardingRules")]
        public Output<ImmutableArray<Outputs.ForwardingRule>> ForwardingRules { get; private set; } = null!;


        /// <summary>
        /// Create a LoadBalancersForwardingRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LoadBalancersForwardingRule(string name, LoadBalancersForwardingRuleArgs args, CustomResourceOptions? options = null)
            : base("digitalocean-native:load_balancers/v2:LoadBalancersForwardingRule", name, args ?? new LoadBalancersForwardingRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LoadBalancersForwardingRule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("digitalocean-native:load_balancers/v2:LoadBalancersForwardingRule", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing LoadBalancersForwardingRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LoadBalancersForwardingRule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LoadBalancersForwardingRule(name, id, options);
        }
    }

    public sealed class LoadBalancersForwardingRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("forwardingRules", required: true)]
        private InputList<Inputs.ForwardingRuleArgs>? _forwardingRules;
        public InputList<Inputs.ForwardingRuleArgs> ForwardingRules
        {
            get => _forwardingRules ?? (_forwardingRules = new InputList<Inputs.ForwardingRuleArgs>());
            set => _forwardingRules = value;
        }

        /// <summary>
        /// A unique identifier for a load balancer.
        /// </summary>
        [Input("lbId")]
        public Input<string>? LbId { get; set; }

        public LoadBalancersForwardingRuleArgs()
        {
        }
        public static new LoadBalancersForwardingRuleArgs Empty => new LoadBalancersForwardingRuleArgs();
    }
}
