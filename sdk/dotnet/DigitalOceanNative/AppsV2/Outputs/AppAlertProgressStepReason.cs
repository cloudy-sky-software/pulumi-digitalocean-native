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

    [OutputType]
    public sealed class AppAlertProgressStepReason
    {
        public readonly string? Code;
        public readonly string? Message;

        [OutputConstructor]
        private AppAlertProgressStepReason(
            string? code,

            string? message)
        {
            Code = code;
            Message = message;
        }
    }
}