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
    [DigitalOceanNativeResourceType("digitalocean-native:load_balancers/v2:LoadBalancersDroplet")]
    public partial class LoadBalancersDroplet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An array containing the IDs of the Droplets assigned to the load balancer.
        /// </summary>
        [Output("dropletIds")]
        public Output<ImmutableArray<int>> DropletIds { get; private set; } = null!;


        /// <summary>
        /// Create a LoadBalancersDroplet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LoadBalancersDroplet(string name, LoadBalancersDropletArgs args, CustomResourceOptions? options = null)
            : base("digitalocean-native:load_balancers/v2:LoadBalancersDroplet", name, args ?? new LoadBalancersDropletArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LoadBalancersDroplet(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("digitalocean-native:load_balancers/v2:LoadBalancersDroplet", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing LoadBalancersDroplet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LoadBalancersDroplet Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LoadBalancersDroplet(name, id, options);
        }
    }

    public sealed class LoadBalancersDropletArgs : global::Pulumi.ResourceArgs
    {
        [Input("dropletIds", required: true)]
        private InputList<int>? _dropletIds;

        /// <summary>
        /// An array containing the IDs of the Droplets assigned to the load balancer.
        /// </summary>
        public InputList<int> DropletIds
        {
            get => _dropletIds ?? (_dropletIds = new InputList<int>());
            set => _dropletIds = value;
        }

        /// <summary>
        /// A unique identifier for a load balancer.
        /// </summary>
        [Input("lbId")]
        public Input<string>? LbId { get; set; }

        public LoadBalancersDropletArgs()
        {
        }
        public static new LoadBalancersDropletArgs Empty => new LoadBalancersDropletArgs();
    }
}
