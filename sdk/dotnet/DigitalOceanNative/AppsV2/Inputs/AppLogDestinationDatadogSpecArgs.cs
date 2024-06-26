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
    /// DataDog configuration.
    /// </summary>
    public sealed class AppLogDestinationDatadogSpecArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Datadog API key.
        /// </summary>
        [Input("apiKey", required: true)]
        public Input<string> ApiKey { get; set; } = null!;

        /// <summary>
        /// Datadog HTTP log intake endpoint.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        public AppLogDestinationDatadogSpecArgs()
        {
        }
        public static new AppLogDestinationDatadogSpecArgs Empty => new AppLogDestinationDatadogSpecArgs();
    }
}
