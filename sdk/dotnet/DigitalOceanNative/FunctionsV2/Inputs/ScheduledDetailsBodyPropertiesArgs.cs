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
    /// Optional data to be sent to function while triggering the function.
    /// </summary>
    public sealed class ScheduledDetailsBodyPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ScheduledDetailsBodyPropertiesArgs()
        {
        }
        public static new ScheduledDetailsBodyPropertiesArgs Empty => new ScheduledDetailsBodyPropertiesArgs();
    }
}