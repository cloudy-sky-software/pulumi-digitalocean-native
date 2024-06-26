// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.DropletsV2
{
    public static class GetDropletAction
    {
        public static Task<GetDropletActionResult> InvokeAsync(GetDropletActionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDropletActionResult>("digitalocean-native:droplets/v2:getDropletAction", args ?? new GetDropletActionArgs(), options.WithDefaults());

        public static Output<GetDropletActionResult> Invoke(GetDropletActionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDropletActionResult>("digitalocean-native:droplets/v2:getDropletAction", args ?? new GetDropletActionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDropletActionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique numeric ID that can be used to identify and reference an action.
        /// </summary>
        [Input("actionId", required: true)]
        public string ActionId { get; set; } = null!;

        /// <summary>
        /// A unique identifier for a Droplet instance.
        /// </summary>
        [Input("dropletId", required: true)]
        public string DropletId { get; set; } = null!;

        public GetDropletActionArgs()
        {
        }
        public static new GetDropletActionArgs Empty => new GetDropletActionArgs();
    }

    public sealed class GetDropletActionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique numeric ID that can be used to identify and reference an action.
        /// </summary>
        [Input("actionId", required: true)]
        public Input<string> ActionId { get; set; } = null!;

        /// <summary>
        /// A unique identifier for a Droplet instance.
        /// </summary>
        [Input("dropletId", required: true)]
        public Input<string> DropletId { get; set; } = null!;

        public GetDropletActionInvokeArgs()
        {
        }
        public static new GetDropletActionInvokeArgs Empty => new GetDropletActionInvokeArgs();
    }


    [OutputType]
    public sealed class GetDropletActionResult
    {
        public readonly Outputs.GetDropletActionProperties Items;

        [OutputConstructor]
        private GetDropletActionResult(Outputs.GetDropletActionProperties items)
        {
            Items = items;
        }
    }
}
