// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.AppsV2.Outputs
{

    /// <summary>
    /// Papertrail configuration.
    /// </summary>
    [OutputType]
    public sealed class AppLogDestinationPapertrailSpec
    {
        /// <summary>
        /// Papertrail syslog endpoint.
        /// </summary>
        public readonly string Endpoint;

        [OutputConstructor]
        private AppLogDestinationPapertrailSpec(string endpoint)
        {
            Endpoint = endpoint;
        }
    }
}