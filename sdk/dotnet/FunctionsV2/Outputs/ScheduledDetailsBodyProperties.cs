// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOceanNative.FunctionsV2.Outputs
{

    /// <summary>
    /// Optional data to be sent to function while triggering the function.
    /// </summary>
    [OutputType]
    public sealed class ScheduledDetailsBodyProperties
    {
        public readonly string? Name;

        [OutputConstructor]
        private ScheduledDetailsBodyProperties(string? name)
        {
            Name = name;
        }
    }
}
