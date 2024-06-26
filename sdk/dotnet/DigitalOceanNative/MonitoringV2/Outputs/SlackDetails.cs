// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.MonitoringV2.Outputs
{

    [OutputType]
    public sealed class SlackDetails
    {
        /// <summary>
        /// Slack channel to notify of an alert trigger.
        /// </summary>
        public readonly string Channel;
        /// <summary>
        /// Slack Webhook URL.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private SlackDetails(
            string channel,

            string url)
        {
            Channel = channel;
            Url = url;
        }
    }
}
