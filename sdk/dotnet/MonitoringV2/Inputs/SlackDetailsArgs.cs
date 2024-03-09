// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.MonitoringV2.Inputs
{

    public sealed class SlackDetailsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Slack channel to notify of an alert trigger.
        /// </summary>
        [Input("channel", required: true)]
        public Input<string> Channel { get; set; } = null!;

        /// <summary>
        /// Slack Webhook URL.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public SlackDetailsArgs()
        {
        }
        public static new SlackDetailsArgs Empty => new SlackDetailsArgs();
    }
}