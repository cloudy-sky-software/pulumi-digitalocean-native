// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.DigitalOceanNative.FunctionsV2.Inputs
{

    /// <summary>
    /// Trigger details for SCHEDULED type, where body is optional.
    /// </summary>
    public sealed class ScheduledDetailsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional data to be sent to function while triggering the function.
        /// </summary>
        [Input("body")]
        public Input<Inputs.ScheduledDetailsBodyPropertiesArgs>? Body { get; set; }

        /// <summary>
        /// valid cron expression string which is required for SCHEDULED type triggers.
        /// </summary>
        [Input("cron", required: true)]
        public Input<string> Cron { get; set; } = null!;

        public ScheduledDetailsArgs()
        {
        }
        public static new ScheduledDetailsArgs Empty => new ScheduledDetailsArgs();
    }
}
