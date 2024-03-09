// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.AppsV2.Inputs
{

    /// <summary>
    /// The redirect configuration for the rule. Only one of `component` or `redirect` may be set.
    /// </summary>
    public sealed class AppIngressSpecRuleRoutingRedirectArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authority/host to redirect to. This can be a hostname or IP address. Note: use `port` to set the port.
        /// </summary>
        [Input("authority")]
        public Input<string>? Authority { get; set; }

        /// <summary>
        /// The port to redirect to.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The redirect code to use. Defaults to `302`. Supported values are 300, 301, 302, 303, 304, 307, 308.
        /// </summary>
        [Input("redirectCode")]
        public Input<int>? RedirectCode { get; set; }

        /// <summary>
        /// The scheme to redirect to. Supported values are `http` or `https`. Default: `https`.
        /// </summary>
        [Input("scheme")]
        public Input<string>? Scheme { get; set; }

        /// <summary>
        /// An optional URI path to redirect to. Note: if this is specified the whole URI of the original request will be overwritten to this value, irrespective of the original request URI being matched.
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public AppIngressSpecRuleRoutingRedirectArgs()
        {
        }
        public static new AppIngressSpecRuleRoutingRedirectArgs Empty => new AppIngressSpecRuleRoutingRedirectArgs();
    }
}