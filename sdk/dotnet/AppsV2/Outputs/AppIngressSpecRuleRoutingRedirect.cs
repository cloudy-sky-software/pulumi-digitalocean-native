// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.AppsV2.Outputs
{

    /// <summary>
    /// The redirect configuration for the rule. Only one of `component` or `redirect` may be set.
    /// </summary>
    [OutputType]
    public sealed class AppIngressSpecRuleRoutingRedirect
    {
        /// <summary>
        /// The authority/host to redirect to. This can be a hostname or IP address. Note: use `port` to set the port.
        /// </summary>
        public readonly string? Authority;
        /// <summary>
        /// The port to redirect to.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// The redirect code to use. Defaults to `302`. Supported values are 300, 301, 302, 303, 304, 307, 308.
        /// </summary>
        public readonly int? RedirectCode;
        /// <summary>
        /// The scheme to redirect to. Supported values are `http` or `https`. Default: `https`.
        /// </summary>
        public readonly string? Scheme;
        /// <summary>
        /// An optional URI path to redirect to. Note: if this is specified the whole URI of the original request will be overwritten to this value, irrespective of the original request URI being matched.
        /// </summary>
        public readonly string? Uri;

        [OutputConstructor]
        private AppIngressSpecRuleRoutingRedirect(
            string? authority,

            int? port,

            int? redirectCode,

            string? scheme,

            string? uri)
        {
            Authority = authority;
            Port = port;
            RedirectCode = redirectCode;
            Scheme = scheme;
            Uri = uri;
        }
    }
}
