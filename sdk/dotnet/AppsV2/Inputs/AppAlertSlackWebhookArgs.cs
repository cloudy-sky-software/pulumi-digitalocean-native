// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.AppsV2.Inputs
{

    public sealed class AppAlertSlackWebhookArgs : global::Pulumi.ResourceArgs
    {
        [Input("channel")]
        public Input<string>? Channel { get; set; }

        [Input("url")]
        public Input<string>? Url { get; set; }

        public AppAlertSlackWebhookArgs()
        {
        }
        public static new AppAlertSlackWebhookArgs Empty => new AppAlertSlackWebhookArgs();
    }
}
