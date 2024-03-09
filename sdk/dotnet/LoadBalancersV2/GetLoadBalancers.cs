// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.LoadBalancersV2
{
    public static class GetLoadBalancers
    {
        public static Task<GetLoadBalancersResult> InvokeAsync(GetLoadBalancersArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLoadBalancersResult>("digitalocean-native:load_balancers/v2:getLoadBalancers", args ?? new GetLoadBalancersArgs(), options.WithDefaults());

        public static Output<GetLoadBalancersResult> Invoke(GetLoadBalancersInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLoadBalancersResult>("digitalocean-native:load_balancers/v2:getLoadBalancers", args ?? new GetLoadBalancersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLoadBalancersArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a load balancer.
        /// </summary>
        [Input("lbId", required: true)]
        public string LbId { get; set; } = null!;

        public GetLoadBalancersArgs()
        {
        }
        public static new GetLoadBalancersArgs Empty => new GetLoadBalancersArgs();
    }

    public sealed class GetLoadBalancersInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for a load balancer.
        /// </summary>
        [Input("lbId", required: true)]
        public Input<string> LbId { get; set; } = null!;

        public GetLoadBalancersInvokeArgs()
        {
        }
        public static new GetLoadBalancersInvokeArgs Empty => new GetLoadBalancersInvokeArgs();
    }


    [OutputType]
    public sealed class GetLoadBalancersResult
    {
        public readonly Outputs.GetLoadBalancersProperties Items;

        [OutputConstructor]
        private GetLoadBalancersResult(Outputs.GetLoadBalancersProperties items)
        {
            Items = items;
        }
    }
}
