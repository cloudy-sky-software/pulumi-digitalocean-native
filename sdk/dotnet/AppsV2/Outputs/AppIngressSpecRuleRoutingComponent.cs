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
    /// The component to route to. Only one of `component` or `redirect` may be set.
    /// </summary>
    [OutputType]
    public sealed class AppIngressSpecRuleRoutingComponent
    {
        /// <summary>
        /// The name of the component to route to.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// An optional flag to preserve the path that is forwarded to the backend service. By default, the HTTP request path will be trimmed from the left when forwarded to the component. For example, a component with `path=/api` will have requests to `/api/list` trimmed to `/list`. If this value is `true`, the path will remain `/api/list`. Note: this is not applicable for Functions Components and is mutually exclusive with `rewrite`.
        /// </summary>
        public readonly string? PreservePathPrefix;
        /// <summary>
        /// An optional field that will rewrite the path of the component to be what is specified here. By default, the HTTP request path will be trimmed from the left when forwarded to the component. For example, a component with `path=/api` will have requests to `/api/list` trimmed to `/list`. If you specified the rewrite to be `/v1/`, requests to `/api/list` would be rewritten to `/v1/list`. Note: this is mutually exclusive with `preserve_path_prefix`.
        /// </summary>
        public readonly string? Rewrite;

        [OutputConstructor]
        private AppIngressSpecRuleRoutingComponent(
            string name,

            string? preservePathPrefix,

            string? rewrite)
        {
            Name = name;
            PreservePathPrefix = preservePathPrefix;
            Rewrite = rewrite;
        }
    }
}
