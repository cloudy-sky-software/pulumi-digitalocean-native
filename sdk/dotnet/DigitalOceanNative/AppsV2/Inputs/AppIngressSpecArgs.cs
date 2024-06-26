// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.AppsV2.Inputs
{

    /// <summary>
    /// Specification for app ingress configurations.
    /// </summary>
    public sealed class AppIngressSpecArgs : global::Pulumi.ResourceArgs
    {
        [Input("rules")]
        private InputList<Inputs.AppIngressSpecRuleArgs>? _rules;

        /// <summary>
        /// Rules for configuring HTTP ingress for component routes, CORS, rewrites, and redirects.
        /// </summary>
        public InputList<Inputs.AppIngressSpecRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.AppIngressSpecRuleArgs>());
            set => _rules = value;
        }

        public AppIngressSpecArgs()
        {
        }
        public static new AppIngressSpecArgs Empty => new AppIngressSpecArgs();
    }
}
